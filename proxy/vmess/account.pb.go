package vmess

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
	protocol "v2ray.com/core/common/protocol"
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

type Account struct {
	// ID of the account, in the form of a UUID, e.g., "66ad4540-b58c-4ad2-9926-ea63445a9b57".
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Number of alternative IDs. Client and server must share the same number.
	AlterId uint32 `protobuf:"varint,2,opt,name=alter_id,json=alterId,proto3" json:"alter_id,omitempty"`
	// Security settings. Only applies to client side.
	SecuritySettings     *protocol.SecurityConfig `protobuf:"bytes,3,opt,name=security_settings,json=securitySettings,proto3" json:"security_settings,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *Account) Reset()         { *m = Account{} }
func (m *Account) String() string { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()    {}
func (*Account) Descriptor() ([]byte, []int) {
	return fileDescriptor_d65dee31e5abbda0, []int{0}
}
func (m *Account) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Account.Unmarshal(m, b)
}
func (m *Account) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Account.Marshal(b, m, deterministic)
}
func (m *Account) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Account.Merge(m, src)
}
func (m *Account) XXX_Size() int {
	return xxx_messageInfo_Account.Size(m)
}
func (m *Account) XXX_DiscardUnknown() {
	xxx_messageInfo_Account.DiscardUnknown(m)
}

var xxx_messageInfo_Account proto.InternalMessageInfo

func (m *Account) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Account) GetAlterId() uint32 {
	if m != nil {
		return m.AlterId
	}
	return 0
}

func (m *Account) GetSecuritySettings() *protocol.SecurityConfig {
	if m != nil {
		return m.SecuritySettings
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "v2ray.core.proxy.vmess.Account")
}

func init() {
	proto.RegisterFile("v2ray.com/core/proxy/vmess/account.proto", fileDescriptor_d65dee31e5abbda0)
}

var fileDescriptor_d65dee31e5abbda0 = []byte{
	// 244 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x8f, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xc9, 0x8a, 0x56, 0x23, 0x8a, 0xee, 0xa1, 0xac, 0x3d, 0x2d, 0x9e, 0x16, 0x91, 0x09,
	0xac, 0x77, 0x41, 0x7b, 0xf2, 0x56, 0xb6, 0x50, 0xc1, 0x4b, 0x89, 0x49, 0xac, 0x81, 0x66, 0xa7,
	0x4c, 0xd2, 0x62, 0x1e, 0xc2, 0x17, 0xf1, 0x29, 0xa5, 0xd9, 0x5d, 0x10, 0xe9, 0x2d, 0x61, 0xbe,
	0xff, 0xfb, 0x67, 0x78, 0xb5, 0xab, 0x49, 0x46, 0x50, 0xe8, 0x84, 0x42, 0x32, 0x62, 0x43, 0xf8,
	0x15, 0xc5, 0xce, 0x19, 0xef, 0x85, 0x54, 0x0a, 0xb7, 0x6d, 0x80, 0x0d, 0x61, 0xc0, 0x7c, 0x3c,
	0x90, 0x64, 0x20, 0x51, 0x90, 0xa8, 0xc9, 0xfd, 0x3f, 0x83, 0x42, 0xe7, 0xb0, 0x15, 0x29, 0xa4,
	0x70, 0x2d, 0x3e, 0x8d, 0xd4, 0x86, 0x7c, 0x67, 0xb9, 0xfd, 0x66, 0x7c, 0xf4, 0xd4, 0x79, 0xf3,
	0x4b, 0x9e, 0x59, 0x5d, 0xb0, 0x92, 0x55, 0x67, 0x4d, 0x66, 0x75, 0x7e, 0xc3, 0x4f, 0xe5, 0x3a,
	0x18, 0x5a, 0x5a, 0x5d, 0x64, 0x25, 0xab, 0x2e, 0x9a, 0x51, 0xfa, 0xbf, 0xe8, 0xfc, 0x95, 0x5f,
	0x7b, 0xa3, 0xb6, 0x64, 0x43, 0x5c, 0x7a, 0x13, 0x82, 0x6d, 0x57, 0xbe, 0x38, 0x2a, 0x59, 0x75,
	0x5e, 0xdf, 0xc1, 0x9f, 0xc5, 0xba, 0x72, 0x18, 0xca, 0x61, 0xde, 0x87, 0xa6, 0xd8, 0x7e, 0xd8,
	0x55, 0x73, 0x35, 0x48, 0xe6, 0xbd, 0xe3, 0xf9, 0x91, 0x4f, 0x14, 0x3a, 0x38, 0x7c, 0xdb, 0x8c,
	0xbd, 0x1d, 0xa7, 0xc7, 0x4f, 0x36, 0x5e, 0xd4, 0x8d, 0x8c, 0x30, 0xdd, 0x13, 0xb3, 0x44, 0x2c,
	0xf6, 0x83, 0xf7, 0x93, 0x54, 0xf5, 0xf0, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x82, 0xf9, 0x1b, 0xdb,
	0x48, 0x01, 0x00, 0x00,
}
