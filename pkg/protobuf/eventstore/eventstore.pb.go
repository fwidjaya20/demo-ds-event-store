// Code generated by protoc-gen-go. DO NOT EDIT.
// source: pkg/protobuf/eventstore/eventstore.proto

package eventstore

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Event struct {
	EventId              string   `protobuf:"bytes,1,opt,name=event_id,json=eventId,proto3" json:"event_id,omitempty"`
	EventType            string   `protobuf:"bytes,2,opt,name=event_type,json=eventType,proto3" json:"event_type,omitempty"`
	AggregateId          string   `protobuf:"bytes,3,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateType        string   `protobuf:"bytes,4,opt,name=aggregate_type,json=aggregateType,proto3" json:"aggregate_type,omitempty"`
	EventData            string   `protobuf:"bytes,5,opt,name=event_data,json=eventData,proto3" json:"event_data,omitempty"`
	Channel              string   `protobuf:"bytes,6,opt,name=channel,proto3" json:"channel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Event) Reset()         { *m = Event{} }
func (m *Event) String() string { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()    {}
func (*Event) Descriptor() ([]byte, []int) {
	return fileDescriptor_00f0fcda0ffd967c, []int{0}
}

func (m *Event) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Event.Unmarshal(m, b)
}
func (m *Event) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Event.Marshal(b, m, deterministic)
}
func (m *Event) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Event.Merge(m, src)
}
func (m *Event) XXX_Size() int {
	return xxx_messageInfo_Event.Size(m)
}
func (m *Event) XXX_DiscardUnknown() {
	xxx_messageInfo_Event.DiscardUnknown(m)
}

var xxx_messageInfo_Event proto.InternalMessageInfo

func (m *Event) GetEventId() string {
	if m != nil {
		return m.EventId
	}
	return ""
}

func (m *Event) GetEventType() string {
	if m != nil {
		return m.EventType
	}
	return ""
}

func (m *Event) GetAggregateId() string {
	if m != nil {
		return m.AggregateId
	}
	return ""
}

func (m *Event) GetAggregateType() string {
	if m != nil {
		return m.AggregateType
	}
	return ""
}

func (m *Event) GetEventData() string {
	if m != nil {
		return m.EventData
	}
	return ""
}

func (m *Event) GetChannel() string {
	if m != nil {
		return m.Channel
	}
	return ""
}

