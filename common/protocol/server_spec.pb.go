package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import v2ray_core_common_net "v2ray.com/core/common/net"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ServerEndpoint struct {
	Address *v2ray_core_common_net.IPOrDomain `protobuf:"bytes,1,opt,name=address" json:"address,omitempty"`
	Port    uint32                            `protobuf:"varint,2,opt,name=port" json:"port,omitempty"`
	User    []*User                           `protobuf:"bytes,3,rep,name=user" json:"user,omitempty"`
}

func (m *ServerEndpoint) Reset()                    { *m = ServerEndpoint{} }
func (m *ServerEndpoint) String() string            { return proto.CompactTextString(m) }
func (*ServerEndpoint) ProtoMessage()               {}
func (*ServerEndpoint) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

func (m *ServerEndpoint) GetAddress() *v2ray_core_common_net.IPOrDomain {
	if m != nil {
		return m.Address
	}
	return nil
}

func (m *ServerEndpoint) GetPort() uint32 {
	if m != nil {
		return m.Port
	}
	return 0
}

func (m *ServerEndpoint) GetUser() []*User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*ServerEndpoint)(nil), "v2ray.core.common.protocol.ServerEndpoint")
}

func init() { proto.RegisterFile("v2ray.com/core/common/protocol/server_spec.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x8f, 0x3f, 0x4f, 0xf3, 0x30,
	0x10, 0xc6, 0x95, 0xb6, 0x7a, 0x5f, 0xe4, 0x0a, 0x90, 0x3c, 0x45, 0x19, 0x50, 0x60, 0x21, 0x2c,
	0x67, 0x14, 0xd8, 0xd8, 0x5a, 0x18, 0x98, 0x88, 0x52, 0x81, 0x10, 0x0b, 0x0a, 0xce, 0x0d, 0x95,
	0xb0, 0xcf, 0x3a, 0x9b, 0x4a, 0xfd, 0x24, 0x7c, 0x07, 0x3e, 0x25, 0x8a, 0xdd, 0x4c, 0xfc, 0xdb,
	0xac, 0xbb, 0xdf, 0xef, 0x9e, 0xc7, 0xe2, 0x7c, 0x53, 0x73, 0xb7, 0x05, 0x4d, 0x46, 0x69, 0x62,
	0x54, 0x9a, 0x8c, 0x21, 0xab, 0x1c, 0x53, 0x20, 0x4d, 0xaf, 0xca, 0x23, 0x6f, 0x90, 0x9f, 0xbd,
	0x43, 0x0d, 0x71, 0x28, 0x8b, 0xd1, 0x60, 0x84, 0x44, 0xc3, 0x48, 0x17, 0xa7, 0xdf, 0x5f, 0xb3,
	0x18, 0x54, 0xd7, 0xf7, 0x8c, 0xde, 0x27, 0xb6, 0x38, 0xfb, 0x23, 0xf6, 0xcd, 0x23, 0x27, 0xf4,
	0xe4, 0x3d, 0x13, 0x07, 0xab, 0xd8, 0xe2, 0xc6, 0xf6, 0x8e, 0xd6, 0x36, 0xc8, 0x2b, 0xf1, 0x7f,
	0x77, 0x2e, 0xcf, 0xca, 0xac, 0x9a, 0xd7, 0xc7, 0xf0, 0xb5, 0x94, 0xc5, 0x00, 0xb7, 0xcd, 0x1d,
	0x5f, 0x93, 0xe9, 0xd6, 0xb6, 0x1d, 0x0d, 0x29, 0xc5, 0xcc, 0x11, 0x87, 0x7c, 0x52, 0x66, 0xd5,
	0x7e, 0x1b, 0xdf, 0xf2, 0x52, 0xcc, 0x86, 0xc4, 0x7c, 0x5a, 0x4e, 0xab, 0x79, 0x5d, 0xc2, 0xcf,
	0x5f, 0x84, 0x7b, 0x8f, 0xdc, 0x46, 0x7a, 0xf1, 0x28, 0x8e, 0x34, 0x99, 0x5f, 0xe0, 0xc5, 0x61,
	0x2a, 0xbe, 0x72, 0xa8, 0x9b, 0x61, 0xf6, 0xb4, 0x37, 0xae, 0x3e, 0x26, 0xc5, 0x43, 0xdd, 0x76,
	0x5b, 0x58, 0x0e, 0xde, 0x32, 0x79, 0xcd, 0x6e, 0xf9, 0xf2, 0x2f, 0x62, 0x17, 0x9f, 0x01, 0x00,
	0x00, 0xff, 0xff, 0x90, 0x81, 0xc4, 0x33, 0x9e, 0x01, 0x00, 0x00,
}
