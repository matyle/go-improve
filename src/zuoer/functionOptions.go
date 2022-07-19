package zuoer

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Option func(*Server)

type Server struct {
	Addr     string
	Port     int
	Protocal string
	Timeout  time.Duration
	Maxconn  int
	Tls      *tls.Config
}

func Protocal(p string) Option {
	return func(s *Server) {
		s.Protocal = p
	}

}

func Timeout(t time.Duration) Option {
	return func(s *Server) {
		s.Timeout = t
	}

}

func Maxconn(m int) Option {
	return func(s *Server) {
		s.Maxconn = m
	}
}

func Tls(t *tls.Config) Option {
	return func(s *Server) {
		s.Tls = t
	}
}

func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {
	ser := Server{
		Addr:     addr,
		Port:     port,
		Protocal: "tcp",
		Timeout:  30 * time.Millisecond,
		Maxconn:  1000,
		Tls:      nil,
	}

	//...

	for _, option := range options {
		option(&ser)
	}

	return &ser, nil
}

func InitOptions() {
	s1, _ := NewServer("localhost", 80)
	s2, _ := NewServer("localhost", 2047, Protocal("udp"))
	s3, _ := NewServer("localhost", 2047, Protocal("udp"), Timeout(300*time.Second), Maxconn(3000))
	s4, _ := NewServer("localhost", 2047, Protocal("udp"), Timeout(300*time.Second), Maxconn(3000), Tls(&tls.Config{}))

	fmt.Printf("s1:%v\n", s1)
	fmt.Printf("s2:%v\n", s2)
	fmt.Printf("s3:%v\n", s3)
	fmt.Printf("s4:%v\n", s4)

}
