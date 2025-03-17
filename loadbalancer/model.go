package loadbalancer

import "github.com/chhod-bhai/mylb/serverpool"

type LoadBalancer struct {
	ServerPool *serverpool.ServerPool
}
