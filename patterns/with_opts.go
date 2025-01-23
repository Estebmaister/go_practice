package server

import "fmt"

type Server struct {
	opts
}

type opts struct {
	Host    string
	Port    int
	TLS     bool
	MaxConn int
}

type Option func(*opts)

func WithHost(host string) Option {
	return func(o *opts) {
		o.Host = host
	}
}

func WithPort(port int) Option {
	return func(o *opts) {
		o.Port = port
	}
}

func WithMaxConn(maxConn int) Option {
	return func(o *opts) {
		o.MaxConn = maxConn
	}
}

func WithTLS(tls bool) Option {
	return func(o *opts) {
		o.TLS = tls
	}
}

func New(options ...Option) *Server {
	o := &opts{
		Host:    "localhost",
		Port:    8080,
		TLS:     false,
		MaxConn: 1000,
	}
	for _, opt := range options {
		opt(o)
	}
	return &Server{*o}
}

func (s *Server) Run() {
	fmt.Println("Server is running on port", s.Port)
	fmt.Printf("%#v", *s)
}
