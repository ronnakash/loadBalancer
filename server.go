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

type simpleServer struct {
	addr string
	port string
	proxy *httputil.ReverseProxy
}

func (s *simpleServer) Address() string { return s.addr }

func (s *simpleServer) IsAlive() bool { return true }

func (s *simpleServer) Serve(rw http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(rw, "Hello from %s%s\n", s.addr , s.port);
}

func newSimpleServer(addr string, port string) *simpleServer {
	serverUrl, err := url.Parse(addr+":"+addr)
	if(err != nil) {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	return &simpleServer{
		addr:  addr,
		port: port,
		proxy: httputil.NewSingleHostReverseProxy(serverUrl),
	}
}