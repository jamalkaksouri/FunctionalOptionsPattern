// Declare a new constructor for each configuration option

package Approache1

import (
	"time"
)

type Server struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

func NewWithTimeout(host string, port int, timeout time.Duration) *Server {
	return &Server{
		Host:    host,
		Port:    port,
		Timeout: timeout,
	}
}

func NewWithTimeoutAndMaxConn(host string, port int, timeout time.Duration, maxConn int) *Server {
	return &Server{
		Host:    host,
		Port:    port,
		Timeout: timeout,
		MaxConn: maxConn,
	}
}

func New(host string, port int) *Server {
	return &Server{
		Host: host,
		Port: port,
	}
}

func (server *Server) Start() error {
	return nil
}
