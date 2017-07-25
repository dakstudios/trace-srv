// Code generated by protoc-gen-go. DO NOT EDIT.
// source: github.com/dakstudios/trace-srv/proto/trace/trace.proto

/*
Package org_dakstudios_srv_trace_trace is a generated protocol buffer package.

It is generated from these files:
	github.com/dakstudios/trace-srv/proto/trace/trace.proto

It has these top-level messages:
	CreateRequest
	CreateResponse
	ReadRequest
	ReadResponse
	DeleteRequest
	DeleteResponse
	SearchRequest
	SearchResponse
	LiveRequest
	LiveResponse
*/
package org_dakstudios_srv_trace_trace

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import go_micro_os_trace "github.com/micro/go-os/trace/proto"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CreateRequest struct {
	Span *go_micro_os_trace.Span `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
}

func (m *CreateRequest) Reset()                    { *m = CreateRequest{} }
func (m *CreateRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateRequest) ProtoMessage()               {}
func (*CreateRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CreateRequest) GetSpan() *go_micro_os_trace.Span {
	if m != nil {
		return m.Span
	}
	return nil
}

type CreateResponse struct {
}

func (m *CreateResponse) Reset()                    { *m = CreateResponse{} }
func (m *CreateResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateResponse) ProtoMessage()               {}
func (*CreateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type ReadRequest struct {
	// trace id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *ReadRequest) Reset()                    { *m = ReadRequest{} }
func (m *ReadRequest) String() string            { return proto.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()               {}
func (*ReadRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *ReadRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type ReadResponse struct {
	Spans []*go_micro_os_trace.Span `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
}

