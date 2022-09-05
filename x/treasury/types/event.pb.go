// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: treasury/event.proto

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

type GasPricesChanged struct {
	GasPrices []*GasPrice `protobuf:"bytes,1,rep,name=gasPrices,proto3" json:"gasPrices,omitempty"`
}

func (m *GasPricesChanged) Reset()         { *m = GasPricesChanged{} }
func (m *GasPricesChanged) String() string { return proto.CompactTextString(m) }
func (*GasPricesChanged) ProtoMessage()    {}
func (*GasPricesChanged) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ee0673910aebc86, []int{0}
}
func (m *GasPricesChanged) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasPricesChanged) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasPricesChanged.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasPricesChanged) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasPricesChanged.Merge(m, src)
}
func (m *GasPricesChanged) XXX_Size() int {
	return m.Size()
}
func (m *GasPricesChanged) XXX_DiscardUnknown() {
	xxx_messageInfo_GasPricesChanged.DiscardUnknown(m)
}

var xxx_messageInfo_GasPricesChanged proto.InternalMessageInfo

func (m *GasPricesChanged) GetGasPrices() []*GasPrice {
	if m != nil {
		return m.GasPrices
	}
	return nil
}

type GasBidExecuted struct {
	BidNumber    uint64 `protobuf:"varint,1,opt,name=bidNumber,proto3" json:"bidNumber,omitempty"`
	Recepient    string `protobuf:"bytes,2,opt,name=recepient,proto3" json:"recepient,omitempty"`
	MintedAmount string `protobuf:"bytes,3,opt,name=mintedAmount,proto3" json:"mintedAmount,omitempty"`
	PaidAmount   string `protobuf:"bytes,4,opt,name=paidAmount,proto3" json:"paidAmount,omitempty"`
	TokenAddress string `protobuf:"bytes,5,opt,name=tokenAddress,proto3" json:"tokenAddress,omitempty"`
	Chain        string `protobuf:"bytes,6,opt,name=chain,proto3" json:"chain,omitempty"`
	GasPrice     string `protobuf:"bytes,7,opt,name=gasPrice,proto3" json:"gasPrice,omitempty"`
}

func (m *GasBidExecuted) Reset()         { *m = GasBidExecuted{} }
func (m *GasBidExecuted) String() string { return proto.CompactTextString(m) }
func (*GasBidExecuted) ProtoMessage()    {}
func (*GasBidExecuted) Descriptor() ([]byte, []int) {
	return fileDescriptor_6ee0673910aebc86, []int{1}
}
func (m *GasBidExecuted) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GasBidExecuted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GasBidExecuted.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GasBidExecuted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GasBidExecuted.Merge(m, src)
}
func (m *GasBidExecuted) XXX_Size() int {
	return m.Size()
}
func (m *GasBidExecuted) XXX_DiscardUnknown() {
	xxx_messageInfo_GasBidExecuted.DiscardUnknown(m)
}

var xxx_messageInfo_GasBidExecuted proto.InternalMessageInfo

func (m *GasBidExecuted) GetBidNumber() uint64 {
	if m != nil {
		return m.BidNumber
	}
	return 0
}

func (m *GasBidExecuted) GetRecepient() string {
	if m != nil {
		return m.Recepient
	}
	return ""
}

func (m *GasBidExecuted) GetMintedAmount() string {
	if m != nil {
		return m.MintedAmount
	}
	return ""
}

func (m *GasBidExecuted) GetPaidAmount() string {
	if m != nil {
		return m.PaidAmount
	}
	return ""
}

func (m *GasBidExecuted) GetTokenAddress() string {
	if m != nil {
		return m.TokenAddress
	}
	return ""
}

func (m *GasBidExecuted) GetChain() string {
	if m != nil {
		return m.Chain
	}
	return ""
}

func (m *GasBidExecuted) GetGasPrice() string {
	if m != nil {
		return m.GasPrice
	}
	return ""
}

func init() {
	proto.RegisterType((*GasPricesChanged)(nil), "dreanity.saturn.treasury.GasPricesChanged")
	proto.RegisterType((*GasBidExecuted)(nil), "dreanity.saturn.treasury.GasBidExecuted")
}

func init() { proto.RegisterFile("treasury/event.proto", fileDescriptor_6ee0673910aebc86) }

