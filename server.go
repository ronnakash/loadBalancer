package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

)

// func main() {
//     http.HandleFunc("/", HelloServer)
//     http.ListenAndServe(":8080", nil)
// }

type Server interface {
	// Address returns the address with which to access the server
	Address() string

	// IsAlive returns true if the server is alive and able to serve requests
	IsAlive() bool

	// Serve uses this server to process the request
	Serve(rw http.ResponseWriter, req *http.Request)
}

// func HelloServer(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
// }

type SimpleServer struct {
	addr 	string
	port 	string
	proxy 	*httputil.ReverseProxy
}

type ServerParams struct {
	Address string
	Port string
}

func (s *SimpleServer) Address() string { return s.addr + ":" + s.port }

func (s *SimpleServer) IsAlive() bool { return true }

func (s *SimpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
    // fmt.Fprintf(rw, "Hello from %s:%s\n", s.addr , s.port);
	s.proxy.ServeHTTP(rw, req)
}

func NewSimpleServer(addr string, port string) *SimpleServer {
	serverUrl, err := url.Parse("http://" + addr + ":" + port)
	if(err != nil) {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return &SimpleServer{
		addr:  addr,
		port: port,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}

func newSimpleServer(params ServerParams) *SimpleServer {
	serverUrl, err := url.Parse("http://" + params.Address +":"+params.Port)
	if(err != nil) {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return &SimpleServer{
		addr:  params.Address,
		port: params.Port,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}