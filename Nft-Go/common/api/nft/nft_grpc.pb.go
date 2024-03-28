// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: nft.proto

package nft

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
	Nft_GetMessageByAddressOrHash_FullMethodName = "/nft.Nft/GetMessageByAddressOrHash"
	Nft_CreateActivity_FullMethodName            = "/nft.Nft/CreateActivity"
	Nft_PrizeDcFromActivity_FullMethodName       = "/nft.Nft/PrizeDcFromActivity"
	Nft_GetActivityPages_FullMethodName          = "/nft.Nft/GetActivityPages"
	Nft_GetDcPages_FullMethodName                = "/nft.Nft/GetDcPages"
	Nft_GetPoolPages_FullMethodName              = "/nft.Nft/GetPoolPages"
	Nft_GetOneActivity_FullMethodName            = "/nft.Nft/GetOneActivity"
	Nft_SearchActivities_FullMethodName          = "/nft.Nft/SearchActivities"
	Nft_GiveDc_FullMethodName                    = "/nft.Nft/GiveDc"
	Nft_SelectDc_FullMethodName                  = "/nft.Nft/SelectDc"
	Nft_GetDcById_FullMethodName                 = "/nft.Nft/GetDcById"
	Nft_GetMyDc_FullMethodName                   = "/nft.Nft/GetMyDc"
	Nft_GetDcHistory_FullMethodName              = "/nft.Nft/GetDcHistory"
	Nft_CreatePool_FullMethodName                = "/nft.Nft/CreatePool"
	Nft_BuyFromPool_FullMethodName               = "/nft.Nft/BuyFromPool"
	Nft_SelectPool_FullMethodName                = "/nft.Nft/SelectPool"
	Nft_GetPoolById_FullMethodName               = "/nft.Nft/GetPoolById"
	Nft_GetMyPool_FullMethodName                 = "/nft.Nft/GetMyPool"
)

// NftClient is the client API for Nft service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NftClient interface {
	GetMessageByAddressOrHash(ctx context.Context, in *GetMessageByAddressOrHashRequest, opts ...grpc.CallOption) (*GetMessageByAddressOrHashDTO, error)
	CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*Response, error)
	PrizeDcFromActivity(ctx context.Context, in *GetDcFromActivityRequest, opts ...grpc.CallOption) (*Response, error)
	GetActivityPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error)
	GetDcPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DcPageVOList, error)
	GetPoolPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PoolPageVOList, error)
	GetOneActivity(ctx context.Context, in *GetOneActivityRequest, opts ...grpc.CallOption) (*ActivityDetailsVO, error)
	SearchActivities(ctx context.Context, in *SearchActivitiesRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error)
	GiveDc(ctx context.Context, in *GiveDcRequest, opts ...grpc.CallOption) (*Response, error)
	SelectDc(ctx context.Context, in *SelectDcRequest, opts ...grpc.CallOption) (*DcPageVOList, error)
	GetDcById(ctx context.Context, in *GetDcByIdRequest, opts ...grpc.CallOption) (*DcDetailVO, error)
	GetMyDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error)
	GetDcHistory(ctx context.Context, in *GetDcHistoryRequest, opts ...grpc.CallOption) (*CollectionMessageOnChainVO, error)
	CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Response, error)
	BuyFromPool(ctx context.Context, in *BuyFromPoolRequest, opts ...grpc.CallOption) (*Response, error)
	SelectPool(ctx context.Context, in *SelectPoolRequest, opts ...grpc.CallOption) (*PoolPageVOList, error)
	GetPoolById(ctx context.Context, in *GetPoolByIdRequest, opts ...grpc.CallOption) (*PoolDetailsVO, error)
	GetMyPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error)
}

type nftClient struct {
	cc grpc.ClientConnInterface
}

func NewNftClient(cc grpc.ClientConnInterface) NftClient {
	return &nftClient{cc}
}