var fileDescriptor_6ee0673910aebc86 = []byte{
	// 313 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0xc7, 0xbb, 0xf6, 0x43, 0xbb, 0x8a, 0xc8, 0xd2, 0xc3, 0x52, 0x64, 0x29, 0x39, 0x15, 0x84,
	0x04, 0xf4, 0x05, 0x6c, 0xa5, 0xf4, 0x26, 0x52, 0x3c, 0x79, 0x91, 0x4d, 0x76, 0x48, 0x17, 0xc9,
	0x26, 0xec, 0x4e, 0xa4, 0x7d, 0x0b, 0x1f, 0xcb, 0x63, 0x8f, 0x1e, 0xa5, 0x05, 0x9f, 0x43, 0x9a,
	0xb0, 0x4d, 0x3d, 0x78, 0x9c, 0xdf, 0xff, 0x03, 0x66, 0x86, 0x0e, 0xd0, 0x82, 0x74, 0xa5, 0x5d,
	0x47, 0xf0, 0x0e, 0x06, 0xc3, 0xc2, 0xe6, 0x98, 0x33, 0xae, 0x2c, 0x48, 0xa3, 0x71, 0x1d, 0x3a,
	0x89, 0xa5, 0x35, 0xa1, 0x77, 0x0d, 0xf9, 0xc1, 0x9f, 0x4a, 0xf7, 0x5a, 0x58, 0x9d, 0x40, 0x9d,
	0x09, 0x9e, 0xe9, 0xd5, 0x5c, 0xba, 0xa7, 0x3d, 0x71, 0x0f, 0x4b, 0x69, 0x52, 0x50, 0xec, 0x9e,
	0xf6, 0x53, 0xcf, 0x38, 0x19, 0xb5, 0xc7, 0xe7, 0xb7, 0x41, 0xf8, 0x5f, 0x77, 0xe8, 0xe3, 0x8b,
	0x26, 0x14, 0xfc, 0x10, 0x7a, 0x39, 0x97, 0x6e, 0xaa, 0xd5, 0x6c, 0x05, 0x49, 0x89, 0xa0, 0xd8,
	0x35, 0xed, 0xc7, 0x5a, 0x3d, 0x96, 0x59, 0x0c, 0x96, 0x93, 0x11, 0x19, 0x77, 0x16, 0x0d, 0xd8,
	0xab, 0x16, 0x12, 0x28, 0x34, 0x18, 0xe4, 0x27, 0x23, 0x32, 0xee, 0x2f, 0x1a, 0xc0, 0x02, 0x7a,
	0x91, 0x69, 0x83, 0xa0, 0x26, 0x59, 0x5e, 0x1a, 0xe4, 0xed, 0xca, 0xf0, 0x87, 0x31, 0x41, 0x69,
	0x21, 0xb5, 0x77, 0x74, 0x2a, 0xc7, 0x11, 0xd9, 0x77, 0x60, 0xfe, 0x06, 0x66, 0xa2, 0x94, 0x05,
	0xe7, 0x78, 0xb7, 0xee, 0x38, 0x66, 0x6c, 0x40, 0xbb, 0xc9, 0x52, 0x6a, 0xc3, 0x7b, 0x95, 0x58,
	0x0f, 0x6c, 0x48, 0xcf, 0xfc, 0x66, 0xfc, 0xb4, 0x12, 0x0e, 0xf3, 0x74, 0xf6, 0xb9, 0x15, 0x64,
	0xb3, 0x15, 0xe4, 0x7b, 0x2b, 0xc8, 0xc7, 0x4e, 0xb4, 0x36, 0x3b, 0xd1, 0xfa, 0xda, 0x89, 0xd6,
	0xcb, 0x4d, 0xaa, 0x71, 0x59, 0xc6, 0x61, 0x92, 0x67, 0x91, 0xbf, 0x5d, 0x54, 0xdf, 0x2e, 0x5a,
	0x45, 0x87, 0x7f, 0xe0, 0xba, 0x00, 0x17, 0xf7, 0xaa, 0x67, 0xdc, 0xfd, 0x06, 0x00, 0x00, 0xff,
	0xff, 0xd5, 0x48, 0xd6, 0xb0, 0xd8, 0x01, 0x00, 0x00,
}

func (m *GasPricesChanged) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasPricesChanged) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasPricesChanged) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for iNdEx := len(m.GasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.GasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintEvent(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GasBidExecuted) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GasBidExecuted) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GasBidExecuted) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.GasPrice) > 0 {
		i -= len(m.GasPrice)
		copy(dAtA[i:], m.GasPrice)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.GasPrice)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Chain) > 0 {
		i -= len(m.Chain)
		copy(dAtA[i:], m.Chain)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Chain)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.TokenAddress) > 0 {
		i -= len(m.TokenAddress)
		copy(dAtA[i:], m.TokenAddress)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.TokenAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PaidAmount) > 0 {
		i -= len(m.PaidAmount)
		copy(dAtA[i:], m.PaidAmount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.PaidAmount)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.MintedAmount) > 0 {
		i -= len(m.MintedAmount)
		copy(dAtA[i:], m.MintedAmount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.MintedAmount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Recepient) > 0 {
		i -= len(m.Recepient)
		copy(dAtA[i:], m.Recepient)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Recepient)))
		i--
		dAtA[i] = 0x12
	}
	if m.BidNumber != 0 {
		i = encodeVarintEvent(dAtA, i, uint64(m.BidNumber))
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
func (m *GasPricesChanged) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.GasPrices) > 0 {
		for _, e := range m.GasPrices {
			l = e.Size()
			n += 1 + l + sovEvent(uint64(l))
		}
	}
	return n
}

func (m *GasBidExecuted) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BidNumber != 0 {
		n += 1 + sovEvent(uint64(m.BidNumber))
	}
	l = len(m.Recepient)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.MintedAmount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.PaidAmount)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.TokenAddress)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Chain)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.GasPrice)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GasPricesChanged) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GasPricesChanged: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasPricesChanged: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.GasPrices = append(m.GasPrices, &GasPrice{})
			if err := m.GasPrices[len(m.GasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
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
func (m *GasBidExecuted) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GasBidExecuted: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GasBidExecuted: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BidNumber", wireType)
			}
			m.BidNumber = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BidNumber |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Recepient", wireType)
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
			m.Recepient = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MintedAmount", wireType)
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
			m.MintedAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaidAmount", wireType)
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
			m.PaidAmount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenAddress", wireType)
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
			m.TokenAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Chain", wireType)
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
			m.Chain = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GasPrice", wireType)
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
			m.GasPrice = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
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