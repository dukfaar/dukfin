// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package dukfin

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

// DukFinClient is the client API for DukFin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DukFinClient interface {
	CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error)
	GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*Account, error)
	DepositToAccount(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*TransactionsReply, error)
	WithDrawFromAccount(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionsReply, error)
	TransferMoney(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionsReply, error)
}

type dukFinClient struct {
	cc grpc.ClientConnInterface
}

func NewDukFinClient(cc grpc.ClientConnInterface) DukFinClient {
	return &dukFinClient{cc}
}

func (c *dukFinClient) CreateAccount(ctx context.Context, in *CreateAccountRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/DukFin/CreateAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dukFinClient) GetAccounts(ctx context.Context, in *GetAccountsRequest, opts ...grpc.CallOption) (*Account, error) {
	out := new(Account)
	err := c.cc.Invoke(ctx, "/DukFin/GetAccounts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dukFinClient) DepositToAccount(ctx context.Context, in *DepositRequest, opts ...grpc.CallOption) (*TransactionsReply, error) {
	out := new(TransactionsReply)
	err := c.cc.Invoke(ctx, "/DukFin/DepositToAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dukFinClient) WithDrawFromAccount(ctx context.Context, in *WithdrawRequest, opts ...grpc.CallOption) (*TransactionsReply, error) {
	out := new(TransactionsReply)
	err := c.cc.Invoke(ctx, "/DukFin/WithDrawFromAccount", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dukFinClient) TransferMoney(ctx context.Context, in *TransferRequest, opts ...grpc.CallOption) (*TransactionsReply, error) {
	out := new(TransactionsReply)
	err := c.cc.Invoke(ctx, "/DukFin/TransferMoney", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DukFinServer is the server API for DukFin service.
// All implementations must embed UnimplementedDukFinServer
// for forward compatibility
type DukFinServer interface {
	CreateAccount(context.Context, *CreateAccountRequest) (*Account, error)
	GetAccounts(context.Context, *GetAccountsRequest) (*Account, error)
	DepositToAccount(context.Context, *DepositRequest) (*TransactionsReply, error)
	WithDrawFromAccount(context.Context, *WithdrawRequest) (*TransactionsReply, error)
	TransferMoney(context.Context, *TransferRequest) (*TransactionsReply, error)
	mustEmbedUnimplementedDukFinServer()
}

// UnimplementedDukFinServer must be embedded to have forward compatible implementations.
type UnimplementedDukFinServer struct {
}

func (UnimplementedDukFinServer) CreateAccount(context.Context, *CreateAccountRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateAccount not implemented")
}
func (UnimplementedDukFinServer) GetAccounts(context.Context, *GetAccountsRequest) (*Account, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccounts not implemented")
}
func (UnimplementedDukFinServer) DepositToAccount(context.Context, *DepositRequest) (*TransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DepositToAccount not implemented")
}
func (UnimplementedDukFinServer) WithDrawFromAccount(context.Context, *WithdrawRequest) (*TransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method WithDrawFromAccount not implemented")
}
func (UnimplementedDukFinServer) TransferMoney(context.Context, *TransferRequest) (*TransactionsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferMoney not implemented")
}
func (UnimplementedDukFinServer) mustEmbedUnimplementedDukFinServer() {}

// UnsafeDukFinServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DukFinServer will
// result in compilation errors.
type UnsafeDukFinServer interface {
	mustEmbedUnimplementedDukFinServer()
}

func RegisterDukFinServer(s grpc.ServiceRegistrar, srv DukFinServer) {
	s.RegisterService(&DukFin_ServiceDesc, srv)
}

func _DukFin_CreateAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateAccountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DukFinServer).CreateAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DukFin/CreateAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DukFinServer).CreateAccount(ctx, req.(*CreateAccountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DukFin_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccountsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DukFinServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DukFin/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DukFinServer).GetAccounts(ctx, req.(*GetAccountsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DukFin_DepositToAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DepositRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DukFinServer).DepositToAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DukFin/DepositToAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DukFinServer).DepositToAccount(ctx, req.(*DepositRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DukFin_WithDrawFromAccount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WithdrawRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DukFinServer).WithDrawFromAccount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DukFin/WithDrawFromAccount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DukFinServer).WithDrawFromAccount(ctx, req.(*WithdrawRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DukFin_TransferMoney_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TransferRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DukFinServer).TransferMoney(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/DukFin/TransferMoney",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DukFinServer).TransferMoney(ctx, req.(*TransferRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// DukFin_ServiceDesc is the grpc.ServiceDesc for DukFin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DukFin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "DukFin",
	HandlerType: (*DukFinServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateAccount",
			Handler:    _DukFin_CreateAccount_Handler,
		},
		{
			MethodName: "GetAccounts",
			Handler:    _DukFin_GetAccounts_Handler,
		},
		{
			MethodName: "DepositToAccount",
			Handler:    _DukFin_DepositToAccount_Handler,
		},
		{
			MethodName: "WithDrawFromAccount",
			Handler:    _DukFin_WithDrawFromAccount_Handler,
		},
		{
			MethodName: "TransferMoney",
			Handler:    _DukFin_TransferMoney_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "dukfin.proto",
}
