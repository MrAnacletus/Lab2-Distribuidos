// Code generated by protoc-gen-go. DO NOT EDIT.
// source: comunicacion.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=Message,proto3" json:"Message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type Jugada struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Jugada               int32    `protobuf:"varint,2,opt,name=Jugada,proto3" json:"Jugada,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jugada) Reset()         { *m = Jugada{} }
func (m *Jugada) String() string { return proto.CompactTextString(m) }
func (*Jugada) ProtoMessage()    {}
func (*Jugada) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{2}
}

func (m *Jugada) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jugada.Unmarshal(m, b)
}
func (m *Jugada) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jugada.Marshal(b, m, deterministic)
}
func (m *Jugada) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jugada.Merge(m, src)
}
func (m *Jugada) XXX_Size() int {
	return xxx_messageInfo_Jugada.Size(m)
}
func (m *Jugada) XXX_DiscardUnknown() {
	xxx_messageInfo_Jugada.DiscardUnknown(m)
}

var xxx_messageInfo_Jugada proto.InternalMessageInfo

func (m *Jugada) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Jugada) GetJugada() int32 {
	if m != nil {
		return m.Jugada
	}
	return 0
}

type Resultado struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Estado               int32    `protobuf:"varint,2,opt,name=Estado,proto3" json:"Estado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Resultado) Reset()         { *m = Resultado{} }
func (m *Resultado) String() string { return proto.CompactTextString(m) }
func (*Resultado) ProtoMessage()    {}
func (*Resultado) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{3}
}

func (m *Resultado) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resultado.Unmarshal(m, b)
}
func (m *Resultado) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resultado.Marshal(b, m, deterministic)
}
func (m *Resultado) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resultado.Merge(m, src)
}
func (m *Resultado) XXX_Size() int {
	return xxx_messageInfo_Resultado.Size(m)
}
func (m *Resultado) XXX_DiscardUnknown() {
	xxx_messageInfo_Resultado.DiscardUnknown(m)
}

var xxx_messageInfo_Resultado proto.InternalMessageInfo

func (m *Resultado) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *Resultado) GetEstado() int32 {
	if m != nil {
		return m.Estado
	}
	return 0
}

type RequestJugadas struct {
	ID                   int32    `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestJugadas) Reset()         { *m = RequestJugadas{} }
func (m *RequestJugadas) String() string { return proto.CompactTextString(m) }
func (*RequestJugadas) ProtoMessage()    {}
func (*RequestJugadas) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{4}
}

func (m *RequestJugadas) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestJugadas.Unmarshal(m, b)
}
func (m *RequestJugadas) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestJugadas.Marshal(b, m, deterministic)
}
func (m *RequestJugadas) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestJugadas.Merge(m, src)
}
func (m *RequestJugadas) XXX_Size() int {
	return xxx_messageInfo_RequestJugadas.Size(m)
}
func (m *RequestJugadas) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestJugadas.DiscardUnknown(m)
}

var xxx_messageInfo_RequestJugadas proto.InternalMessageInfo

func (m *RequestJugadas) GetID() int32 {
	if m != nil {
		return m.ID
	}
	return 0
}

type RequestPozoActual struct {
	Pozo                 int32    `protobuf:"varint,1,opt,name=Pozo,proto3" json:"Pozo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RequestPozoActual) Reset()         { *m = RequestPozoActual{} }
func (m *RequestPozoActual) String() string { return proto.CompactTextString(m) }
func (*RequestPozoActual) ProtoMessage()    {}
func (*RequestPozoActual) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{5}
}

func (m *RequestPozoActual) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RequestPozoActual.Unmarshal(m, b)
}
func (m *RequestPozoActual) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RequestPozoActual.Marshal(b, m, deterministic)
}
func (m *RequestPozoActual) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPozoActual.Merge(m, src)
}
func (m *RequestPozoActual) XXX_Size() int {
	return xxx_messageInfo_RequestPozoActual.Size(m)
}
func (m *RequestPozoActual) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPozoActual.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPozoActual proto.InternalMessageInfo

func (m *RequestPozoActual) GetPozo() int32 {
	if m != nil {
		return m.Pozo
	}
	return 0
}

type ResponsePozoActual struct {
	Pozo                 int32    `protobuf:"varint,1,opt,name=Pozo,proto3" json:"Pozo,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponsePozoActual) Reset()         { *m = ResponsePozoActual{} }
func (m *ResponsePozoActual) String() string { return proto.CompactTextString(m) }
func (*ResponsePozoActual) ProtoMessage()    {}
func (*ResponsePozoActual) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{6}
}

func (m *ResponsePozoActual) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponsePozoActual.Unmarshal(m, b)
}
func (m *ResponsePozoActual) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponsePozoActual.Marshal(b, m, deterministic)
}
func (m *ResponsePozoActual) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePozoActual.Merge(m, src)
}
func (m *ResponsePozoActual) XXX_Size() int {
	return xxx_messageInfo_ResponsePozoActual.Size(m)
}
func (m *ResponsePozoActual) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePozoActual.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePozoActual proto.InternalMessageInfo

func (m *ResponsePozoActual) GetPozo() int32 {
	if m != nil {
		return m.Pozo
	}
	return 0
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "grpc.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "grpc.HelloReply")
	proto.RegisterType((*Jugada)(nil), "grpc.Jugada")
	proto.RegisterType((*Resultado)(nil), "grpc.Resultado")
	proto.RegisterType((*RequestJugadas)(nil), "grpc.RequestJugadas")
	proto.RegisterType((*RequestPozoActual)(nil), "grpc.RequestPozoActual")
	proto.RegisterType((*ResponsePozoActual)(nil), "grpc.ResponsePozoActual")
}

