package serverpool

import (
	"fmt"
	"sync/atomic"

	"github.com/chhod-bhai/mylb/server"
)

func New(urls []string) *ServerPool {
	return &ServerPool{
		Servers: server.NewServerListFromStrs(urls),
		current: 0,
	}
}

func (sp *ServerPool) Healthcheck() {
	for _, s := range sp.Servers {
		status := "down"
		isUp := s.CheckServerStatus()
		s.SetAlive(isUp)
		if isUp {
			status = "up"
		}
		fmt.Printf("[INFO]: Server (%s) is (%s)", s.Url.Host, status)
	}
}

func (s *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.Servers)))
}

func (s *ServerPool) GetNextPeer() *server.Server {
	next := s.NextIndex()
	serversLen := len(s.Servers)
	end := next + serversLen
	for i := next; i < end; i++ {
		idx := i % serversLen
		selectedServer := s.Servers[idx]
		if selectedServer.IsAlive() {
			if i != next {
				atomic.StoreUint64(&s.current, uint64(idx))
			}
			return selectedServer
		}
	}
	return nil
}
