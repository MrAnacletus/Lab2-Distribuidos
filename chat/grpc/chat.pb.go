// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat.proto

package grpc

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

// The request message containing the user's name.
type Jugada struct {
	Id                   int32    `protobuf:"varint,1,opt,name=Id,proto3" json:"Id,omitempty"`
	Jugador              int32    `protobuf:"varint,2,opt,name=Jugador,proto3" json:"Jugador,omitempty"`
	Carta                int32    `protobuf:"varint,3,opt,name=Carta,proto3" json:"Carta,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Jugada) Reset()         { *m = Jugada{} }
func (m *Jugada) String() string { return proto.CompactTextString(m) }
func (*Jugada) ProtoMessage()    {}
func (*Jugada) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{0}
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

func (m *Jugada) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Jugada) GetJugador() int32 {
	if m != nil {
		return m.Jugador
	}
	return 0
}

func (m *Jugada) GetCarta() int32 {
	if m != nil {
		return m.Carta
	}
	return 0
}

// The response message containing the greetings
type Recibo struct {
	Estado               int32    `protobuf:"varint,1,opt,name=estado,proto3" json:"estado,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Recibo) Reset()         { *m = Recibo{} }
func (m *Recibo) String() string { return proto.CompactTextString(m) }
func (*Recibo) ProtoMessage()    {}
func (*Recibo) Descriptor() ([]byte, []int) {
	return fileDescriptor_8c585a45e2093e54, []int{1}
}

func (m *Recibo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Recibo.Unmarshal(m, b)
}
func (m *Recibo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Recibo.Marshal(b, m, deterministic)
}
func (m *Recibo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Recibo.Merge(m, src)
}
func (m *Recibo) XXX_Size() int {
	return xxx_messageInfo_Recibo.Size(m)
}
func (m *Recibo) XXX_DiscardUnknown() {
	xxx_messageInfo_Recibo.DiscardUnknown(m)
}

var xxx_messageInfo_Recibo proto.InternalMessageInfo

func (m *Recibo) GetEstado() int32 {
	if m != nil {
		return m.Estado
	}
	return 0
}

func init() {
	proto.RegisterType((*Jugada)(nil), "grpc.Jugada")
	proto.RegisterType((*Recibo)(nil), "grpc.Recibo")
}

func init() { proto.RegisterFile("chat.proto", fileDescriptor_8c585a45e2093e54) }

var fileDescriptor_8c585a45e2093e54 = []byte{
	// 213 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x34, 0x8f, 0xbf, 0x4b, 0x04, 0x31,
	0x10, 0x85, 0xbd, 0xd5, 0x5b, 0x61, 0x38, 0x2d, 0x82, 0xc8, 0x62, 0x75, 0x6c, 0x65, 0xa1, 0x09,
	0xac, 0xb6, 0x16, 0x7a, 0x16, 0x9e, 0x68, 0x73, 0x76, 0x76, 0x93, 0x1f, 0xec, 0x05, 0xce, 0xcb,
	0x32, 0x99, 0xf8, 0xf7, 0x4b, 0xb2, 0xb1, 0x0a, 0xdf, 0x0b, 0xf3, 0xf1, 0x1e, 0x80, 0xd9, 0x23,
	0xcb, 0x89, 0x02, 0x07, 0x71, 0x36, 0xd2, 0x64, 0xfa, 0x37, 0x68, 0xdf, 0xd3, 0x88, 0x16, 0xc5,
	0x25, 0x34, 0x5b, 0xdb, 0x2d, 0xd6, 0x8b, 0xdb, 0xe5, 0xae, 0xd9, 0x5a, 0xd1, 0xc1, 0x79, 0xf9,
	0x09, 0xd4, 0x35, 0x25, 0xfc, 0x47, 0x71, 0x05, 0xcb, 0x0d, 0x12, 0x63, 0x77, 0x5a, 0xf2, 0x19,
	0xfa, 0x35, 0xb4, 0x3b, 0x67, 0xbc, 0x0e, 0xe2, 0x1a, 0x5a, 0x17, 0x19, 0x6d, 0xa8, 0xb6, 0x4a,
	0xc3, 0x13, 0x5c, 0x6c, 0x0e, 0xde, 0x1d, 0xf9, 0xcb, 0xd1, 0xaf, 0x37, 0x4e, 0xdc, 0xc1, 0xaa,
	0x3a, 0xf3, 0x43, 0x62, 0x25, 0x73, 0x27, 0x39, 0x17, 0xba, 0xa9, 0x34, 0x4b, 0xfb, 0x93, 0x97,
	0xc7, 0xef, 0x61, 0xf4, 0xbc, 0x4f, 0x5a, 0x9a, 0xf0, 0xa3, 0x3e, 0xe9, 0xf9, 0x88, 0xe6, 0xe0,
	0x38, 0x45, 0xf5, 0x81, 0x7a, 0xb8, 0x7f, 0xf5, 0x91, 0xc9, 0xeb, 0xe4, 0x6d, 0x88, 0x2a, 0xef,
	0x54, 0xf9, 0x5c, 0xb7, 0x65, 0xed, 0xc3, 0x5f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf3, 0x3a, 0x7e,
	0x8f, 0xfb, 0x00, 0x00, 0x00,
}
