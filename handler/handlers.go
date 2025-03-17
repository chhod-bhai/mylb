package handler

import (
	"net/http"

	"github.com/chhod-bhai/mylb/loadbalancer"
)

func Serve(lb *loadbalancer.LoadBalancer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		server := lb.ServerPool.GetNextPeer()
		if server == nil {
			http.Error(w, "Service not available", http.StatusServiceUnavailable)
			return
		}
		server.ReverseProxy.ServeHTTP(w, r)
	}
}
