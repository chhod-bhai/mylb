package serverpool

import (
	"github.com/chhod-bhai/mylb/server"
)

type ServerPool struct {
	Servers []*server.Server
	current uint64
}
