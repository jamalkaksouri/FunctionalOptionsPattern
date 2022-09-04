// Functional Options Pattern

package Approach3

import (
	"time"
)

type Server struct {
	host    string
	port    int
	timeout time.Duration
	maxConn int
}

type Options func(*Server)

func New(options ...Options) *Server {
	svr := &Server{}
	for _, s := range options {
		s(svr)
	}
	return svr
}

func (server *Server) Start() error {
	return nil
}

func WithHost(host string) Options {
	return func(server *Server) {
		server.host = host
	}
}

func WithPort(port int) Options {
	return func(server *Server) {
		server.port = port
	}
}

func WithTimeout(timeout time.Duration) Options {
	return func(server *Server) {
		server.timeout = timeout
	}
}

func WithMaxConn(maxConn int) Options {
	return func(server *Server) {
		server.maxConn = maxConn
	}
}
