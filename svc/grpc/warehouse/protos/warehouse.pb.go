// Code generated by protoc-gen-go. DO NOT EDIT.
// source: warehouse.proto

package warehouse

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
	return fileDescriptor_warehouse_cdf0fe9ada95e2f0, []int{0}
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
	return fileDescriptor_warehouse_cdf0fe9ada95e2f0, []int{1}
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

type ClaimRequest struct {
	Products             []*ProductRequest `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	CartID               int64             `protobuf:"varint,2,opt,name=cartID,proto3" json:"cartID,omitempty"`
	OrderID              int64             `protobuf:"varint,3,opt,name=orderID,proto3" json:"orderID,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ClaimRequest) Reset()         { *m = ClaimRequest{} }
func (m *ClaimRequest) String() string { return proto.CompactTextString(m) }
func (*ClaimRequest) ProtoMessage()    {}
func (*ClaimRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_warehouse_cdf0fe9ada95e2f0, []int{2}
}
func (m *ClaimRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimRequest.Unmarshal(m, b)
}
func (m *ClaimRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimRequest.Marshal(b, m, deterministic)
}
func (dst *ClaimRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimRequest.Merge(dst, src)
}
func (m *ClaimRequest) XXX_Size() int {
	return xxx_messageInfo_ClaimRequest.Size(m)
}
func (m *ClaimRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimRequest proto.InternalMessageInfo

func (m *ClaimRequest) GetProducts() []*ProductRequest {
	if m != nil {
		return m.Products
	}
	return nil
}

func (m *ClaimRequest) GetCartID() int64 {
	if m != nil {
		return m.CartID
	}
	return 0
}

func (m *ClaimRequest) GetOrderID() int64 {
	if m != nil {
		return m.OrderID
	}
	return 0
}

type ClaimResponse struct {
	Products             []*ProductResponse `protobuf:"bytes,1,rep,name=products,proto3" json:"products,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *ClaimResponse) Reset()         { *m = ClaimResponse{} }
func (m *ClaimResponse) String() string { return proto.CompactTextString(m) }
func (*ClaimResponse) ProtoMessage()    {}
func (*ClaimResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_warehouse_cdf0fe9ada95e2f0, []int{3}
}
func (m *ClaimResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ClaimResponse.Unmarshal(m, b)
}
func (m *ClaimResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ClaimResponse.Marshal(b, m, deterministic)
}
func (dst *ClaimResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimResponse.Merge(dst, src)
}
func (m *ClaimResponse) XXX_Size() int {
	return xxx_messageInfo_ClaimResponse.Size(m)
}
func (m *ClaimResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimResponse proto.InternalMessageInfo

func (m *ClaimResponse) GetProducts() []*ProductResponse {
	if m != nil {
		return m.Products
	}
	return nil
}

func init() {
	proto.RegisterType((*ProductRequest)(nil), "warehouse.ProductRequest")
	proto.RegisterType((*ProductResponse)(nil), "warehouse.ProductResponse")
	proto.RegisterType((*ClaimRequest)(nil), "warehouse.ClaimRequest")
	proto.RegisterType((*ClaimResponse)(nil), "warehouse.ClaimResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WarehouseClient is the client API for Warehouse service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WarehouseClient interface {
	ValidateStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	ClaimProduct(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error)
}

type warehouseClient struct {
	cc *grpc.ClientConn
}

func NewWarehouseClient(cc *grpc.ClientConn) WarehouseClient {
	return &warehouseClient{cc}
}

func (c *warehouseClient) ValidateStock(ctx context.Context, in *ProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	out := new(ProductResponse)
	err := c.cc.Invoke(ctx, "/warehouse.Warehouse/ValidateStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *warehouseClient) ClaimProduct(ctx context.Context, in *ClaimRequest, opts ...grpc.CallOption) (*ClaimResponse, error) {
	out := new(ClaimResponse)
	err := c.cc.Invoke(ctx, "/warehouse.Warehouse/ClaimProduct", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WarehouseServer is the server API for Warehouse service.
type WarehouseServer interface {
	ValidateStock(context.Context, *ProductRequest) (*ProductResponse, error)
	ClaimProduct(context.Context, *ClaimRequest) (*ClaimResponse, error)
}

func RegisterWarehouseServer(s *grpc.Server, srv WarehouseServer) {
	s.RegisterService(&_Warehouse_serviceDesc, srv)
}

func _Warehouse_ValidateStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProductRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).ValidateStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warehouse.Warehouse/ValidateStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).ValidateStock(ctx, req.(*ProductRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Warehouse_ClaimProduct_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ClaimRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WarehouseServer).ClaimProduct(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/warehouse.Warehouse/ClaimProduct",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WarehouseServer).ClaimProduct(ctx, req.(*ClaimRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Warehouse_serviceDesc = grpc.ServiceDesc{
	ServiceName: "warehouse.Warehouse",
	HandlerType: (*WarehouseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateStock",
			Handler:    _Warehouse_ValidateStock_Handler,
		},
		{
			MethodName: "ClaimProduct",
			Handler:    _Warehouse_ClaimProduct_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "warehouse.proto",
}

func init() { proto.RegisterFile("warehouse.proto", fileDescriptor_warehouse_cdf0fe9ada95e2f0) }

var fileDescriptor_warehouse_cdf0fe9ada95e2f0 = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2f, 0x4f, 0x2c, 0x4a,
	0xcd, 0xc8, 0x2f, 0x2d, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x84, 0x0b, 0x28,
	0xb9, 0x71, 0xf1, 0x05, 0x14, 0xe5, 0xa7, 0x94, 0x26, 0x97, 0x04, 0xa5, 0x16, 0x96, 0xa6, 0x16,
	0x97, 0x08, 0xc9, 0x70, 0x71, 0x16, 0x40, 0x44, 0x3c, 0x5d, 0x24, 0x18, 0x15, 0x18, 0x35, 0x98,
	0x83, 0x10, 0x02, 0x42, 0x62, 0x5c, 0x6c, 0x29, 0xa9, 0xb9, 0x89, 0x79, 0x29, 0x12, 0x4c, 0x60,
	0x29, 0x28, 0x4f, 0x29, 0x9e, 0x8b, 0x1f, 0x6e, 0x4e, 0x71, 0x41, 0x7e, 0x5e, 0x71, 0x2a, 0x79,
	0x06, 0x81, 0xc4, 0x8b, 0x4b, 0x12, 0x4b, 0x4a, 0x8b, 0x25, 0x98, 0x21, 0xe2, 0x10, 0x9e, 0x52,
	0x39, 0x17, 0x8f, 0x73, 0x4e, 0x62, 0x66, 0x2e, 0xcc, 0x99, 0xa6, 0x5c, 0x1c, 0x50, 0xc3, 0x8a,
	0x25, 0x18, 0x15, 0x98, 0x35, 0xb8, 0x8d, 0x24, 0xf5, 0x10, 0xfe, 0x44, 0xf5, 0x53, 0x10, 0x5c,
	0x29, 0xc8, 0xf8, 0xe4, 0xc4, 0x22, 0x90, 0x8b, 0xa0, 0xd6, 0x42, 0x78, 0x42, 0x12, 0x5c, 0xec,
	0xf9, 0x45, 0x29, 0xa9, 0x45, 0x9e, 0x2e, 0x50, 0x7b, 0x61, 0x5c, 0x25, 0x77, 0x2e, 0x5e, 0xa8,
	0xc5, 0x50, 0x7f, 0x99, 0x61, 0xd8, 0x2c, 0x85, 0xcd, 0x66, 0x88, 0x6a, 0x84, 0xd5, 0x46, 0xb3,
	0x18, 0xb9, 0x38, 0xc3, 0x61, 0xea, 0x84, 0x3c, 0xb8, 0x78, 0xc3, 0x12, 0x73, 0x32, 0x53, 0x12,
	0x4b, 0x52, 0x83, 0x4b, 0xf2, 0x93, 0xb3, 0x85, 0x70, 0x3b, 0x5f, 0x0a, 0x8f, 0xf9, 0x4a, 0x0c,
	0x42, 0xce, 0xd0, 0x90, 0x81, 0xca, 0x08, 0x89, 0x23, 0xa9, 0x46, 0x0e, 0x32, 0x29, 0x09, 0x4c,
	0x09, 0x98, 0x21, 0x4e, 0xdc, 0x51, 0x88, 0x44, 0x91, 0xc4, 0x06, 0x4e, 0x26, 0xc6, 0x80, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x21, 0xfb, 0x03, 0x55, 0x39, 0x02, 0x00, 0x00,
}
