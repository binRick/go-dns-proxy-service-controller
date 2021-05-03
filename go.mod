module github.com/binrick/go-dns-proxy-service-controller

go 1.16

require (
	github.com/binrick/go-dns-proxy-service v0.0.0-20210503182841-74f0bdb2918e
	github.com/kr/text v0.2.0 // indirect
	github.com/ryboe/q v1.0.13
	golang.org/x/net v0.0.0-20210503060351-7fd8e65b6420 // indirect
	golang.org/x/sys v0.0.0-20210503173754-0981d6026fa6 // indirect
)

replace github.com/binrick/go-dns-proxy-service => ../go-dns-proxy-service
