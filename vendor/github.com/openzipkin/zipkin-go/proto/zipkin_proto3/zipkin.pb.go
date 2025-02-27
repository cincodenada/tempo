// Copyright 2021 The OpenZipkin Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.0
// source: proto/zipkin_proto3/zipkin.proto

// This is the package for using protobuf with Zipkin API V2, but for historical
// reasons uses the protoc syntax version instead.

package zipkin_proto3

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// When present, kind clarifies timestamp, duration and remote_endpoint. When
// absent, the span is local or incomplete. Unlike client and server, there
// is no direct critical path latency relationship between producer and
// consumer spans.
type Span_Kind int32

const (
	// Default value interpreted as absent.
	Span_SPAN_KIND_UNSPECIFIED Span_Kind = 0
	// The span represents the client side of an RPC operation, implying the
	// following:
	//
	// timestamp is the moment a request was sent to the server.
	// duration is the delay until a response or an error was received.
	// remote_endpoint is the server.
	Span_CLIENT Span_Kind = 1
	// The span represents the server side of an RPC operation, implying the
	// following:
	//
	// timestamp is the moment a client request was received.
	// duration is the delay until a response was sent or an error.
	// remote_endpoint is the client.
	Span_SERVER Span_Kind = 2
	// The span represents production of a message to a remote broker, implying
	// the following:
	//
	// timestamp is the moment a message was sent to a destination.
	// duration is the delay sending the message, such as batching.
	// remote_endpoint is the broker.
	Span_PRODUCER Span_Kind = 3
	// The span represents consumption of a message from a remote broker, not
	// time spent servicing it. For example, a message processor would be an
	// in-process child span of a consumer. Consumer spans imply the following:
	//
	// timestamp is the moment a message was received from an origin.
	// duration is the delay consuming the message, such as from backlog.
	// remote_endpoint is the broker.
	Span_CONSUMER Span_Kind = 4
)

// Enum value maps for Span_Kind.
var (
	Span_Kind_name = map[int32]string{
		0: "SPAN_KIND_UNSPECIFIED",
		1: "CLIENT",
		2: "SERVER",
		3: "PRODUCER",
		4: "CONSUMER",
	}
	Span_Kind_value = map[string]int32{
		"SPAN_KIND_UNSPECIFIED": 0,
		"CLIENT":                1,
		"SERVER":                2,
		"PRODUCER":              3,
		"CONSUMER":              4,
	}
)

func (x Span_Kind) Enum() *Span_Kind {
	p := new(Span_Kind)
	*p = x
	return p
}

func (x Span_Kind) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Span_Kind) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_zipkin_proto3_zipkin_proto_enumTypes[0].Descriptor()
}

func (Span_Kind) Type() protoreflect.EnumType {
	return &file_proto_zipkin_proto3_zipkin_proto_enumTypes[0]
}

func (x Span_Kind) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Span_Kind.Descriptor instead.
func (Span_Kind) EnumDescriptor() ([]byte, []int) {
	return file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP(), []int{0, 0}
}

