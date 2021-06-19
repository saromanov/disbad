// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package master

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

// MasterClient is the client API for Master service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MasterClient interface {
	JoinExistingCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error)
	JoinMaster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error)
	UpdateMaster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error)
	LeaveCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error)
	GetMaster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Node, error)
	Init(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error)
}

type masterClient struct {
	cc grpc.ClientConnInterface
}

func NewMasterClient(cc grpc.ClientConnInterface) MasterClient {
	return &masterClient{cc}
}

func (c *masterClient) JoinExistingCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/master.Master/JoinExistingCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) JoinMaster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Cluster, error) {
	out := new(Cluster)
	err := c.cc.Invoke(ctx, "/master.Master/JoinMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) UpdateMaster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/master.Master/UpdateMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) LeaveCluster(ctx context.Context, in *Node, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/master.Master/LeaveCluster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) GetMaster(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Node, error) {
	out := new(Node)
	err := c.cc.Invoke(ctx, "/master.Master/GetMaster", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *masterClient) Init(ctx context.Context, in *Cluster, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/master.Master/Init", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MasterServer is the server API for Master service.
// All implementations must embed UnimplementedMasterServer
// for forward compatibility
type MasterServer interface {
	JoinExistingCluster(context.Context, *Node) (*Cluster, error)
	JoinMaster(context.Context, *Node) (*Cluster, error)
	UpdateMaster(context.Context, *Node) (*Response, error)
	LeaveCluster(context.Context, *Node) (*Response, error)
	GetMaster(context.Context, *Cluster) (*Node, error)
	Init(context.Context, *Cluster) (*Response, error)
	mustEmbedUnimplementedMasterServer()
}

// UnimplementedMasterServer must be embedded to have forward compatible implementations.
type UnimplementedMasterServer struct {
}

func (UnimplementedMasterServer) JoinExistingCluster(context.Context, *Node) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinExistingCluster not implemented")
}
func (UnimplementedMasterServer) JoinMaster(context.Context, *Node) (*Cluster, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinMaster not implemented")
}
func (UnimplementedMasterServer) UpdateMaster(context.Context, *Node) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateMaster not implemented")
}
func (UnimplementedMasterServer) LeaveCluster(context.Context, *Node) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LeaveCluster not implemented")
}
func (UnimplementedMasterServer) GetMaster(context.Context, *Cluster) (*Node, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMaster not implemented")
}
func (UnimplementedMasterServer) Init(context.Context, *Cluster) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Init not implemented")
}
func (UnimplementedMasterServer) mustEmbedUnimplementedMasterServer() {}

// UnsafeMasterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MasterServer will
// result in compilation errors.
type UnsafeMasterServer interface {
	mustEmbedUnimplementedMasterServer()
}

func RegisterMasterServer(s grpc.ServiceRegistrar, srv MasterServer) {
	s.RegisterService(&Master_ServiceDesc, srv)
}

func _Master_JoinExistingCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).JoinExistingCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/JoinExistingCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).JoinExistingCluster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_JoinMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).JoinMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/JoinMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).JoinMaster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_UpdateMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).UpdateMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/UpdateMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).UpdateMaster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_LeaveCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Node)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).LeaveCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/LeaveCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).LeaveCluster(ctx, req.(*Node))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_GetMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).GetMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/GetMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).GetMaster(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

func _Master_Init_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Cluster)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MasterServer).Init(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/master.Master/Init",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MasterServer).Init(ctx, req.(*Cluster))
	}
	return interceptor(ctx, in, info, handler)
}

// Master_ServiceDesc is the grpc.ServiceDesc for Master service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Master_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "master.Master",
	HandlerType: (*MasterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "JoinExistingCluster",
			Handler:    _Master_JoinExistingCluster_Handler,
		},
		{
			MethodName: "JoinMaster",
			Handler:    _Master_JoinMaster_Handler,
		},
		{
			MethodName: "UpdateMaster",
			Handler:    _Master_UpdateMaster_Handler,
		},
		{
			MethodName: "LeaveCluster",
			Handler:    _Master_LeaveCluster_Handler,
		},
		{
			MethodName: "GetMaster",
			Handler:    _Master_GetMaster_Handler,
		},
		{
			MethodName: "Init",
			Handler:    _Master_Init_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "master.proto",
}
