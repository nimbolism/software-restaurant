// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v5.26.1
// source: card_service.proto

package proto

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

const (
	CardService_GetCardInfo_FullMethodName    = "/CardService/GetCardInfo"
	CardService_UpdateReserves_FullMethodName = "/CardService/UpdateReserves"
)

// CardServiceClient is the client API for CardService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardServiceClient interface {
	// GetCardInfo returns information about a card, including access level and blacklisted status.
	GetCardInfo(ctx context.Context, in *GetCardInfoRequest, opts ...grpc.CallOption) (*CardInfoResponse, error)
	// UpdateReserves updates the reserves count of a card.
	UpdateReserves(ctx context.Context, in *UpdateReservesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type cardServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCardServiceClient(cc grpc.ClientConnInterface) CardServiceClient {
	return &cardServiceClient{cc}
}

func (c *cardServiceClient) GetCardInfo(ctx context.Context, in *GetCardInfoRequest, opts ...grpc.CallOption) (*CardInfoResponse, error) {
	out := new(CardInfoResponse)
	err := c.cc.Invoke(ctx, CardService_GetCardInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cardServiceClient) UpdateReserves(ctx context.Context, in *UpdateReservesRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, CardService_UpdateReserves_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardServiceServer is the server API for CardService service.
// All implementations must embed UnimplementedCardServiceServer
// for forward compatibility
type CardServiceServer interface {
	// GetCardInfo returns information about a card, including access level and blacklisted status.
	GetCardInfo(context.Context, *GetCardInfoRequest) (*CardInfoResponse, error)
	// UpdateReserves updates the reserves count of a card.
	UpdateReserves(context.Context, *UpdateReservesRequest) (*emptypb.Empty, error)
	mustEmbedUnimplementedCardServiceServer()
}

// UnimplementedCardServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCardServiceServer struct {
}

func (UnimplementedCardServiceServer) GetCardInfo(context.Context, *GetCardInfoRequest) (*CardInfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCardInfo not implemented")
}
func (UnimplementedCardServiceServer) UpdateReserves(context.Context, *UpdateReservesRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateReserves not implemented")
}
func (UnimplementedCardServiceServer) mustEmbedUnimplementedCardServiceServer() {}

// UnsafeCardServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardServiceServer will
// result in compilation errors.
type UnsafeCardServiceServer interface {
	mustEmbedUnimplementedCardServiceServer()
}

func RegisterCardServiceServer(s grpc.ServiceRegistrar, srv CardServiceServer) {
	s.RegisterService(&CardService_ServiceDesc, srv)
}

func _CardService_GetCardInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCardInfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).GetCardInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardService_GetCardInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).GetCardInfo(ctx, req.(*GetCardInfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CardService_UpdateReserves_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateReservesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardServiceServer).UpdateReserves(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardService_UpdateReserves_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardServiceServer).UpdateReserves(ctx, req.(*UpdateReservesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardService_ServiceDesc is the grpc.ServiceDesc for CardService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "CardService",
	HandlerType: (*CardServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCardInfo",
			Handler:    _CardService_GetCardInfo_Handler,
		},
		{
			MethodName: "UpdateReserves",
			Handler:    _CardService_UpdateReserves_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "card_service.proto",
}
