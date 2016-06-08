package point

import (
	"github.com/v2ray/v2ray-core/app/dns"
	"github.com/v2ray/v2ray-core/app/router"
	"github.com/v2ray/v2ray-core/common/log"
	v2net "github.com/v2ray/v2ray-core/common/net"
	"github.com/v2ray/v2ray-core/transport"
)

type InboundConnectionConfig struct {
	Port     v2net.Port
	ListenOn v2net.Address
	Protocol string
	Settings []byte
}

type OutboundConnectionConfig struct {
	Protocol    string
	SendThrough v2net.Address
	Settings    []byte
}

type LogConfig struct {
	AccessLog string
	ErrorLog  string
	LogLevel  log.LogLevel
}

const (
	AllocationStrategyAlways   = "always"
	AllocationStrategyRandom   = "random"
	AllocationStrategyExternal = "external"
)

type InboundDetourAllocationConfig struct {
	Strategy    string // Allocation strategy of this inbound detour.
	Concurrency int    // Number of handlers (ports) running in parallel.
	Refresh     int    // Number of minutes before a handler is regenerated.
}

type InboundDetourConfig struct {
	Protocol   string
	PortRange  v2net.PortRange
	ListenOn   v2net.Address
	Tag        string
	Allocation *InboundDetourAllocationConfig
	Settings   []byte
}

type OutboundDetourConfig struct {
	Protocol    string
	SendThrough v2net.Address
	Tag         string
	Settings    []byte
}

type Config struct {
	Port            v2net.Port
	LogConfig       *LogConfig
	RouterConfig    *router.Config
	DNSConfig       *dns.Config
	InboundConfig   *InboundConnectionConfig
	OutboundConfig  *OutboundConnectionConfig
	InboundDetours  []*InboundDetourConfig
	OutboundDetours []*OutboundDetourConfig
	TransportConfig *transport.Config
}

type ConfigLoader func(init string) (*Config, error)

var (
	configLoader ConfigLoader
)

func LoadConfig(init string) (*Config, error) {
	if configLoader == nil {
		return nil, ErrorBadConfiguration
	}
	return configLoader(init)
}
