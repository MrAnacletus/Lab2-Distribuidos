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

type Resultados struct {
	Lista                []*Resultado `protobuf:"bytes,1,rep,name=Lista,proto3" json:"Lista,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Resultados) Reset()         { *m = Resultados{} }
func (m *Resultados) String() string { return proto.CompactTextString(m) }
func (*Resultados) ProtoMessage()    {}
func (*Resultados) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{4}
}

func (m *Resultados) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Resultados.Unmarshal(m, b)
}
func (m *Resultados) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Resultados.Marshal(b, m, deterministic)
}
func (m *Resultados) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Resultados.Merge(m, src)
}
func (m *Resultados) XXX_Size() int {
	return xxx_messageInfo_Resultados.Size(m)
}
func (m *Resultados) XXX_DiscardUnknown() {
	xxx_messageInfo_Resultados.DiscardUnknown(m)
}

var xxx_messageInfo_Resultados proto.InternalMessageInfo

func (m *Resultados) GetLista() []*Resultado {
	if m != nil {
		return m.Lista
	}
	return nil
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
	proto.RegisterType((*Resultados)(nil), "grpc.Resultados")
	proto.RegisterType((*RequestPozoActual)(nil), "grpc.RequestPozoActual")
	proto.RegisterType((*ResponsePozoActual)(nil), "grpc.ResponsePozoActual")
}

func init() { proto.RegisterFile("comunicacion.proto", fileDescriptor_5e958f10881bcae0) }

var fileDescriptor_5e958f10881bcae0 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0x4d, 0x4f, 0xc2, 0x40,
	0x14, 0x4c, 0x11, 0x50, 0x1e, 0xf8, 0xb5, 0x07, 0x6d, 0x38, 0x91, 0x4d, 0x54, 0x12, 0x63, 0x4b,
	0x4a, 0xbc, 0x8b, 0xc1, 0x20, 0x06, 0xd4, 0x94, 0x9b, 0xb7, 0xed, 0xb2, 0xa9, 0x1b, 0xdb, 0x6e,
	0xed, 0xee, 0x9a, 0xe0, 0x0f, 0xf6, 0x77, 0x98, 0x7e, 0x21, 0x12, 0xc2, 0x41, 0x4f, 0xdd, 0x79,
	0x6f, 0x26, 0x6f, 0x66, 0x52, 0x40, 0x54, 0x84, 0x3a, 0xe2, 0x94, 0x50, 0x2e, 0x22, 0x2b, 0x4e,
	0x84, 0x12, 0xa8, 0xea, 0x27, 0x31, 0xc5, 0x18, 0x5a, 0xf7, 0x2c, 0x08, 0x84, 0xcb, 0xde, 0x35,
	0x93, 0x0a, 0x21, 0xa8, 0x3e, 0x92, 0x90, 0x99, 0x46, 0xc7, 0xe8, 0x36, 0xdc, 0xec, 0x8d, 0xcf,
	0x01, 0x0a, 0x4e, 0x1c, 0x2c, 0x90, 0x09, 0xbb, 0x53, 0x26, 0x25, 0xf1, 0x4b, 0x52, 0x09, 0x71,
	0x0f, 0xea, 0x0f, 0xda, 0x27, 0x73, 0x82, 0x0e, 0xa0, 0x32, 0x1e, 0x66, 0xeb, 0x9a, 0x5b, 0x19,
	0x0f, 0xd1, 0x49, 0xb9, 0x31, 0x2b, 0xd9, 0xac, 0x40, 0xb8, 0x0f, 0x0d, 0x97, 0x49, 0x1d, 0x28,
	0x32, 0x17, 0x9b, 0x44, 0x77, 0x32, 0xdd, 0x94, 0xa2, 0x1c, 0xe1, 0x3e, 0xc0, 0x52, 0x24, 0xd1,
	0x19, 0xd4, 0x26, 0x5c, 0x2a, 0x62, 0x1a, 0x9d, 0x9d, 0x6e, 0xd3, 0x39, 0xb4, 0xd2, 0x58, 0xd6,
	0x92, 0xe0, 0xe6, 0x5b, 0x7c, 0x01, 0xc7, 0x45, 0xc4, 0x67, 0xf1, 0x29, 0x06, 0x54, 0x69, 0x12,
	0xa4, 0x61, 0x53, 0x54, 0xdc, 0xcc, 0xde, 0xb8, 0x0b, 0xc8, 0x65, 0x32, 0x16, 0x91, 0x64, 0xdb,
	0x99, 0xce, 0x97, 0x51, 0x74, 0x37, 0x63, 0xc9, 0x07, 0xa7, 0x0c, 0xf5, 0x60, 0x6f, 0x46, 0x16,
	0xd9, 0x08, 0xa1, 0xdc, 0xc7, 0x6a, 0xb7, 0xed, 0xa3, 0x5f, 0xb3, 0xb4, 0xcb, 0x1b, 0x68, 0xae,
	0xb8, 0x42, 0xa7, 0xa5, 0xf9, 0x35, 0xa3, 0x6d, 0x73, 0x99, 0x6a, 0xdd, 0xd8, 0x25, 0x34, 0x46,
	0x4c, 0x15, 0xb5, 0xb7, 0x72, 0x5a, 0x8e, 0x36, 0x9c, 0xbb, 0x86, 0xfd, 0x11, 0x53, 0x2b, 0xe5,
	0x6d, 0x71, 0xf9, 0xc3, 0x72, 0x9e, 0xa0, 0x99, 0x5e, 0x2c, 0x63, 0xfe, 0xdb, 0xf4, 0x6d, 0xf8,
	0xf2, 0xe6, 0x73, 0xf5, 0xaa, 0x3d, 0x8b, 0x8a, 0xd0, 0x9e, 0x26, 0x83, 0x88, 0xd0, 0x80, 0x29,
	0x2d, 0xed, 0x09, 0xf1, 0x9c, 0xab, 0x21, 0x97, 0x2a, 0xe1, 0x9e, 0xe6, 0x73, 0x21, 0xed, 0x38,
	0xd1, 0xcc, 0x23, 0x76, 0xf6, 0xcf, 0xda, 0x7f, 0x52, 0x79, 0xf5, 0xec, 0xd3, 0xff, 0x0e, 0x00,
	0x00, 0xff, 0xff, 0xa3, 0x33, 0x0a, 0x09, 0x06, 0x03, 0x00, 0x00,
}
