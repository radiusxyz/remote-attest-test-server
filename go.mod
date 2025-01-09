module remote-attest-test-server-go

go 1.23.4

replace remote-attest-test-server-go/proto => ./proto

require (
	github.com/google/go-tdx-guest v0.3.1
	google.golang.org/grpc v1.69.2
	google.golang.org/protobuf v1.35.1
)

require (
	github.com/google/go-configfs-tsm v0.2.2 // indirect
	github.com/google/logger v1.1.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.28.0 // indirect
	golang.org/x/net v0.30.0 // indirect
	golang.org/x/sys v0.26.0 // indirect
	golang.org/x/text v0.19.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241015192408-796eee8c2d53 // indirect
)
