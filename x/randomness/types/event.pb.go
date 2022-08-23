// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: randomness/event.proto

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

type ProvenRandomnessCreated struct {
	Round             uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	Randomness        string `protobuf:"bytes,2,opt,name=randomness,proto3" json:"randomness,omitempty"`
	Signature         string `protobuf:"bytes,3,opt,name=signature,proto3" json:"signature,omitempty"`
	PreviousSignature string `protobuf:"bytes,4,opt,name=previousSignature,proto3" json:"previousSignature,omitempty"`
	RoundTime         uint64 `protobuf:"varint,6,opt,name=roundTime,proto3" json:"roundTime,omitempty"`
}

func (m *ProvenRandomnessCreated) Reset()         { *m = ProvenRandomnessCreated{} }
func (m *ProvenRandomnessCreated) String() string { return proto.CompactTextString(m) }
func (*ProvenRandomnessCreated) ProtoMessage()    {}
func (*ProvenRandomnessCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d0e7b92baa19ab, []int{0}
}
func (m *ProvenRandomnessCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ProvenRandomnessCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ProvenRandomnessCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ProvenRandomnessCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProvenRandomnessCreated.Merge(m, src)
}
func (m *ProvenRandomnessCreated) XXX_Size() int {
	return m.Size()
}
func (m *ProvenRandomnessCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_ProvenRandomnessCreated.DiscardUnknown(m)
}

var xxx_messageInfo_ProvenRandomnessCreated proto.InternalMessageInfo

func (m *ProvenRandomnessCreated) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *ProvenRandomnessCreated) GetRandomness() string {
	if m != nil {
		return m.Randomness
	}
	return ""
}

func (m *ProvenRandomnessCreated) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func (m *ProvenRandomnessCreated) GetPreviousSignature() string {
	if m != nil {
		return m.PreviousSignature
	}
	return ""
}

func (m *ProvenRandomnessCreated) GetRoundTime() uint64 {
	if m != nil {
		return m.RoundTime
	}
	return 0
}

type UnprovenRandomnessCreated struct {
	Round     uint64 `protobuf:"varint,1,opt,name=round,proto3" json:"round,omitempty"`
	RoundTime uint64 `protobuf:"varint,2,opt,name=roundTime,proto3" json:"roundTime,omitempty"`
}

func (m *UnprovenRandomnessCreated) Reset()         { *m = UnprovenRandomnessCreated{} }
func (m *UnprovenRandomnessCreated) String() string { return proto.CompactTextString(m) }
func (*UnprovenRandomnessCreated) ProtoMessage()    {}
func (*UnprovenRandomnessCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_84d0e7b92baa19ab, []int{1}
}
func (m *UnprovenRandomnessCreated) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *UnprovenRandomnessCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_UnprovenRandomnessCreated.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *UnprovenRandomnessCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UnprovenRandomnessCreated.Merge(m, src)
}
func (m *UnprovenRandomnessCreated) XXX_Size() int {
	return m.Size()
}
func (m *UnprovenRandomnessCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_UnprovenRandomnessCreated.DiscardUnknown(m)
}

var xxx_messageInfo_UnprovenRandomnessCreated proto.InternalMessageInfo

func (m *UnprovenRandomnessCreated) GetRound() uint64 {
	if m != nil {
		return m.Round
	}
	return 0
}

func (m *UnprovenRandomnessCreated) GetRoundTime() uint64 {
	if m != nil {
		return m.RoundTime
	}
	return 0
}

func init() {
	proto.RegisterType((*ProvenRandomnessCreated)(nil), "saturn.randomness.ProvenRandomnessCreated")
	proto.RegisterType((*UnprovenRandomnessCreated)(nil), "saturn.randomness.UnprovenRandomnessCreated")
}

func init() { proto.RegisterFile("randomness/event.proto", fileDescriptor_84d0e7b92baa19ab) }

