// Code generated by protoc-gen-go. DO NOT EDIT.
// source: greet/greet.proto

package greet

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

type Greeting struct {
	FirstName            string   `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName             string   `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Greeting) Reset()         { *m = Greeting{} }
func (m *Greeting) String() string { return proto.CompactTextString(m) }
func (*Greeting) ProtoMessage()    {}
func (*Greeting) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{0}
}

func (m *Greeting) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Greeting.Unmarshal(m, b)
}
func (m *Greeting) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Greeting.Marshal(b, m, deterministic)
}
func (m *Greeting) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Greeting.Merge(m, src)
}
func (m *Greeting) XXX_Size() int {
	return xxx_messageInfo_Greeting.Size(m)
}
func (m *Greeting) XXX_DiscardUnknown() {
	xxx_messageInfo_Greeting.DiscardUnknown(m)
}

var xxx_messageInfo_Greeting proto.InternalMessageInfo

func (m *Greeting) GetFirstName() string {
	if m != nil {
		return m.FirstName
	}
	return ""
}

func (m *Greeting) GetLastName() string {
	if m != nil {
		return m.LastName
	}
	return ""
}

type GreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetRequest) Reset()         { *m = GreetRequest{} }
func (m *GreetRequest) String() string { return proto.CompactTextString(m) }
func (*GreetRequest) ProtoMessage()    {}
func (*GreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{1}
}

