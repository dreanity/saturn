// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: giveaway/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCreateGiveaway struct {
	Creator  string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Duration int32  `protobuf:"varint,2,opt,name=duration,proto3" json:"duration,omitempty"`
	Name     string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
	Prize    *Prize `protobuf:"bytes,5,opt,name=prize,proto3" json:"prize,omitempty"`
}

func (m *MsgCreateGiveaway) Reset()         { *m = MsgCreateGiveaway{} }
func (m *MsgCreateGiveaway) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGiveaway) ProtoMessage()    {}
func (*MsgCreateGiveaway) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ad54ab600213fb, []int{0}
}
func (m *MsgCreateGiveaway) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGiveaway) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGiveaway.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGiveaway) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGiveaway.Merge(m, src)
}
func (m *MsgCreateGiveaway) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGiveaway) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGiveaway.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGiveaway proto.InternalMessageInfo

func (m *MsgCreateGiveaway) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCreateGiveaway) GetDuration() int32 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *MsgCreateGiveaway) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MsgCreateGiveaway) GetPrize() *Prize {
	if m != nil {
		return m.Prize
	}
	return nil
}

type MsgCreateGiveawayResponse struct {
}

func (m *MsgCreateGiveawayResponse) Reset()         { *m = MsgCreateGiveawayResponse{} }
func (m *MsgCreateGiveawayResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateGiveawayResponse) ProtoMessage()    {}
func (*MsgCreateGiveawayResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_97ad54ab600213fb, []int{1}
}
func (m *MsgCreateGiveawayResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateGiveawayResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateGiveawayResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateGiveawayResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateGiveawayResponse.Merge(m, src)
}
func (m *MsgCreateGiveawayResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateGiveawayResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateGiveawayResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateGiveawayResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateGiveaway)(nil), "dreanity.saturn.giveaway.MsgCreateGiveaway")
	proto.RegisterType((*MsgCreateGiveawayResponse)(nil), "dreanity.saturn.giveaway.MsgCreateGiveawayResponse")
}

func init() { proto.RegisterFile("giveaway/tx.proto", fileDescriptor_97ad54ab600213fb) }

var fileDescriptor_97ad54ab600213fb = []byte{
	// 268 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4c, 0xcf, 0x2c, 0x4b,
	0x4d, 0x2c, 0x4f, 0xac, 0xd4, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x48,
	0x29, 0x4a, 0x4d, 0xcc, 0xcb, 0x2c, 0xa9, 0xd4, 0x2b, 0x4e, 0x2c, 0x29, 0x2d, 0xca, 0xd3, 0x83,
	0x29, 0x91, 0x12, 0x81, 0x2b, 0x2e, 0x28, 0xca, 0xac, 0x4a, 0x85, 0xa8, 0x57, 0x9a, 0xc2, 0xc8,
	0x25, 0xe8, 0x5b, 0x9c, 0xee, 0x5c, 0x94, 0x9a, 0x58, 0x92, 0xea, 0x0e, 0x55, 0x21, 0x24, 0xc1,
	0xc5, 0x9e, 0x0c, 0x12, 0xc9, 0x2f, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x71, 0x85,
	0xa4, 0xb8, 0x38, 0x52, 0x4a, 0x8b, 0x12, 0x4b, 0x32, 0xf3, 0xf3, 0x24, 0x98, 0x14, 0x18, 0x35,
	0x58, 0x83, 0xe0, 0x7c, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09, 0x16, 0xb0, 0x16,
	0x30, 0x5b, 0xc8, 0x94, 0x8b, 0x15, 0x6c, 0x9d, 0x04, 0xab, 0x02, 0xa3, 0x06, 0xb7, 0x91, 0xbc,
	0x1e, 0x2e, 0xf7, 0xe9, 0x05, 0x80, 0x94, 0x05, 0x41, 0x54, 0x2b, 0x49, 0x73, 0x49, 0x62, 0xb8,
	0x2a, 0x28, 0xb5, 0xb8, 0x20, 0x3f, 0xaf, 0x38, 0xd5, 0xa8, 0x92, 0x8b, 0xd9, 0xb7, 0x38, 0x5d,
	0xa8, 0x88, 0x8b, 0x0f, 0xcd, 0xd9, 0xda, 0xb8, 0x4d, 0xc7, 0x30, 0x4d, 0xca, 0x98, 0x04, 0xc5,
	0x30, 0xab, 0x9d, 0x5c, 0x4f, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x09, 0x8f, 0xe5, 0x18, 0x2e, 0x3c, 0x96, 0x63, 0xb8, 0xf1, 0x58, 0x8e, 0x21, 0x4a, 0x3b,
	0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x66, 0xb0, 0x3e, 0xc4, 0x60,
	0xfd, 0x0a, 0x7d, 0x44, 0x44, 0x55, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0x03, 0xdf, 0x18, 0x10,
	0x00, 0x00, 0xff, 0xff, 0xe9, 0xe4, 0xc8, 0x79, 0xc1, 0x01, 0x00, 0x00,
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
	CreateGiveaway(ctx context.Context, in *MsgCreateGiveaway, opts ...grpc.CallOption) (*MsgCreateGiveawayResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateGiveaway(ctx context.Context, in *MsgCreateGiveaway, opts ...grpc.CallOption) (*MsgCreateGiveawayResponse, error) {
	out := new(MsgCreateGiveawayResponse)
	err := c.cc.Invoke(ctx, "/dreanity.saturn.giveaway.Msg/CreateGiveaway", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateGiveaway(context.Context, *MsgCreateGiveaway) (*MsgCreateGiveawayResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateGiveaway(ctx context.Context, req *MsgCreateGiveaway) (*MsgCreateGiveawayResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateGiveaway not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateGiveaway_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateGiveaway)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateGiveaway(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dreanity.saturn.giveaway.Msg/CreateGiveaway",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateGiveaway(ctx, req.(*MsgCreateGiveaway))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dreanity.saturn.giveaway.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateGiveaway",
			Handler:    _Msg_CreateGiveaway_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "giveaway/tx.proto",
}

func (m *MsgCreateGiveaway) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGiveaway) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGiveaway) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Prize != nil {
		{
			size, err := m.Prize.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x22
	}
	if m.Duration != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Duration))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateGiveawayResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateGiveawayResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateGiveawayResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateGiveaway) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Duration != 0 {
		n += 1 + sovTx(uint64(m.Duration))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Prize != nil {
		l = m.Prize.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateGiveawayResponse) Size() (n int) {
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
func (m *MsgCreateGiveaway) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateGiveaway: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGiveaway: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Duration", wireType)
			}
			m.Duration = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Duration |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prize", wireType)
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
			if m.Prize == nil {
				m.Prize = &Prize{}
			}
			if err := m.Prize.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *MsgCreateGiveawayResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateGiveawayResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateGiveawayResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
