// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ligato/vpp/l2/xconnect.proto

package vpp_l2

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

type XConnectPair struct {
	ReceiveInterface     string   `protobuf:"bytes,1,opt,name=receive_interface,json=receiveInterface,proto3" json:"receive_interface,omitempty"`
	TransmitInterface    string   `protobuf:"bytes,2,opt,name=transmit_interface,json=transmitInterface,proto3" json:"transmit_interface,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *XConnectPair) Reset()         { *m = XConnectPair{} }
func (m *XConnectPair) String() string { return proto.CompactTextString(m) }
func (*XConnectPair) ProtoMessage()    {}
func (*XConnectPair) Descriptor() ([]byte, []int) {
	return fileDescriptor_d21cc032e061c819, []int{0}
}

func (m *XConnectPair) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_XConnectPair.Unmarshal(m, b)
}
func (m *XConnectPair) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_XConnectPair.Marshal(b, m, deterministic)
}
func (m *XConnectPair) XXX_Merge(src proto.Message) {
	xxx_messageInfo_XConnectPair.Merge(m, src)
}
func (m *XConnectPair) XXX_Size() int {
	return xxx_messageInfo_XConnectPair.Size(m)
}
func (m *XConnectPair) XXX_DiscardUnknown() {
	xxx_messageInfo_XConnectPair.DiscardUnknown(m)
}

var xxx_messageInfo_XConnectPair proto.InternalMessageInfo

func (m *XConnectPair) GetReceiveInterface() string {
	if m != nil {
		return m.ReceiveInterface
	}
	return ""
}

func (m *XConnectPair) GetTransmitInterface() string {
	if m != nil {
		return m.TransmitInterface
	}
	return ""
}

func init() {
	proto.RegisterType((*XConnectPair)(nil), "ligato.vpp.l2.XConnectPair")
}

func init() { proto.RegisterFile("ligato/vpp/l2/xconnect.proto", fileDescriptor_d21cc032e061c819) }

var fileDescriptor_d21cc032e061c819 = []byte{
	// 173 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xc9, 0x4c, 0x4f,
	0x2c, 0xc9, 0xd7, 0x2f, 0x2b, 0x28, 0xd0, 0xcf, 0x31, 0xd2, 0xaf, 0x48, 0xce, 0xcf, 0xcb, 0x4b,
	0x4d, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x85, 0xc8, 0xea, 0x95, 0x15, 0x14,
	0xe8, 0xe5, 0x18, 0x29, 0x65, 0x71, 0xf1, 0x44, 0x38, 0x43, 0x14, 0x04, 0x24, 0x66, 0x16, 0x09,
	0x69, 0x73, 0x09, 0x16, 0xa5, 0x26, 0xa7, 0x66, 0x96, 0xa5, 0xc6, 0x67, 0xe6, 0x95, 0xa4, 0x16,
	0xa5, 0x25, 0x26, 0xa7, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x09, 0x40, 0x25, 0x3c, 0x61,
	0xe2, 0x42, 0xba, 0x5c, 0x42, 0x25, 0x45, 0x89, 0x79, 0xc5, 0xb9, 0x99, 0x25, 0x48, 0xaa, 0x99,
	0xc0, 0xaa, 0x05, 0x61, 0x32, 0x70, 0xe5, 0x4e, 0x66, 0x51, 0x26, 0xe9, 0xf9, 0x7a, 0x50, 0xfb,
	0x33, 0xc1, 0x0e, 0xd4, 0x4d, 0x4c, 0x4f, 0xcd, 0x2b, 0xd1, 0x2f, 0x33, 0xd6, 0x07, 0xbb, 0x4e,
	0x1f, 0xc5, 0xe9, 0xd6, 0x65, 0x05, 0x05, 0xf1, 0x39, 0x46, 0x49, 0x6c, 0x60, 0x39, 0x63, 0x40,
	0x00, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x24, 0x86, 0xb6, 0xd9, 0x00, 0x00, 0x00,
}
