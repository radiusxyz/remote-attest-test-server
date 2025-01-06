use jsonrpsee::proc_macros::rpc;
use jsonrpsee::core::{async_trait, RpcResult};
use jsonrpsee::server::ServerBuilder;
use std::net::SocketAddr;
use tdx_attest_rs::*;
use jsonrpsee_types::error::ErrorObject;

#[rpc(server)]
pub trait MyRpc {
    #[method(name = "hello")]
    async fn hello(&self) -> RpcResult<String>;

    #[method(name = "get_quote")]
    async fn get_quote(&self) -> RpcResult<String>;
}

pub struct MyRpcImpl;

#[async_trait]
impl MyRpcServer for MyRpcImpl {
    async fn hello(&self) -> RpcResult<String> {
        Ok("Hello from Rust + jsonrpsee!".to_string())
    }

    async fn get_quote(&self) -> RpcResult<String> {
        let report_data_bytes = [0u8; 64];

        let mut tdx_report_data = tdx_report_data_t { d: [0; 64usize] };
        tdx_report_data.d.copy_from_slice(&report_data_bytes);

        let (error, quote) = tdx_att_get_quote(Some(&tdx_report_data), None, None, 0);

        println!("{}", format!("error: {:#X}", error as u32));
        println!("{}", format!("quote: {:?}", quote));

        if error == tdx_attest_error_t::TDX_ATTEST_SUCCESS {
            if let Some(quote) = quote {    
                Ok(String::from_utf8(quote).unwrap_or_else(|_| "Failed to convert quote to string".to_string()))
            } else {
                Ok("tdx_att_get_quote: No quote returned".to_string())
            }
        } else {
            Err(ErrorObject::owned((error as u32).try_into().unwrap(), "Failed to get quote", None::<()>))
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
