// Use a custom Config struct

package Approach2

import "time"

type Config struct {
	Host    string
	Port    int
	Timeout time.Duration
	MaxConn int
}

type Server struct {
	cfg Config
}

func New(cfg Config) *Server {
	return &Server{cfg}
}

func (s *Server) Start() error {
	return nil
}
