// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: user_svc_public.proto

package gen

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
	UserServicePublic_Register_FullMethodName          = "/proto.UserServicePublic/Register"
	UserServicePublic_Login_FullMethodName             = "/proto.UserServicePublic/Login"
	UserServicePublic_DeactivateProfile_FullMethodName = "/proto.UserServicePublic/DeactivateProfile"
)

// UserServicePublicClient is the client API for UserServicePublic service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServicePublicClient interface {
	// Register a new user
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
	// Login a user
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	// Delete a user's profile
	DeactivateProfile(ctx context.Context, in *DeactivateProfileRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error)
}

type userServicePublicClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServicePublicClient(cc grpc.ClientConnInterface) UserServicePublicClient {
	return &userServicePublicClient{cc}
}

func (c *userServicePublicClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	out := new(RegisterResponse)
	err := c.cc.Invoke(ctx, UserServicePublic_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicePublicClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, UserServicePublic_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServicePublicClient) DeactivateProfile(ctx context.Context, in *DeactivateProfileRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error) {
	out := new(SuccessMessageResponse)
	err := c.cc.Invoke(ctx, UserServicePublic_DeactivateProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServicePublicServer is the server API for UserServicePublic service.
// All implementations should embed UnimplementedUserServicePublicServer
// for forward compatibility
type UserServicePublicServer interface {
	// Register a new user
	Register(context.Context, *RegisterRequest) (*RegisterResponse, error)
	// Login a user
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	// Delete a user's profile
	DeactivateProfile(context.Context, *DeactivateProfileRequest) (*SuccessMessageResponse, error)
}

// UnimplementedUserServicePublicServer should be embedded to have forward compatible implementations.
type UnimplementedUserServicePublicServer struct {
}

func (UnimplementedUserServicePublicServer) Register(context.Context, *RegisterRequest) (*RegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServicePublicServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServicePublicServer) DeactivateProfile(context.Context, *DeactivateProfileRequest) (*SuccessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeactivateProfile not implemented")
}

// UnsafeUserServicePublicServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServicePublicServer will
// result in compilation errors.
type UnsafeUserServicePublicServer interface {
	mustEmbedUnimplementedUserServicePublicServer()
}

func RegisterUserServicePublicServer(s grpc.ServiceRegistrar, srv UserServicePublicServer) {
	s.RegisterService(&UserServicePublic_ServiceDesc, srv)
}

func _UserServicePublic_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicePublicServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServicePublic_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicePublicServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServicePublic_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicePublicServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServicePublic_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicePublicServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserServicePublic_DeactivateProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeactivateProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServicePublicServer).DeactivateProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: UserServicePublic_DeactivateProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServicePublicServer).DeactivateProfile(ctx, req.(*DeactivateProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserServicePublic_ServiceDesc is the grpc.ServiceDesc for UserServicePublic service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserServicePublic_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.UserServicePublic",
	HandlerType: (*UserServicePublicServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _UserServicePublic_Register_Handler,
		},
		{
			MethodName: "Login",
			Handler:    _UserServicePublic_Login_Handler,
		},
		{
			MethodName: "DeactivateProfile",
			Handler:    _UserServicePublic_DeactivateProfile_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user_svc_public.proto",
}