// A span is a single-host view of an operation. A trace is a series of spans
// (often RPC calls) which nest to form a latency tree. Spans are in the same
// trace when they share the same trace ID. The parent_id field establishes the
// position of one span in the tree.
//
// The root span is where parent_id is Absent and usually has the longest
// duration in the trace. However, nested asynchronous work can materialize as
// child spans whose duration exceed the root span.
//
// Spans usually represent remote activity such as RPC calls, or messaging
// producers and consumers. However, they can also represent in-process
// activity in any position of the trace. For example, a root span could
// represent a server receiving an initial client request. A root span could
// also represent a scheduled job that has no remote context.
//
// Encoding notes:
//
// Epoch timestamp are encoded fixed64 as varint would also be 8 bytes, and more
// expensive to encode and size. Duration is stored uint64, as often the numbers
// are quite small.
//
// Default values are ok, as only natural numbers are used. For example, zero is
// an invalid timestamp and an invalid duration, false values for debug or shared
// are ignorable, and zero-length strings also coerce to null.
//
// The next id is 14.
//
// Note fields up to 15 take 1 byte to encode. Take care when adding new fields
// https://developers.google.com/protocol-buffers/docs/proto3#assigning-tags
type Span struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Randomly generated, unique identifier for a trace, set on all spans within
	// it.
	//
	// This field is required and encoded as 8 or 16 opaque bytes.
	TraceId []byte `protobuf:"bytes,1,opt,name=trace_id,json=traceId,proto3" json:"trace_id,omitempty"`
	// The parent span ID or absent if this the root span in a trace.
	ParentId []byte `protobuf:"bytes,2,opt,name=parent_id,json=parentId,proto3" json:"parent_id,omitempty"`
	// Unique identifier for this operation within the trace.
	//
	// This field is required and encoded as 8 opaque bytes.
	Id []byte `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	// When present, used to interpret remote_endpoint
	Kind Span_Kind `protobuf:"varint,4,opt,name=kind,proto3,enum=zipkin.proto3.Span_Kind" json:"kind,omitempty"`
	// The logical operation this span represents in lowercase (e.g. rpc method).
	// Leave absent if unknown.
	//
	// As these are lookup labels, take care to ensure names are low cardinality.
	// For example, do not embed variables into the name.
	Name string `protobuf:"bytes,5,opt,name=name,proto3" json:"name,omitempty"`
	// Epoch microseconds of the start of this span, possibly absent if
	// incomplete.
	//
	// For example, 1502787600000000 corresponds to 2017-08-15 09:00 UTC
	//
	// This value should be set directly by instrumentation, using the most
	// precise value possible. For example, gettimeofday or multiplying epoch
	// millis by 1000.
	//
	// There are three known edge-cases where this could be reported absent.
	// - A span was allocated but never started (ex not yet received a timestamp)
	// - The span's start event was lost
	// - Data about a completed span (ex tags) were sent after the fact
	Timestamp uint64 `protobuf:"fixed64,6,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Duration in microseconds of the critical path, if known. Durations of less
	// than one are rounded up. Duration of children can be longer than their
	// parents due to asynchronous operations.
	//
	// For example 150 milliseconds is 150000 microseconds.
	Duration uint64 `protobuf:"varint,7,opt,name=duration,proto3" json:"duration,omitempty"`
	// The host that recorded this span, primarily for query by service name.
	//
	// Instrumentation should always record this. Usually, absent implies late
	// data. The IP address corresponding to this is usually the site local or
	// advertised service address. When present, the port indicates the listen
	// port.
	LocalEndpoint *Endpoint `protobuf:"bytes,8,opt,name=local_endpoint,json=localEndpoint,proto3" json:"local_endpoint,omitempty"`
	// When an RPC (or messaging) span, indicates the other side of the
	// connection.
	//
	// By recording the remote endpoint, your trace will contain network context
	// even if the peer is not tracing. For example, you can record the IP from
	// the "X-Forwarded-For" header or the service name and socket of a remote
	// peer.
	RemoteEndpoint *Endpoint `protobuf:"bytes,9,opt,name=remote_endpoint,json=remoteEndpoint,proto3" json:"remote_endpoint,omitempty"`
	// Associates events that explain latency with the time they happened.
	Annotations []*Annotation `protobuf:"bytes,10,rep,name=annotations,proto3" json:"annotations,omitempty"`
	// Tags give your span context for search, viewing and analysis.
	//
	// For example, a key "your_app.version" would let you lookup traces by
	// version. A tag "sql.query" isn't searchable, but it can help in debugging
	// when viewing a trace.
	Tags map[string]string `protobuf:"bytes,11,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// True is a request to store this span even if it overrides sampling policy.
	//
	// This is true when the "X-B3-Flags" header has a value of 1.
	Debug bool `protobuf:"varint,12,opt,name=debug,proto3" json:"debug,omitempty"`
	// True if we are contributing to a span started by another tracer (ex on a
	// different host).
	Shared bool `protobuf:"varint,13,opt,name=shared,proto3" json:"shared,omitempty"`
}

func (x *Span) Reset() {
	*x = Span{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Span) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Span) ProtoMessage() {}

func (x *Span) ProtoReflect() protoreflect.Message {
	mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Span.ProtoReflect.Descriptor instead.
func (*Span) Descriptor() ([]byte, []int) {
	return file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP(), []int{0}
}

func (x *Span) GetTraceId() []byte {
	if x != nil {
		return x.TraceId
	}
	return nil
}

func (x *Span) GetParentId() []byte {
	if x != nil {
		return x.ParentId
	}
	return nil
}

func (x *Span) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *Span) GetKind() Span_Kind {
	if x != nil {
		return x.Kind
	}
	return Span_SPAN_KIND_UNSPECIFIED
}

func (x *Span) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Span) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Span) GetDuration() uint64 {
	if x != nil {
		return x.Duration
	}
	return 0
}

func (x *Span) GetLocalEndpoint() *Endpoint {
	if x != nil {
		return x.LocalEndpoint
	}
	return nil
}

func (x *Span) GetRemoteEndpoint() *Endpoint {
	if x != nil {
		return x.RemoteEndpoint
	}
	return nil
}

func (x *Span) GetAnnotations() []*Annotation {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Span) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Span) GetDebug() bool {
	if x != nil {
		return x.Debug
	}
	return false
}

func (x *Span) GetShared() bool {
	if x != nil {
		return x.Shared
	}
	return false
}

// The network context of a node in the service graph.
//
// The next id is 5.
type Endpoint struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Lower-case label of this node in the service graph, such as "favstar".
	// Leave absent if unknown.
	//
	// This is a primary label for trace lookup and aggregation, so it should be
	// intuitive and consistent. Many use a name from service discovery.
	ServiceName string `protobuf:"bytes,1,opt,name=service_name,json=serviceName,proto3" json:"service_name,omitempty"`
	// 4 byte representation of the primary IPv4 address associated with this
	// connection. Absent if unknown.
	Ipv4 []byte `protobuf:"bytes,2,opt,name=ipv4,proto3" json:"ipv4,omitempty"`
	// 16 byte representation of the primary IPv6 address associated with this
	// connection. Absent if unknown.
	//
	// Prefer using the ipv4 field for mapped addresses.
	Ipv6 []byte `protobuf:"bytes,3,opt,name=ipv6,proto3" json:"ipv6,omitempty"`
	// Depending on context, this could be a listen port or the client-side of a
	// socket. Absent if unknown.
	Port int32 `protobuf:"varint,4,opt,name=port,proto3" json:"port,omitempty"`
}

func (x *Endpoint) Reset() {
	*x = Endpoint{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Endpoint) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Endpoint) ProtoMessage() {}

func (x *Endpoint) ProtoReflect() protoreflect.Message {
	mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Endpoint.ProtoReflect.Descriptor instead.
func (*Endpoint) Descriptor() ([]byte, []int) {
	return file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP(), []int{1}
}

func (x *Endpoint) GetServiceName() string {
	if x != nil {
		return x.ServiceName
	}
	return ""
}

func (x *Endpoint) GetIpv4() []byte {
	if x != nil {
		return x.Ipv4
	}
	return nil
}

func (x *Endpoint) GetIpv6() []byte {
	if x != nil {
		return x.Ipv6
	}
	return nil
}

func (x *Endpoint) GetPort() int32 {
	if x != nil {
		return x.Port
	}
	return 0
}

// Associates an event that explains latency with a timestamp.
// Unlike log statements, annotations are often codes. Ex. "ws" for WireSend
//
// The next id is 3.
type Annotation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Epoch microseconds of this event.
	//
	// For example, 1502787600000000 corresponds to 2017-08-15 09:00 UTC
	//
	// This value should be set directly by instrumentation, using the most
	// precise value possible. For example, gettimeofday or multiplying epoch
	// millis by 1000.
	Timestamp uint64 `protobuf:"fixed64,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Usually a short tag indicating an event, like "error"
	//
	// While possible to add larger data, such as garbage collection details, low
	// cardinality event names both keep the size of spans down and also are easy
	// to search against.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Annotation) Reset() {
	*x = Annotation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Annotation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Annotation) ProtoMessage() {}