func init() { proto.RegisterFile("comunicacion.proto", fileDescriptor_5e958f10881bcae0) }

var fileDescriptor_5e958f10881bcae0 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0xd1, 0xae, 0xd2, 0x40,
	0x10, 0x86, 0x43, 0x83, 0x28, 0x03, 0x01, 0xdd, 0x18, 0x6d, 0xb8, 0x22, 0x7b, 0xa1, 0x24, 0x62,
	0x4b, 0x8a, 0x0f, 0x20, 0x06, 0xa3, 0x18, 0x40, 0x53, 0xee, 0xbc, 0xdb, 0x6e, 0x27, 0x75, 0x63,
	0xdb, 0xad, 0xdd, 0xae, 0x09, 0xbe, 0x82, 0x2f, 0x6d, 0xda, 0x6e, 0x11, 0xac, 0xe1, 0xe4, 0x9c,
	0x73, 0xc5, 0xce, 0xcc, 0xf7, 0xd3, 0xe9, 0xd7, 0x16, 0x08, 0x97, 0x89, 0x4e, 0x05, 0x67, 0x5c,
	0xc8, 0xd4, 0xc9, 0x72, 0x59, 0x48, 0xd2, 0x8d, 0xf2, 0x8c, 0x53, 0x0a, 0xc3, 0x8f, 0x18, 0xc7,
	0xd2, 0xc7, 0x1f, 0x1a, 0x55, 0x41, 0x08, 0x74, 0xf7, 0x2c, 0x41, 0xbb, 0x33, 0xed, 0xcc, 0xfa,
	0x7e, 0x75, 0xa6, 0x2f, 0x00, 0x0c, 0x93, 0xc5, 0x47, 0x62, 0xc3, 0xc3, 0x1d, 0x2a, 0xc5, 0xa2,
	0x06, 0x6a, 0x4a, 0xba, 0x80, 0xde, 0x27, 0x1d, 0xb1, 0x90, 0x91, 0x11, 0x58, 0x9b, 0x75, 0x35,
	0x7e, 0xe0, 0x5b, 0x9b, 0x35, 0x79, 0xd6, 0x4c, 0x6c, 0xab, 0xea, 0x99, 0x8a, 0x2e, 0xa1, 0xef,
	0xa3, 0xd2, 0x71, 0xc1, 0x42, 0xf9, 0xbf, 0xd0, 0x7b, 0x55, 0x4e, 0x9a, 0x50, 0x5d, 0xd1, 0x29,
	0x8c, 0xcc, 0xb6, 0xf5, 0xbf, 0xa8, 0x7f, 0x93, 0xf4, 0x25, 0x3c, 0x31, 0xc4, 0x17, 0xf9, 0x4b,
	0xae, 0x78, 0xa1, 0x59, 0x5c, 0xde, 0x59, 0x59, 0x19, 0xac, 0x3a, 0xd3, 0x19, 0x10, 0x1f, 0x55,
	0x26, 0x53, 0x85, 0xd7, 0x49, 0xef, 0xb7, 0x05, 0xc3, 0xad, 0x08, 0x31, 0x3f, 0x60, 0xfe, 0x53,
	0x70, 0x24, 0x0b, 0x78, 0x74, 0x60, 0xc7, 0xca, 0x0b, 0x21, 0x4e, 0xe9, 0xd2, 0x39, 0x17, 0x39,
	0x79, 0x7c, 0xd1, 0x2b, 0xc5, 0xbd, 0x85, 0xc1, 0xd9, 0x56, 0xe4, 0x79, 0x0d, 0xb4, 0x16, 0x9d,
	0xd8, 0xcd, 0xa0, 0xb5, 0xd8, 0x2b, 0x80, 0x03, 0xa6, 0xa1, 0x91, 0x3c, 0xac, 0xb9, 0xba, 0x9a,
	0x8c, 0x4f, 0x29, 0xa3, 0x73, 0x0e, 0x83, 0xbf, 0xb0, 0x77, 0x2b, 0x7a, 0x79, 0x03, 0xed, 0x69,
	0x18, 0x97, 0x6f, 0xc6, 0x5e, 0x86, 0xd8, 0xf8, 0x98, 0x5f, 0xd9, 0xad, 0xed, 0xe2, 0x0d, 0xc0,
	0x07, 0x3c, 0x3d, 0xbf, 0xa7, 0x17, 0x2a, 0x4c, 0xb7, 0x9d, 0xf2, 0x3e, 0xc3, 0xa0, 0xb4, 0xd1,
	0x5c, 0xf2, 0xde, 0x42, 0xdf, 0x25, 0x5f, 0xbf, 0x47, 0xa2, 0xf8, 0xa6, 0x03, 0x87, 0xcb, 0xc4,
	0xdd, 0xe5, 0xab, 0x94, 0xf1, 0x18, 0x0b, 0xad, 0xdc, 0x2d, 0x0b, 0xbc, 0xd7, 0x6b, 0xa1, 0x8a,
	0x5c, 0x04, 0x5a, 0x84, 0x52, 0xb9, 0x59, 0xae, 0x31, 0x60, 0x6e, 0xf5, 0xf1, 0xb8, 0x77, 0x4a,
	0x05, 0xbd, 0xea, 0x67, 0xf9, 0x27, 0x00, 0x00, 0xff, 0xff, 0x4c, 0xd3, 0x7d, 0x8c, 0x8f, 0x03,
	0x00, 0x00,
}
