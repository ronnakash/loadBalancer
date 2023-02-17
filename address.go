package main

import (
	"fmt"
	"net/url"
	"strconv"
	// "strconv"
)

type ServerAddress interface {
	
	String() string

	Url() *url.URL

	BaseAddress() string

	Port() int

	PortString() string
}




type ServerAddressImpl struct {
	address	string
	port	int
}


func (s *ServerAddressImpl) String() string {
    return fmt.Sprintf("%s:%d", s.address, s.port)
}

func (s *ServerAddressImpl) Url() *url.URL {
    return &url.URL{
        Scheme: "http",
        Host:   fmt.Sprintf("%s:%d", s.address, s.port),
    }
}

func (s *ServerAddressImpl) BaseAddress() string {
    return s.address
}

func (s *ServerAddressImpl) Port() int {
    return s.port
}

func (s *ServerAddressImpl) PortString() string {
    return strconv.Itoa(s.port)
}

func NewServerAddress(sp *ServerParams) *ServerAddressImpl {
	port, err := strconv.Atoi(sp.port)
	if err != nil {
		return nil
	}
	return &ServerAddressImpl{
		address: sp.address,
		port:    port,
	}
}