func (m *ReadResponse) Reset()                    { *m = ReadResponse{} }
func (m *ReadResponse) String() string            { return proto.CompactTextString(m) }
func (*ReadResponse) ProtoMessage()               {}
func (*ReadResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *ReadResponse) GetSpans() []*go_micro_os_trace.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type DeleteRequest struct {
	// trace id
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *DeleteRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type DeleteResponse struct {
}

func (m *DeleteResponse) Reset()                    { *m = DeleteResponse{} }
func (m *DeleteResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteResponse) ProtoMessage()               {}
func (*DeleteResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type SearchRequest struct {
	Name    string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Limit   int64  `protobuf:"varint,2,opt,name=limit" json:"limit,omitempty"`
	Offset  int64  `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Reverse bool   `protobuf:"varint,4,opt,name=reverse" json:"reverse,omitempty"`
}

func (m *SearchRequest) Reset()                    { *m = SearchRequest{} }
func (m *SearchRequest) String() string            { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()               {}
func (*SearchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *SearchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *SearchRequest) GetLimit() int64 {
	if m != nil {
		return m.Limit
	}
	return 0
}

func (m *SearchRequest) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *SearchRequest) GetReverse() bool {
	if m != nil {
		return m.Reverse
	}
	return false
}

type SearchResponse struct {
	Spans []*go_micro_os_trace.Span `protobuf:"bytes,1,rep,name=spans" json:"spans,omitempty"`
}

func (m *SearchResponse) Reset()                    { *m = SearchResponse{} }
func (m *SearchResponse) String() string            { return proto.CompactTextString(m) }
func (*SearchResponse) ProtoMessage()               {}
func (*SearchResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *SearchResponse) GetSpans() []*go_micro_os_trace.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type LiveRequest struct {
	// could be empty
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *LiveRequest) Reset()                    { *m = LiveRequest{} }
func (m *LiveRequest) String() string            { return proto.CompactTextString(m) }
func (*LiveRequest) ProtoMessage()               {}
func (*LiveRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *LiveRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type LiveResponse struct {
	Span *go_micro_os_trace.Span `protobuf:"bytes,1,opt,name=span" json:"span,omitempty"`
}

func (m *LiveResponse) Reset()                    { *m = LiveResponse{} }
func (m *LiveResponse) String() string            { return proto.CompactTextString(m) }
func (*LiveResponse) ProtoMessage()               {}
func (*LiveResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *LiveResponse) GetSpan() *go_micro_os_trace.Span {
	if m != nil {
		return m.Span
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateRequest)(nil), "org.dakstudios.srv.trace.trace.CreateRequest")
	proto.RegisterType((*CreateResponse)(nil), "org.dakstudios.srv.trace.trace.CreateResponse")
	proto.RegisterType((*ReadRequest)(nil), "org.dakstudios.srv.trace.trace.ReadRequest")
	proto.RegisterType((*ReadResponse)(nil), "org.dakstudios.srv.trace.trace.ReadResponse")
	proto.RegisterType((*DeleteRequest)(nil), "org.dakstudios.srv.trace.trace.DeleteRequest")
	proto.RegisterType((*DeleteResponse)(nil), "org.dakstudios.srv.trace.trace.DeleteResponse")
	proto.RegisterType((*SearchRequest)(nil), "org.dakstudios.srv.trace.trace.SearchRequest")
	proto.RegisterType((*SearchResponse)(nil), "org.dakstudios.srv.trace.trace.SearchResponse")
	proto.RegisterType((*LiveRequest)(nil), "org.dakstudios.srv.trace.trace.LiveRequest")
	proto.RegisterType((*LiveResponse)(nil), "org.dakstudios.srv.trace.trace.LiveResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Trace service

type TraceClient interface {
	Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error)
	Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error)
	Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error)
	Live(ctx context.Context, in *LiveRequest, opts ...client.CallOption) (Trace_LiveClient, error)
}

type traceClient struct {
	c           client.Client
	serviceName string
}

func NewTraceClient(serviceName string, c client.Client) TraceClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "org.dakstudios.srv.trace.trace"
	}
	return &traceClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *traceClient) Create(ctx context.Context, in *CreateRequest, opts ...client.CallOption) (*CreateResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Trace.Create", in)
	out := new(CreateResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceClient) Read(ctx context.Context, in *ReadRequest, opts ...client.CallOption) (*ReadResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Trace.Read", in)
	out := new(ReadResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceClient) Delete(ctx context.Context, in *DeleteRequest, opts ...client.CallOption) (*DeleteResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Trace.Delete", in)
	out := new(DeleteResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceClient) Search(ctx context.Context, in *SearchRequest, opts ...client.CallOption) (*SearchResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Trace.Search", in)
	out := new(SearchResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *traceClient) Live(ctx context.Context, in *LiveRequest, opts ...client.CallOption) (Trace_LiveClient, error) {
	req := c.c.NewRequest(c.serviceName, "Trace.Live", &LiveRequest{})
	stream, err := c.c.Stream(ctx, req, opts...)
	if err != nil {
		return nil, err
	}
	if err := stream.Send(in); err != nil {
		return nil, err
	}
	return &traceLiveClient{stream}, nil
}

type Trace_LiveClient interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Recv() (*LiveResponse, error)
}

type traceLiveClient struct {
	stream client.Streamer
}

func (x *traceLiveClient) Close() error {
	return x.stream.Close()
}

func (x *traceLiveClient) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceLiveClient) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceLiveClient) Recv() (*LiveResponse, error) {
	m := new(LiveResponse)
	err := x.stream.Recv(m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for Trace service

type TraceHandler interface {
	Create(context.Context, *CreateRequest, *CreateResponse) error
	Read(context.Context, *ReadRequest, *ReadResponse) error
	Delete(context.Context, *DeleteRequest, *DeleteResponse) error
	Search(context.Context, *SearchRequest, *SearchResponse) error
	Live(context.Context, *LiveRequest, Trace_LiveStream) error
}

func RegisterTraceHandler(s server.Server, hdlr TraceHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Trace{hdlr}, opts...))
}

type Trace struct {
	TraceHandler
}

func (h *Trace) Create(ctx context.Context, in *CreateRequest, out *CreateResponse) error {
	return h.TraceHandler.Create(ctx, in, out)
}

func (h *Trace) Read(ctx context.Context, in *ReadRequest, out *ReadResponse) error {
	return h.TraceHandler.Read(ctx, in, out)
}

func (h *Trace) Delete(ctx context.Context, in *DeleteRequest, out *DeleteResponse) error {
	return h.TraceHandler.Delete(ctx, in, out)
}

func (h *Trace) Search(ctx context.Context, in *SearchRequest, out *SearchResponse) error {
	return h.TraceHandler.Search(ctx, in, out)
}

func (h *Trace) Live(ctx context.Context, stream server.Streamer) error {
	m := new(LiveRequest)
	if err := stream.Recv(m); err != nil {
		return err
	}
	return h.TraceHandler.Live(ctx, m, &traceLiveStream{stream})
}

type Trace_LiveStream interface {
	SendMsg(interface{}) error
	RecvMsg(interface{}) error
	Close() error
	Send(*LiveResponse) error
}

type traceLiveStream struct {
	stream server.Streamer
}

func (x *traceLiveStream) Close() error {
	return x.stream.Close()
}

func (x *traceLiveStream) SendMsg(m interface{}) error {
	return x.stream.Send(m)
}

func (x *traceLiveStream) RecvMsg(m interface{}) error {
	return x.stream.Recv(m)
}

func (x *traceLiveStream) Send(m *LiveResponse) error {
	return x.stream.Send(m)
}

func init() {
	proto.RegisterFile("github.com/dakstudios/trace-srv/proto/trace/trace.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 417 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd4, 0x30,
	0x10, 0xdd, 0xec, 0x66, 0x03, 0xcc, 0x76, 0xa3, 0xca, 0x42, 0x10, 0xad, 0x04, 0x04, 0x9f, 0x56,
	0x2a, 0x71, 0x50, 0x39, 0x70, 0x00, 0xc4, 0x01, 0x8e, 0x9c, 0x52, 0x7e, 0xc0, 0x4d, 0xa6, 0xa9,
	0xd5, 0x26, 0x0e, 0xb6, 0x37, 0x5f, 0xc2, 0x07, 0xa3, 0xd8, 0x09, 0x4d, 0x7a, 0xd8, 0x2c, 0x5c,
	0xac, 0x8c, 0xfd, 0xe6, 0x3d, 0x8f, 0xdf, 0x53, 0xe0, 0x63, 0x29, 0xcc, 0xed, 0xe1, 0x9a, 0xe5,
	0xb2, 0x4a, 0x0b, 0x7e, 0xa7, 0xcd, 0xa1, 0x10, 0x52, 0xa7, 0x46, 0xf1, 0x1c, 0x13, 0xad, 0xda,
	0xb4, 0x51, 0xd2, 0x48, 0x57, 0xbb, 0x95, 0xd9, 0x1d, 0xf2, 0x5a, 0xaa, 0x92, 0x3d, 0x74, 0x30,
	0xad, 0x5a, 0xe6, 0xce, 0xed, 0xba, 0x63, 0x23, 0xe2, 0x4a, 0xe4, 0x4a, 0xa6, 0xa5, 0x4c, 0x06,
	0xe6, 0x31, 0xab, 0xe3, 0xa3, 0x9f, 0x61, 0xfb, 0x4d, 0x21, 0x37, 0x98, 0xe1, 0xaf, 0x03, 0x6a,
	0x43, 0x2e, 0xc0, 0xd7, 0x0d, 0xaf, 0x23, 0x2f, 0xf6, 0xf6, 0x9b, 0xcb, 0x97, 0xac, 0x94, 0xcc,
	0xf2, 0x30, 0xa9, 0x7b, 0xa1, 0xab, 0x86, 0xd7, 0x99, 0x05, 0xd1, 0x73, 0x08, 0x87, 0x6e, 0xdd,
	0xc8, 0x5a, 0x23, 0x7d, 0x05, 0x9b, 0x0c, 0x79, 0x31, 0xb0, 0x85, 0xb0, 0x14, 0x85, 0xe5, 0x7a,
	0x96, 0x2d, 0x45, 0x41, 0xbf, 0xc0, 0x99, 0x3b, 0x76, 0x70, 0x92, 0xc0, 0xba, 0x23, 0xd2, 0x91,
	0x17, 0xaf, 0x8e, 0xc9, 0x39, 0x14, 0x7d, 0x03, 0xdb, 0xef, 0x78, 0x8f, 0x0f, 0xb7, 0x7d, 0xcc,
	0x7f, 0x0e, 0xe1, 0x00, 0xe8, 0x2f, 0x74, 0x07, 0xdb, 0x2b, 0xe4, 0x2a, 0xbf, 0x1d, 0x5a, 0x08,
	0xf8, 0x35, 0xaf, 0xb0, 0x6f, 0xb2, 0xdf, 0xe4, 0x39, 0xac, 0xef, 0x45, 0x25, 0x4c, 0xb4, 0x8c,
	0xbd, 0xfd, 0x2a, 0x73, 0x05, 0x79, 0x01, 0x81, 0xbc, 0xb9, 0xd1, 0x68, 0xa2, 0x95, 0xdd, 0xee,
	0x2b, 0x12, 0xc1, 0x13, 0x85, 0x2d, 0x2a, 0x8d, 0x91, 0x1f, 0x7b, 0xfb, 0xa7, 0xd9, 0x50, 0xd2,
	0xaf, 0x10, 0x0e, 0x62, 0xff, 0x37, 0xe0, 0x5b, 0xd8, 0xfc, 0x10, 0x2d, 0x1e, 0xb9, 0x2b, 0xfd,
	0x04, 0x67, 0x0e, 0xd2, 0x2b, 0xfc, 0x8b, 0x61, 0x97, 0xbf, 0x7d, 0x58, 0xff, 0xec, 0x36, 0x89,
	0x80, 0xc0, 0x59, 0x47, 0x12, 0x76, 0x3c, 0x53, 0x6c, 0x12, 0x90, 0x1d, 0x3b, 0x15, 0xde, 0x1b,
	0xb0, 0x20, 0x39, 0xf8, 0x9d, 0xe9, 0xe4, 0x62, 0xae, 0x73, 0x94, 0x9c, 0xdd, 0xbb, 0xd3, 0xc0,
	0x7f, 0x45, 0x04, 0x04, 0xce, 0xf9, 0xf9, 0x79, 0x26, 0x11, 0x9a, 0x9f, 0xe7, 0x51, 0xa0, 0xac,
	0x94, 0x73, 0x79, 0x5e, 0x6a, 0x12, 0xbd, 0x79, 0xa9, 0x69, 0x78, 0xe8, 0x82, 0x20, 0xf8, 0x9d,
	0xd9, 0xf3, 0x4f, 0x37, 0x4a, 0xcd, 0xfc, 0xd3, 0x8d, 0xf3, 0x43, 0x17, 0xef, 0xbd, 0xeb, 0xc0,
	0xfe, 0x0c, 0x3e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x57, 0xf5, 0x94, 0x91, 0x97, 0x04, 0x00,
	0x00,
}