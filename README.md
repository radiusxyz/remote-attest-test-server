### In TD
```bash
cd ~/remote-attest-test-server-go
go build -o myserver server/main.go
sudo ./myserver
```

### In Local
```bash
grpcurl -plaintext \
  -d '{"report_data": "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=="}' \
  <td-ip>:<port> \
  attest.Attest/Attest

# { "quote": ... }
```

