### In TD
```bash
cargo run
```

### In Local
```bash
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{"jsonrpc":"2.0","method":"get_quote","params":[[0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0]],"id":1}' \
  "http://<td-ip>:3000"

# {"jsonrpc":"2.0","id":1,"error":{"code":10,"message":"Failed to get quote"}}
```