func (m *GreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetRequest.Unmarshal(m, b)
}
func (m *GreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetRequest.Marshal(b, m, deterministic)
}
func (m *GreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetRequest.Merge(m, src)
}
func (m *GreetRequest) XXX_Size() int {
	return xxx_messageInfo_GreetRequest.Size(m)
}
func (m *GreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetRequest proto.InternalMessageInfo

func (m *GreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetResponse) Reset()         { *m = GreetResponse{} }
func (m *GreetResponse) String() string { return proto.CompactTextString(m) }
func (*GreetResponse) ProtoMessage()    {}
func (*GreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{2}
}

func (m *GreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetResponse.Unmarshal(m, b)
}
func (m *GreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetResponse.Marshal(b, m, deterministic)
}
func (m *GreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetResponse.Merge(m, src)
}
func (m *GreetResponse) XXX_Size() int {
	return xxx_messageInfo_GreetResponse.Size(m)
}
func (m *GreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetResponse proto.InternalMessageInfo

func (m *GreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetManyTimesRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetManyTimesRequest) Reset()         { *m = GreetManyTimesRequest{} }
func (m *GreetManyTimesRequest) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesRequest) ProtoMessage()    {}
func (*GreetManyTimesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{3}
}

func (m *GreetManyTimesRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesRequest.Unmarshal(m, b)
}
func (m *GreetManyTimesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesRequest.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesRequest.Merge(m, src)
}
func (m *GreetManyTimesRequest) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesRequest.Size(m)
}
func (m *GreetManyTimesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesRequest proto.InternalMessageInfo

func (m *GreetManyTimesRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetManyTimesResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetManyTimesResponse) Reset()         { *m = GreetManyTimesResponse{} }
func (m *GreetManyTimesResponse) String() string { return proto.CompactTextString(m) }
func (*GreetManyTimesResponse) ProtoMessage()    {}
func (*GreetManyTimesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{4}
}

func (m *GreetManyTimesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetManyTimesResponse.Unmarshal(m, b)
}
func (m *GreetManyTimesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetManyTimesResponse.Marshal(b, m, deterministic)
}
func (m *GreetManyTimesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetManyTimesResponse.Merge(m, src)
}
func (m *GreetManyTimesResponse) XXX_Size() int {
	return xxx_messageInfo_GreetManyTimesResponse.Size(m)
}
func (m *GreetManyTimesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetManyTimesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetManyTimesResponse proto.InternalMessageInfo

func (m *GreetManyTimesResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type LongGreetRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *LongGreetRequest) Reset()         { *m = LongGreetRequest{} }
func (m *LongGreetRequest) String() string { return proto.CompactTextString(m) }
func (*LongGreetRequest) ProtoMessage()    {}
func (*LongGreetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{5}
}

func (m *LongGreetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetRequest.Unmarshal(m, b)
}
func (m *LongGreetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetRequest.Marshal(b, m, deterministic)
}
func (m *LongGreetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetRequest.Merge(m, src)
}
func (m *LongGreetRequest) XXX_Size() int {
	return xxx_messageInfo_LongGreetRequest.Size(m)
}
func (m *LongGreetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetRequest proto.InternalMessageInfo

func (m *LongGreetRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type LongGreetResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LongGreetResponse) Reset()         { *m = LongGreetResponse{} }
func (m *LongGreetResponse) String() string { return proto.CompactTextString(m) }
func (*LongGreetResponse) ProtoMessage()    {}
func (*LongGreetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{6}
}

func (m *LongGreetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LongGreetResponse.Unmarshal(m, b)
}
func (m *LongGreetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LongGreetResponse.Marshal(b, m, deterministic)
}
func (m *LongGreetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LongGreetResponse.Merge(m, src)
}
func (m *LongGreetResponse) XXX_Size() int {
	return xxx_messageInfo_LongGreetResponse.Size(m)
}
func (m *LongGreetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LongGreetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LongGreetResponse proto.InternalMessageInfo

func (m *LongGreetResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetEveryoneRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetEveryoneRequest) Reset()         { *m = GreetEveryoneRequest{} }
func (m *GreetEveryoneRequest) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneRequest) ProtoMessage()    {}
func (*GreetEveryoneRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{7}
}

func (m *GreetEveryoneRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneRequest.Unmarshal(m, b)
}
func (m *GreetEveryoneRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneRequest.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneRequest.Merge(m, src)
}
func (m *GreetEveryoneRequest) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneRequest.Size(m)
}
func (m *GreetEveryoneRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneRequest proto.InternalMessageInfo

func (m *GreetEveryoneRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetEveryoneResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetEveryoneResponse) Reset()         { *m = GreetEveryoneResponse{} }
func (m *GreetEveryoneResponse) String() string { return proto.CompactTextString(m) }
func (*GreetEveryoneResponse) ProtoMessage()    {}
func (*GreetEveryoneResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{8}
}

func (m *GreetEveryoneResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetEveryoneResponse.Unmarshal(m, b)
}
func (m *GreetEveryoneResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetEveryoneResponse.Marshal(b, m, deterministic)
}
func (m *GreetEveryoneResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetEveryoneResponse.Merge(m, src)
}
func (m *GreetEveryoneResponse) XXX_Size() int {
	return xxx_messageInfo_GreetEveryoneResponse.Size(m)
}
func (m *GreetEveryoneResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetEveryoneResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetEveryoneResponse proto.InternalMessageInfo

func (m *GreetEveryoneResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type GreetWithDeadlineRequest struct {
	Greeting             *Greeting `protobuf:"bytes,1,opt,name=greeting,proto3" json:"greeting,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *GreetWithDeadlineRequest) Reset()         { *m = GreetWithDeadlineRequest{} }
func (m *GreetWithDeadlineRequest) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineRequest) ProtoMessage()    {}
func (*GreetWithDeadlineRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{9}
}

func (m *GreetWithDeadlineRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineRequest.Unmarshal(m, b)
}
func (m *GreetWithDeadlineRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineRequest.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineRequest.Merge(m, src)
}
func (m *GreetWithDeadlineRequest) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineRequest.Size(m)
}
func (m *GreetWithDeadlineRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineRequest proto.InternalMessageInfo

func (m *GreetWithDeadlineRequest) GetGreeting() *Greeting {
	if m != nil {
		return m.Greeting
	}
	return nil
}

type GreetWithDeadlineResponse struct {
	Result               string   `protobuf:"bytes,1,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GreetWithDeadlineResponse) Reset()         { *m = GreetWithDeadlineResponse{} }
func (m *GreetWithDeadlineResponse) String() string { return proto.CompactTextString(m) }
func (*GreetWithDeadlineResponse) ProtoMessage()    {}
func (*GreetWithDeadlineResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_285dce36ca2d93c1, []int{10}
}

func (m *GreetWithDeadlineResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GreetWithDeadlineResponse.Unmarshal(m, b)
}
func (m *GreetWithDeadlineResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GreetWithDeadlineResponse.Marshal(b, m, deterministic)
}
func (m *GreetWithDeadlineResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GreetWithDeadlineResponse.Merge(m, src)
}
func (m *GreetWithDeadlineResponse) XXX_Size() int {
	return xxx_messageInfo_GreetWithDeadlineResponse.Size(m)
}
func (m *GreetWithDeadlineResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GreetWithDeadlineResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GreetWithDeadlineResponse proto.InternalMessageInfo

func (m *GreetWithDeadlineResponse) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

func init() {
	proto.RegisterType((*Greeting)(nil), "greet.Greeting")
	proto.RegisterType((*GreetRequest)(nil), "greet.GreetRequest")
	proto.RegisterType((*GreetResponse)(nil), "greet.GreetResponse")
	proto.RegisterType((*GreetManyTimesRequest)(nil), "greet.GreetManyTimesRequest")
	proto.RegisterType((*GreetManyTimesResponse)(nil), "greet.GreetManyTimesResponse")
	proto.RegisterType((*LongGreetRequest)(nil), "greet.LongGreetRequest")
	proto.RegisterType((*LongGreetResponse)(nil), "greet.LongGreetResponse")
	proto.RegisterType((*GreetEveryoneRequest)(nil), "greet.GreetEveryoneRequest")
	proto.RegisterType((*GreetEveryoneResponse)(nil), "greet.GreetEveryoneResponse")
	proto.RegisterType((*GreetWithDeadlineRequest)(nil), "greet.GreetWithDeadlineRequest")
	proto.RegisterType((*GreetWithDeadlineResponse)(nil), "greet.GreetWithDeadlineResponse")
}

func init() { proto.RegisterFile("greet/greet.proto", fileDescriptor_285dce36ca2d93c1) }

var fileDescriptor_285dce36ca2d93c1 = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xd1, 0x4b, 0xc3, 0x30,
	0x10, 0x87, 0xad, 0xb0, 0xd1, 0x9e, 0xca, 0x5c, 0x36, 0xc7, 0x54, 0x04, 0xe9, 0x8b, 0xc2, 0x60,
	0x93, 0xcd, 0x37, 0x1f, 0x04, 0xad, 0xee, 0x45, 0x45, 0xaa, 0xe0, 0xa3, 0xc4, 0x79, 0xd6, 0x40,
	0x9b, 0xcc, 0x24, 0x1b, 0xec, 0xbf, 0x97, 0x5e, 0x33, 0x9d, 0x22, 0x16, 0xfa, 0x52, 0x7a, 0x77,
	0xbf, 0x7c, 0x5f, 0x5a, 0x0e, 0x9a, 0x89, 0x46, 0xb4, 0x03, 0x7a, 0xf6, 0xa7, 0x5a, 0x59, 0xc5,
	0x6a, 0x54, 0x84, 0xd7, 0xe0, 0x8f, 0xf3, 0x17, 0x21, 0x13, 0x76, 0x00, 0xf0, 0x26, 0xb4, 0xb1,
	0xcf, 0x92, 0x67, 0xd8, 0xf5, 0x0e, 0xbd, 0xe3, 0x20, 0x0e, 0xa8, 0x73, 0xc7, 0x33, 0x64, 0xfb,
	0x10, 0xa4, 0x7c, 0x39, 0x5d, 0xa7, 0xa9, 0x9f, 0x37, 0xf2, 0x61, 0x78, 0x06, 0x9b, 0xc4, 0x89,
	0xf1, 0x63, 0x86, 0xc6, 0xb2, 0x1e, 0xf8, 0x89, 0xe3, 0x12, 0x69, 0x63, 0xd8, 0xe8, 0x17, 0xfa,
	0xa5, 0x2e, 0xfe, 0x0a, 0x84, 0x47, 0xb0, 0xe5, 0x0e, 0x9b, 0xa9, 0x92, 0x06, 0x59, 0x07, 0xea,
	0x1a, 0xcd, 0x2c, 0xb5, 0xee, 0x16, 0xae, 0x0a, 0x23, 0xd8, 0xa1, 0xe0, 0x2d, 0x97, 0x8b, 0x47,
	0x91, 0xa1, 0xa9, 0xa4, 0x3b, 0x81, 0xce, 0x6f, 0x4a, 0x89, 0xf7, 0x1c, 0xb6, 0x6f, 0x94, 0x4c,
	0xaa, 0x7f, 0x61, 0x0f, 0x9a, 0x2b, 0x80, 0x12, 0xdb, 0x25, 0xb4, 0x29, 0x78, 0x35, 0x47, 0xbd,
	0x50, 0x12, 0x2b, 0x19, 0x07, 0xee, 0x57, 0x7d, 0x43, 0x4a, 0xac, 0x63, 0xe8, 0xd2, 0x81, 0x27,
	0x61, 0xdf, 0x23, 0xe4, 0xaf, 0xa9, 0xa8, 0x68, 0x1e, 0xc1, 0xee, 0x1f, 0xa0, 0xff, 0xed, 0xc3,
	0xc8, 0xed, 0xcf, 0x03, 0xea, 0xb9, 0x98, 0x20, 0x3b, 0x85, 0x1a, 0xd5, 0xac, 0xb5, 0x2a, 0x72,
	0xf7, 0xd9, 0x6b, 0xff, 0x6c, 0x16, 0xec, 0x70, 0xed, 0xa2, 0x05, 0x8d, 0x89, 0xca, 0x8a, 0x0d,
	0x2f, 0x22, 0xf7, 0xde, 0x4b, 0x9d, 0xca, 0xd1, 0x67, 0x00, 0x00, 0x00, 0xff, 0xff, 0x17, 0x86,
	0x97, 0xfe, 0x05, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// GreetServiceClient is the client API for GreetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type GreetServiceClient interface {
	// Unary
	// 接口声明, 类似http 的一来一回短消息
	Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error)
}

type greetServiceClient struct {
	cc *grpc.ClientConn
}

func NewGreetServiceClient(cc *grpc.ClientConn) GreetServiceClient {
	return &greetServiceClient{cc}
}

func (c *greetServiceClient) Greet(ctx context.Context, in *GreetRequest, opts ...grpc.CallOption) (*GreetResponse, error) {
	out := new(GreetResponse)
	err := c.cc.Invoke(ctx, "/greet.GreetService/Greet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GreetServiceServer is the server API for GreetService service.
type GreetServiceServer interface {
	// Unary
	// 接口声明, 类似http 的一来一回短消息
	Greet(context.Context, *GreetRequest) (*GreetResponse, error)
}

// UnimplementedGreetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedGreetServiceServer struct {
}

func (*UnimplementedGreetServiceServer) Greet(ctx context.Context, req *GreetRequest) (*GreetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Greet not implemented")
}

func RegisterGreetServiceServer(s *grpc.Server, srv GreetServiceServer) {
	s.RegisterService(&_GreetService_serviceDesc, srv)
}

func _GreetService_Greet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GreetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GreetServiceServer).Greet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/greet.GreetService/Greet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GreetServiceServer).Greet(ctx, req.(*GreetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _GreetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "greet.GreetService",
	HandlerType: (*GreetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Greet",
			Handler:    _GreetService_Greet_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "greet/greet.proto",
}