// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: interfaces.proto

package interfaces

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import namespace "github.com/ligato/vpp-agent/plugins/linuxv2/model/namespace"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type LinuxInterface_Type int32

const (
	LinuxInterface_UNDEFINED LinuxInterface_Type = 0
	LinuxInterface_VETH      LinuxInterface_Type = 1
	LinuxInterface_AUTO_TAP  LinuxInterface_Type = 2
)

var LinuxInterface_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "VETH",
	2: "AUTO_TAP",
}
var LinuxInterface_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"VETH":      1,
	"AUTO_TAP":  2,
}

func (x LinuxInterface_Type) String() string {
	return proto.EnumName(LinuxInterface_Type_name, int32(x))
}
func (LinuxInterface_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e2fdd9dd69733505, []int{0, 0}
}

type LinuxInterface struct {
	Name        string                       `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Type        LinuxInterface_Type          `protobuf:"varint,2,opt,name=type,proto3,enum=interfaces.LinuxInterface_Type" json:"type,omitempty"`
	Namespace   *namespace.LinuxNetNamespace `protobuf:"bytes,3,opt,name=namespace" json:"namespace,omitempty"`
	HostIfName  string                       `protobuf:"bytes,4,opt,name=host_if_name,json=hostIfName,proto3" json:"host_if_name,omitempty"`
	Enabled     bool                         `protobuf:"varint,5,opt,name=enabled,proto3" json:"enabled,omitempty"`
	IpAddresses []string                     `protobuf:"bytes,6,rep,name=ip_addresses,json=ipAddresses" json:"ip_addresses,omitempty"`
	PhysAddress string                       `protobuf:"bytes,7,opt,name=phys_address,json=physAddress,proto3" json:"phys_address,omitempty"`
	Mtu         uint32                       `protobuf:"varint,8,opt,name=mtu,proto3" json:"mtu,omitempty"`
	// Types that are valid to be assigned to Link:
	//	*LinuxInterface_Veth
	//	*LinuxInterface_AutoTap
	Link                 isLinuxInterface_Link `protobuf_oneof:"link"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LinuxInterface) Reset()         { *m = LinuxInterface{} }
func (m *LinuxInterface) String() string { return proto.CompactTextString(m) }
func (*LinuxInterface) ProtoMessage()    {}
func (*LinuxInterface) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e2fdd9dd69733505, []int{0}
}
func (m *LinuxInterface) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterface.Unmarshal(m, b)
}
func (m *LinuxInterface) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterface.Marshal(b, m, deterministic)
}
func (dst *LinuxInterface) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterface.Merge(dst, src)
}
func (m *LinuxInterface) XXX_Size() int {
	return xxx_messageInfo_LinuxInterface.Size(m)
}
func (m *LinuxInterface) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterface.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterface proto.InternalMessageInfo

type isLinuxInterface_Link interface {
	isLinuxInterface_Link()
}

type LinuxInterface_Veth struct {
	Veth *LinuxInterface_VethLink `protobuf:"bytes,9,opt,name=veth,oneof"`
}
type LinuxInterface_AutoTap struct {
	AutoTap *LinuxInterface_AutoTapLink `protobuf:"bytes,10,opt,name=autoTap,oneof"`
}

func (*LinuxInterface_Veth) isLinuxInterface_Link()    {}
func (*LinuxInterface_AutoTap) isLinuxInterface_Link() {}

func (m *LinuxInterface) GetLink() isLinuxInterface_Link {
	if m != nil {
		return m.Link
	}
	return nil
}

func (m *LinuxInterface) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LinuxInterface) GetType() LinuxInterface_Type {
	if m != nil {
		return m.Type
	}
	return LinuxInterface_UNDEFINED
}

func (m *LinuxInterface) GetNamespace() *namespace.LinuxNetNamespace {
	if m != nil {
		return m.Namespace
	}
	return nil
}

