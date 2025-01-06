use serde::de::{Deserializer, Error as SerdeError};
use serde::Deserialize;
use jsonrpsee::proc_macros::rpc;
use jsonrpsee::core::{async_trait, RpcResult};
use jsonrpsee::server::ServerBuilder;
use std::net::SocketAddr;
use tdx_attest_rs::*;
use jsonrpsee_types::error::ErrorObject;

fn deserialize_array_64<'de, D>(deserializer: D) -> Result<[u8; 64], D::Error>
where
    D: Deserializer<'de>,
{
    let data: Vec<u8> = Deserialize::deserialize(deserializer)?;
    let len = data.len();
    if len != 64 {
        return Err(SerdeError::invalid_length(len, &"64"));
    }
    let mut arr = [0; 64];
    arr.copy_from_slice(&data);
    Ok(arr)
}

#[derive(Deserialize)]
#[serde(transparent)]
pub struct Bytes64Param(
    #[serde(deserialize_with = "deserialize_array_64")]
    pub [u8; 64]
);

#[rpc(server)]
pub trait MyRpc {
    #[method(name = "hello")]
    async fn hello(&self) -> RpcResult<String>;

    #[method(name = "get_quote")]
    async fn get_quote(&self, param: Bytes64Param) -> RpcResult<String>;
}

pub struct MyRpcImpl;

#[async_trait]
impl MyRpcServer for MyRpcImpl {
    async fn hello(&self) -> RpcResult<String> {
        Ok("Hello from Rust + jsonrpsee!".to_string())
    }

    async fn get_quote(&self, param: Bytes64Param) -> RpcResult<String> {
        let mut tdx_report_data = tdx_report_data_t { d: [0; 64] };
        tdx_report_data.d.copy_from_slice(&param.0);
        let (error, quote) = tdx_att_get_quote(Some(&tdx_report_data), None, None, 0);

        println!("{}", format!("error: {:#X}", error as u32));
        println!("{}", format!("quote: {:?}", quote));

        if error == tdx_attest_error_t::TDX_ATTEST_SUCCESS {
            if let Some(quote) = quote {
                Ok(String::from_utf8(quote)
                    .unwrap_or_else(|_| "Failed to convert quote to string".to_string()))
            } else {
                Ok("tdx_att_get_quote: No quote returned".to_string())
            }
        } else {
            Err(ErrorObject::owned(
                (error as u32).try_into().unwrap(),
                "Failed to get quote",
                None::<()>,
            ))
        }
    }
}

#[tokio::main]
async fn main() -> anyhow::Result<()> {
    let addr = SocketAddr::from(([0, 0, 0, 0], 3000));
    let server = ServerBuilder::default().build(addr).await?;
    let handle = server.start(MyRpcImpl.into_rpc());
    println!("JSON-RPC server running at http://{}", addr);
    handle.stopped().await;
    Ok(())
}
