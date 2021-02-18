// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package sources

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SourcesServiceClient is the client API for SourcesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SourcesServiceClient interface {
	// HealthCheck is a basic healthcheck endpoint
	HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// DetermineLinkSource determines the source of a link
	DetermineLinkSource(ctx context.Context, in *DetermineLinkSourceRequest, opts ...grpc.CallOption) (*DetermineLinkSourceResponse, error)
	//GetSource gets a source
	GetSource(ctx context.Context, in *GetSourceRequest, opts ...grpc.CallOption) (*GetSourceResponse, error)
	// CreateSource creates a source
	CreateSource(ctx context.Context, in *CreateSourceRequest, opts ...grpc.CallOption) (*CreateSourceRequest, error)
}

type sourcesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSourcesServiceClient(cc grpc.ClientConnInterface) SourcesServiceClient {
	return &sourcesServiceClient{cc}
}

func (c *sourcesServiceClient) HealthCheck(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/srcabl.sources.SourcesService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcesServiceClient) DetermineLinkSource(ctx context.Context, in *DetermineLinkSourceRequest, opts ...grpc.CallOption) (*DetermineLinkSourceResponse, error) {
	out := new(DetermineLinkSourceResponse)
	err := c.cc.Invoke(ctx, "/srcabl.sources.SourcesService/DetermineLinkSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcesServiceClient) GetSource(ctx context.Context, in *GetSourceRequest, opts ...grpc.CallOption) (*GetSourceResponse, error) {
	out := new(GetSourceResponse)
	err := c.cc.Invoke(ctx, "/srcabl.sources.SourcesService/GetSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sourcesServiceClient) CreateSource(ctx context.Context, in *CreateSourceRequest, opts ...grpc.CallOption) (*CreateSourceRequest, error) {
	out := new(CreateSourceRequest)
	err := c.cc.Invoke(ctx, "/srcabl.sources.SourcesService/CreateSource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SourcesServiceServer is the server API for SourcesService service.
// All implementations must embed UnimplementedSourcesServiceServer
// for forward compatibility
type SourcesServiceServer interface {
	// HealthCheck is a basic healthcheck endpoint
	HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	// DetermineLinkSource determines the source of a link
	DetermineLinkSource(context.Context, *DetermineLinkSourceRequest) (*DetermineLinkSourceResponse, error)
	//GetSource gets a source
	GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error)
	// CreateSource creates a source
	CreateSource(context.Context, *CreateSourceRequest) (*CreateSourceRequest, error)
	mustEmbedUnimplementedSourcesServiceServer()
}

// UnimplementedSourcesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSourcesServiceServer struct {
}

func (UnimplementedSourcesServiceServer) HealthCheck(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedSourcesServiceServer) DetermineLinkSource(context.Context, *DetermineLinkSourceRequest) (*DetermineLinkSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DetermineLinkSource not implemented")
}
func (UnimplementedSourcesServiceServer) GetSource(context.Context, *GetSourceRequest) (*GetSourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSource not implemented")
}
func (UnimplementedSourcesServiceServer) CreateSource(context.Context, *CreateSourceRequest) (*CreateSourceRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSource not implemented")
}
func (UnimplementedSourcesServiceServer) mustEmbedUnimplementedSourcesServiceServer() {}

// UnsafeSourcesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SourcesServiceServer will
// result in compilation errors.
type UnsafeSourcesServiceServer interface {
	mustEmbedUnimplementedSourcesServiceServer()
}

func RegisterSourcesServiceServer(s grpc.ServiceRegistrar, srv SourcesServiceServer) {
	s.RegisterService(&SourcesService_ServiceDesc, srv)
}

func _SourcesService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcesServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srcabl.sources.SourcesService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcesServiceServer).HealthCheck(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourcesService_DetermineLinkSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetermineLinkSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcesServiceServer).DetermineLinkSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srcabl.sources.SourcesService/DetermineLinkSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcesServiceServer).DetermineLinkSource(ctx, req.(*DetermineLinkSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourcesService_GetSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcesServiceServer).GetSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srcabl.sources.SourcesService/GetSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcesServiceServer).GetSource(ctx, req.(*GetSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SourcesService_CreateSource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SourcesServiceServer).CreateSource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/srcabl.sources.SourcesService/CreateSource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SourcesServiceServer).CreateSource(ctx, req.(*CreateSourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SourcesService_ServiceDesc is the grpc.ServiceDesc for SourcesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SourcesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "srcabl.sources.SourcesService",
	HandlerType: (*SourcesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _SourcesService_HealthCheck_Handler,
		},
		{
			MethodName: "DetermineLinkSource",
			Handler:    _SourcesService_DetermineLinkSource_Handler,
		},
		{
			MethodName: "GetSource",
			Handler:    _SourcesService_GetSource_Handler,
		},
		{
			MethodName: "CreateSource",
			Handler:    _SourcesService_CreateSource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sources/service.proto",
}