func (c *nftClient) GetMessageByAddressOrHash(ctx context.Context, in *GetMessageByAddressOrHashRequest, opts ...grpc.CallOption) (*GetMessageByAddressOrHashDTO, error) {
	out := new(GetMessageByAddressOrHashDTO)
	err := c.cc.Invoke(ctx, Nft_GetMessageByAddressOrHash_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) CreateActivity(ctx context.Context, in *CreateActivityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Nft_CreateActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) PrizeDcFromActivity(ctx context.Context, in *GetDcFromActivityRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Nft_PrizeDcFromActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetActivityPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	out := new(ActivityPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetActivityPages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDcPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetDcPages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetPoolPages(ctx context.Context, in *PageRequest, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetPoolPages_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetOneActivity(ctx context.Context, in *GetOneActivityRequest, opts ...grpc.CallOption) (*ActivityDetailsVO, error) {
	out := new(ActivityDetailsVO)
	err := c.cc.Invoke(ctx, Nft_GetOneActivity_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SearchActivities(ctx context.Context, in *SearchActivitiesRequest, opts ...grpc.CallOption) (*ActivityPageVOList, error) {
	out := new(ActivityPageVOList)
	err := c.cc.Invoke(ctx, Nft_SearchActivities_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GiveDc(ctx context.Context, in *GiveDcRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Nft_GiveDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SelectDc(ctx context.Context, in *SelectDcRequest, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_SelectDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDcById(ctx context.Context, in *GetDcByIdRequest, opts ...grpc.CallOption) (*DcDetailVO, error) {
	out := new(DcDetailVO)
	err := c.cc.Invoke(ctx, Nft_GetDcById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetMyDc(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*DcPageVOList, error) {
	out := new(DcPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetMyDc_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetDcHistory(ctx context.Context, in *GetDcHistoryRequest, opts ...grpc.CallOption) (*CollectionMessageOnChainVO, error) {
	out := new(CollectionMessageOnChainVO)
	err := c.cc.Invoke(ctx, Nft_GetDcHistory_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) CreatePool(ctx context.Context, in *CreatePoolRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Nft_CreatePool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) BuyFromPool(ctx context.Context, in *BuyFromPoolRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, Nft_BuyFromPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) SelectPool(ctx context.Context, in *SelectPoolRequest, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_SelectPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetPoolById(ctx context.Context, in *GetPoolByIdRequest, opts ...grpc.CallOption) (*PoolDetailsVO, error) {
	out := new(PoolDetailsVO)
	err := c.cc.Invoke(ctx, Nft_GetPoolById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nftClient) GetMyPool(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*PoolPageVOList, error) {
	out := new(PoolPageVOList)
	err := c.cc.Invoke(ctx, Nft_GetMyPool_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NftServer is the server API for Nft service.
// All implementations must embed UnimplementedNftServer
// for forward compatibility
type NftServer interface {
	GetMessageByAddressOrHash(context.Context, *GetMessageByAddressOrHashRequest) (*GetMessageByAddressOrHashDTO, error)
	CreateActivity(context.Context, *CreateActivityRequest) (*Response, error)
	PrizeDcFromActivity(context.Context, *GetDcFromActivityRequest) (*Response, error)
	GetActivityPages(context.Context, *PageRequest) (*ActivityPageVOList, error)
	GetDcPages(context.Context, *PageRequest) (*DcPageVOList, error)
	GetPoolPages(context.Context, *PageRequest) (*PoolPageVOList, error)
	GetOneActivity(context.Context, *GetOneActivityRequest) (*ActivityDetailsVO, error)
	SearchActivities(context.Context, *SearchActivitiesRequest) (*ActivityPageVOList, error)
	GiveDc(context.Context, *GiveDcRequest) (*Response, error)
	SelectDc(context.Context, *SelectDcRequest) (*DcPageVOList, error)
	GetDcById(context.Context, *GetDcByIdRequest) (*DcDetailVO, error)
	GetMyDc(context.Context, *Empty) (*DcPageVOList, error)
	GetDcHistory(context.Context, *GetDcHistoryRequest) (*CollectionMessageOnChainVO, error)
	CreatePool(context.Context, *CreatePoolRequest) (*Response, error)
	BuyFromPool(context.Context, *BuyFromPoolRequest) (*Response, error)
	SelectPool(context.Context, *SelectPoolRequest) (*PoolPageVOList, error)
	GetPoolById(context.Context, *GetPoolByIdRequest) (*PoolDetailsVO, error)
	GetMyPool(context.Context, *Empty) (*PoolPageVOList, error)
	mustEmbedUnimplementedNftServer()
}

// UnimplementedNftServer must be embedded to have forward compatible implementations.
type UnimplementedNftServer struct {
}

func (UnimplementedNftServer) GetMessageByAddressOrHash(context.Context, *GetMessageByAddressOrHashRequest) (*GetMessageByAddressOrHashDTO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMessageByAddressOrHash not implemented")
}
func (UnimplementedNftServer) CreateActivity(context.Context, *CreateActivityRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateActivity not implemented")
}
func (UnimplementedNftServer) PrizeDcFromActivity(context.Context, *GetDcFromActivityRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrizeDcFromActivity not implemented")
}
func (UnimplementedNftServer) GetActivityPages(context.Context, *PageRequest) (*ActivityPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetActivityPages not implemented")
}
func (UnimplementedNftServer) GetDcPages(context.Context, *PageRequest) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDcPages not implemented")
}
func (UnimplementedNftServer) GetPoolPages(context.Context, *PageRequest) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolPages not implemented")
}
func (UnimplementedNftServer) GetOneActivity(context.Context, *GetOneActivityRequest) (*ActivityDetailsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOneActivity not implemented")
}
func (UnimplementedNftServer) SearchActivities(context.Context, *SearchActivitiesRequest) (*ActivityPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchActivities not implemented")
}
func (UnimplementedNftServer) GiveDc(context.Context, *GiveDcRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GiveDc not implemented")
}
func (UnimplementedNftServer) SelectDc(context.Context, *SelectDcRequest) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectDc not implemented")
}
func (UnimplementedNftServer) GetDcById(context.Context, *GetDcByIdRequest) (*DcDetailVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDcById not implemented")
}
func (UnimplementedNftServer) GetMyDc(context.Context, *Empty) (*DcPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyDc not implemented")
}
func (UnimplementedNftServer) GetDcHistory(context.Context, *GetDcHistoryRequest) (*CollectionMessageOnChainVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDcHistory not implemented")
}
func (UnimplementedNftServer) CreatePool(context.Context, *CreatePoolRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePool not implemented")
}
func (UnimplementedNftServer) BuyFromPool(context.Context, *BuyFromPoolRequest) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BuyFromPool not implemented")
}
func (UnimplementedNftServer) SelectPool(context.Context, *SelectPoolRequest) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectPool not implemented")
}
func (UnimplementedNftServer) GetPoolById(context.Context, *GetPoolByIdRequest) (*PoolDetailsVO, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPoolById not implemented")
}
func (UnimplementedNftServer) GetMyPool(context.Context, *Empty) (*PoolPageVOList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyPool not implemented")
}
func (UnimplementedNftServer) mustEmbedUnimplementedNftServer() {}

// UnsafeNftServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NftServer will
// result in compilation errors.
type UnsafeNftServer interface {
	mustEmbedUnimplementedNftServer()
}

func RegisterNftServer(s grpc.ServiceRegistrar, srv NftServer) {
	s.RegisterService(&Nft_ServiceDesc, srv)
}

func _Nft_GetMessageByAddressOrHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMessageByAddressOrHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMessageByAddressOrHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMessageByAddressOrHash_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMessageByAddressOrHash(ctx, req.(*GetMessageByAddressOrHashRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_CreateActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).CreateActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_CreateActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).CreateActivity(ctx, req.(*CreateActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_PrizeDcFromActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDcFromActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).PrizeDcFromActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_PrizeDcFromActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).PrizeDcFromActivity(ctx, req.(*GetDcFromActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetActivityPages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetActivityPages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetActivityPages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetActivityPages(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDcPages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDcPages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDcPages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDcPages(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetPoolPages_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetPoolPages(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetPoolPages_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetPoolPages(ctx, req.(*PageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetOneActivity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetOneActivityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetOneActivity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetOneActivity_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetOneActivity(ctx, req.(*GetOneActivityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SearchActivities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchActivitiesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SearchActivities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SearchActivities_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SearchActivities(ctx, req.(*SearchActivitiesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GiveDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GiveDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GiveDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GiveDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GiveDc(ctx, req.(*GiveDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SelectDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectDcRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SelectDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SelectDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SelectDc(ctx, req.(*SelectDcRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDcById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDcByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDcById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDcById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDcById(ctx, req.(*GetDcByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetMyDc_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMyDc(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMyDc_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMyDc(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetDcHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDcHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetDcHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetDcHistory_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetDcHistory(ctx, req.(*GetDcHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_CreatePool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).CreatePool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_CreatePool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).CreatePool(ctx, req.(*CreatePoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_BuyFromPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuyFromPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).BuyFromPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_BuyFromPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).BuyFromPool(ctx, req.(*BuyFromPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_SelectPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectPoolRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).SelectPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_SelectPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).SelectPool(ctx, req.(*SelectPoolRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetPoolById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPoolByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetPoolById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetPoolById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetPoolById(ctx, req.(*GetPoolByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Nft_GetMyPool_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NftServer).GetMyPool(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Nft_GetMyPool_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NftServer).GetMyPool(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// Nft_ServiceDesc is the grpc.ServiceDesc for Nft service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Nft_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "nft.Nft",
	HandlerType: (*NftServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetMessageByAddressOrHash",
			Handler:    _Nft_GetMessageByAddressOrHash_Handler,
		},
		{
			MethodName: "CreateActivity",
			Handler:    _Nft_CreateActivity_Handler,
		},
		{
			MethodName: "PrizeDcFromActivity",
			Handler:    _Nft_PrizeDcFromActivity_Handler,
		},
		{
			MethodName: "GetActivityPages",
			Handler:    _Nft_GetActivityPages_Handler,
		},
		{
			MethodName: "GetDcPages",
			Handler:    _Nft_GetDcPages_Handler,
		},
		{
			MethodName: "GetPoolPages",
			Handler:    _Nft_GetPoolPages_Handler,
		},
		{
			MethodName: "GetOneActivity",
			Handler:    _Nft_GetOneActivity_Handler,
		},
		{
			MethodName: "SearchActivities",
			Handler:    _Nft_SearchActivities_Handler,
		},
		{
			MethodName: "GiveDc",
			Handler:    _Nft_GiveDc_Handler,
		},
		{
			MethodName: "SelectDc",
			Handler:    _Nft_SelectDc_Handler,
		},
		{
			MethodName: "GetDcById",
			Handler:    _Nft_GetDcById_Handler,
		},
		{
			MethodName: "GetMyDc",
			Handler:    _Nft_GetMyDc_Handler,
		},
		{
			MethodName: "GetDcHistory",
			Handler:    _Nft_GetDcHistory_Handler,
		},
		{
			MethodName: "CreatePool",
			Handler:    _Nft_CreatePool_Handler,
		},
		{
			MethodName: "BuyFromPool",
			Handler:    _Nft_BuyFromPool_Handler,
		},
		{
			MethodName: "SelectPool",
			Handler:    _Nft_SelectPool_Handler,
		},
		{
			MethodName: "GetPoolById",
			Handler:    _Nft_GetPoolById_Handler,
		},
		{
			MethodName: "GetMyPool",
			Handler:    _Nft_GetMyPool_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "nft.proto",
}
