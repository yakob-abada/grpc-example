// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.3
// source: proto/explore-service.proto

package explore

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
	ExploreService_ListLikedYou_FullMethodName    = "/explore.ExploreService/ListLikedYou"
	ExploreService_ListNewLikedYou_FullMethodName = "/explore.ExploreService/ListNewLikedYou"
	ExploreService_CountLikedYou_FullMethodName   = "/explore.ExploreService/CountLikedYou"
	ExploreService_PutDecision_FullMethodName     = "/explore.ExploreService/PutDecision"
)

// ExploreServiceClient is the client API for ExploreService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExploreServiceClient interface {
	ListLikedYou(ctx context.Context, in *ListLikedYouRequest, opts ...grpc.CallOption) (*ListLikedYouResponse, error)
	ListNewLikedYou(ctx context.Context, in *ListLikedYouRequest, opts ...grpc.CallOption) (*ListLikedYouResponse, error)
	CountLikedYou(ctx context.Context, in *CountLikedYouRequest, opts ...grpc.CallOption) (*CountLikedYouResponse, error)
	PutDecision(ctx context.Context, in *PutDecisionRequest, opts ...grpc.CallOption) (*PutDecisionResponse, error)
}

type exploreServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewExploreServiceClient(cc grpc.ClientConnInterface) ExploreServiceClient {
	return &exploreServiceClient{cc}
}

func (c *exploreServiceClient) ListLikedYou(ctx context.Context, in *ListLikedYouRequest, opts ...grpc.CallOption) (*ListLikedYouResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLikedYouResponse)
	err := c.cc.Invoke(ctx, ExploreService_ListLikedYou_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreServiceClient) ListNewLikedYou(ctx context.Context, in *ListLikedYouRequest, opts ...grpc.CallOption) (*ListLikedYouResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListLikedYouResponse)
	err := c.cc.Invoke(ctx, ExploreService_ListNewLikedYou_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreServiceClient) CountLikedYou(ctx context.Context, in *CountLikedYouRequest, opts ...grpc.CallOption) (*CountLikedYouResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CountLikedYouResponse)
	err := c.cc.Invoke(ctx, ExploreService_CountLikedYou_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *exploreServiceClient) PutDecision(ctx context.Context, in *PutDecisionRequest, opts ...grpc.CallOption) (*PutDecisionResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutDecisionResponse)
	err := c.cc.Invoke(ctx, ExploreService_PutDecision_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExploreServiceServer is the server API for ExploreService service.
// All implementations must embed UnimplementedExploreServiceServer
// for forward compatibility.
type ExploreServiceServer interface {
	ListLikedYou(context.Context, *ListLikedYouRequest) (*ListLikedYouResponse, error)
	ListNewLikedYou(context.Context, *ListLikedYouRequest) (*ListLikedYouResponse, error)
	CountLikedYou(context.Context, *CountLikedYouRequest) (*CountLikedYouResponse, error)
	PutDecision(context.Context, *PutDecisionRequest) (*PutDecisionResponse, error)
	mustEmbedUnimplementedExploreServiceServer()
}

// UnimplementedExploreServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedExploreServiceServer struct{}

func (UnimplementedExploreServiceServer) ListLikedYou(context.Context, *ListLikedYouRequest) (*ListLikedYouResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLikedYou not implemented")
}
func (UnimplementedExploreServiceServer) ListNewLikedYou(context.Context, *ListLikedYouRequest) (*ListLikedYouResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNewLikedYou not implemented")
}
func (UnimplementedExploreServiceServer) CountLikedYou(context.Context, *CountLikedYouRequest) (*CountLikedYouResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CountLikedYou not implemented")
}
func (UnimplementedExploreServiceServer) PutDecision(context.Context, *PutDecisionRequest) (*PutDecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutDecision not implemented")
}
func (UnimplementedExploreServiceServer) mustEmbedUnimplementedExploreServiceServer() {}
func (UnimplementedExploreServiceServer) testEmbeddedByValue()                        {}

// UnsafeExploreServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExploreServiceServer will
// result in compilation errors.
type UnsafeExploreServiceServer interface {
	mustEmbedUnimplementedExploreServiceServer()
}

func RegisterExploreServiceServer(s grpc.ServiceRegistrar, srv ExploreServiceServer) {
	// If the following call pancis, it indicates UnimplementedExploreServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ExploreService_ServiceDesc, srv)
}

func _ExploreService_ListLikedYou_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLikedYouRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExploreServiceServer).ListLikedYou(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExploreService_ListLikedYou_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExploreServiceServer).ListLikedYou(ctx, req.(*ListLikedYouRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExploreService_ListNewLikedYou_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLikedYouRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExploreServiceServer).ListNewLikedYou(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExploreService_ListNewLikedYou_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExploreServiceServer).ListNewLikedYou(ctx, req.(*ListLikedYouRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExploreService_CountLikedYou_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CountLikedYouRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExploreServiceServer).CountLikedYou(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExploreService_CountLikedYou_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExploreServiceServer).CountLikedYou(ctx, req.(*CountLikedYouRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ExploreService_PutDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutDecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExploreServiceServer).PutDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ExploreService_PutDecision_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExploreServiceServer).PutDecision(ctx, req.(*PutDecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ExploreService_ServiceDesc is the grpc.ServiceDesc for ExploreService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ExploreService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "explore.ExploreService",
	HandlerType: (*ExploreServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLikedYou",
			Handler:    _ExploreService_ListLikedYou_Handler,
		},
		{
			MethodName: "ListNewLikedYou",
			Handler:    _ExploreService_ListNewLikedYou_Handler,
		},
		{
			MethodName: "CountLikedYou",
			Handler:    _ExploreService_CountLikedYou_Handler,
		},
		{
			MethodName: "PutDecision",
			Handler:    _ExploreService_PutDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/explore-service.proto",
}
