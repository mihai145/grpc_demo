// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

package app

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

type KeyValuePair struct {
	Key                  string   `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Value                string   `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *KeyValuePair) Reset()         { *m = KeyValuePair{} }
func (m *KeyValuePair) String() string { return proto.CompactTextString(m) }
func (*KeyValuePair) ProtoMessage()    {}
func (*KeyValuePair) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{0}
}

func (m *KeyValuePair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_KeyValuePair.Unmarshal(m, b)
}
func (m *KeyValuePair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_KeyValuePair.Marshal(b, m, deterministic)
}
func (m *KeyValuePair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_KeyValuePair.Merge(m, src)
}
func (m *KeyValuePair) XXX_Size() int {
	return xxx_messageInfo_KeyValuePair.Size(m)
}
func (m *KeyValuePair) XXX_DiscardUnknown() {
	xxx_messageInfo_KeyValuePair.DiscardUnknown(m)
}

var xxx_messageInfo_KeyValuePair proto.InternalMessageInfo

func (m *KeyValuePair) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *KeyValuePair) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

type ListPairs struct {
	Pairs                []*KeyValuePair `protobuf:"bytes,1,rep,name=pairs,proto3" json:"pairs,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ListPairs) Reset()         { *m = ListPairs{} }
func (m *ListPairs) String() string { return proto.CompactTextString(m) }
func (*ListPairs) ProtoMessage()    {}
func (*ListPairs) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{1}
}

func (m *ListPairs) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ListPairs.Unmarshal(m, b)
}
func (m *ListPairs) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ListPairs.Marshal(b, m, deterministic)
}
func (m *ListPairs) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ListPairs.Merge(m, src)
}
func (m *ListPairs) XXX_Size() int {
	return xxx_messageInfo_ListPairs.Size(m)
}
func (m *ListPairs) XXX_DiscardUnknown() {
	xxx_messageInfo_ListPairs.DiscardUnknown(m)
}

var xxx_messageInfo_ListPairs proto.InternalMessageInfo

func (m *ListPairs) GetPairs() []*KeyValuePair {
	if m != nil {
		return m.Pairs
	}
	return nil
}

type StatusResponse struct {
	Ok                   bool     `protobuf:"varint,1,opt,name=ok,proto3" json:"ok,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusResponse) Reset()         { *m = StatusResponse{} }
func (m *StatusResponse) String() string { return proto.CompactTextString(m) }
func (*StatusResponse) ProtoMessage()    {}
func (*StatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{2}
}

func (m *StatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusResponse.Unmarshal(m, b)
}
func (m *StatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusResponse.Marshal(b, m, deterministic)
}
func (m *StatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusResponse.Merge(m, src)
}
func (m *StatusResponse) XXX_Size() int {
	return xxx_messageInfo_StatusResponse.Size(m)
}
func (m *StatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StatusResponse proto.InternalMessageInfo

func (m *StatusResponse) GetOk() bool {
	if m != nil {
		return m.Ok
	}
	return false
}

type Empty struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Empty) Reset()         { *m = Empty{} }
func (m *Empty) String() string { return proto.CompactTextString(m) }
func (*Empty) ProtoMessage()    {}
func (*Empty) Descriptor() ([]byte, []int) {
	return fileDescriptor_e0f9056a14b86d47, []int{3}
}

func (m *Empty) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Empty.Unmarshal(m, b)
}
func (m *Empty) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Empty.Marshal(b, m, deterministic)
}
func (m *Empty) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Empty.Merge(m, src)
}
func (m *Empty) XXX_Size() int {
	return xxx_messageInfo_Empty.Size(m)
}
func (m *Empty) XXX_DiscardUnknown() {
	xxx_messageInfo_Empty.DiscardUnknown(m)
}

var xxx_messageInfo_Empty proto.InternalMessageInfo

func init() {
	proto.RegisterType((*KeyValuePair)(nil), "KeyValuePair")
	proto.RegisterType((*ListPairs)(nil), "ListPairs")
	proto.RegisterType((*StatusResponse)(nil), "StatusResponse")
	proto.RegisterType((*Empty)(nil), "Empty")
}

func init() { proto.RegisterFile("app.proto", fileDescriptor_e0f9056a14b86d47) }

var fileDescriptor_e0f9056a14b86d47 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4c, 0x2c, 0x28, 0xd0,
	0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x57, 0x32, 0xe3, 0xe2, 0xf1, 0x4e, 0xad, 0x0c, 0x4b, 0xcc, 0x29,
	0x4d, 0x0d, 0x48, 0xcc, 0x2c, 0x12, 0x12, 0xe0, 0x62, 0xce, 0x4e, 0xad, 0x94, 0x60, 0x54, 0x60,
	0xd4, 0xe0, 0x0c, 0x02, 0x31, 0x85, 0x44, 0xb8, 0x58, 0xcb, 0x40, 0xd2, 0x12, 0x4c, 0x60, 0x31,
	0x08, 0x47, 0xc9, 0x80, 0x8b, 0xd3, 0x27, 0xb3, 0xb8, 0x04, 0xa4, 0xa7, 0x58, 0x48, 0x99, 0x8b,
	0xb5, 0x00, 0xc4, 0x90, 0x60, 0x54, 0x60, 0xd6, 0xe0, 0x36, 0xe2, 0xd5, 0x43, 0x36, 0x32, 0x08,
	0x22, 0xa7, 0xa4, 0xc0, 0xc5, 0x17, 0x5c, 0x92, 0x58, 0x52, 0x5a, 0x1c, 0x94, 0x5a, 0x5c, 0x90,
	0x9f, 0x57, 0x9c, 0x2a, 0xc4, 0xc7, 0xc5, 0x94, 0x9f, 0x0d, 0xb6, 0x8a, 0x23, 0x88, 0x29, 0x3f,
	0x5b, 0x89, 0x9d, 0x8b, 0xd5, 0x35, 0xb7, 0xa0, 0xa4, 0xd2, 0x28, 0x82, 0x8b, 0xdb, 0x25, 0x33,
	0xb9, 0x24, 0x38, 0xb5, 0xa8, 0x2c, 0x33, 0x39, 0x55, 0x48, 0x93, 0x8b, 0xdd, 0x31, 0x25, 0x05,
	0xec, 0x3c, 0x54, 0xa3, 0xa5, 0xf8, 0xf5, 0xd0, 0x8c, 0x94, 0xe3, 0xe2, 0x70, 0x4f, 0x85, 0xba,
	0x8a, 0x4d, 0x0f, 0x6c, 0x9a, 0x14, 0x97, 0x1e, 0xdc, 0xa5, 0x49, 0x6c, 0x60, 0x5f, 0x1b, 0x03,
	0x02, 0x00, 0x00, 0xff, 0xff, 0x79, 0x64, 0xc3, 0xe7, 0x02, 0x01, 0x00, 0x00,
}
