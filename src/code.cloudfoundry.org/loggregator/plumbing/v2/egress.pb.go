// Code generated by protoc-gen-go.
// source: egress.proto
// DO NOT EDIT!

package loggregator_v2

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EgressRequest struct {
	ShardId string  `protobuf:"bytes,1,opt,name=shard_id,json=shardId" json:"shard_id,omitempty"`
	Filter  *Filter `protobuf:"bytes,2,opt,name=filter" json:"filter,omitempty"`
	// TODO: This can be removed once the envelope.deprecated_tags is removed.
	UsePreferredTags bool `protobuf:"varint,3,opt,name=use_preferred_tags,json=usePreferredTags" json:"use_preferred_tags,omitempty"`
}

func (m *EgressRequest) Reset()                    { *m = EgressRequest{} }
func (m *EgressRequest) String() string            { return proto.CompactTextString(m) }
func (*EgressRequest) ProtoMessage()               {}
func (*EgressRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *EgressRequest) GetShardId() string {
	if m != nil {
		return m.ShardId
	}
	return ""
}

func (m *EgressRequest) GetFilter() *Filter {
	if m != nil {
		return m.Filter
	}
	return nil
}

func (m *EgressRequest) GetUsePreferredTags() bool {
	if m != nil {
		return m.UsePreferredTags
	}
	return false
}

