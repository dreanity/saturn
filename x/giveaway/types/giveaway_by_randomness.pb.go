// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giveaway/giveaway_by_randomness.proto

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

type GiveawayByRandomness struct {
	Round   uint64   `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Indexes []uint32 `protobuf:"varint,2,rep,packed,name=indexes,proto3" json:"indexes,omitempty"`
}

func (m *GiveawayByRandomness) Reset()         { *m = GiveawayByRandomness{} }
func (m *GiveawayByRandomness) String() string { return proto.CompactTextString(m) }
func (*GiveawayByRandomness) ProtoMessage()    {}
func (*GiveawayByRandomness) Descriptor() ([]byte, []int) {
	return fileDescriptor_0c31911c8004ec46, []int{0}
}
func (m *GiveawayByRandomness) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GiveawayByRandomness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GiveawayByRandomness.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GiveawayByRandomness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GiveawayByRandomness.Merge(m, src)
}
func (m *GiveawayByRandomness) XXX_Size() int {
	return m.Size()
}
func (m *GiveawayByRandomness) XXX_DiscardUnknown() {
	xxx_messageInfo_GiveawayByRandomness.DiscardUnknown(m)
}

var xxx_messageInfo_GiveawayByRandomness proto.InternalMessageInfo

func (m *GiveawayByRandomness) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *GiveawayByRandomness) GetIndexes() []uint32 {
	if m != nil {
		return m.Indexes
	}
	return nil
}

func init() {
	proto.RegisterType((*GiveawayByRandomness)(nil), "dreanity.saturn.giveaway.GiveawayByRandomness")
}

func init() {
	proto.RegisterFile("giveaway/giveaway_by_randomness.proto", fileDescriptor_0c31911c8004ec46)
}

var fileDescriptor_0c31911c8004ec46 = []byte{
	// 188 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcf, 0x2c, 0x4b,
	0x4d, 0x2c, 0x4f, 0xac, 0xd4, 0x87, 0x31, 0xe2, 0x93, 0x2a, 0xe3, 0x8b, 0x12, 0xf3, 0x52, 0xf2,
	0x73, 0xf3, 0x52, 0x8b, 0x8b, 0xf5, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0x24, 0x52, 0x8a, 0x52,
	0x13, 0xf3, 0x32, 0x4b, 0x2a, 0xf5, 0x8a, 0x13, 0x4b, 0x4a, 0x8b, 0xf2, 0xf4, 0x60, 0xaa, 0x95,
	0xdc, 0xb8, 0x44, 0xdc, 0xa1, 0x6c, 0xa7, 0xca, 0x20, 0xb8, 0x3e, 0x21, 0x11, 0x2e, 0xd6, 0xa2,
	0xfc, 0xd2, 0xbc, 0x14, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x08, 0x47, 0x48, 0x82, 0x8b,
	0x3d, 0x33, 0x2f, 0x25, 0xb5, 0x22, 0xb5, 0x58, 0x82, 0x49, 0x81, 0x59, 0x83, 0x37, 0x08, 0xc6,
	0x75, 0x72, 0x3d, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xed, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0x7d, 0x98, 0x33, 0xf4, 0x21, 0xce, 0xd0, 0xaf,
	0x80, 0x3b, 0x5b, 0xbf, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x5e, 0x63, 0x40, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x18, 0x02, 0xbf, 0x0d, 0xd8, 0x00, 0x00, 0x00,
}

func (m *GiveawayByRandomness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GiveawayByRandomness) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GiveawayByRandomness) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Indexes) > 0 {
		dAtA2 := make([]byte, len(m.Indexes)*10)
		var j1 int
		for _, num := range m.Indexes {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		i -= j1
		copy(dAtA[i:], dAtA2[:j1])
		i = encodeVarintGiveawayByRandomness(dAtA, i, uint64(j1))
		i--
		dAtA[i] = 0x12
	}
	if m.Round != 0 {
		i = encodeVarintGiveawayByRandomness(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintGiveawayByRandomness(dAtA []byte, offset int, v uint64) int {
	offset -= sovGiveawayByRandomness(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GiveawayByRandomness) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovGiveawayByRandomness(uint64(m.Round))
	}
	if len(m.Indexes) > 0 {
		l = 0
		for _, e := range m.Indexes {
			l += sovGiveawayByRandomness(uint64(e))
		}
		n += 1 + sovGiveawayByRandomness(uint64(l)) + l
	}
	return n
}

func sovGiveawayByRandomness(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGiveawayByRandomness(x uint64) (n int) {
	return sovGiveawayByRandomness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GiveawayByRandomness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGiveawayByRandomness
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
			return fmt.Errorf("proto: GiveawayByRandomness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GiveawayByRandomness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGiveawayByRandomness
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
		case 2:
			if wireType == 0 {
				var v uint32
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGiveawayByRandomness
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Indexes = append(m.Indexes, v)
			} else if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGiveawayByRandomness
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthGiveawayByRandomness
				}
				postIndex := iNdEx + packedLen
				if postIndex < 0 {
					return ErrInvalidLengthGiveawayByRandomness
				}
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				var elementCount int
				var count int
				for _, integer := range dAtA[iNdEx:postIndex] {
					if integer < 128 {
						count++
					}
				}
				elementCount = count
				if elementCount != 0 && len(m.Indexes) == 0 {
					m.Indexes = make([]uint32, 0, elementCount)
				}
				for iNdEx < postIndex {
					var v uint32
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGiveawayByRandomness
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= uint32(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Indexes = append(m.Indexes, v)
				}
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Indexes", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGiveawayByRandomness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGiveawayByRandomness
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
func skipGiveawayByRandomness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGiveawayByRandomness
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
					return 0, ErrIntOverflowGiveawayByRandomness
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
					return 0, ErrIntOverflowGiveawayByRandomness
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
				return 0, ErrInvalidLengthGiveawayByRandomness
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGiveawayByRandomness
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGiveawayByRandomness
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGiveawayByRandomness        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGiveawayByRandomness          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGiveawayByRandomness = fmt.Errorf("proto: unexpected end of group")
)