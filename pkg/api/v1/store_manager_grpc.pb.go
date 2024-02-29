// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: api/proto/v1/store_manager.proto

// This document describes the store manager API for NBN23 SwishPay platform.

package _go

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

const (
	PayService_GetUser_FullMethodName         = "/store.manager.v1.PayService/GetUser"
	PayService_GetSubscription_FullMethodName = "/store.manager.v1.PayService/GetSubscription"
	PayService_UpsertUser_FullMethodName      = "/store.manager.v1.PayService/UpsertUser"
	PayService_AddMember_FullMethodName       = "/store.manager.v1.PayService/AddMember"
	PayService_RemoveMember_FullMethodName    = "/store.manager.v1.PayService/RemoveMember"
)

// PayServiceClient is the client API for PayService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PayServiceClient interface {
	// GetUser
	//
	// GetUser returns the user with the given id.
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// GetSubscription
	//
	// GetSubscription returns the subscription with the given id.
	GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error)
	// UpsertUser
	//
	// UpsertUser creates or updates a user.
	UpsertUser(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error)
	// AddMember
	//
	// AddMember adds a member to a subscription.
	AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error)
	// RemoveMember
	//
	// RemoveMember removes a member from a subscription.
	RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*RemoveMemberResponse, error)
}

type payServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPayServiceClient(cc grpc.ClientConnInterface) PayServiceClient {
	return &payServiceClient{cc}
}

