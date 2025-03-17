package server

import (
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
	Url          *url.URL
	Alive        bool
	mux          sync.RWMutex
	ReverseProxy httputil.ReverseProxy
}
