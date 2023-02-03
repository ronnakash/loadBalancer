package main

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	port            string
	roundRobinCount int
	servers         []Server
	algorithm		string
}

func NewLoadBalancer(config Config) *LoadBalancer {
	var serverParams []ServerParams = config.Servers
	var servers []Server

	// Initialize each server
    for i := range serverParams {
		fmt.Printf("server %d: %s:%s\n", i, serverParams[i].Address, serverParams[i].Port )
        servers = append(servers, NewSimpleServer(serverParams[i]))
    }

	return &LoadBalancer{
		port:            	config.Port,
		roundRobinCount:	0,
		servers:        	servers,
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
    for _, s := range lb.servers {
		conns:= s.GetConnections()
        if minConns == -1 || conns < minConns {
            server = s
            minConns = conns
        }
    }

    return server
}


func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	// could optionally log stuff about the request here!
	fmt.Printf("forwarding request to address %s\n", targetServer.Address())
	targetServer.IncrementConnections()
	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
	// fmt.Printf("done processing request at %s\n", targetServer.Address())
	targetServer.DecrementConnections()
}

func (lb *LoadBalancer) ChangeAlgorithm(algo string) {
	if algo == "round-robin" || algo == "least-connections" {
		lb.algorithm = algo
	} else {
		fmt.Printf("No algorithm %s \n", algo)
	}
}


func (lb *LoadBalancer) AddServer(params []string) {
	//TODO: ensure params are of length 2, contain valid address and port
	newServer := NewSimpleServer(ServerParams {Address : params[0], Port : params[1]})
	lb.servers = append(lb.servers, newServer)
}

func (lb *LoadBalancer) RemoveServer(params []string) {
	var removeIndex int
	toRemove := NewSimpleServer(ServerParams {Address : params[0], Port : params[1]})
	for i, server := range lb.servers {
		if (server.Address() == toRemove.Address()){
			removeIndex = i
			break
		}
	}
	lb.servers = append(lb.servers[:removeIndex], lb.servers[removeIndex+1:]...)

}

