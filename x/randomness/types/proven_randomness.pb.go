// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: randomness/proven_randomness.proto

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

type ProvenRandomness struct {
	Round             uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Randomness        string `protobuf:"bytes,3,opt,name=randomness,proto3" json:"randomness,omitempty"`
	Signature         string `protobuf:"bytes,4,opt,name=signature,proto3" json:"signature,omitempty"`
	PreviousSignature string `protobuf:"bytes,5,opt,name=previousSignature,proto3" json:"previousSignature,omitempty"`
}

func (m *ProvenRandomness) Reset()         { *m = ProvenRandomness{} }
func (m *ProvenRandomness) String() string { return proto.CompactTextString(m) }
func (*ProvenRandomness) ProtoMessage()    {}
func (*ProvenRandomness) Descriptor() ([]byte, []int) {
	return fileDescriptor_61df09207b669f89, []int{0}
}
func (m *ProvenRandomness) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvenRandomness) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvenRandomness.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvenRandomness) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvenRandomness.Merge(m, src)
}
func (m *ProvenRandomness) XXX_Size() int {
	return m.Size()
}
func (m *ProvenRandomness) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvenRandomness.DiscardUnknown(m)
}

var xxx_messageInfo_ProvenRandomness proto.InternalMessageInfo

func (m *ProvenRandomness) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ProvenRandomness) GetRandomness() string {
	if m != nil {
		return m.Randomness
	}
	return ""
}

func (m *ProvenRandomness) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ProvenRandomness) GetPreviousSignature() string {
	if m != nil {
		return m.PreviousSignature
	}
	return ""
}

func init() {
	proto.RegisterType((*ProvenRandomness)(nil), "saturn.randomness.ProvenRandomness")
}

func init() {
	proto.RegisterFile("randomness/proven_randomness.proto", fileDescriptor_61df09207b669f89)
}

var fileDescriptor_61df09207b669f89 = []byte{
	// 191 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2a, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xcd, 0x4b, 0x2d, 0x2e, 0xd6, 0x2f, 0x28, 0xca, 0x2f, 0x4b, 0xcd, 0x8b, 0x47, 0x88,
	0xe8, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b, 0x09, 0x16, 0x27, 0x96, 0x94, 0x16, 0xe5, 0xe9, 0x21,
	0x24, 0x94, 0xa6, 0x30, 0x72, 0x09, 0x04, 0x80, 0x95, 0x07, 0xc1, 0x05, 0x85, 0x44, 0xb8, 0x58,
	0x8b, 0xf2, 0x4b, 0xf3, 0x52, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x20, 0x1c, 0x21, 0x39,
	0x2e, 0x2e, 0x84, 0x46, 0x09, 0x66, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x24, 0x11, 0x21, 0x19, 0x2e,
	0xce, 0xe2, 0xcc, 0xf4, 0x3c, 0x90, 0x15, 0xa9, 0x12, 0x2c, 0x60, 0x69, 0x84, 0x80, 0x90, 0x0e,
	0x97, 0x60, 0x41, 0x51, 0x6a, 0x59, 0x66, 0x7e, 0x69, 0x71, 0x30, 0x5c, 0x15, 0x2b, 0x58, 0x15,
	0xa6, 0x84, 0x93, 0xf1, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7,
	0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x49, 0x42,
	0xfc, 0xa0, 0x5f, 0xa1, 0x8f, 0xe4, 0xe1, 0x92, 0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x2f,
	0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xc8, 0x32, 0x5b, 0xbf, 0x0b, 0x01, 0x00, 0x00,
}

func (m *ProvenRandomness) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvenRandomness) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProvenRandomness) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PreviousSignature) > 0 {
		i -= len(m.PreviousSignature)
		copy(dAtA[i:], m.PreviousSignature)
		i = encodeVarintProvenRandomness(dAtA, i, uint64(len(m.PreviousSignature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintProvenRandomness(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Randomness) > 0 {
		i -= len(m.Randomness)
		copy(dAtA[i:], m.Randomness)
		i = encodeVarintProvenRandomness(dAtA, i, uint64(len(m.Randomness)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Round != 0 {
		i = encodeVarintProvenRandomness(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProvenRandomness(dAtA []byte, offset int, v uint64) int {
	offset -= sovProvenRandomness(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProvenRandomness) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovProvenRandomness(uint64(m.Round))
	}
	l = len(m.Randomness)
	if l > 0 {
		n += 1 + l + sovProvenRandomness(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovProvenRandomness(uint64(l))
	}
	l = len(m.PreviousSignature)
	if l > 0 {
		n += 1 + l + sovProvenRandomness(uint64(l))
	}
	return n
}

func sovProvenRandomness(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProvenRandomness(x uint64) (n int) {
	return sovProvenRandomness(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProvenRandomness) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProvenRandomness
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
			return fmt.Errorf("proto: ProvenRandomness: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvenRandomness: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvenRandomness
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Randomness", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvenRandomness
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
				return ErrInvalidLengthProvenRandomness
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvenRandomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Randomness = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvenRandomness
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
				return ErrInvalidLengthProvenRandomness
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvenRandomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProvenRandomness
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
				return ErrInvalidLengthProvenRandomness
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProvenRandomness
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProvenRandomness(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProvenRandomness
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
func skipProvenRandomness(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProvenRandomness
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
					return 0, ErrIntOverflowProvenRandomness
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
					return 0, ErrIntOverflowProvenRandomness
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
				return 0, ErrInvalidLengthProvenRandomness
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProvenRandomness
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProvenRandomness
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProvenRandomness        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProvenRandomness          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProvenRandomness = fmt.Errorf("proto: unexpected end of group")
)
