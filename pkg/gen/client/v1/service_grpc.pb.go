// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: client/v1/service.proto

package v1

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

// TailscaleEphemeralExitNodesServiceClient is the client API for TailscaleEphemeralExitNodesService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TailscaleEphemeralExitNodesServiceClient interface {
	HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error)
	GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error)
	GetDefaultProvider(ctx context.Context, in *GetDefaultProviderRequest, opts ...grpc.CallOption) (*GetDefaultProviderResponse, error)
	ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error)
	GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error)
	StartNode(ctx context.Context, in *StartNodeRequest, opts ...grpc.CallOption) (*StartNodeResponse, error)
	StopNode(ctx context.Context, in *StopNodeRequest, opts ...grpc.CallOption) (*StopNodeResponse, error)
	ProvisionNode(ctx context.Context, in *ProvisionNodeRequest, opts ...grpc.CallOption) (*ProvisionNodeResponse, error)
	DeprovisionNode(ctx context.Context, in *DeprovisionNodeRequest, opts ...grpc.CallOption) (*DeprovisionNodeResponse, error)
	GetExecution(ctx context.Context, in *GetExecutionRequest, opts ...grpc.CallOption) (*GetExecutionResponse, error)
}

type tailscaleEphemeralExitNodesServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTailscaleEphemeralExitNodesServiceClient(cc grpc.ClientConnInterface) TailscaleEphemeralExitNodesServiceClient {
	return &tailscaleEphemeralExitNodesServiceClient{cc}
}

