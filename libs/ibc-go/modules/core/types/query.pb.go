// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc/core/types/v1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/codec/types"
	_ "github.com/furyaxyz/fuxchain/libs/cosmos-sdk/types/query"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryIbcParamsRequest is the request type for the Query/IbcParams RPC
// method.
type QueryIbcParamsRequest struct {
}

func (m *QueryIbcParamsRequest) Reset()         { *m = QueryIbcParamsRequest{} }
func (m *QueryIbcParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIbcParamsRequest) ProtoMessage()    {}
func (*QueryIbcParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc0ad6869acad8f, []int{0}
}
func (m *QueryIbcParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIbcParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIbcParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIbcParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIbcParamsRequest.Merge(m, src)
}
func (m *QueryIbcParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIbcParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIbcParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIbcParamsRequest proto.InternalMessageInfo

// QueryIbcParamsResponse is the response type for the Query/IbcParams RPC
// method.
type QueryIbcParamsResponse struct {
	// params defines the parameters of the module.
	Params *Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
}

func (m *QueryIbcParamsResponse) Reset()         { *m = QueryIbcParamsResponse{} }
func (m *QueryIbcParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIbcParamsResponse) ProtoMessage()    {}
func (*QueryIbcParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8cc0ad6869acad8f, []int{1}
}
func (m *QueryIbcParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIbcParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIbcParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIbcParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIbcParamsResponse.Merge(m, src)
}
func (m *QueryIbcParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIbcParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIbcParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIbcParamsResponse proto.InternalMessageInfo

func (m *QueryIbcParamsResponse) GetParams() *Params {
	if m != nil {
		return m.Params
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryIbcParamsRequest)(nil), "ibc.core.v1.QueryIbcParamsRequest")
	proto.RegisterType((*QueryIbcParamsResponse)(nil), "ibc.core.v1.QueryIbcParamsResponse")
}

func init() { proto.RegisterFile("ibc/core/types/v1/query.proto", fileDescriptor_8cc0ad6869acad8f) }

var fileDescriptor_8cc0ad6869acad8f = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x51, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x6e, 0x04, 0x0b, 0xa6, 0x20, 0x12, 0xb5, 0x6a, 0xd0, 0x28, 0xf1, 0x22, 0x8a, 0x3b, 0xa4,
	0xbe, 0x81, 0x20, 0xe8, 0x4d, 0x3d, 0x7a, 0xdb, 0x5d, 0xc7, 0x75, 0xa1, 0xd9, 0x49, 0xb3, 0x9b,
	0x40, 0xaf, 0x3e, 0x81, 0xe0, 0x4b, 0x79, 0x2c, 0x78, 0xf1, 0x28, 0xad, 0x0f, 0x22, 0xf9, 0x51,
	0xea, 0x0f, 0xde, 0x26, 0xf3, 0xfd, 0xe4, 0xdb, 0x6f, 0xfc, 0x1d, 0x2d, 0x24, 0x48, 0xca, 0x11,
	0xdc, 0x38, 0x43, 0x0b, 0x65, 0x02, 0xa3, 0x02, 0xf3, 0x31, 0xcb, 0x72, 0x72, 0x14, 0xf4, 0xb4,
	0x90, 0xac, 0x82, 0x59, 0x99, 0x84, 0x87, 0x92, 0x6c, 0x4a, 0x16, 0x04, 0xb7, 0xd8, 0xb0, 0xa0,
	0x4c, 0x04, 0x3a, 0x9e, 0x40, 0xc6, 0x95, 0x36, 0xdc, 0x69, 0x32, 0x8d, 0x30, 0xdc, 0xfd, 0xed,
	0xab, 0xd0, 0xa0, 0xd5, 0xb6, 0x25, 0x6c, 0x29, 0x22, 0x35, 0x44, 0xa8, 0xbf, 0x44, 0x71, 0x07,
	0xdc, 0xb4, 0x3f, 0x0d, 0xb7, 0x5b, 0x88, 0x67, 0x1a, 0xb8, 0x31, 0xe4, 0x6a, 0xe3, 0x4f, 0xe1,
	0x9a, 0x22, 0x45, 0xf5, 0x08, 0xd5, 0xd4, 0x6c, 0xe3, 0x0d, 0x7f, 0xfd, 0xaa, 0x4a, 0x74, 0x21,
	0xe4, 0x25, 0xcf, 0x79, 0x6a, 0xaf, 0x71, 0x54, 0xa0, 0x75, 0xf1, 0x99, 0xdf, 0xff, 0x09, 0xd8,
	0x8c, 0x8c, 0xc5, 0xe0, 0xc8, 0xef, 0x66, 0xf5, 0x66, 0xd3, 0xdb, 0xf3, 0x0e, 0x7a, 0x83, 0x55,
	0x36, 0xf7, 0x58, 0xd6, 0x92, 0x5b, 0xca, 0xa0, 0xf0, 0x17, 0x6b, 0x9b, 0x60, 0xe8, 0x2f, 0x7d,
	0x59, 0x05, 0xf1, 0x37, 0xc9, 0x9f, 0x01, 0xc2, 0xfd, 0x7f, 0x39, 0x4d, 0x96, 0xb8, 0xff, 0xf0,
	0xf2, 0xfe, 0xb4, 0xb0, 0x12, 0x2c, 0x43, 0xd5, 0x5b, 0x59, 0x15, 0x5a, 0xe1, 0xa7, 0xe7, 0xcf,
	0xd3, 0xc8, 0x9b, 0x4c, 0x23, 0xef, 0x6d, 0x1a, 0x79, 0x8f, 0xb3, 0xa8, 0x33, 0x99, 0x45, 0x9d,
	0xd7, 0x59, 0xd4, 0xb9, 0x61, 0x4a, 0xbb, 0xfb, 0x42, 0x30, 0x49, 0x29, 0xb4, 0x77, 0xd1, 0x42,
	0x1e, 0x2b, 0x82, 0x72, 0x00, 0x29, 0xdd, 0x16, 0x43, 0xb4, 0x73, 0x07, 0x10, 0xdd, 0xba, 0xa7,
	0x93, 0x8f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x1e, 0x7a, 0x77, 0xf1, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// ClientParams queries all parameters of the ibc module.
	IbcParams(ctx context.Context, in *QueryIbcParamsRequest, opts ...grpc.CallOption) (*QueryIbcParamsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IbcParams(ctx context.Context, in *QueryIbcParamsRequest, opts ...grpc.CallOption) (*QueryIbcParamsResponse, error) {
	out := new(QueryIbcParamsResponse)
	err := c.cc.Invoke(ctx, "/ibc.core.v1.Query/IbcParams", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// ClientParams queries all parameters of the ibc module.
	IbcParams(context.Context, *QueryIbcParamsRequest) (*QueryIbcParamsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IbcParams(ctx context.Context, req *QueryIbcParamsRequest) (*QueryIbcParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IbcParams not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IbcParams_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIbcParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IbcParams(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ibc.core.v1.Query/IbcParams",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IbcParams(ctx, req.(*QueryIbcParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ibc.core.v1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IbcParams",
			Handler:    _Query_IbcParams_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc/core/types/v1/query.proto",
}

func (m *QueryIbcParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIbcParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIbcParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryIbcParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIbcParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIbcParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Params != nil {
		{
			size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryIbcParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryIbcParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Params != nil {
		l = m.Params.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryIbcParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIbcParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIbcParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryIbcParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryIbcParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIbcParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Params == nil {
				m.Params = &Params{}
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