func (m *LinuxInterface) GetHostIfName() string {
	if m != nil {
		return m.HostIfName
	}
	return ""
}

func (m *LinuxInterface) GetEnabled() bool {
	if m != nil {
		return m.Enabled
	}
	return false
}

func (m *LinuxInterface) GetIpAddresses() []string {
	if m != nil {
		return m.IpAddresses
	}
	return nil
}

func (m *LinuxInterface) GetPhysAddress() string {
	if m != nil {
		return m.PhysAddress
	}
	return ""
}

func (m *LinuxInterface) GetMtu() uint32 {
	if m != nil {
		return m.Mtu
	}
	return 0
}

func (m *LinuxInterface) GetVeth() *LinuxInterface_VethLink {
	if x, ok := m.GetLink().(*LinuxInterface_Veth); ok {
		return x.Veth
	}
	return nil
}

func (m *LinuxInterface) GetAutoTap() *LinuxInterface_AutoTapLink {
	if x, ok := m.GetLink().(*LinuxInterface_AutoTap); ok {
		return x.AutoTap
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LinuxInterface) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LinuxInterface_OneofMarshaler, _LinuxInterface_OneofUnmarshaler, _LinuxInterface_OneofSizer, []interface{}{
		(*LinuxInterface_Veth)(nil),
		(*LinuxInterface_AutoTap)(nil),
	}
}

func _LinuxInterface_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LinuxInterface)
	// link
	switch x := m.Link.(type) {
	case *LinuxInterface_Veth:
		_ = b.EncodeVarint(9<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Veth); err != nil {
			return err
		}
	case *LinuxInterface_AutoTap:
		_ = b.EncodeVarint(10<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.AutoTap); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("LinuxInterface.Link has unexpected type %T", x)
	}
	return nil
}

func _LinuxInterface_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LinuxInterface)
	switch tag {
	case 9: // link.veth
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinuxInterface_VethLink)
		err := b.DecodeMessage(msg)
		m.Link = &LinuxInterface_Veth{msg}
		return true, err
	case 10: // link.autoTap
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LinuxInterface_AutoTapLink)
		err := b.DecodeMessage(msg)
		m.Link = &LinuxInterface_AutoTap{msg}
		return true, err
	default:
		return false, nil
	}
}

func _LinuxInterface_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LinuxInterface)
	// link
	switch x := m.Link.(type) {
	case *LinuxInterface_Veth:
		s := proto.Size(x.Veth)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LinuxInterface_AutoTap:
		s := proto.Size(x.AutoTap)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type LinuxInterface_VethLink struct {
	PeerIfName           string   `protobuf:"bytes,1,opt,name=peer_if_name,json=peerIfName,proto3" json:"peer_if_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxInterface_VethLink) Reset()         { *m = LinuxInterface_VethLink{} }
func (m *LinuxInterface_VethLink) String() string { return proto.CompactTextString(m) }
func (*LinuxInterface_VethLink) ProtoMessage()    {}
func (*LinuxInterface_VethLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e2fdd9dd69733505, []int{0, 0}
}
func (m *LinuxInterface_VethLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterface_VethLink.Unmarshal(m, b)
}
func (m *LinuxInterface_VethLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterface_VethLink.Marshal(b, m, deterministic)
}
func (dst *LinuxInterface_VethLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterface_VethLink.Merge(dst, src)
}
func (m *LinuxInterface_VethLink) XXX_Size() int {
	return xxx_messageInfo_LinuxInterface_VethLink.Size(m)
}
func (m *LinuxInterface_VethLink) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterface_VethLink.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterface_VethLink proto.InternalMessageInfo

func (m *LinuxInterface_VethLink) GetPeerIfName() string {
	if m != nil {
		return m.PeerIfName
	}
	return ""
}

type LinuxInterface_AutoTapLink struct {
	TempIfName           string   `protobuf:"bytes,1,opt,name=temp_if_name,json=tempIfName,proto3" json:"temp_if_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LinuxInterface_AutoTapLink) Reset()         { *m = LinuxInterface_AutoTapLink{} }
func (m *LinuxInterface_AutoTapLink) String() string { return proto.CompactTextString(m) }
func (*LinuxInterface_AutoTapLink) ProtoMessage()    {}
func (*LinuxInterface_AutoTapLink) Descriptor() ([]byte, []int) {
	return fileDescriptor_interfaces_e2fdd9dd69733505, []int{0, 1}
}
func (m *LinuxInterface_AutoTapLink) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LinuxInterface_AutoTapLink.Unmarshal(m, b)
}
func (m *LinuxInterface_AutoTapLink) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LinuxInterface_AutoTapLink.Marshal(b, m, deterministic)
}
func (dst *LinuxInterface_AutoTapLink) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LinuxInterface_AutoTapLink.Merge(dst, src)
}
func (m *LinuxInterface_AutoTapLink) XXX_Size() int {
	return xxx_messageInfo_LinuxInterface_AutoTapLink.Size(m)
}
func (m *LinuxInterface_AutoTapLink) XXX_DiscardUnknown() {
	xxx_messageInfo_LinuxInterface_AutoTapLink.DiscardUnknown(m)
}

