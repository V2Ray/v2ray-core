package net_test

import (
	"net"
	"testing"

	v2net "github.com/v2ray/v2ray-core/common/net"
	"github.com/v2ray/v2ray-core/testing/assert"
)

func TestIPv4Address(t *testing.T) {
	assert := assert.On(t)

	ip := []byte{byte(1), byte(2), byte(3), byte(4)}
	addr := v2net.IPAddress(ip)

	assert.Address(addr).IsIPv4()
	assert.Address(addr).IsNotIPv6()
	assert.Address(addr).IsNotDomain()
	assert.Bytes(addr.IP()).Equals(ip)
	assert.Address(addr).EqualsString("1.2.3.4")
}

func TestIPv6Address(t *testing.T) {
	assert := assert.On(t)

	ip := []byte{
		byte(1), byte(2), byte(3), byte(4),
		byte(1), byte(2), byte(3), byte(4),
		byte(1), byte(2), byte(3), byte(4),
		byte(1), byte(2), byte(3), byte(4),
	}
	addr := v2net.IPAddress(ip)

	assert.Address(addr).IsIPv6()
	assert.Address(addr).IsNotIPv4()
	assert.Address(addr).IsNotDomain()
	assert.IP(addr.IP()).Equals(net.IP(ip))
	assert.Address(addr).EqualsString("[102:304:102:304:102:304:102:304]")
}

func TestIPv4Asv6(t *testing.T) {
	assert := assert.On(t)
	ip := []byte{
		byte(0), byte(0), byte(0), byte(0),
		byte(0), byte(0), byte(0), byte(0),
		byte(0), byte(0), byte(255), byte(255),
		byte(1), byte(2), byte(3), byte(4),
	}
	addr := v2net.IPAddress(ip)
	assert.Address(addr).EqualsString("1.2.3.4")
}

func TestDomainAddress(t *testing.T) {
	assert := assert.On(t)

	domain := "v2ray.com"
	addr := v2net.DomainAddress(domain)

	assert.Address(addr).IsDomain()
	assert.Address(addr).IsNotIPv6()
	assert.Address(addr).IsNotIPv4()
	assert.String(addr.Domain()).Equals(domain)
	assert.Address(addr).EqualsString("v2ray.com")
}

func TestNetIPv4Address(t *testing.T) {
	assert := assert.On(t)

	ip := net.IPv4(1, 2, 3, 4)
	addr := v2net.IPAddress(ip)
	assert.Address(addr).IsIPv4()
	assert.Address(addr).EqualsString("1.2.3.4")
}

func TestIPv4AddressEquals(t *testing.T) {
	assert := assert.On(t)

	addr := v2net.IPAddress([]byte{1, 2, 3, 4})
	assert.Address(addr).NotEquals(nil)

	addr2 := v2net.IPAddress([]byte{1, 2, 3, 4})
	assert.Address(addr).Equals(addr2)

	addr3 := v2net.IPAddress([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6})
	assert.Address(addr).NotEquals(addr3)

	addr4 := v2net.IPAddress([]byte{1, 2, 3, 5})
	assert.Address(addr).NotEquals(addr4)
}

func TestIPv6AddressEquals(t *testing.T) {
	assert := assert.On(t)

	addr := v2net.IPAddress([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6})
	assert.Bool(addr.Equals(nil)).IsFalse()

	addr2 := v2net.IPAddress([]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6})
	assert.Bool(addr.Equals(addr2)).IsTrue()

	addr3 := v2net.IPAddress([]byte{1, 2, 3, 4})
	assert.Bool(addr.Equals(addr3)).IsFalse()

	addr4 := v2net.IPAddress([]byte{1, 3, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6})
	assert.Bool(addr.Equals(addr4)).IsFalse()
}

func TestDomainAddressEquals(t *testing.T) {
	assert := assert.On(t)

	addr := v2net.DomainAddress("v2ray.com")
	assert.Bool(addr.Equals(nil)).IsFalse()

	addr2 := v2net.DomainAddress("v2ray.com")
	assert.Bool(addr.Equals(addr2)).IsTrue()

	addr3 := v2net.DomainAddress("www.v2ray.com")
	assert.Bool(addr.Equals(addr3)).IsFalse()

	addr4 := v2net.IPAddress([]byte{1, 3, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6})
	assert.Bool(addr.Equals(addr4)).IsFalse()
}
