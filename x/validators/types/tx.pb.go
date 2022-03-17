// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: validators/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgCreateValidator struct {
	OperatorAddress string      `protobuf:"bytes,1,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	ConsensusPubkey string      `protobuf:"bytes,2,opt,name=consensusPubkey,proto3" json:"consensusPubkey,omitempty"`
	Description     Description `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
}

func (m *MsgCreateValidator) Reset()         { *m = MsgCreateValidator{} }
func (m *MsgCreateValidator) String() string { return proto.CompactTextString(m) }
func (*MsgCreateValidator) ProtoMessage()    {}
func (*MsgCreateValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{0}
}
func (m *MsgCreateValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateValidator.Merge(m, src)
}
func (m *MsgCreateValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateValidator proto.InternalMessageInfo

func (m *MsgCreateValidator) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *MsgCreateValidator) GetConsensusPubkey() string {
	if m != nil {
		return m.ConsensusPubkey
	}
	return ""
}

func (m *MsgCreateValidator) GetDescription() Description {
	if m != nil {
		return m.Description
	}
	return Description{}
}

type MsgCreateValidatorResponse struct {
}

func (m *MsgCreateValidatorResponse) Reset()         { *m = MsgCreateValidatorResponse{} }
func (m *MsgCreateValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateValidatorResponse) ProtoMessage()    {}
func (*MsgCreateValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{1}
}
func (m *MsgCreateValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateValidatorResponse.Merge(m, src)
}
func (m *MsgCreateValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateValidatorResponse proto.InternalMessageInfo

type MsgUpdateValidator struct {
	OperatorAddress string      `protobuf:"bytes,1,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
	Description     Description `protobuf:"bytes,3,opt,name=description,proto3" json:"description"`
}

func (m *MsgUpdateValidator) Reset()         { *m = MsgUpdateValidator{} }
func (m *MsgUpdateValidator) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateValidator) ProtoMessage()    {}
func (*MsgUpdateValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{2}
}
func (m *MsgUpdateValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateValidator.Merge(m, src)
}
func (m *MsgUpdateValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateValidator proto.InternalMessageInfo

func (m *MsgUpdateValidator) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

func (m *MsgUpdateValidator) GetDescription() Description {
	if m != nil {
		return m.Description
	}
	return Description{}
}

type MsgUpdateValidatorResponse struct {
}

func (m *MsgUpdateValidatorResponse) Reset()         { *m = MsgUpdateValidatorResponse{} }
func (m *MsgUpdateValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpdateValidatorResponse) ProtoMessage()    {}
func (*MsgUpdateValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{3}
}
func (m *MsgUpdateValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpdateValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpdateValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpdateValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpdateValidatorResponse.Merge(m, src)
}
func (m *MsgUpdateValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpdateValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpdateValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpdateValidatorResponse proto.InternalMessageInfo

type MsgDeleteValidator struct {
	OperatorAddress string `protobuf:"bytes,1,opt,name=operatorAddress,proto3" json:"operatorAddress,omitempty"`
}

func (m *MsgDeleteValidator) Reset()         { *m = MsgDeleteValidator{} }
func (m *MsgDeleteValidator) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteValidator) ProtoMessage()    {}
func (*MsgDeleteValidator) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{4}
}
func (m *MsgDeleteValidator) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteValidator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteValidator.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteValidator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteValidator.Merge(m, src)
}
func (m *MsgDeleteValidator) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteValidator) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteValidator.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteValidator proto.InternalMessageInfo

func (m *MsgDeleteValidator) GetOperatorAddress() string {
	if m != nil {
		return m.OperatorAddress
	}
	return ""
}

type MsgDeleteValidatorResponse struct {
}

func (m *MsgDeleteValidatorResponse) Reset()         { *m = MsgDeleteValidatorResponse{} }
func (m *MsgDeleteValidatorResponse) String() string { return proto.CompactTextString(m) }
func (*MsgDeleteValidatorResponse) ProtoMessage()    {}
func (*MsgDeleteValidatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_69b7604ee268048e, []int{5}
}
func (m *MsgDeleteValidatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgDeleteValidatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgDeleteValidatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgDeleteValidatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgDeleteValidatorResponse.Merge(m, src)
}
func (m *MsgDeleteValidatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgDeleteValidatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgDeleteValidatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgDeleteValidatorResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateValidator)(nil), "threefold.threefoldhub.validators.MsgCreateValidator")
	proto.RegisterType((*MsgCreateValidatorResponse)(nil), "threefold.threefoldhub.validators.MsgCreateValidatorResponse")
	proto.RegisterType((*MsgUpdateValidator)(nil), "threefold.threefoldhub.validators.MsgUpdateValidator")
	proto.RegisterType((*MsgUpdateValidatorResponse)(nil), "threefold.threefoldhub.validators.MsgUpdateValidatorResponse")
	proto.RegisterType((*MsgDeleteValidator)(nil), "threefold.threefoldhub.validators.MsgDeleteValidator")
	proto.RegisterType((*MsgDeleteValidatorResponse)(nil), "threefold.threefoldhub.validators.MsgDeleteValidatorResponse")
}

func init() { proto.RegisterFile("validators/tx.proto", fileDescriptor_69b7604ee268048e) }

var fileDescriptor_69b7604ee268048e = []byte{
	// 389 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2e, 0x4b, 0xcc, 0xc9,
	0x4c, 0x49, 0x2c, 0xc9, 0x2f, 0x2a, 0xd6, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17,
	0x52, 0x2c, 0xc9, 0x28, 0x4a, 0x4d, 0x4d, 0xcb, 0xcf, 0x49, 0xd1, 0x83, 0xb3, 0x32, 0x4a, 0x93,
	0xf4, 0x10, 0x6a, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52, 0xf5, 0xc1, 0x1a, 0x92, 0x4a,
	0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0xba, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x4c, 0x7d,
	0x10, 0x0b, 0x2a, 0x2a, 0x99, 0x9c, 0x5f, 0x9c, 0x9b, 0x5f, 0x1c, 0x0f, 0x91, 0x80, 0x70, 0xa0,
	0x52, 0x52, 0x48, 0x6e, 0x80, 0x33, 0x21, 0x72, 0x4a, 0x07, 0x18, 0xb9, 0x84, 0x7c, 0x8b, 0xd3,
	0x9d, 0x8b, 0x52, 0x13, 0x4b, 0x52, 0xc3, 0x60, 0x92, 0x42, 0x1a, 0x5c, 0xfc, 0xf9, 0x05, 0xa9,
	0x45, 0x20, 0xb6, 0x63, 0x4a, 0x4a, 0x51, 0x6a, 0x71, 0xb1, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67,
	0x10, 0xba, 0x30, 0x48, 0x65, 0x72, 0x7e, 0x5e, 0x71, 0x6a, 0x5e, 0x71, 0x69, 0x71, 0x40, 0x69,
	0x52, 0x76, 0x6a, 0xa5, 0x04, 0x13, 0x44, 0x25, 0x9a, 0xb0, 0x50, 0x18, 0x17, 0x77, 0x4a, 0x6a,
	0x71, 0x72, 0x51, 0x66, 0x41, 0x49, 0x66, 0x7e, 0x9e, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0xb7, 0x91,
	0x9e, 0x1e, 0xc1, 0xb0, 0xd0, 0x73, 0x41, 0xe8, 0x72, 0x62, 0x39, 0x71, 0x4f, 0x9e, 0x21, 0x08,
	0xd9, 0x20, 0x25, 0x19, 0x2e, 0x29, 0x4c, 0x1f, 0x04, 0xa5, 0x16, 0x17, 0x80, 0xec, 0x57, 0x9a,
	0x06, 0xf1, 0x60, 0x68, 0x41, 0x0a, 0x99, 0x1e, 0xa4, 0xad, 0xb3, 0xd1, 0xdc, 0x05, 0x77, 0xb6,
	0x1d, 0xd8, 0xd5, 0x2e, 0xa9, 0x39, 0xa9, 0x64, 0xb9, 0x1a, 0x6a, 0x3a, 0x9a, 0x7e, 0x98, 0xe9,
	0x46, 0x8b, 0x99, 0xb9, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xda, 0x19, 0xb9, 0xf8, 0xd1, 0xa3, 0xde,
	0x94, 0x08, 0xaf, 0x61, 0x86, 0xb7, 0x94, 0x2d, 0x59, 0xda, 0x60, 0x2e, 0x02, 0xbb, 0x04, 0x3d,
	0x8e, 0x88, 0x74, 0x09, 0x9a, 0x36, 0x62, 0x5d, 0x82, 0x23, 0xe4, 0xc1, 0x2e, 0x41, 0x0f, 0x77,
	0x22, 0x5d, 0x82, 0xa6, 0x8d, 0x58, 0x97, 0xe0, 0x88, 0x25, 0x27, 0xff, 0x13, 0x8f, 0xe4, 0x18,
	0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5,
	0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0x32, 0x4d, 0xcf, 0x2c, 0x01, 0x99, 0x93, 0x9c, 0x9f, 0xab,
	0x0f, 0x37, 0x18, 0xc1, 0x8a, 0xcf, 0x28, 0x4d, 0xd2, 0xaf, 0xd0, 0x47, 0x2e, 0x7a, 0x2a, 0x0b,
	0x52, 0x8b, 0x93, 0xd8, 0xc0, 0x79, 0xde, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xd5, 0x5d,
	0x71, 0x95, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error)
	UpdateValidator(ctx context.Context, in *MsgUpdateValidator, opts ...grpc.CallOption) (*MsgUpdateValidatorResponse, error)
	DeleteValidator(ctx context.Context, in *MsgDeleteValidator, opts ...grpc.CallOption) (*MsgDeleteValidatorResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateValidator(ctx context.Context, in *MsgCreateValidator, opts ...grpc.CallOption) (*MsgCreateValidatorResponse, error) {
	out := new(MsgCreateValidatorResponse)
	err := c.cc.Invoke(ctx, "/threefold.threefoldhub.validators.Msg/CreateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpdateValidator(ctx context.Context, in *MsgUpdateValidator, opts ...grpc.CallOption) (*MsgUpdateValidatorResponse, error) {
	out := new(MsgUpdateValidatorResponse)
	err := c.cc.Invoke(ctx, "/threefold.threefoldhub.validators.Msg/UpdateValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) DeleteValidator(ctx context.Context, in *MsgDeleteValidator, opts ...grpc.CallOption) (*MsgDeleteValidatorResponse, error) {
	out := new(MsgDeleteValidatorResponse)
	err := c.cc.Invoke(ctx, "/threefold.threefoldhub.validators.Msg/DeleteValidator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateValidator(context.Context, *MsgCreateValidator) (*MsgCreateValidatorResponse, error)
	UpdateValidator(context.Context, *MsgUpdateValidator) (*MsgUpdateValidatorResponse, error)
	DeleteValidator(context.Context, *MsgDeleteValidator) (*MsgDeleteValidatorResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateValidator(ctx context.Context, req *MsgCreateValidator) (*MsgCreateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateValidator not implemented")
}
func (*UnimplementedMsgServer) UpdateValidator(ctx context.Context, req *MsgUpdateValidator) (*MsgUpdateValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateValidator not implemented")
}
func (*UnimplementedMsgServer) DeleteValidator(ctx context.Context, req *MsgDeleteValidator) (*MsgDeleteValidatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteValidator not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/threefold.threefoldhub.validators.Msg/CreateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateValidator(ctx, req.(*MsgCreateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpdateValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpdateValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpdateValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/threefold.threefoldhub.validators.Msg/UpdateValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpdateValidator(ctx, req.(*MsgUpdateValidator))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_DeleteValidator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgDeleteValidator)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).DeleteValidator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/threefold.threefoldhub.validators.Msg/DeleteValidator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).DeleteValidator(ctx, req.(*MsgDeleteValidator))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "threefold.threefoldhub.validators.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateValidator",
			Handler:    _Msg_CreateValidator_Handler,
		},
		{
			MethodName: "UpdateValidator",
			Handler:    _Msg_UpdateValidator_Handler,
		},
		{
			MethodName: "DeleteValidator",
			Handler:    _Msg_DeleteValidator_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "validators/tx.proto",
}

func (m *MsgCreateValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.ConsensusPubkey) > 0 {
		i -= len(m.ConsensusPubkey)
		copy(dAtA[i:], m.ConsensusPubkey)
		i = encodeVarintTx(dAtA, i, uint64(len(m.ConsensusPubkey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpdateValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Description.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgUpdateValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpdateValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpdateValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgDeleteValidator) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteValidator) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteValidator) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OperatorAddress) > 0 {
		i -= len(m.OperatorAddress)
		copy(dAtA[i:], m.OperatorAddress)
		i = encodeVarintTx(dAtA, i, uint64(len(m.OperatorAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgDeleteValidatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgDeleteValidatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgDeleteValidatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCreateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.ConsensusPubkey)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCreateValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpdateValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Description.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgUpdateValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgDeleteValidator) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.OperatorAddress)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgDeleteValidatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCreateValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ConsensusPubkey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ConsensusPubkey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgCreateValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgCreateValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Description.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgUpdateValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgUpdateValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpdateValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteValidator) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteValidator: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteValidator: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OperatorAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *MsgDeleteValidatorResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgDeleteValidatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgDeleteValidatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
