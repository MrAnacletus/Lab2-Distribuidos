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

type Jugada2 struct {
	ID1                  []int32  `protobuf:"varint,1,rep,packed,name=ID1,proto3" json:"ID1,omitempty"`
	ID2                  []int32  `protobuf:"varint,2,rep,packed,name=ID2,proto3" json:"ID2,omitempty"`
	Jugada1              int32    `protobuf:"varint,3,opt,name=Jugada1,proto3" json:"Jugada1,omitempty"`
	Jugada2              int32    `protobuf:"varint,4,opt,name=Jugada2,proto3" json:"Jugada2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jugada2) Reset()         { *m = Jugada2{} }
func (m *Jugada2) String() string { return proto.CompactTextString(m) }
func (*Jugada2) ProtoMessage()    {}
func (*Jugada2) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{4}
}

func (m *Jugada2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jugada2.Unmarshal(m, b)
}
func (m *Jugada2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jugada2.Marshal(b, m, deterministic)
}
func (m *Jugada2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jugada2.Merge(m, src)
}
func (m *Jugada2) XXX_Size() int {
	return xxx_messageInfo_Jugada2.Size(m)
}
func (m *Jugada2) XXX_DiscardUnknown() {
	xxx_messageInfo_Jugada2.DiscardUnknown(m)
}

var xxx_messageInfo_Jugada2 proto.InternalMessageInfo

func (m *Jugada2) GetID1() []int32 {
	if m != nil {
		return m.ID1
	}
	return nil
}

func (m *Jugada2) GetID2() []int32 {
	if m != nil {
		return m.ID2
	}
	return nil
}

func (m *Jugada2) GetJugada1() int32 {
	if m != nil {
		return m.Jugada1
	}
	return 0
}

func (m *Jugada2) GetJugada2() int32 {
	if m != nil {
		return m.Jugada2
	}
	return 0
}

type Jugada3 struct {
	ID1                  int32    `protobuf:"varint,1,opt,name=ID1,proto3" json:"ID1,omitempty"`
	ID2                  int32    `protobuf:"varint,2,opt,name=ID2,proto3" json:"ID2,omitempty"`
	Jugada1              int32    `protobuf:"varint,3,opt,name=Jugada1,proto3" json:"Jugada1,omitempty"`
	Jugada2              int32    `protobuf:"varint,4,opt,name=Jugada2,proto3" json:"Jugada2,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jugada3) Reset()         { *m = Jugada3{} }
func (m *Jugada3) String() string { return proto.CompactTextString(m) }
func (*Jugada3) ProtoMessage()    {}
func (*Jugada3) Descriptor() ([]byte, []int) {
	return fileDescriptor_5e958f10881bcae0, []int{5}
}

func (m *Jugada3) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Jugada3.Unmarshal(m, b)
}
func (m *Jugada3) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Jugada3.Marshal(b, m, deterministic)
}
func (m *Jugada3) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Jugada3.Merge(m, src)
}
func (m *Jugada3) XXX_Size() int {
	return xxx_messageInfo_Jugada3.Size(m)
}
func (m *Jugada3) XXX_DiscardUnknown() {
	xxx_messageInfo_Jugada3.DiscardUnknown(m)
}

var xxx_messageInfo_Jugada3 proto.InternalMessageInfo

func (m *Jugada3) GetID1() int32 {
	if m != nil {
		return m.ID1
	}
	return 0
}

func (m *Jugada3) GetID2() int32 {
	if m != nil {
		return m.ID2
	}
	return 0
}

func (m *Jugada3) GetJugada1() int32 {
	if m != nil {
		return m.Jugada1
	}
	return 0
}

