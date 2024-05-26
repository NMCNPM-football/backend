// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             (unknown)
// source: summary_svc.proto

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
	SummaryService_CreateSummary_FullMethodName = "/proto.SummaryService/CreateSummary"
	SummaryService_CreateSeason_FullMethodName  = "/proto.SummaryService/CreateSeason"
)

// SummaryServiceClient is the client API for SummaryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SummaryServiceClient interface {
	CreateSummary(ctx context.Context, in *CreateSummaryRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error)
	CreateSeason(ctx context.Context, in *CreateSeasonRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error)
}

type summaryServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSummaryServiceClient(cc grpc.ClientConnInterface) SummaryServiceClient {
	return &summaryServiceClient{cc}
}

func (c *summaryServiceClient) CreateSummary(ctx context.Context, in *CreateSummaryRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error) {
	out := new(SuccessMessageResponse)
	err := c.cc.Invoke(ctx, SummaryService_CreateSummary_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *summaryServiceClient) CreateSeason(ctx context.Context, in *CreateSeasonRequest, opts ...grpc.CallOption) (*SuccessMessageResponse, error) {
	out := new(SuccessMessageResponse)
	err := c.cc.Invoke(ctx, SummaryService_CreateSeason_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SummaryServiceServer is the server API for SummaryService service.
// All implementations should embed UnimplementedSummaryServiceServer
// for forward compatibility
type SummaryServiceServer interface {
	CreateSummary(context.Context, *CreateSummaryRequest) (*SuccessMessageResponse, error)
	CreateSeason(context.Context, *CreateSeasonRequest) (*SuccessMessageResponse, error)
}

// UnimplementedSummaryServiceServer should be embedded to have forward compatible implementations.
type UnimplementedSummaryServiceServer struct {
}

func (UnimplementedSummaryServiceServer) CreateSummary(context.Context, *CreateSummaryRequest) (*SuccessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSummary not implemented")
}
func (UnimplementedSummaryServiceServer) CreateSeason(context.Context, *CreateSeasonRequest) (*SuccessMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSeason not implemented")
}

// UnsafeSummaryServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SummaryServiceServer will
// result in compilation errors.
type UnsafeSummaryServiceServer interface {
	mustEmbedUnimplementedSummaryServiceServer()
}

func RegisterSummaryServiceServer(s grpc.ServiceRegistrar, srv SummaryServiceServer) {
	s.RegisterService(&SummaryService_ServiceDesc, srv)
}

func _SummaryService_CreateSummary_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSummaryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryServiceServer).CreateSummary(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SummaryService_CreateSummary_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryServiceServer).CreateSummary(ctx, req.(*CreateSummaryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SummaryService_CreateSeason_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSeasonRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SummaryServiceServer).CreateSeason(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SummaryService_CreateSeason_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SummaryServiceServer).CreateSeason(ctx, req.(*CreateSeasonRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SummaryService_ServiceDesc is the grpc.ServiceDesc for SummaryService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SummaryService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.SummaryService",
	HandlerType: (*SummaryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSummary",
			Handler:    _SummaryService_CreateSummary_Handler,
		},
		{
			MethodName: "CreateSeason",
			Handler:    _SummaryService_CreateSeason_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "summary_svc.proto",
}
