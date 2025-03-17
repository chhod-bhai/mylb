package app

import (
	"fmt"
	"net/http"

	"github.com/chhod-bhai/mylb/handler"
	"github.com/chhod-bhai/mylb/loadbalancer"
)

func Start() {
	serverStrs := []string{
		"http://localhost:8081",
		"http://localhost:8082",
		"http://localhost:8083",
		"http://localhost:8084",
	}

	// Initialise load balancer with serverpool
	lb := loadbalancer.New(serverStrs)

	go lb.PollHealth()

	httpServer := http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: http.HandlerFunc(handler.Serve(lb)),
	}

	if err := httpServer.ListenAndServe(); err != nil {
		fmt.Printf("[ERROR]: error listening on port %d - %v\n", 8080, err)
		return
	}
}
