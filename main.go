package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	dns_proxy "github.com/binrick/go-dns-proxy-service/server"
	"github.com/ryboe/q"
)

type WireguardUIResolverService struct {
	InternalProxyService *dns_proxy.DNSProxyServer
}

var (
	dns_proxy_shutdown_channel   = make(chan bool, 1)
	TOGGLE_STATE_INTERVAL        = 5 * time.Second
	DNS_PROXY_PORT               = 8384
	DNS_PROXY_PROTOCOL           = `udp4`
	DNS_PROXY_UPSTREAM_RESOLVERS = []string{`8.8.8.8:53`, `1.1.1.1:53`}
	ProxyService                 = &dns_proxy.DNSProxyServer{
		Port:      DNS_PROXY_PORT,
		Proto:     DNS_PROXY_PROTOCOL,
		Upstreams: DNS_PROXY_UPSTREAM_RESOLVERS,
		Shutdown:  dns_proxy_shutdown_channel,
	}
	dns_proxy_server = &dns_proxy.DNSProxyServer{}
)

func jsonPrettyPrint(in string) string {
	var out bytes.Buffer
	err := json.Indent(&out, []byte(in), "", "  ")
	if err != nil {
		return in
	}
	return out.String()
}

func init() {
}

func stats() {
	for {
		//pp.Print(dns_proxy_server)
		q.Q(123, `abc`)
		q.Q(dns_proxy_server)
		dns_proxy_server.Stats()
		//		pp.Print(dns_proxy_server)
		/*
			server_s, _ := json.Marshal(dns_proxy_server)
			server_p := string(jsonPrettyPrint(string(server_s)))
			fmt.Print(server_p)
		*/
		time.Sleep(TOGGLE_STATE_INTERVAL)
	}
}

func main() {

	dns_proxy_server = ProxyService.New()
	dns_proxy_shutdown_channel <- false
	go stats()
	for {
		time.Sleep(TOGGLE_STATE_INTERVAL)
		fmt.Printf("Stopping DNS Proxy\n")
		dns_proxy_shutdown_channel <- true

		time.Sleep(TOGGLE_STATE_INTERVAL)
		fmt.Printf("Starting DNS Proxy\n")
		dns_proxy_shutdown_channel <- false
	}

}
