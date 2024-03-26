// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: club_svc.proto

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
	ClubService_GetClubProfile_FullMethodName = "/proto.ClubService/GetClubProfile"
	ClubService_UpdateClub_FullMethodName     = "/proto.ClubService/UpdateClub"
)

// ClubServiceClient is the client API for ClubService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ClubServiceClient interface {
	// Get a club's information
	GetClubProfile(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ClubProfileResponse, error)
	UpdateClub(ctx context.Context, in *ClubProfileRequest, opts ...grpc.CallOption) (*ClubProfileResponse, error)
}

type clubServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewClubServiceClient(cc grpc.ClientConnInterface) ClubServiceClient {
	return &clubServiceClient{cc}
}

func (c *clubServiceClient) GetClubProfile(ctx context.Context, in *EmptyRequest, opts ...grpc.CallOption) (*ClubProfileResponse, error) {
	out := new(ClubProfileResponse)
	err := c.cc.Invoke(ctx, ClubService_GetClubProfile_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *clubServiceClient) UpdateClub(ctx context.Context, in *ClubProfileRequest, opts ...grpc.CallOption) (*ClubProfileResponse, error) {
	out := new(ClubProfileResponse)
	err := c.cc.Invoke(ctx, ClubService_UpdateClub_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ClubServiceServer is the server API for ClubService service.
// All implementations should embed UnimplementedClubServiceServer
// for forward compatibility
type ClubServiceServer interface {
	// Get a club's information
	GetClubProfile(context.Context, *EmptyRequest) (*ClubProfileResponse, error)
	UpdateClub(context.Context, *ClubProfileRequest) (*ClubProfileResponse, error)
}

// UnimplementedClubServiceServer should be embedded to have forward compatible implementations.
type UnimplementedClubServiceServer struct {
}

func (UnimplementedClubServiceServer) GetClubProfile(context.Context, *EmptyRequest) (*ClubProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetClubProfile not implemented")
}
func (UnimplementedClubServiceServer) UpdateClub(context.Context, *ClubProfileRequest) (*ClubProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateClub not implemented")
}

// UnsafeClubServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ClubServiceServer will
// result in compilation errors.
type UnsafeClubServiceServer interface {
	mustEmbedUnimplementedClubServiceServer()
}

func RegisterClubServiceServer(s grpc.ServiceRegistrar, srv ClubServiceServer) {
	s.RegisterService(&ClubService_ServiceDesc, srv)
}

func _ClubService_GetClubProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EmptyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).GetClubProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClubService_GetClubProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).GetClubProfile(ctx, req.(*EmptyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ClubService_UpdateClub_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClubProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ClubServiceServer).UpdateClub(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ClubService_UpdateClub_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ClubServiceServer).UpdateClub(ctx, req.(*ClubProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ClubService_ServiceDesc is the grpc.ServiceDesc for ClubService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ClubService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ClubService",
	HandlerType: (*ClubServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetClubProfile",
			Handler:    _ClubService_GetClubProfile_Handler,
		},
		{
			MethodName: "UpdateClub",
			Handler:    _ClubService_UpdateClub_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "club_svc.proto",
}