func (x *Annotation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Annotation.ProtoReflect.Descriptor instead.
func (*Annotation) Descriptor() ([]byte, []int) {
	return file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP(), []int{2}
}

func (x *Annotation) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Annotation) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// A list of spans with possibly different trace ids, in no particular order.
//
// This is used for all transports: POST, Kafka messages etc. No other fields
// are expected, This message facilitates the mechanics of encoding a list, as
// a field number is required. The name of this type is the same in the OpenApi
// aka Swagger specification. https://zipkin.io/zipkin-api/#/default/post_spans
type ListOfSpans struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spans []*Span `protobuf:"bytes,1,rep,name=spans,proto3" json:"spans,omitempty"`
}

func (x *ListOfSpans) Reset() {
	*x = ListOfSpans{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListOfSpans) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListOfSpans) ProtoMessage() {}

func (x *ListOfSpans) ProtoReflect() protoreflect.Message {
	mi := &file_proto_zipkin_proto3_zipkin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListOfSpans.ProtoReflect.Descriptor instead.
func (*ListOfSpans) Descriptor() ([]byte, []int) {
	return file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP(), []int{3}
}

func (x *ListOfSpans) GetSpans() []*Span {
	if x != nil {
		return x.Spans
	}
	return nil
}

var File_proto_zipkin_proto3_zipkin_proto protoreflect.FileDescriptor

var file_proto_zipkin_proto3_zipkin_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0d, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33, 0x22, 0xfa, 0x04, 0x0a, 0x04, 0x53, 0x70, 0x61, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x07, 0x74, 0x72,
	0x61, 0x63, 0x65, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x2c, 0x0a, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x18, 0x2e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
	0x2e, 0x53, 0x70, 0x61, 0x6e, 0x2e, 0x4b, 0x69, 0x6e, 0x64, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x06, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x3e,
	0x0a, 0x0e, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x52,
	0x0d, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x40,
	0x0a, 0x0f, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x5f, 0x65, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x52, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x12, 0x3b, 0x0a, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x0b, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x31, 0x0a,
	0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x7a, 0x69,
	0x70, 0x6b, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x6e,
	0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73,
	0x12, 0x14, 0x0a, 0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x05, 0x64, 0x65, 0x62, 0x75, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64,
	0x18, 0x0d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x68, 0x61, 0x72, 0x65, 0x64, 0x1a, 0x37,
	0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x55, 0x0a, 0x04, 0x4b, 0x69, 0x6e, 0x64, 0x12,
	0x19, 0x0a, 0x15, 0x53, 0x50, 0x41, 0x4e, 0x5f, 0x4b, 0x49, 0x4e, 0x44, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x43, 0x4c,
	0x49, 0x45, 0x4e, 0x54, 0x10, 0x01, 0x12, 0x0a, 0x0a, 0x06, 0x53, 0x45, 0x52, 0x56, 0x45, 0x52,
	0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x52, 0x4f, 0x44, 0x55, 0x43, 0x45, 0x52, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x53, 0x55, 0x4d, 0x45, 0x52, 0x10, 0x04, 0x22, 0x69,
	0x0a, 0x08, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a,
	0x04, 0x69, 0x70, 0x76, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x69, 0x70, 0x76,
	0x34, 0x12, 0x12, 0x0a, 0x04, 0x69, 0x70, 0x76, 0x36, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x69, 0x70, 0x76, 0x36, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x04, 0x70, 0x6f, 0x72, 0x74, 0x22, 0x40, 0x0a, 0x0a, 0x41, 0x6e, 0x6e,
	0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x38, 0x0a, 0x0b, 0x4c,
	0x69, 0x73, 0x74, 0x4f, 0x66, 0x53, 0x70, 0x61, 0x6e, 0x73, 0x12, 0x29, 0x0a, 0x05, 0x73, 0x70,
	0x61, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x7a, 0x69, 0x70, 0x6b,
	0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x2e, 0x53, 0x70, 0x61, 0x6e, 0x52, 0x05,
	0x73, 0x70, 0x61, 0x6e, 0x73, 0x42, 0x47, 0x0a, 0x0e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x32,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x50, 0x01, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6f, 0x70, 0x65, 0x6e, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e,
	0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x7a, 0x69, 0x70, 0x6b, 0x69, 0x6e, 0x5f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_zipkin_proto3_zipkin_proto_rawDescOnce sync.Once
	file_proto_zipkin_proto3_zipkin_proto_rawDescData = file_proto_zipkin_proto3_zipkin_proto_rawDesc
)

