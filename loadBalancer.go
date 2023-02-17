package main

import (
	"fmt"
	"net/http"
)

type LoadBalancer struct {
	port            int
	roundRobinCount int
	servers         []Server
	algorithm		string
}

func NewLoadBalancer(config Config) *LoadBalancer {
	var serverParams []ServerParams = config.servers
	var servers []Server

	// Initialize each server
    for i := range serverParams {
        servers = append(servers, NewSimpleServer(&serverParams[i]))
    }

	return &LoadBalancer{
		port:            	config.port,
		roundRobinCount:	0,
		servers:        	servers,
		algorithm: 			config.algorithm,

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
	fmt.Printf("forwarding request to address %s\n", (*targetServer.Address()).String())
	targetServer.IncrementConnections()
	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
	// fmt.Printf("done processing request at %s\n", targetServer.Address())
	targetServer.DecrementConnections()
}

func (lb *LoadBalancer) ChangeAlgorithm(algo string) bool{
	if algo == "round-robin" || algo == "least-connections" {
		lb.algorithm = algo
	} else {
		fmt.Printf("Algorithm %s is not supported\n", algo)
		return false
	}
	return true
}


func (lb *LoadBalancer) AddServer(params []string) bool{
	//TODO: ensure params are of length 2, contain valid address and port
	newServer := NewSimpleServer(&ServerParams {address : params[0], port : params[1]})
	if newServer != nil{
		lb.servers = append(lb.servers, newServer)
	}
	return newServer != nil
}

func (lb *LoadBalancer) RemoveServer(params []string) bool{
	toRemove := NewSimpleServer(&ServerParams {address : params[0], port : params[1]})
	for i, server := range lb.servers {
		if server.Address() == toRemove.Address(){
			lb.servers = append(lb.servers[:i], lb.servers[i+1:]...)
			return true
		}
	}
	return false
}

