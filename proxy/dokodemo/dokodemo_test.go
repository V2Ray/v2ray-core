package dokodemo_test

import (
	"net"
	"testing"

	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/app/dispatcher"
	dispatchers "github.com/v2ray/v2ray-core/app/dispatcher/impl"
	"github.com/v2ray/v2ray-core/app/proxyman"
	v2net "github.com/v2ray/v2ray-core/common/net"
	v2nettesting "github.com/v2ray/v2ray-core/common/net/testing"
	"github.com/v2ray/v2ray-core/proxy"
	. "github.com/v2ray/v2ray-core/proxy/dokodemo"
	"github.com/v2ray/v2ray-core/proxy/freedom"
	"github.com/v2ray/v2ray-core/testing/assert"
	"github.com/v2ray/v2ray-core/testing/servers/tcp"
	"github.com/v2ray/v2ray-core/testing/servers/udp"
)

func TestDokodemoTCP(t *testing.T) {
	assert := assert.On(t)

	tcpServer := &tcp.Server{
		Port: v2nettesting.PickPort(),
		MsgProcessor: func(data []byte) []byte {
			buffer := make([]byte, 0, 2048)
			buffer = append(buffer, []byte("Processed: ")...)
			buffer = append(buffer, data...)
			return buffer
		},
	}
	_, err := tcpServer.Start()
	assert.Error(err).IsNil()

	defer tcpServer.Close()

	space := app.NewSpace()
	space.BindApp(dispatcher.APP_ID, dispatchers.NewDefaultDispatcher(space))
	ohm := proxyman.NewDefaultOutboundHandlerManager()
	ohm.SetDefaultHandler(
		freedom.NewFreedomConnection(
			&freedom.Config{}, space, &proxy.OutboundHandlerMeta{Address: v2net.LocalHostIP}))
	space.BindApp(proxyman.APP_ID_OUTBOUND_MANAGER, ohm)

	data2Send := "Data to be sent to remote."

	port := v2nettesting.PickPort()
	dokodemo := NewDokodemoDoor(&Config{
		Address: v2net.LocalHostIP,
		Port:    tcpServer.Port,
		Network: v2net.TCPNetwork.AsList(),
		Timeout: 600,
	}, space, &proxy.InboundHandlerMeta{Address: v2net.LocalHostIP, Port: port})
	defer dokodemo.Close()

	assert.Error(space.Initialize()).IsNil()

	err = dokodemo.Start()
	assert.Error(err).IsNil()
	assert.Port(port).Equals(dokodemo.Port())

	tcpClient, err := net.DialTCP("tcp", nil, &net.TCPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(port),
		Zone: "",
	})
	assert.Error(err).IsNil()

	tcpClient.Write([]byte(data2Send))
	tcpClient.CloseWrite()

	response := make([]byte, 1024)
	nBytes, err := tcpClient.Read(response)
	assert.Error(err).IsNil()
	tcpClient.Close()

	assert.String("Processed: " + data2Send).Equals(string(response[:nBytes]))
}

func TestDokodemoUDP(t *testing.T) {
	assert := assert.On(t)

	udpServer := &udp.Server{
		Port: v2nettesting.PickPort(),
		MsgProcessor: func(data []byte) []byte {
			buffer := make([]byte, 0, 2048)
			buffer = append(buffer, []byte("Processed: ")...)
			buffer = append(buffer, data...)
			return buffer
		},
	}
	_, err := udpServer.Start()
	assert.Error(err).IsNil()

	defer udpServer.Close()

	space := app.NewSpace()
	space.BindApp(dispatcher.APP_ID, dispatchers.NewDefaultDispatcher(space))
	ohm := proxyman.NewDefaultOutboundHandlerManager()
	ohm.SetDefaultHandler(
		freedom.NewFreedomConnection(
			&freedom.Config{}, space, &proxy.OutboundHandlerMeta{Address: v2net.AnyIP}))
	space.BindApp(proxyman.APP_ID_OUTBOUND_MANAGER, ohm)

	data2Send := "Data to be sent to remote."

	port := v2nettesting.PickPort()
	dokodemo := NewDokodemoDoor(&Config{
		Address: v2net.LocalHostIP,
		Port:    udpServer.Port,
		Network: v2net.UDPNetwork.AsList(),
		Timeout: 600,
	}, space, &proxy.InboundHandlerMeta{Address: v2net.LocalHostIP, Port: port})
	defer dokodemo.Close()

	assert.Error(space.Initialize()).IsNil()

	err = dokodemo.Start()
	assert.Error(err).IsNil()
	assert.Port(port).Equals(dokodemo.Port())

	udpClient, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   []byte{127, 0, 0, 1},
		Port: int(port),
		Zone: "",
	})
	assert.Error(err).IsNil()
	defer udpClient.Close()

	udpClient.Write([]byte(data2Send))

	response := make([]byte, 1024)
	nBytes, addr, err := udpClient.ReadFromUDP(response)
	assert.Error(err).IsNil()
	assert.IP(addr.IP).Equals(v2net.LocalHostIP.IP())
	assert.Bytes(response[:nBytes]).Equals([]byte("Processed: " + data2Send))
}