func (c *tailscaleEphemeralExitNodesServiceClient) HealthCheck(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/HealthCheck", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error) {
	out := new(ListProvidersResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error) {
	out := new(GetProviderResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/GetProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) GetDefaultProvider(ctx context.Context, in *GetDefaultProviderRequest, opts ...grpc.CallOption) (*GetDefaultProviderResponse, error) {
	out := new(GetDefaultProviderResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/GetDefaultProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) ListNodes(ctx context.Context, in *ListNodesRequest, opts ...grpc.CallOption) (*ListNodesResponse, error) {
	out := new(ListNodesResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/ListNodes", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) GetNode(ctx context.Context, in *GetNodeRequest, opts ...grpc.CallOption) (*GetNodeResponse, error) {
	out := new(GetNodeResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/GetNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) StartNode(ctx context.Context, in *StartNodeRequest, opts ...grpc.CallOption) (*StartNodeResponse, error) {
	out := new(StartNodeResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/StartNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) StopNode(ctx context.Context, in *StopNodeRequest, opts ...grpc.CallOption) (*StopNodeResponse, error) {
	out := new(StopNodeResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/StopNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) ProvisionNode(ctx context.Context, in *ProvisionNodeRequest, opts ...grpc.CallOption) (*ProvisionNodeResponse, error) {
	out := new(ProvisionNodeResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/ProvisionNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) DeprovisionNode(ctx context.Context, in *DeprovisionNodeRequest, opts ...grpc.CallOption) (*DeprovisionNodeResponse, error) {
	out := new(DeprovisionNodeResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/DeprovisionNode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *tailscaleEphemeralExitNodesServiceClient) GetExecution(ctx context.Context, in *GetExecutionRequest, opts ...grpc.CallOption) (*GetExecutionResponse, error) {
	out := new(GetExecutionResponse)
	err := c.cc.Invoke(ctx, "/client.v1.TailscaleEphemeralExitNodesService/GetExecution", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TailscaleEphemeralExitNodesServiceServer is the server API for TailscaleEphemeralExitNodesService service.
// All implementations must embed UnimplementedTailscaleEphemeralExitNodesServiceServer
// for forward compatibility
type TailscaleEphemeralExitNodesServiceServer interface {
	HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
	ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error)
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error)
	GetDefaultProvider(context.Context, *GetDefaultProviderRequest) (*GetDefaultProviderResponse, error)
	ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error)
	GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error)
	StartNode(context.Context, *StartNodeRequest) (*StartNodeResponse, error)
	StopNode(context.Context, *StopNodeRequest) (*StopNodeResponse, error)
	ProvisionNode(context.Context, *ProvisionNodeRequest) (*ProvisionNodeResponse, error)
	DeprovisionNode(context.Context, *DeprovisionNodeRequest) (*DeprovisionNodeResponse, error)
	GetExecution(context.Context, *GetExecutionRequest) (*GetExecutionResponse, error)
	mustEmbedUnimplementedTailscaleEphemeralExitNodesServiceServer()
}

// UnimplementedTailscaleEphemeralExitNodesServiceServer must be embedded to have forward compatible implementations.
type UnimplementedTailscaleEphemeralExitNodesServiceServer struct {
}

func (UnimplementedTailscaleEphemeralExitNodesServiceServer) HealthCheck(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method HealthCheck not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvider not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) GetDefaultProvider(context.Context, *GetDefaultProviderRequest) (*GetDefaultProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDefaultProvider not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) ListNodes(context.Context, *ListNodesRequest) (*ListNodesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListNodes not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) GetNode(context.Context, *GetNodeRequest) (*GetNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNode not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) StartNode(context.Context, *StartNodeRequest) (*StartNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StartNode not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) StopNode(context.Context, *StopNodeRequest) (*StopNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StopNode not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) ProvisionNode(context.Context, *ProvisionNodeRequest) (*ProvisionNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProvisionNode not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) DeprovisionNode(context.Context, *DeprovisionNodeRequest) (*DeprovisionNodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeprovisionNode not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) GetExecution(context.Context, *GetExecutionRequest) (*GetExecutionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetExecution not implemented")
}
func (UnimplementedTailscaleEphemeralExitNodesServiceServer) mustEmbedUnimplementedTailscaleEphemeralExitNodesServiceServer() {
}

// UnsafeTailscaleEphemeralExitNodesServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TailscaleEphemeralExitNodesServiceServer will
// result in compilation errors.
type UnsafeTailscaleEphemeralExitNodesServiceServer interface {
	mustEmbedUnimplementedTailscaleEphemeralExitNodesServiceServer()
}

func RegisterTailscaleEphemeralExitNodesServiceServer(s grpc.ServiceRegistrar, srv TailscaleEphemeralExitNodesServiceServer) {
	s.RegisterService(&TailscaleEphemeralExitNodesService_ServiceDesc, srv)
}

func _TailscaleEphemeralExitNodesService_HealthCheck_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).HealthCheck(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/HealthCheck",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).HealthCheck(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ListProviders(ctx, req.(*ListProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/GetProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetProvider(ctx, req.(*GetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_GetDefaultProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDefaultProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetDefaultProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/GetDefaultProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetDefaultProvider(ctx, req.(*GetDefaultProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_ListNodes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListNodesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ListNodes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/ListNodes",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ListNodes(ctx, req.(*ListNodesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_GetNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/GetNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetNode(ctx, req.(*GetNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_StartNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StartNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).StartNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/StartNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).StartNode(ctx, req.(*StartNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_StopNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StopNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).StopNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/StopNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).StopNode(ctx, req.(*StopNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_ProvisionNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProvisionNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ProvisionNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/ProvisionNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).ProvisionNode(ctx, req.(*ProvisionNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_DeprovisionNode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeprovisionNodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).DeprovisionNode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/DeprovisionNode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).DeprovisionNode(ctx, req.(*DeprovisionNodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TailscaleEphemeralExitNodesService_GetExecution_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetExecutionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetExecution(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/client.v1.TailscaleEphemeralExitNodesService/GetExecution",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TailscaleEphemeralExitNodesServiceServer).GetExecution(ctx, req.(*GetExecutionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TailscaleEphemeralExitNodesService_ServiceDesc is the grpc.ServiceDesc for TailscaleEphemeralExitNodesService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TailscaleEphemeralExitNodesService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "client.v1.TailscaleEphemeralExitNodesService",
	HandlerType: (*TailscaleEphemeralExitNodesServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "HealthCheck",
			Handler:    _TailscaleEphemeralExitNodesService_HealthCheck_Handler,
		},
		{
			MethodName: "ListProviders",
			Handler:    _TailscaleEphemeralExitNodesService_ListProviders_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _TailscaleEphemeralExitNodesService_GetProvider_Handler,
		},
		{
			MethodName: "GetDefaultProvider",
			Handler:    _TailscaleEphemeralExitNodesService_GetDefaultProvider_Handler,
		},
		{
			MethodName: "ListNodes",
			Handler:    _TailscaleEphemeralExitNodesService_ListNodes_Handler,
		},
		{
			MethodName: "GetNode",
			Handler:    _TailscaleEphemeralExitNodesService_GetNode_Handler,
		},
		{
			MethodName: "StartNode",
			Handler:    _TailscaleEphemeralExitNodesService_StartNode_Handler,
		},
		{
			MethodName: "StopNode",
			Handler:    _TailscaleEphemeralExitNodesService_StopNode_Handler,
		},
		{
			MethodName: "ProvisionNode",
			Handler:    _TailscaleEphemeralExitNodesService_ProvisionNode_Handler,
		},
		{
			MethodName: "DeprovisionNode",
			Handler:    _TailscaleEphemeralExitNodesService_DeprovisionNode_Handler,
		},
		{
			MethodName: "GetExecution",
			Handler:    _TailscaleEphemeralExitNodesService_GetExecution_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "client/v1/service.proto",
}