var xxx_messageInfo_LinuxInterface_AutoTapLink proto.InternalMessageInfo

func (m *LinuxInterface_AutoTapLink) GetTempIfName() string {
	if m != nil {
		return m.TempIfName
	}
	return ""
}

func init() {
	proto.RegisterType((*LinuxInterface)(nil), "interfaces.LinuxInterface")
	proto.RegisterType((*LinuxInterface_VethLink)(nil), "interfaces.LinuxInterface.VethLink")
	proto.RegisterType((*LinuxInterface_AutoTapLink)(nil), "interfaces.LinuxInterface.AutoTapLink")
	proto.RegisterEnum("interfaces.LinuxInterface_Type", LinuxInterface_Type_name, LinuxInterface_Type_value)
}

func init() { proto.RegisterFile("interfaces.proto", fileDescriptor_interfaces_e2fdd9dd69733505) }

var fileDescriptor_interfaces_e2fdd9dd69733505 = []byte{
	// 419 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xc1, 0x6b, 0xdb, 0x30,
	0x14, 0xc6, 0xeb, 0xc6, 0x4b, 0xec, 0x97, 0xb4, 0x18, 0x9d, 0x44, 0x18, 0xcc, 0xeb, 0x60, 0xf8,
	0xb0, 0xda, 0x90, 0x9e, 0xb6, 0x5b, 0x4a, 0x33, 0x1a, 0x56, 0xb2, 0x61, 0xdc, 0x5e, 0x83, 0x92,
	0xbc, 0xc4, 0xa2, 0xb6, 0x2c, 0x22, 0x39, 0x2c, 0xff, 0xdc, 0xfe, 0xb6, 0x21, 0xb9, 0xae, 0xbb,
	0x1d, 0x72, 0x7b, 0xfe, 0xf8, 0x7e, 0x9f, 0x9e, 0x3f, 0x09, 0x02, 0x2e, 0x34, 0xee, 0xb7, 0x6c,
	0x8d, 0x2a, 0x96, 0xfb, 0x4a, 0x57, 0x04, 0x3a, 0x65, 0xfc, 0x63, 0xc7, 0x75, 0x5e, 0xaf, 0xe2,
	0x75, 0x55, 0x26, 0x05, 0xdf, 0x31, 0x5d, 0x25, 0x07, 0x29, 0xaf, 0xd9, 0x0e, 0x85, 0x4e, 0x64,
	0x51, 0xef, 0xb8, 0x50, 0x49, 0xc1, 0x45, 0xfd, 0xfb, 0x30, 0x49, 0xca, 0x6a, 0x83, 0x45, 0x22,
	0x58, 0x89, 0x4a, 0xb2, 0x35, 0x76, 0x53, 0x13, 0x7c, 0xf5, 0xc7, 0x85, 0xcb, 0x07, 0xe3, 0x9d,
	0xb7, 0x07, 0x10, 0x02, 0xae, 0x71, 0x51, 0x27, 0x74, 0x22, 0x3f, 0xb5, 0x33, 0xb9, 0x01, 0x57,
	0x1f, 0x25, 0xd2, 0xf3, 0xd0, 0x89, 0x2e, 0x27, 0x1f, 0xe2, 0x37, 0x0b, 0xfe, 0x4b, 0xc7, 0xd9,
	0x51, 0x62, 0x6a, 0xcd, 0xe4, 0x1b, 0xf8, 0xaf, 0xc7, 0xd1, 0x5e, 0xe8, 0x44, 0xc3, 0xc9, 0xfb,
	0xb8, 0x5b, 0xc0, 0x82, 0x0b, 0xd4, 0x8b, 0x56, 0x49, 0x3b, 0x3b, 0x09, 0x61, 0x94, 0x57, 0x4a,
	0x2f, 0xf9, 0x76, 0x69, 0x97, 0x71, 0xed, 0x32, 0x60, 0xb4, 0xf9, 0xd6, 0x10, 0x84, 0xc2, 0x00,
	0x05, 0x5b, 0x15, 0xb8, 0xa1, 0xef, 0x42, 0x27, 0xf2, 0xd2, 0xf6, 0x93, 0x7c, 0x84, 0x11, 0x97,
	0x4b, 0xb6, 0xd9, 0xec, 0x51, 0x29, 0x54, 0xb4, 0x1f, 0xf6, 0x22, 0x3f, 0x1d, 0x72, 0x39, 0x6d,
	0x25, 0x63, 0x91, 0xf9, 0x51, 0xb5, 0x26, 0x3a, 0xb0, 0xf1, 0x43, 0xa3, 0xbd, 0x98, 0x48, 0x00,
	0xbd, 0x52, 0xd7, 0xd4, 0x0b, 0x9d, 0xe8, 0x22, 0x35, 0x23, 0xf9, 0x0a, 0xee, 0x01, 0x75, 0x4e,
	0x7d, 0xfb, 0x2b, 0x9f, 0x4e, 0x94, 0xf0, 0x84, 0x3a, 0x7f, 0xe0, 0xe2, 0xf9, 0xfe, 0x2c, 0xb5,
	0x08, 0xb9, 0x85, 0x01, 0xab, 0x75, 0x95, 0x31, 0x49, 0xc1, 0xd2, 0x9f, 0x4f, 0xd0, 0xd3, 0xc6,
	0xf9, 0x12, 0xd0, 0x82, 0xe3, 0x2f, 0xe0, 0xb5, 0xb9, 0xa6, 0x1e, 0x89, 0xb8, 0x7f, 0xad, 0xa7,
	0xb9, 0x2b, 0x30, 0x5a, 0x53, 0xcf, 0x38, 0x81, 0xe1, 0x9b, 0x1c, 0x03, 0x68, 0x2c, 0xe5, 0xff,
	0x80, 0xd1, 0x1a, 0xe0, 0xea, 0x1a, 0x5c, 0x73, 0x77, 0xe4, 0x02, 0xfc, 0xc7, 0xc5, 0xdd, 0xec,
	0xfb, 0x7c, 0x31, 0xbb, 0x0b, 0xce, 0x88, 0x07, 0xee, 0xd3, 0x2c, 0xbb, 0x0f, 0x1c, 0x32, 0x02,
	0x6f, 0xfa, 0x98, 0xfd, 0x5c, 0x66, 0xd3, 0x5f, 0xc1, 0xf9, 0x6d, 0x1f, 0xdc, 0x82, 0x8b, 0xe7,
	0x55, 0xdf, 0xbe, 0xa3, 0x9b, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x45, 0x81, 0x47, 0x49, 0xb4,
	0x02, 0x00, 0x00,
}