type Filter struct {
	SourceId string `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	// Types that are valid to be assigned to Message:
	//	*Filter_Log
	Message isFilter_Message `protobuf_oneof:"Message"`
}

func (m *Filter) Reset()                    { *m = Filter{} }
func (m *Filter) String() string            { return proto.CompactTextString(m) }
func (*Filter) ProtoMessage()               {}
func (*Filter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

type isFilter_Message interface {
	isFilter_Message()
}

type Filter_Log struct {
	Log *LogFilter `protobuf:"bytes,2,opt,name=log,oneof"`
}

func (*Filter_Log) isFilter_Message() {}

func (m *Filter) GetMessage() isFilter_Message {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Filter) GetSourceId() string {
	if m != nil {
		return m.SourceId
	}
	return ""
}

func (m *Filter) GetLog() *LogFilter {
	if x, ok := m.GetMessage().(*Filter_Log); ok {
		return x.Log
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Filter) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Filter_OneofMarshaler, _Filter_OneofUnmarshaler, _Filter_OneofSizer, []interface{}{
		(*Filter_Log)(nil),
	}
}

func _Filter_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Filter)
	// Message
	switch x := m.Message.(type) {
	case *Filter_Log:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Log); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Filter.Message has unexpected type %T", x)
	}
	return nil
}

func _Filter_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Filter)
	switch tag {
	case 2: // Message.log
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LogFilter)
		err := b.DecodeMessage(msg)
		m.Message = &Filter_Log{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Filter_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Filter)
	// Message
	switch x := m.Message.(type) {
	case *Filter_Log:
		s := proto.Size(x.Log)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LogFilter struct {
}

func (m *LogFilter) Reset()                    { *m = LogFilter{} }
func (m *LogFilter) String() string            { return proto.CompactTextString(m) }
func (*LogFilter) ProtoMessage()               {}
func (*LogFilter) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func init() {
	proto.RegisterType((*EgressRequest)(nil), "loggregator.v2.EgressRequest")
	proto.RegisterType((*Filter)(nil), "loggregator.v2.Filter")
	proto.RegisterType((*LogFilter)(nil), "loggregator.v2.LogFilter")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Egress service

type EgressClient interface {
	Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error)
}

type egressClient struct {
	cc *grpc.ClientConn
}

func NewEgressClient(cc *grpc.ClientConn) EgressClient {
	return &egressClient{cc}
}

func (c *egressClient) Receiver(ctx context.Context, in *EgressRequest, opts ...grpc.CallOption) (Egress_ReceiverClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Egress_serviceDesc.Streams[0], c.cc, "/loggregator.v2.Egress/Receiver", opts...)
	if err != nil {
		return nil, err
	}
	x := &egressReceiverClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Egress_ReceiverClient interface {
	Recv() (*Envelope, error)
	grpc.ClientStream
}

type egressReceiverClient struct {
	grpc.ClientStream
}

func (x *egressReceiverClient) Recv() (*Envelope, error) {
	m := new(Envelope)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Egress service

type EgressServer interface {
	Receiver(*EgressRequest, Egress_ReceiverServer) error
}

func RegisterEgressServer(s *grpc.Server, srv EgressServer) {
	s.RegisterService(&_Egress_serviceDesc, srv)
}

func _Egress_Receiver_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EgressRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EgressServer).Receiver(m, &egressReceiverServer{stream})
}

type Egress_ReceiverServer interface {
	Send(*Envelope) error
	grpc.ServerStream
}

type egressReceiverServer struct {
	grpc.ServerStream
}

func (x *egressReceiverServer) Send(m *Envelope) error {
	return x.ServerStream.SendMsg(m)
}

var _Egress_serviceDesc = grpc.ServiceDesc{
	ServiceName: "loggregator.v2.Egress",
	HandlerType: (*EgressServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Receiver",
			Handler:       _Egress_Receiver_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "egress.proto",
}

func init() { proto.RegisterFile("egress.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x17, 0x07, 0x5d, 0x7b, 0xa7, 0x43, 0xf2, 0x20, 0xdd, 0x44, 0x28, 0x7d, 0xea, 0x83,
	0x16, 0xa9, 0xdf, 0x40, 0xf0, 0xcf, 0x40, 0x41, 0x83, 0xef, 0xa5, 0x2e, 0x77, 0xb1, 0x50, 0x4c,
	0xbd, 0x37, 0xed, 0x67, 0xf0, 0x63, 0x8b, 0x4d, 0x55, 0x36, 0x1f, 0x93, 0xdf, 0x39, 0x87, 0x73,
	0x0f, 0x1c, 0xa2, 0x21, 0x64, 0xce, 0x5b, 0xb2, 0xce, 0xca, 0x45, 0x63, 0x8d, 0x21, 0x34, 0x95,
	0xb3, 0x94, 0xf7, 0xc5, 0x6a, 0x81, 0xef, 0x3d, 0x36, 0xb6, 0x45, 0xcf, 0xd3, 0x4f, 0x01, 0x47,
	0x37, 0x83, 0x41, 0xe1, 0x47, 0x87, 0xec, 0xe4, 0x12, 0x42, 0x7e, 0xab, 0x48, 0x97, 0xb5, 0x8e,
	0x45, 0x22, 0xb2, 0x48, 0xcd, 0x86, 0xf7, 0x5a, 0xcb, 0x1c, 0x82, 0x6d, 0xdd, 0x38, 0xa4, 0xf8,
	0x20, 0x11, 0xd9, 0xbc, 0x38, 0xc9, 0x77, 0xd3, 0xf3, 0xdb, 0x81, 0xaa, 0x51, 0x25, 0xcf, 0x41,
	0x76, 0x8c, 0x65, 0x4b, 0xb8, 0x45, 0x22, 0xd4, 0xa5, 0xab, 0x0c, 0xc7, 0xd3, 0x44, 0x64, 0xa1,
	0x3a, 0xee, 0x18, 0x9f, 0x7e, 0xc0, 0x4b, 0x65, 0x38, 0x2d, 0x21, 0xf0, 0x7e, 0x79, 0x0a, 0x11,
	0xdb, 0x8e, 0x36, 0xf8, 0xd7, 0x21, 0xf4, 0x1f, 0x6b, 0x2d, 0x2f, 0x60, 0xda, 0x58, 0x33, 0x36,
	0x58, 0xee, 0x37, 0x78, 0xb0, 0xc6, 0x87, 0xdc, 0x4f, 0xd4, 0xb7, 0xee, 0x3a, 0x82, 0xd9, 0x23,
	0x32, 0x57, 0x06, 0xd3, 0x39, 0x44, 0xbf, 0xb8, 0x78, 0x86, 0xc0, 0xdf, 0x2d, 0xef, 0x20, 0x54,
	0xb8, 0xc1, 0xba, 0x47, 0x92, 0x67, 0xfb, 0x79, 0x3b, 0xdb, 0xac, 0xe2, 0x7f, 0x78, 0x5c, 0x33,
	0x9d, 0x5c, 0x8a, 0xd7, 0x60, 0x98, 0xf4, 0xea, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x69, 0xb4, 0x70,
	0x7c, 0x82, 0x01, 0x00, 0x00,
}