func file_proto_zipkin_proto3_zipkin_proto_rawDescGZIP() []byte {
	file_proto_zipkin_proto3_zipkin_proto_rawDescOnce.Do(func() {
		file_proto_zipkin_proto3_zipkin_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_zipkin_proto3_zipkin_proto_rawDescData)
	})
	return file_proto_zipkin_proto3_zipkin_proto_rawDescData
}

var file_proto_zipkin_proto3_zipkin_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_zipkin_proto3_zipkin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_zipkin_proto3_zipkin_proto_goTypes = []interface{}{
	(Span_Kind)(0),      // 0: zipkin.proto3.Span.Kind
	(*Span)(nil),        // 1: zipkin.proto3.Span
	(*Endpoint)(nil),    // 2: zipkin.proto3.Endpoint
	(*Annotation)(nil),  // 3: zipkin.proto3.Annotation
	(*ListOfSpans)(nil), // 4: zipkin.proto3.ListOfSpans
	nil,                 // 5: zipkin.proto3.Span.TagsEntry
}
var file_proto_zipkin_proto3_zipkin_proto_depIdxs = []int32{
	0, // 0: zipkin.proto3.Span.kind:type_name -> zipkin.proto3.Span.Kind
	2, // 1: zipkin.proto3.Span.local_endpoint:type_name -> zipkin.proto3.Endpoint
	2, // 2: zipkin.proto3.Span.remote_endpoint:type_name -> zipkin.proto3.Endpoint
	3, // 3: zipkin.proto3.Span.annotations:type_name -> zipkin.proto3.Annotation
	5, // 4: zipkin.proto3.Span.tags:type_name -> zipkin.proto3.Span.TagsEntry
	1, // 5: zipkin.proto3.ListOfSpans.spans:type_name -> zipkin.proto3.Span
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_proto_zipkin_proto3_zipkin_proto_init() }
func file_proto_zipkin_proto3_zipkin_proto_init() {
	if File_proto_zipkin_proto3_zipkin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_zipkin_proto3_zipkin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Span); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_zipkin_proto3_zipkin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Endpoint); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_zipkin_proto3_zipkin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Annotation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_proto_zipkin_proto3_zipkin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListOfSpans); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_zipkin_proto3_zipkin_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_zipkin_proto3_zipkin_proto_goTypes,
		DependencyIndexes: file_proto_zipkin_proto3_zipkin_proto_depIdxs,
		EnumInfos:         file_proto_zipkin_proto3_zipkin_proto_enumTypes,
		MessageInfos:      file_proto_zipkin_proto3_zipkin_proto_msgTypes,
	}.Build()
	File_proto_zipkin_proto3_zipkin_proto = out.File
	file_proto_zipkin_proto3_zipkin_proto_rawDesc = nil
	file_proto_zipkin_proto3_zipkin_proto_goTypes = nil
	file_proto_zipkin_proto3_zipkin_proto_depIdxs = nil
}