type EventResponse struct {
	Event                *Event   `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EventResponse) Reset()         { *m = EventResponse{} }
func (m *EventResponse) String() string { return proto.CompactTextString(m) }
func (*EventResponse) ProtoMessage()    {}
func (*EventResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_00f0fcda0ffd967c, []int{1}
}

func (m *EventResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EventResponse.Unmarshal(m, b)
}
func (m *EventResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EventResponse.Marshal(b, m, deterministic)
}
func (m *EventResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EventResponse.Merge(m, src)
}
func (m *EventResponse) XXX_Size() int {
	return xxx_messageInfo_EventResponse.Size(m)
}
func (m *EventResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EventResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EventResponse proto.InternalMessageInfo

func (m *EventResponse) GetEvent() *Event {
	if m != nil {
		return m.Event
	}
	return nil
}

func init() {
	proto.RegisterType((*Event)(nil), "eventstore.Event")
	proto.RegisterType((*EventResponse)(nil), "eventstore.EventResponse")
}

func init() {
	proto.RegisterFile("pkg/protobuf/eventstore/eventstore.proto", fileDescriptor_00f0fcda0ffd967c)
}

var fileDescriptor_00f0fcda0ffd967c = []byte{
	// 243 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x50, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0xf6, 0xd4, 0x4b, 0xcc, 0x9c, 0x11, 0x9c, 0x6a, 0x23, 0x08, 0x7a, 0x20, 0xa6, 0x4a, 0x20,
	0x36, 0x36, 0x56, 0x6a, 0x11, 0xec, 0xa2, 0x7d, 0x98, 0x78, 0xe3, 0x2a, 0xca, 0xee, 0xb2, 0x37,
	0x0a, 0x79, 0x3f, 0x1f, 0x2c, 0xdc, 0x6c, 0xb8, 0x1c, 0xa4, 0xdb, 0xef, 0x6f, 0x66, 0xe7, 0x83,
	0x71, 0xf8, 0xb6, 0xd3, 0x10, 0xbd, 0xf8, 0xd5, 0xef, 0xc7, 0x94, 0xff, 0xd8, 0x49, 0x2d, 0x3e,
	0x72, 0xe7, 0x39, 0x51, 0x19, 0x61, 0xc7, 0x94, 0xff, 0x19, 0xe4, 0xcf, 0x0d, 0xc4, 0x11, 0x9c,
	0x28, 0xbf, 0xfc, 0xaa, 0x4c, 0x76, 0x95, 0x8d, 0x07, 0x8b, 0xbe, 0xe2, 0x79, 0x85, 0x97, 0x90,
	0x22, 0x4b, 0x59, 0x07, 0x36, 0x87, 0x2a, 0x0e, 0x94, 0x79, 0x5b, 0x07, 0xc6, 0x6b, 0x38, 0x25,
	0x6b, 0x23, 0x5b, 0x12, 0x6e, 0xd2, 0x47, 0x6a, 0x28, 0x5a, 0x6e, 0x5e, 0xe1, 0x0d, 0x9c, 0xed,
	0x2c, 0x3a, 0xe5, 0x58, 0x4d, 0xc3, 0x96, 0xd5, 0x49, 0xed, 0xa2, 0x8a, 0x84, 0x4c, 0xde, 0x59,
	0xf4, 0x44, 0x42, 0x68, 0xa0, 0xff, 0xfe, 0x49, 0xce, 0xf1, 0x8f, 0xe9, 0xa5, 0x1f, 0x6e, 0x61,
	0x79, 0x0f, 0x43, 0xbd, 0x62, 0xc1, 0x75, 0xf0, 0xae, 0x66, 0xbc, 0x85, 0x5c, 0x73, 0x7a, 0x4a,
	0x31, 0x3b, 0x9f, 0x74, 0x5a, 0x48, 0xce, 0xa4, 0xcf, 0x5e, 0x00, 0x14, 0xbf, 0x36, 0x12, 0x3e,
	0x40, 0xf1, 0x18, 0x99, 0x84, 0x53, 0x27, 0xfb, 0xb1, 0x8b, 0xd1, 0xfe, 0xa4, 0xed, 0xce, 0xf2,
	0x60, 0xd5, 0xd3, 0x82, 0xef, 0x36, 0x01, 0x00, 0x00, 0xff, 0xff, 0x43, 0x8d, 0x45, 0x68, 0x8c,
	0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// EventStoreClient is the client API for EventStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type EventStoreClient interface {
	CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventResponse, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) CreateEvent(ctx context.Context, in *Event, opts ...grpc.CallOption) (*EventResponse, error) {
	out := new(EventResponse)
	err := c.cc.Invoke(ctx, "/eventstore.EventStore/CreateEvent", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EventStoreServer is the server API for EventStore service.
type EventStoreServer interface {
	CreateEvent(context.Context, *Event) (*EventResponse, error)
}

// UnimplementedEventStoreServer can be embedded to have forward compatible implementations.
type UnimplementedEventStoreServer struct {
}

func (*UnimplementedEventStoreServer) CreateEvent(ctx context.Context, req *Event) (*EventResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateEvent not implemented")
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_CreateEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Event)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).CreateEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventstore.EventStore/CreateEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).CreateEvent(ctx, req.(*Event))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventstore.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateEvent",
			Handler:    _EventStore_CreateEvent_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pkg/protobuf/eventstore/eventstore.proto",
}