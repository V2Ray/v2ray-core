// +build !linux

package udp

import (
	"v2ray.com/core/common/net"
)

func SetOriginalDestOptions(fd int) error {
	return nil
}

func RetrieveOriginalDest(oob []byte) net.Destination {
	return net.Destination{}
}

func ReadUDPMsg(conn *net.UDPConn, payload []byte, oob []byte) (int, int, int, *net.UDPAddr, error) {
	nBytes, addr, err := conn.ReadFromUDP(payload)
	return nBytes, 0, 0, addr, err
}
