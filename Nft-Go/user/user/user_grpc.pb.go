// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: user.proto

package user

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
	User_Register_FullMethodName         = "/user.User/register"
	User_Login_FullMethodName            = "/user.User/login"
	User_Logout_FullMethodName           = "/user.User/logout"
	User_RefreshTokens_FullMethodName    = "/user.User/refreshTokens"
	User_Upload_FullMethodName           = "/user.User/upload"
	User_GetAllNotices_FullMethodName    = "/user.User/getAllNotices"
	User_GetNoticeByTitle_FullMethodName = "/user.User/getNoticeByTitle"
	User_GetNoticeById_FullMethodName    = "/user.User/getNoticeById"
)

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error)
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error)
	Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	RefreshTokens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Upload(ctx context.Context, opts ...grpc.CallOption) (User_UploadClient, error)
	GetAllNotices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NoticeList, error)
	GetNoticeByTitle(ctx context.Context, in *TitleNoticeRequest, opts ...grpc.CallOption) (*NoticeList, error)
	GetNoticeById(ctx context.Context, in *IdNoticeRequest, opts ...grpc.CallOption) (*Notice, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_Register_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_Login_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Logout(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_Logout_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) RefreshTokens(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, User_RefreshTokens_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) Upload(ctx context.Context, opts ...grpc.CallOption) (User_UploadClient, error) {
	stream, err := c.cc.NewStream(ctx, &User_ServiceDesc.Streams[0], User_Upload_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &userUploadClient{stream}
	return x, nil
}

type User_UploadClient interface {
	Send(*UploadRequest) error
	CloseAndRecv() (*Response, error)
	grpc.ClientStream
}

type userUploadClient struct {
	grpc.ClientStream
}

func (x *userUploadClient) Send(m *UploadRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *userUploadClient) CloseAndRecv() (*Response, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *userClient) GetAllNotices(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*NoticeList, error) {
	out := new(NoticeList)
	err := c.cc.Invoke(ctx, User_GetAllNotices_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetNoticeByTitle(ctx context.Context, in *TitleNoticeRequest, opts ...grpc.CallOption) (*NoticeList, error) {
	out := new(NoticeList)
	err := c.cc.Invoke(ctx, User_GetNoticeByTitle_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userClient) GetNoticeById(ctx context.Context, in *IdNoticeRequest, opts ...grpc.CallOption) (*Notice, error) {
	out := new(Notice)
	err := c.cc.Invoke(ctx, User_GetNoticeById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
// All implementations must embed UnimplementedUserServer
// for forward compatibility
type UserServer interface {
	Register(context.Context, *RegisterRequest) (*Response, error)
	Login(context.Context, *LoginRequest) (*Response, error)
	Logout(context.Context, *Empty) (*Response, error)
	RefreshTokens(context.Context, *Empty) (*Response, error)
	Upload(User_UploadServer) error
	GetAllNotices(context.Context, *Empty) (*NoticeList, error)
	GetNoticeByTitle(context.Context, *TitleNoticeRequest) (*NoticeList, error)
	GetNoticeById(context.Context, *IdNoticeRequest) (*Notice, error)
	mustEmbedUnimplementedUserServer()
}

// UnimplementedUserServer must be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (UnimplementedUserServer) Register(context.Context, *RegisterRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}
func (UnimplementedUserServer) Login(context.Context, *LoginRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedUserServer) Logout(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Logout not implemented")
}
func (UnimplementedUserServer) RefreshTokens(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokens not implemented")
}
func (UnimplementedUserServer) Upload(User_UploadServer) error {
	return status.Errorf(codes.Unimplemented, "method Upload not implemented")
}
func (UnimplementedUserServer) GetAllNotices(context.Context, *Empty) (*NoticeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllNotices not implemented")
}
func (UnimplementedUserServer) GetNoticeByTitle(context.Context, *TitleNoticeRequest) (*NoticeList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoticeByTitle not implemented")
}
func (UnimplementedUserServer) GetNoticeById(context.Context, *IdNoticeRequest) (*Notice, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetNoticeById not implemented")
}
func (UnimplementedUserServer) mustEmbedUnimplementedUserServer() {}

// UnsafeUserServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServer will
// result in compilation errors.
type UnsafeUserServer interface {
	mustEmbedUnimplementedUserServer()
}

func RegisterUserServer(s grpc.ServiceRegistrar, srv UserServer) {
	s.RegisterService(&User_ServiceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Register_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*RegisterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Logout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Logout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_Logout_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Logout(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_RefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).RefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_RefreshTokens_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).RefreshTokens(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_Upload_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(UserServer).Upload(&userUploadServer{stream})
}

type User_UploadServer interface {
	SendAndClose(*Response) error
	Recv() (*UploadRequest, error)
	grpc.ServerStream
}

type userUploadServer struct {
	grpc.ServerStream
}

func (x *userUploadServer) SendAndClose(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func (x *userUploadServer) Recv() (*UploadRequest, error) {
	m := new(UploadRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _User_GetAllNotices_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetAllNotices(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetAllNotices_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetAllNotices(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetNoticeByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TitleNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetNoticeByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetNoticeByTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetNoticeByTitle(ctx, req.(*TitleNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _User_GetNoticeById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IdNoticeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).GetNoticeById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: User_GetNoticeById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).GetNoticeById(ctx, req.(*IdNoticeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// User_ServiceDesc is the grpc.ServiceDesc for User service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var User_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "user.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "register",
			Handler:    _User_Register_Handler,
		},
		{
			MethodName: "login",
			Handler:    _User_Login_Handler,
		},
		{
			MethodName: "logout",
			Handler:    _User_Logout_Handler,
		},
		{
			MethodName: "refreshTokens",
			Handler:    _User_RefreshTokens_Handler,
		},
		{
			MethodName: "getAllNotices",
			Handler:    _User_GetAllNotices_Handler,
		},
		{
			MethodName: "getNoticeByTitle",
			Handler:    _User_GetNoticeByTitle_Handler,
		},
		{
			MethodName: "getNoticeById",
			Handler:    _User_GetNoticeById_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "upload",
			Handler:       _User_Upload_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "user.proto",
}