package rules

import (
	"errors"

	"github.com/v2ray/v2ray-core/app"
	"github.com/v2ray/v2ray-core/app/dns"
	"github.com/v2ray/v2ray-core/app/router"
	"github.com/v2ray/v2ray-core/common/log"
	v2net "github.com/v2ray/v2ray-core/common/net"
)

var (
	ErrorInvalidRule      = errors.New("Invalid Rule")
	ErrorNoRuleApplicable = errors.New("No rule applicable")
)

type Router struct {
	config    *RouterRuleConfig
	cache     *RoutingTable
	dnsServer dns.Server
}

func NewRouter(config *RouterRuleConfig, space app.Space) *Router {
	r := &Router{
		config: config,
		cache:  NewRoutingTable(),
	}
	space.InitializeApplication(func() error {
		if !space.HasApp(dns.APP_ID) {
			log.Error("DNS: Router is not found in the space.")
			return app.ErrorMissingApplication
		}
		r.dnsServer = space.GetApp(dns.APP_ID).(dns.Server)
		return nil
	})
	return r
}

func (this *Router) Release() {

}

// @Private
func (this *Router) ResolveIP(dest v2net.Destination) []v2net.Destination {
	ips := this.dnsServer.Get(dest.Address().Domain())
	if len(ips) == 0 {
		return nil
	}
	dests := make([]v2net.Destination, len(ips))
	for idx, ip := range ips {
		if dest.IsTCP() {
			dests[idx] = v2net.TCPDestination(v2net.IPAddress(ip), dest.Port())
		} else {
			dests[idx] = v2net.UDPDestination(v2net.IPAddress(ip), dest.Port())
		}
	}
	return dests
}

func (this *Router) takeDetourWithoutCache(dest v2net.Destination) (string, error) {
	for _, rule := range this.config.Rules {
		if rule.Apply(dest) {
			return rule.Tag, nil
		}
	}
	if this.config.DomainStrategy == UseIPIfNonMatch && dest.Address().IsDomain() {
		log.Info("Router: Looking up IP for ", dest)
		ipDests := this.ResolveIP(dest)
		if ipDests != nil {
			for _, ipDest := range ipDests {
				log.Info("Router: Trying IP ", ipDest)
				for _, rule := range this.config.Rules {
					if rule.Apply(ipDest) {
						return rule.Tag, nil
					}
				}
			}
		}
	}

	return "", ErrorNoRuleApplicable
}

func (this *Router) TakeDetour(dest v2net.Destination) (string, error) {
	destStr := dest.String()
	found, tag, err := this.cache.Get(destStr)
	if !found {
		tag, err := this.takeDetourWithoutCache(dest)
		this.cache.Set(destStr, tag, err)
		return tag, err
	}
	return tag, err
}

type RouterFactory struct {
}

func (this *RouterFactory) Create(rawConfig interface{}, space app.Space) (router.Router, error) {
	return NewRouter(rawConfig.(*RouterRuleConfig), space), nil
}

func init() {
	router.RegisterRouter("rules", &RouterFactory{})
}
