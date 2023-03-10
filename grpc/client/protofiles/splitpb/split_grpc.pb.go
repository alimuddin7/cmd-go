// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: split.proto

package splitpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SplitServiceClient is the client API for SplitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SplitServiceClient interface {
	Split(ctx context.Context, in *SplitReq, opts ...grpc.CallOption) (*SplitRes, error)
}

type splitServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSplitServiceClient(cc grpc.ClientConnInterface) SplitServiceClient {
	return &splitServiceClient{cc}
}

func (c *splitServiceClient) Split(ctx context.Context, in *SplitReq, opts ...grpc.CallOption) (*SplitRes, error) {
	out := new(SplitRes)
	err := c.cc.Invoke(ctx, "/protofiles.SplitService/Split", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SplitServiceServer is the server API for SplitService service.
// All implementations should embed UnimplementedSplitServiceServer
// for forward compatibility
type SplitServiceServer interface {
	Split(context.Context, *SplitReq) (*SplitRes, error)
}

// UnimplementedSplitServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSplitServiceServer struct {
}

func (UnimplementedSplitServiceServer) Split(context.Context, *SplitReq) (*SplitRes, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Split not implemented")
}

// UnsafeSplitServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SplitServiceServer will
// result in compilation errors.
type UnsafeSplitServiceServer interface {
	mustEmbedUnimplementedSplitServiceServer()
}

func RegisterSplitServiceServer(s grpc.ServiceRegistrar, srv SplitServiceServer) {
	s.RegisterService(&SplitService_ServiceDesc, srv)
}

func _SplitService_Split_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SplitReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SplitServiceServer).Split(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protofiles.SplitService/Split",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SplitServiceServer).Split(ctx, req.(*SplitReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SplitService_ServiceDesc is the grpc.ServiceDesc for SplitService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SplitService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protofiles.SplitService",
	HandlerType: (*SplitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Split",
			Handler:    _SplitService_Split_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "split.proto",
}