func (m *Jugada3) GetJugada2() int32 {
	if m != nil {
		return m.Jugada2
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
	return fileDescriptor_5e958f10881bcae0, []int{6}
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
	return fileDescriptor_5e958f10881bcae0, []int{7}
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
	return fileDescriptor_5e958f10881bcae0, []int{8}
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
	proto.RegisterType((*Jugada2)(nil), "grpc.Jugada2")
	proto.RegisterType((*Jugada3)(nil), "grpc.Jugada3")
	proto.RegisterType((*RequestJugadas)(nil), "grpc.RequestJugadas")
	proto.RegisterType((*RequestPozoActual)(nil), "grpc.RequestPozoActual")
	proto.RegisterType((*ResponsePozoActual)(nil), "grpc.ResponsePozoActual")
}

func init() { proto.RegisterFile("comunicacion.proto", fileDescriptor_5e958f10881bcae0) }

var fileDescriptor_5e958f10881bcae0 = []byte{
	// 448 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xcd, 0x6e, 0xd3, 0x40,
	0x10, 0xc7, 0x65, 0x27, 0x6d, 0xe9, 0x24, 0xb4, 0x65, 0x85, 0xc0, 0xca, 0x29, 0xf2, 0x01, 0x22,
	0x41, 0xe3, 0x76, 0xcd, 0x03, 0x50, 0x14, 0x04, 0x41, 0x6d, 0x41, 0xce, 0x8d, 0xdb, 0xda, 0x1e,
	0x05, 0x0b, 0xc7, 0x6b, 0xbc, 0x5e, 0xa4, 0xf2, 0x0e, 0xbc, 0x33, 0xda, 0xaf, 0xc6, 0xc1, 0xa8,
	0x48, 0xed, 0xc9, 0xf3, 0xf1, 0x1b, 0xcf, 0xf8, 0x3f, 0x23, 0x03, 0xc9, 0xf8, 0x46, 0x56, 0x45,
	0xc6, 0xb2, 0x82, 0x57, 0xf3, 0xba, 0xe1, 0x2d, 0x27, 0xc3, 0x75, 0x53, 0x67, 0x61, 0x08, 0xe3,
	0x8f, 0x58, 0x96, 0x3c, 0xc1, 0x1f, 0x12, 0x45, 0x4b, 0x08, 0x0c, 0xaf, 0xd9, 0x06, 0x03, 0x6f,
	0xea, 0xcd, 0x0e, 0x13, 0x6d, 0x87, 0x2f, 0x00, 0x2c, 0x53, 0x97, 0x37, 0x24, 0x80, 0x83, 0x2b,
	0x14, 0x82, 0xad, 0x1d, 0xe4, 0xdc, 0xf0, 0x0c, 0xf6, 0x3f, 0xc9, 0x35, 0xcb, 0x19, 0x39, 0x02,
	0x7f, 0xb9, 0xd0, 0xe9, 0xbd, 0xc4, 0x5f, 0x2e, 0xc8, 0x33, 0x97, 0x09, 0x7c, 0x1d, 0xb3, 0x5e,
	0x18, 0xc3, 0x61, 0x82, 0x42, 0x96, 0x2d, 0xcb, 0xf9, 0xbf, 0x8a, 0xde, 0x0b, 0x95, 0x71, 0x45,
	0xc6, 0x0b, 0x19, 0x1c, 0x98, 0x72, 0x4a, 0x4e, 0x60, 0xb0, 0x5c, 0x9c, 0x07, 0xde, 0x74, 0x30,
	0xdb, 0x4b, 0x94, 0x69, 0x22, 0x34, 0xf0, 0x5d, 0x84, 0xaa, 0x79, 0x0d, 0x7e, 0x1e, 0x0c, 0xf4,
	0x7b, 0x9c, 0xbb, 0xcd, 0xd0, 0x60, 0xd8, 0xcd, 0xd0, 0x6d, 0x8b, 0x78, 0xdb, 0xc2, 0xeb, 0xb5,
	0xf0, 0x1e, 0xd2, 0x62, 0x0a, 0x47, 0x56, 0x73, 0x13, 0x11, 0x7f, 0x7f, 0x7f, 0xf8, 0x12, 0x9e,
	0x58, 0xe2, 0x0b, 0xff, 0xc5, 0x2f, 0xb2, 0x56, 0xb2, 0x52, 0xed, 0x47, 0x79, 0x16, 0xd3, 0x76,
	0x38, 0x03, 0x92, 0xa0, 0xa8, 0x79, 0x25, 0xf0, 0x6e, 0x92, 0xfe, 0xf6, 0x61, 0x7c, 0x59, 0xe4,
	0xd8, 0xac, 0xb0, 0xf9, 0x59, 0x64, 0x48, 0xce, 0xe0, 0xd1, 0x8a, 0xdd, 0xe8, 0xed, 0x12, 0x32,
	0x57, 0x17, 0x31, 0xef, 0x9e, 0xc3, 0xe4, 0x64, 0x27, 0xa6, 0xd6, 0xff, 0x16, 0x46, 0x9d, 0xa9,
	0xc8, 0x73, 0x03, 0xf4, 0x06, 0x9d, 0x04, 0x2e, 0xd1, 0x1b, 0xec, 0x15, 0xc0, 0x0a, 0xab, 0xdc,
	0x9e, 0xca, 0xd8, 0x70, 0xc6, 0x9b, 0x1c, 0xdf, 0x56, 0xd9, 0xa3, 0x38, 0x85, 0xd1, 0x16, 0xa6,
	0xe4, 0x71, 0x97, 0xa6, 0xff, 0xc1, 0xe3, 0x5d, 0x3c, 0xee, 0xe1, 0x54, 0xc2, 0xb1, 0xba, 0xf0,
	0x6b, 0x9e, 0xa3, 0x53, 0xe4, 0xf5, 0x1d, 0xd3, 0xf5, 0xd5, 0x78, 0x03, 0xf0, 0x01, 0x6f, 0x37,
	0xf8, 0x74, 0x47, 0x0c, 0x1b, 0xed, 0x57, 0xd1, 0xcf, 0x30, 0x52, 0x7a, 0xb8, 0x96, 0x0f, 0x96,
	0xf4, 0xdd, 0xe6, 0xeb, 0xf7, 0x75, 0xd1, 0x7e, 0x93, 0xe9, 0x3c, 0xe3, 0x9b, 0xe8, 0xaa, 0xb9,
	0xa8, 0x58, 0x56, 0x62, 0x2b, 0x45, 0x74, 0xc9, 0x52, 0x7a, 0xba, 0x28, 0x44, 0xdb, 0x14, 0xa9,
	0x2c, 0x72, 0x2e, 0xa2, 0xba, 0x91, 0x98, 0xb2, 0x48, 0xff, 0x04, 0xa2, 0x7b, 0x55, 0xa5, 0xfb,
	0xfa, 0x11, 0xff, 0x09, 0x00, 0x00, 0xff, 0xff, 0x32, 0x15, 0x83, 0x9d, 0x57, 0x04, 0x00, 0x00,
}
