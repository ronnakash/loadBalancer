package main

import (
	"fmt"
	"net/http"
	"sync"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
	mu      		sync.Mutex
	conns  		 	[]int	
	algorithm		string
}

func NewLoadBalancer(config Config) *LoadBalancer {
	return &LoadBalancer{
		port:            	config.Port,
		roundRobinCount:	0,
		servers:        	config.Servers,
		algorithm: 			config.Algorithm,

	}
}


func (lb *LoadBalancer) getNextAvailableServer() Server {
	if (lb.algorithm == "round-robin"){
		return lb.getNextAvailableServerRoundRobin()
	}
	return lb.getLeastConnectedServer()
}


// getNextServerAddr returns the address of the next available
// server to send a request to, using round-robin algorithm
func (lb *LoadBalancer) getNextAvailableServerRoundRobin() Server {
	server := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) getLeastConnectedServer() (Server) {
    var server Server
    var minConns = -1
    for i, s := range lb.servers {
        if minConns == -1 || lb.conns[i] < minConns {
            server = s
            minConns = lb.conns[i]
        }
    }

    return server
}


func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	// could optionally log stuff about the request here!
	fmt.Printf("forwarding request to address %s\n", targetServer.Address())

	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
}