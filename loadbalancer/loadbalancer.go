package loadbalancer

import (
	"fmt"
	"time"

	"github.com/chhod-bhai/mylb/serverpool"
)

func New(serverUrls []string) *LoadBalancer {
	return &LoadBalancer{
		ServerPool: serverpool.New(serverUrls),
	}
}

func (lb *LoadBalancer) PollHealth() {
	t := time.NewTicker(time.Second * 20)
	for range t.C {
		fmt.Println("[INFO]: Starting health check...")
		lb.ServerPool.Healthcheck()
		fmt.Println("[INFO]: Health check completed")
	}
}
