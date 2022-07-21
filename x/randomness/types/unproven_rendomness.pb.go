// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: randomness/unproven_rendomness.proto

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

type UnprovenRendomness struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
	Round uint64 `protobuf:"varint,2,opt,name=round,proto3" json:"round,omitempty"`
	Time  uint64 `protobuf:"varint,3,opt,name=time,proto3" json:"time,omitempty"`
}

func (m *UnprovenRendomness) Reset()         { *m = UnprovenRendomness{} }
func (m *UnprovenRendomness) String() string { return proto.CompactTextString(m) }
func (*UnprovenRendomness) ProtoMessage()    {}
func (*UnprovenRendomness) Descriptor() ([]byte, []int) {
	return fileDescriptor_435119f23e7377b1, []int{0}
}
func (m *UnprovenRendomness) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnprovenRendomness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnprovenRendomness.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnprovenRendomness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnprovenRendomness.Merge(m, src)
}
func (m *UnprovenRendomness) XXX_Size() int {
	return m.Size()
}
func (m *UnprovenRendomness) XXX_DiscardUnknown() {
	xxx_messageInfo_UnprovenRendomness.DiscardUnknown(m)
}

var xxx_messageInfo_UnprovenRendomness proto.InternalMessageInfo

func (m *UnprovenRendomness) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

func (m *UnprovenRendomness) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *UnprovenRendomness) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func init() {
	proto.RegisterType((*UnprovenRendomness)(nil), "saturn.randomness.UnprovenRendomness")
}

func init() {
	proto.RegisterFile("randomness/unproven_rendomness.proto", fileDescriptor_435119f23e7377b1)
}

var fileDescriptor_435119f23e7377b1 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x29, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xcd, 0x4b, 0x2d, 0x2e, 0xd6, 0x2f, 0xcd, 0x2b, 0x28, 0xca, 0x2f, 0x4b, 0xcd, 0x8b,
	0x2f, 0x4a, 0x85, 0x89, 0xe9, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x16, 0x27, 0x96, 0x94,
	0x16, 0xe5, 0xe9, 0x21, 0x14, 0x2b, 0x85, 0x70, 0x09, 0x85, 0x42, 0xd5, 0x07, 0xc1, 0x95, 0x0b,
	0x89, 0x70, 0xb1, 0x66, 0xe6, 0xa5, 0xa4, 0x56, 0x48, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x41,
	0x38, 0x20, 0xd1, 0xa2, 0xfc, 0xd2, 0xbc, 0x14, 0x09, 0x26, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08,
	0x47, 0x48, 0x88, 0x8b, 0xa5, 0x24, 0x33, 0x37, 0x55, 0x82, 0x19, 0x2c, 0x08, 0x66, 0x3b, 0x19,
	0x9f, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb,
	0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94, 0x24, 0xc4, 0x09, 0xfa, 0x15,
	0xfa, 0x48, 0x2e, 0x2e, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0x3b, 0xd2, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0xdc, 0x39, 0x9a, 0x9a, 0xcc, 0x00, 0x00, 0x00,
}

func (m *UnprovenRendomness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnprovenRendomness) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnprovenRendomness) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Time != 0 {
		i = encodeVarintUnprovenRendomness(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x18
	}
	if m.Round != 0 {
		i = encodeVarintUnprovenRendomness(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintUnprovenRendomness(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintUnprovenRendomness(dAtA []byte, offset int, v uint64) int {
	offset -= sovUnprovenRendomness(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *UnprovenRendomness) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovUnprovenRendomness(uint64(l))
	}
	if m.Round != 0 {
		n += 1 + sovUnprovenRendomness(uint64(m.Round))
	}
	if m.Time != 0 {
		n += 1 + sovUnprovenRendomness(uint64(m.Time))
	}
	return n
}

func sovUnprovenRendomness(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozUnprovenRendomness(x uint64) (n int) {
	return sovUnprovenRendomness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *UnprovenRendomness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowUnprovenRendomness
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
			return fmt.Errorf("proto: UnprovenRendomness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnprovenRendomness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnprovenRendomness
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
				return ErrInvalidLengthUnprovenRendomness
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthUnprovenRendomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnprovenRendomness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Round |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowUnprovenRendomness
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipUnprovenRendomness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthUnprovenRendomness
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
func skipUnprovenRendomness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowUnprovenRendomness
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
					return 0, ErrIntOverflowUnprovenRendomness
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
					return 0, ErrIntOverflowUnprovenRendomness
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
				return 0, ErrInvalidLengthUnprovenRendomness
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupUnprovenRendomness
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthUnprovenRendomness
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthUnprovenRendomness        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowUnprovenRendomness          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupUnprovenRendomness = fmt.Errorf("proto: unexpected end of group")
)