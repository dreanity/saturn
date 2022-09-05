// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: treasury/gas_price.proto

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

type GasPrice struct {
	Chain        string `protobuf:"bytes,1,opt,name=chain,proto3" json:"chain,omitempty"`
	TokenAddress string `protobuf:"bytes,2,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	TokenSymbol  string `protobuf:"bytes,3,opt,name=tokenSymbol,proto3" json:"tokenSymbol,omitempty"`
	// the amount of uhydrogen per token
	Value string `protobuf:"bytes,4,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GasPrice) Reset()         { *m = GasPrice{} }
func (m *GasPrice) String() string { return proto.CompactTextString(m) }
func (*GasPrice) ProtoMessage()    {}
func (*GasPrice) Descriptor() ([]byte, []int) {
	return fileDescriptor_1d3094db5d93f7c6, []int{0}
}
func (m *GasPrice) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPrice) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPrice.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPrice) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPrice.Merge(m, src)
}
func (m *GasPrice) XXX_Size() int {
	return m.Size()
}
func (m *GasPrice) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPrice.DiscardUnknown(m)
}

var xxx_messageInfo_GasPrice proto.InternalMessageInfo

func (m *GasPrice) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GasPrice) GetTokenAddress() string {
	if m != nil {
		return m.TokenAddress
	}
	return ""
}

func (m *GasPrice) GetTokenSymbol() string {
	if m != nil {
		return m.TokenSymbol
	}
	return ""
}

func (m *GasPrice) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func init() {
	proto.RegisterType((*GasPrice)(nil), "dreanity.saturn.treasury.GasPrice")
}

func init() { proto.RegisterFile("treasury/gas_price.proto", fileDescriptor_1d3094db5d93f7c6) }

var fileDescriptor_1d3094db5d93f7c6 = []byte{
	// 212 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x28, 0x29, 0x4a, 0x4d,
	0x2c, 0x2e, 0x2d, 0xaa, 0xd4, 0x4f, 0x4f, 0x2c, 0x8e, 0x2f, 0x28, 0xca, 0x4c, 0x4e, 0xd5, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48, 0x29, 0x4a, 0x4d, 0xcc, 0xcb, 0x2c, 0xa9, 0xd4, 0x2b,
	0x4e, 0x2c, 0x29, 0x2d, 0xca, 0xd3, 0x83, 0xa9, 0x54, 0xaa, 0xe1, 0xe2, 0x70, 0x4f, 0x2c, 0x0e,
	0x00, 0xa9, 0x15, 0x12, 0xe1, 0x62, 0x4d, 0xce, 0x48, 0xcc, 0xcc, 0x93, 0x60, 0x54, 0x60, 0xd4,
	0xe0, 0x0c, 0x82, 0x70, 0x84, 0x94, 0xb8, 0x78, 0x4a, 0xf2, 0xb3, 0x53, 0xf3, 0x1c, 0x53, 0x52,
	0x8a, 0x52, 0x8b, 0x8b, 0x25, 0x98, 0xc0, 0x92, 0x28, 0x62, 0x42, 0x0a, 0x5c, 0xdc, 0x60, 0x7e,
	0x70, 0x65, 0x6e, 0x52, 0x7e, 0x8e, 0x04, 0x33, 0x58, 0x09, 0xb2, 0x10, 0xc8, 0xec, 0xb2, 0xc4,
	0x9c, 0xd2, 0x54, 0x09, 0x16, 0x88, 0xd9, 0x60, 0x8e, 0x93, 0xeb, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0x69, 0xa7, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0xc3, 0x1c, 0xaf, 0x0f, 0x71, 0xbc, 0x7e, 0x85, 0x3e, 0xdc, 0xa3, 0x25, 0x95, 0x05, 0xa9,
	0xc5, 0x49, 0x6c, 0x60, 0x5f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb4, 0x76, 0xc9, 0x57,
	0x01, 0x01, 0x00, 0x00,
}

func (m *GasPrice) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPrice) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPrice) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Value) > 0 {
		i -= len(m.Value)
		copy(dAtA[i:], m.Value)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Value)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.TokenSymbol) > 0 {
		i -= len(m.TokenSymbol)
		copy(dAtA[i:], m.TokenSymbol)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.TokenSymbol)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.TokenAddress) > 0 {
		i -= len(m.TokenAddress)
		copy(dAtA[i:], m.TokenAddress)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.TokenAddress)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintGasPrice(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGasPrice(dAtA []byte, offset int, v uint64) int {
	offset -= sovGasPrice(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GasPrice) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	l = len(m.TokenAddress)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	l = len(m.TokenSymbol)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	l = len(m.Value)
	if l > 0 {
		n += 1 + l + sovGasPrice(uint64(l))
	}
	return n
}

func sovGasPrice(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGasPrice(x uint64) (n int) {
	return sovGasPrice(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GasPrice) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGasPrice
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
			return fmt.Errorf("proto: GasPrice: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPrice: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenSymbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenSymbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Value", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGasPrice
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
				return ErrInvalidLengthGasPrice
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGasPrice
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Value = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGasPrice(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGasPrice
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
func skipGasPrice(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGasPrice
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
					return 0, ErrIntOverflowGasPrice
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
					return 0, ErrIntOverflowGasPrice
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
				return 0, ErrInvalidLengthGasPrice
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGasPrice
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGasPrice
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGasPrice        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGasPrice          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGasPrice = fmt.Errorf("proto: unexpected end of group")
)
