// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v4.25.1
// source: attest/attest.proto

package attest

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	Attest_Attest_FullMethodName = "/attest.Attest/Attest"
)

// AttestClient is the client API for Attest service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AttestClient interface {
	Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestReply, error)
}

type attestClient struct {
	cc grpc.ClientConnInterface
}

func NewAttestClient(cc grpc.ClientConnInterface) AttestClient {
	return &attestClient{cc}
}

func (c *attestClient) Attest(ctx context.Context, in *AttestRequest, opts ...grpc.CallOption) (*AttestReply, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AttestReply)
	err := c.cc.Invoke(ctx, Attest_Attest_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AttestServer is the server API for Attest service.
// All implementations must embed UnimplementedAttestServer
// for forward compatibility.
type AttestServer interface {
	Attest(context.Context, *AttestRequest) (*AttestReply, error)
	mustEmbedUnimplementedAttestServer()
}

// UnimplementedAttestServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAttestServer struct{}

func (UnimplementedAttestServer) Attest(context.Context, *AttestRequest) (*AttestReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Attest not implemented")
}
func (UnimplementedAttestServer) mustEmbedUnimplementedAttestServer() {}
func (UnimplementedAttestServer) testEmbeddedByValue()                {}

// UnsafeAttestServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AttestServer will
// result in compilation errors.
type UnsafeAttestServer interface {
	mustEmbedUnimplementedAttestServer()
}

func RegisterAttestServer(s grpc.ServiceRegistrar, srv AttestServer) {
	// If the following call pancis, it indicates UnimplementedAttestServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Attest_ServiceDesc, srv)
}

func _Attest_Attest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AttestRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AttestServer).Attest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Attest_Attest_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AttestServer).Attest(ctx, req.(*AttestRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Attest_ServiceDesc is the grpc.ServiceDesc for Attest service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Attest_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "attest.Attest",
	HandlerType: (*AttestServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Attest",
			Handler:    _Attest_Attest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "attest/attest.proto",
}
