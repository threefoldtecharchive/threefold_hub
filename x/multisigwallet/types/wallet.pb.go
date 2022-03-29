// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: multisigwallet/wallet.proto

package types

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
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

type Wallet struct {
	Index    string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Name     string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Address  string `protobuf:"bytes,3,opt,name=address,proto3" json:"address,omitempty"`
	Members  string `protobuf:"bytes,4,opt,name=members,proto3" json:"members,omitempty"`
	MinSigns string `protobuf:"bytes,5,opt,name=minSigns,proto3" json:"minSigns,omitempty"`
}

func (m *Wallet) Reset()         { *m = Wallet{} }
func (m *Wallet) String() string { return proto.CompactTextString(m) }
func (*Wallet) ProtoMessage()    {}
func (*Wallet) Descriptor() ([]byte, []int) {
	return fileDescriptor_f10865d32f0fdd0b, []int{0}
}
func (m *Wallet) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Wallet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Wallet.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Wallet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Wallet.Merge(m, src)
}
func (m *Wallet) XXX_Size() int {
	return m.Size()
}
func (m *Wallet) XXX_DiscardUnknown() {
	xxx_messageInfo_Wallet.DiscardUnknown(m)
}

var xxx_messageInfo_Wallet proto.InternalMessageInfo

func (m *Wallet) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *Wallet) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Wallet) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Wallet) GetMembers() string {
	if m != nil {
		return m.Members
	}
	return ""
}

func (m *Wallet) GetMinSigns() string {
	if m != nil {
		return m.MinSigns
	}
	return ""
}

func init() {
	proto.RegisterType((*Wallet)(nil), "threefoldtech.threefoldhub.multisigwallet.Wallet")
}

func init() { proto.RegisterFile("multisigwallet/wallet.proto", fileDescriptor_f10865d32f0fdd0b) }

var fileDescriptor_f10865d32f0fdd0b = []byte{
	// 223 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x2d, 0xcd, 0x29,
	0xc9, 0x2c, 0xce, 0x4c, 0x2f, 0x4f, 0xcc, 0xc9, 0x49, 0x2d, 0xd1, 0x87, 0x50, 0x7a, 0x05, 0x45,
	0xf9, 0x25, 0xf9, 0x42, 0x9a, 0x25, 0x19, 0x45, 0xa9, 0xa9, 0x69, 0xf9, 0x39, 0x29, 0x25, 0xa9,
	0xc9, 0x19, 0x7a, 0x70, 0x5e, 0x46, 0x69, 0x92, 0x1e, 0xaa, 0x3e, 0xa5, 0x26, 0x46, 0x2e, 0xb6,
	0x70, 0x30, 0x53, 0x48, 0x84, 0x8b, 0x35, 0x33, 0x2f, 0x25, 0xb5, 0x42, 0x82, 0x51, 0x81, 0x51,
	0x83, 0x33, 0x08, 0xc2, 0x11, 0x12, 0xe2, 0x62, 0xc9, 0x4b, 0xcc, 0x4d, 0x95, 0x60, 0x02, 0x0b,
	0x82, 0xd9, 0x42, 0x12, 0x5c, 0xec, 0x89, 0x29, 0x29, 0x45, 0xa9, 0xc5, 0xc5, 0x12, 0xcc, 0x60,
	0x61, 0x18, 0x17, 0x24, 0x93, 0x9b, 0x9a, 0x9b, 0x94, 0x5a, 0x54, 0x2c, 0xc1, 0x02, 0x91, 0x81,
	0x72, 0x85, 0xa4, 0xb8, 0x38, 0x72, 0x33, 0xf3, 0x82, 0x33, 0xd3, 0xf3, 0x8a, 0x25, 0x58, 0xc1,
	0x52, 0x70, 0xbe, 0x53, 0xf8, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24,
	0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xd9,
	0xa6, 0x67, 0x96, 0x80, 0x5c, 0x9d, 0x9c, 0x9f, 0xab, 0x8f, 0xe2, 0x29, 0x04, 0x2f, 0x3e, 0xa3,
	0x34, 0x49, 0xbf, 0x42, 0x1f, 0x2d, 0x3c, 0x4a, 0x2a, 0x0b, 0x52, 0x8b, 0x93, 0xd8, 0xc0, 0xe1,
	0x61, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x88, 0xd5, 0xa7, 0x7f, 0x2e, 0x01, 0x00, 0x00,
}

func (m *Wallet) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Wallet) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Wallet) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinSigns) > 0 {
		i -= len(m.MinSigns)
		copy(dAtA[i:], m.MinSigns)
		i = encodeVarintWallet(dAtA, i, uint64(len(m.MinSigns)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Members) > 0 {
		i -= len(m.Members)
		copy(dAtA[i:], m.Members)
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Members)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintWallet(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWallet(dAtA []byte, offset int, v uint64) int {
	offset -= sovWallet(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Wallet) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.Members)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	l = len(m.MinSigns)
	if l > 0 {
		n += 1 + l + sovWallet(uint64(l))
	}
	return n
}

func sovWallet(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWallet(x uint64) (n int) {
	return sovWallet(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Wallet) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWallet
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
			return fmt.Errorf("proto: Wallet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Wallet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinSigns", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWallet
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
				return ErrInvalidLengthWallet
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthWallet
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinSigns = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWallet(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthWallet
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
func skipWallet(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
					return 0, ErrIntOverflowWallet
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
				return 0, ErrInvalidLengthWallet
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupWallet
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthWallet
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthWallet        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWallet          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupWallet = fmt.Errorf("proto: unexpected end of group")
)
