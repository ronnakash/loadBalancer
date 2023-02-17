package main

import (
	// "fmt"
	"net/http"
	"net/http/httputil"
	"strings"

	// "net/url"
	"sync"
	// "strconv"
)


type Server interface {
	// Address returns the address with which to access the server
	Address() *ServerAddress
	// Address() ServerAddress

	// IsAlive returns true if the server is alive and able to serve requests
	IsAlive() bool

	//get number of connections to the server
	GetConnections() int

	//increment number of connections to the server
	IncrementConnections() 

	//decrement number of connections to the server
	DecrementConnections() 

	// Serve uses this server to process the request
	Serve(rw http.ResponseWriter, req *http.Request)

}


type SimpleServer struct {
	address 	ServerAddress
	proxy 		*httputil.ReverseProxy
	mutex 		sync.Mutex
	connections	int
}


type ServerParams struct {
	Address string
	Port string
}

func NewServerParams(address string) *ServerParams {
	parts := strings.Split(address, ":")
	return &ServerParams{
		Address: parts[0],
		Port: parts[1],
	}
}

func (s *SimpleServer) Address() *ServerAddress { return &s.address }

func (s *SimpleServer) IsAlive() bool { return true }

func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
	s.proxy.ServeHTTP(rw, req)
}

func (s *SimpleServer) GetConnections() int {
	return s.connections
}

func (s *SimpleServer) IncrementConnections() {
	s.mutex.Lock()
	s.connections++
	s.mutex.Unlock()
}

func (s *SimpleServer) DecrementConnections() {
	s.mutex.Lock()
	s.connections--
	s.mutex.Unlock()
}


func NewSimpleServer(params *ServerParams) *SimpleServer {
	address := NewServerAddress(params)
	return &SimpleServer{
		address: address,
		proxy:   httputil.NewSingleHostReverseProxy(address.Url()),
	}
}
