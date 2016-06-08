package dns_test

import (
	"net"
	"testing"

	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/app/dispatcher"
	dispatchers "github.com/v2ray/v2ray-core/app/dispatcher/impl"
	. "github.com/v2ray/v2ray-core/app/dns"
	"github.com/v2ray/v2ray-core/app/proxyman"
	v2net "github.com/v2ray/v2ray-core/common/net"
	"github.com/v2ray/v2ray-core/proxy"
	"github.com/v2ray/v2ray-core/proxy/freedom"
	"github.com/v2ray/v2ray-core/testing/assert"
)

func TestDnsAdd(t *testing.T) {
	assert := assert.On(t)

	space := app.NewSpace()

	outboundHandlerManager := proxyman.NewDefaultOutboundHandlerManager()
	outboundHandlerManager.SetDefaultHandler(freedom.NewFreedomConnection(&freedom.Config{}, space, &proxy.OutboundHandlerMeta{Address: v2net.AnyIP}))
	space.BindApp(proxyman.APP_ID_OUTBOUND_MANAGER, outboundHandlerManager)
	space.BindApp(dispatcher.APP_ID, dispatchers.NewDefaultDispatcher(space))

	domain := "local.v2ray.com"
	server := NewCacheServer(space, &Config{
		NameServers: []v2net.Destination{
			v2net.UDPDestination(v2net.IPAddress([]byte{8, 8, 8, 8}), v2net.Port(53)),
		},
	})
	space.BindApp(APP_ID, server)
	space.Initialize()

	ips := server.Get(domain)
	assert.Int(len(ips)).Equals(1)
	assert.IP(ips[0].To4()).Equals(net.IP([]byte{127, 0, 0, 1}))
}
