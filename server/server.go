package server

import (
	"fmt"
	"net"
	"net/http/httputil"
	"net/url"
	"time"
)

const (
	Retry int = iota
	Attempt
)

func NewServer(urlStr string) *Server {
	endpoint, err := url.ParseRequestURI(urlStr)
	if err != nil {
		fmt.Printf("[ERROR]: error parsing uri string - %v\n", err)
		return nil
	}

	server := &Server{
		Url:   endpoint,
		Alive: true,
	}
	server.ReverseProxy = *httputil.NewSingleHostReverseProxy(server.Url)

	return server

}

func NewServerListFromStrs(urls []string) (servers []*Server) {
	for _, serverStr := range urls {
		if server := NewServer(serverStr); server != nil {
			servers = append(servers, server)
		}
	}
	return
}

func (s *Server) IsAlive() (isAlive bool) {
	s.mux.RLock()
	isAlive = s.Alive
	s.mux.RUnlock()
	return
}

func (s *Server) SetAlive(isAlive bool) {
	s.mux.Lock()
	s.Alive = isAlive
	s.mux.Unlock()
}

func (s *Server) CheckServerStatus() bool {
	conn, err := net.DialTimeout("tcp", s.Url.Host, 2*time.Second)
	if err != nil {
		fmt.Printf("[ERROR]: Error checking server status for %s", s.Url.Host)
		return false
	}
	_ = conn.Close()
	return true
}