var fileDescriptor_84d0e7b92baa19ab = []byte{
	// 247 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2b, 0x4a, 0xcc, 0x4b,
	0xc9, 0xcf, 0xcd, 0x4b, 0x2d, 0x2e, 0xd6, 0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x12, 0x2c, 0x4e, 0x2c, 0x29, 0x2d, 0xca, 0xd3, 0x43, 0x48, 0x2b, 0xed, 0x64,
	0xe4, 0x12, 0x0f, 0x28, 0xca, 0x2f, 0x4b, 0xcd, 0x0b, 0x82, 0x0b, 0x3a, 0x17, 0xa5, 0x26, 0x96,
	0xa4, 0xa6, 0x08, 0x89, 0x70, 0xb1, 0x16, 0xe5, 0x97, 0xe6, 0xa5, 0x48, 0x30, 0x2a, 0x30, 0x6a,
	0xb0, 0x04, 0x41, 0x38, 0x42, 0x72, 0x5c, 0x5c, 0x08, 0xfd, 0x12, 0x4c, 0x0a, 0x8c, 0x1a, 0x9c,
	0x41, 0x48, 0x22, 0x42, 0x32, 0x5c, 0x9c, 0xc5, 0x99, 0xe9, 0x79, 0x20, 0x9b, 0x52, 0x25, 0x98,
	0xc1, 0xd2, 0x08, 0x01, 0x21, 0x1d, 0x2e, 0xc1, 0x82, 0xa2, 0xd4, 0xb2, 0xcc, 0xfc, 0xd2, 0xe2,
	0x60, 0xb8, 0x2a, 0x16, 0xb0, 0x2a, 0x4c, 0x09, 0x90, 0x59, 0x60, 0x4b, 0x43, 0x32, 0x73, 0x53,
	0x25, 0xd8, 0xc0, 0xae, 0x40, 0x08, 0x28, 0xf9, 0x73, 0x49, 0x86, 0xe6, 0x15, 0x90, 0xe4, 0x78,
	0x14, 0x03, 0x99, 0xd0, 0x0c, 0x74, 0x72, 0x3f, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6,
	0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39,
	0x86, 0x28, 0xdd, 0xf4, 0xcc, 0x92, 0x8c, 0xd2, 0x24, 0xbd, 0xe4, 0xfc, 0x5c, 0xfd, 0x94, 0xa2,
	0xd4, 0xc4, 0xbc, 0xcc, 0x92, 0x4a, 0x7d, 0x48, 0x68, 0xea, 0x57, 0xe8, 0x23, 0x05, 0x77, 0x49,
	0x65, 0x41, 0x6a, 0x71, 0x12, 0x1b, 0x38, 0xbc, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x72,
	0x46, 0xc3, 0x3c, 0x89, 0x01, 0x00, 0x00,
}

func (m *ProvenRandomnessCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ProvenRandomnessCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ProvenRandomnessCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoundTime != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RoundTime))
		i--
		dAtA[i] = 0x30
	}
	if len(m.PreviousSignature) > 0 {
		i -= len(m.PreviousSignature)
		copy(dAtA[i:], m.PreviousSignature)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PreviousSignature)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Randomness) > 0 {
		i -= len(m.Randomness)
		copy(dAtA[i:], m.Randomness)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Randomness)))
		i--
		dAtA[i] = 0x12
	}
	if m.Round != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *UnprovenRandomnessCreated) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *UnprovenRandomnessCreated) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *UnprovenRandomnessCreated) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.RoundTime != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.RoundTime))
		i--
		dAtA[i] = 0x10
	}
	if m.Round != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.Round))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ProvenRandomnessCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovEvent(uint64(m.Round))
	}
	l = len(m.Randomness)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PreviousSignature)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.RoundTime != 0 {
		n += 1 + sovEvent(uint64(m.RoundTime))
	}
	return n
}

func (m *UnprovenRandomnessCreated) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Round != 0 {
		n += 1 + sovEvent(uint64(m.Round))
	}
	if m.RoundTime != 0 {
		n += 1 + sovEvent(uint64(m.RoundTime))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ProvenRandomnessCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: ProvenRandomnessCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ProvenRandomnessCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Randomness", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Randomness = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousSignature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousSignature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundTime", wireType)
			}
			m.RoundTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *UnprovenRandomnessCreated) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: UnprovenRandomnessCreated: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: UnprovenRandomnessCreated: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Round", wireType)
			}
			m.Round = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field RoundTime", wireType)
			}
			m.RoundTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.RoundTime |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)