func (c *payServiceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, PayService_GetUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) GetSubscription(ctx context.Context, in *GetSubscriptionRequest, opts ...grpc.CallOption) (*GetSubscriptionResponse, error) {
	out := new(GetSubscriptionResponse)
	err := c.cc.Invoke(ctx, PayService_GetSubscription_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) UpsertUser(ctx context.Context, in *UpsertUserRequest, opts ...grpc.CallOption) (*UpsertUserResponse, error) {
	out := new(UpsertUserResponse)
	err := c.cc.Invoke(ctx, PayService_UpsertUser_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) AddMember(ctx context.Context, in *AddMemberRequest, opts ...grpc.CallOption) (*AddMemberResponse, error) {
	out := new(AddMemberResponse)
	err := c.cc.Invoke(ctx, PayService_AddMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *payServiceClient) RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*RemoveMemberResponse, error) {
	out := new(RemoveMemberResponse)
	err := c.cc.Invoke(ctx, PayService_RemoveMember_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PayServiceServer is the server API for PayService service.
// All implementations must embed UnimplementedPayServiceServer
// for forward compatibility
type PayServiceServer interface {
	// GetUser
	//
	// GetUser returns the user with the given id.
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// GetSubscription
	//
	// GetSubscription returns the subscription with the given id.
	GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error)
	// UpsertUser
	//
	// UpsertUser creates or updates a user.
	UpsertUser(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error)
	// AddMember
	//
	// AddMember adds a member to a subscription.
	AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error)
	// RemoveMember
	//
	// RemoveMember removes a member from a subscription.
	RemoveMember(context.Context, *RemoveMemberRequest) (*RemoveMemberResponse, error)
	mustEmbedUnimplementedPayServiceServer()
}

// UnimplementedPayServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPayServiceServer struct {
}

func (UnimplementedPayServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedPayServiceServer) GetSubscription(context.Context, *GetSubscriptionRequest) (*GetSubscriptionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSubscription not implemented")
}
func (UnimplementedPayServiceServer) UpsertUser(context.Context, *UpsertUserRequest) (*UpsertUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertUser not implemented")
}
func (UnimplementedPayServiceServer) AddMember(context.Context, *AddMemberRequest) (*AddMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddMember not implemented")
}
func (UnimplementedPayServiceServer) RemoveMember(context.Context, *RemoveMemberRequest) (*RemoveMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMember not implemented")
}
func (UnimplementedPayServiceServer) mustEmbedUnimplementedPayServiceServer() {}

// UnsafePayServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PayServiceServer will
// result in compilation errors.
type UnsafePayServiceServer interface {
	mustEmbedUnimplementedPayServiceServer()
}

func RegisterPayServiceServer(s grpc.ServiceRegistrar, srv PayServiceServer) {
	s.RegisterService(&PayService_ServiceDesc, srv)
}

func _PayService_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayService_GetUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_GetSubscription_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSubscriptionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).GetSubscription(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayService_GetSubscription_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).GetSubscription(ctx, req.(*GetSubscriptionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_UpsertUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).UpsertUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayService_UpsertUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).UpsertUser(ctx, req.(*UpsertUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_AddMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).AddMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayService_AddMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).AddMember(ctx, req.(*AddMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PayService_RemoveMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PayServiceServer).RemoveMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PayService_RemoveMember_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PayServiceServer).RemoveMember(ctx, req.(*RemoveMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PayService_ServiceDesc is the grpc.ServiceDesc for PayService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PayService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.manager.v1.PayService",
	HandlerType: (*PayServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _PayService_GetUser_Handler,
		},
		{
			MethodName: "GetSubscription",
			Handler:    _PayService_GetSubscription_Handler,
		},
		{
			MethodName: "UpsertUser",
			Handler:    _PayService_UpsertUser_Handler,
		},
		{
			MethodName: "AddMember",
			Handler:    _PayService_AddMember_Handler,
		},
		{
			MethodName: "RemoveMember",
			Handler:    _PayService_RemoveMember_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/store_manager.proto",
}

const (
	SubscriptionEventService_OnSubscriptionEvent_FullMethodName = "/store.manager.v1.SubscriptionEventService/OnSubscriptionEvent"
)

// SubscriptionEventServiceClient is the client API for SubscriptionEventService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SubscriptionEventServiceClient interface {
	// OnSubscriptionEvent
	//
	// OnSubscriptionEvent handles a subscription event.
	OnSubscriptionEvent(ctx context.Context, in *OnSubscriptionEventRequest, opts ...grpc.CallOption) (*OnSubscriptionEventResponse, error)
}

type subscriptionEventServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSubscriptionEventServiceClient(cc grpc.ClientConnInterface) SubscriptionEventServiceClient {
	return &subscriptionEventServiceClient{cc}
}

func (c *subscriptionEventServiceClient) OnSubscriptionEvent(ctx context.Context, in *OnSubscriptionEventRequest, opts ...grpc.CallOption) (*OnSubscriptionEventResponse, error) {
	out := new(OnSubscriptionEventResponse)
	err := c.cc.Invoke(ctx, SubscriptionEventService_OnSubscriptionEvent_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SubscriptionEventServiceServer is the server API for SubscriptionEventService service.
// All implementations must embed UnimplementedSubscriptionEventServiceServer
// for forward compatibility
type SubscriptionEventServiceServer interface {
	// OnSubscriptionEvent
	//
	// OnSubscriptionEvent handles a subscription event.
	OnSubscriptionEvent(context.Context, *OnSubscriptionEventRequest) (*OnSubscriptionEventResponse, error)
	mustEmbedUnimplementedSubscriptionEventServiceServer()
}

// UnimplementedSubscriptionEventServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSubscriptionEventServiceServer struct {
}

func (UnimplementedSubscriptionEventServiceServer) OnSubscriptionEvent(context.Context, *OnSubscriptionEventRequest) (*OnSubscriptionEventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OnSubscriptionEvent not implemented")
}
func (UnimplementedSubscriptionEventServiceServer) mustEmbedUnimplementedSubscriptionEventServiceServer() {
}

// UnsafeSubscriptionEventServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SubscriptionEventServiceServer will
// result in compilation errors.
type UnsafeSubscriptionEventServiceServer interface {
	mustEmbedUnimplementedSubscriptionEventServiceServer()
}

func RegisterSubscriptionEventServiceServer(s grpc.ServiceRegistrar, srv SubscriptionEventServiceServer) {
	s.RegisterService(&SubscriptionEventService_ServiceDesc, srv)
}

func _SubscriptionEventService_OnSubscriptionEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OnSubscriptionEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SubscriptionEventServiceServer).OnSubscriptionEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SubscriptionEventService_OnSubscriptionEvent_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SubscriptionEventServiceServer).OnSubscriptionEvent(ctx, req.(*OnSubscriptionEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SubscriptionEventService_ServiceDesc is the grpc.ServiceDesc for SubscriptionEventService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SubscriptionEventService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "store.manager.v1.SubscriptionEventService",
	HandlerType: (*SubscriptionEventServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "OnSubscriptionEvent",
			Handler:    _SubscriptionEventService_OnSubscriptionEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/proto/v1/store_manager.proto",
}
