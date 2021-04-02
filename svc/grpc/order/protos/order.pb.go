// Code generated by protoc-gen-go. DO NOT EDIT.
// source: order.proto

package order

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProductRequest struct {
	ProductID            int64    `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Demand               int64    `protobuf:"varint,2,opt,name=demand,proto3" json:"demand,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductRequest) Reset()         { *m = ProductRequest{} }
func (m *ProductRequest) String() string { return proto.CompactTextString(m) }
func (*ProductRequest) ProtoMessage()    {}
func (*ProductRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_9340e08afa767c13, []int{0}
}
func (m *ProductRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductRequest.Unmarshal(m, b)
}
func (m *ProductRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductRequest.Marshal(b, m, deterministic)
}
func (dst *ProductRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductRequest.Merge(dst, src)
}
func (m *ProductRequest) XXX_Size() int {
	return xxx_messageInfo_ProductRequest.Size(m)
}
func (m *ProductRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProductRequest proto.InternalMessageInfo

func (m *ProductRequest) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *ProductRequest) GetDemand() int64 {
	if m != nil {
		return m.Demand
	}
	return 0
}

type ProductResponse struct {
	ProductID            int64    `protobuf:"varint,1,opt,name=productID,proto3" json:"productID,omitempty"`
	Demand               int64    `protobuf:"varint,2,opt,name=demand,proto3" json:"demand,omitempty"`
	Status               int64    `protobuf:"varint,3,opt,name=status,proto3" json:"status,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProductResponse) Reset()         { *m = ProductResponse{} }
func (m *ProductResponse) String() string { return proto.CompactTextString(m) }
func (*ProductResponse) ProtoMessage()    {}
func (*ProductResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_9340e08afa767c13, []int{1}
}
func (m *ProductResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProductResponse.Unmarshal(m, b)
}
func (m *ProductResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProductResponse.Marshal(b, m, deterministic)
}
func (dst *ProductResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProductResponse.Merge(dst, src)
}
func (m *ProductResponse) XXX_Size() int {
	return xxx_messageInfo_ProductResponse.Size(m)
}
func (m *ProductResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProductResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProductResponse proto.InternalMessageInfo

func (m *ProductResponse) GetProductID() int64 {
	if m != nil {
		return m.ProductID
	}
	return 0
}

func (m *ProductResponse) GetDemand() int64 {
	if m != nil {
		return m.Demand
	}
	return 0
}

func (m *ProductResponse) GetStatus() int64 {
	if m != nil {
		return m.Status
	}
	return 0
}

type OrderRequest struct {
	Products             []*ProductRequest `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	CartID               int64             `protobuf:"varint,2,opt,name=cartID,proto3" json:"cartID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_9340e08afa767c13, []int{2}
}
func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (dst *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(dst, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetProducts() []*ProductRequest {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *OrderRequest) GetCartID() int64 {
	if m != nil {
		return m.CartID
	}
	return 0
}

type OrderResponse struct {
	Products             []*ProductResponse `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	IsSuccess            bool               `protobuf:"varint,2,opt,name=isSuccess,proto3" json:"isSuccess,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_order_9340e08afa767c13, []int{3}
}
func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (dst *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(dst, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func (m *OrderResponse) GetProducts() []*ProductResponse {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *OrderResponse) GetIsSuccess() bool {
	if m != nil {
		return m.IsSuccess
	}
	return false
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "order.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "order.ProductResponse")
	proto.RegisterType((*OrderRequest)(nil), "order.OrderRequest")
	proto.RegisterType((*OrderResponse)(nil), "order.OrderResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	Checkout(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderClient struct {
	cc *grpc.ClientConn
}

func NewOrderClient(cc *grpc.ClientConn) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) Checkout(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/order.Order/Checkout", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	Checkout(context.Context, *OrderRequest) (*OrderResponse, error)
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_Checkout_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).Checkout(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/order.Order/Checkout",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).Checkout(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "order.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Checkout",
			Handler:    _Order_Checkout_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}

func init() { proto.RegisterFile("order.proto", fileDescriptor_order_9340e08afa767c13) }

var fileDescriptor_order_9340e08afa767c13 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcf, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x5d, 0x97, 0xd6, 0xed, 0xd4, 0x1f, 0x10, 0xb5, 0x04, 0xe9, 0xa1, 0xe4, 0xd4, 0x53,
	0xc1, 0xf5, 0xe0, 0x55, 0x54, 0x04, 0x4f, 0x4a, 0x3c, 0xe9, 0x45, 0x62, 0x12, 0x50, 0xc4, 0x66,
	0xcd, 0x24, 0xff, 0xbf, 0x24, 0x13, 0xdb, 0xae, 0xde, 0x7a, 0x7c, 0x5f, 0xc2, 0x37, 0x6f, 0x18,
	0x18, 0x3b, 0x6f, 0xac, 0x5f, 0x74, 0xde, 0x05, 0xc7, 0x06, 0x39, 0x88, 0x3b, 0x38, 0x7c, 0xf4,
	0xce, 0x44, 0x1d, 0xa4, 0xfd, 0x8e, 0x16, 0x03, 0x9b, 0xc2, 0xa8, 0x23, 0x72, 0x7f, 0xcb, 0xab,
	0x59, 0x35, 0xaf, 0xe5, 0x1a, 0xb0, 0x09, 0x0c, 0x8d, 0xfd, 0x52, 0x4b, 0xc3, 0x77, 0xf3, 0x53,
	0x49, 0xe2, 0x15, 0x8e, 0x56, 0x1e, 0xec, 0xdc, 0x12, 0xed, 0x76, 0xa2, 0xc4, 0x31, 0xa8, 0x10,
	0x91, 0xd7, 0xc4, 0x29, 0x89, 0x67, 0xd8, 0x7f, 0x48, 0x8d, 0x7f, 0x6b, 0x9e, 0x43, 0x53, 0x64,
	0xc8, 0xab, 0x59, 0x3d, 0x1f, 0xb7, 0xa7, 0x0b, 0xda, 0xaf, 0xbf, 0x8f, 0x5c, 0x7d, 0x4b, 0x6a,
	0xad, 0x7c, 0x6a, 0x53, 0x46, 0x52, 0x12, 0x0a, 0x0e, 0x8a, 0xba, 0x34, 0x6f, 0xff, 0xb9, 0x27,
	0x7f, 0xdd, 0xf4, 0x73, 0x43, 0x3e, 0x85, 0xd1, 0x07, 0x3e, 0x45, 0xad, 0x2d, 0x62, 0xf6, 0x37,
	0x72, 0x0d, 0xda, 0x2b, 0x18, 0xe4, 0x11, 0xec, 0x12, 0x9a, 0x9b, 0x77, 0xab, 0x3f, 0x5d, 0x0c,
	0xec, 0xb8, 0x48, 0x37, 0xf7, 0x3a, 0x3b, 0xe9, 0x43, 0x9a, 0x23, 0x76, 0xae, 0xf7, 0x5e, 0xe8,
	0x62, 0x6f, 0xc3, 0x7c, 0xbf, 0x8b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x77, 0x3e, 0xaa, 0xd0,
	0xce, 0x01, 0x00, 0x00,
}
