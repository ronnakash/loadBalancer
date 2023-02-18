package main

import (
	"fmt"
	"net/http"
	"strconv"
)

type LoadBalancer struct {
	port            int
	logger 			*Logger
	roundRobinCount int
	servers         *[]Server
	algorithm		string
}

func NewLoadBalancer(config Config) *LoadBalancer {
	var serverParams []ServerParams = config.Servers
	var servers []Server
	logger := NewLogger(config.Logging)

	// Initialize each server
    for i := range serverParams {
		server := NewSimpleServer(&serverParams[i])
        servers = append(servers, server)
    }

	lb := &LoadBalancer{
		port:            	config.Port,
		roundRobinCount:	0,
		servers:        	&servers,
		algorithm: 			config.Algorithm,
		logger: 			logger,
	}
	lb.PrintServerList()
	return lb
}

func (lb *LoadBalancer) PrintServerList() {
	for i, server := range *lb.servers {
		serverAddr := *server.Address()
		lb.logger.Print(fmt.Sprintf("Server %d at %s\n", i, serverAddr.String()))
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
	servers := *lb.servers
	server := servers[lb.roundRobinCount%len(servers)]
	for !server.IsAlive() {
		lb.roundRobinCount++
		server = servers[lb.roundRobinCount%len(servers)]
	}
	lb.roundRobinCount++
	return server
}

func (lb *LoadBalancer) getLeastConnectedServer() Server {
    var server Server
    var minConns = -1
    for _, s := range *lb.servers {
		conns:= s.GetConnections()
        if conns < minConns {
            server = s
            minConns = conns
        }
    }
    return server
}


func (lb *LoadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {
	targetServer := lb.getNextAvailableServer()

	// could optionally log stuff about the request here!
	lb.logger.Print(fmt.Sprintf("Forwarding request to address %s with %d connection \n", (*targetServer.Address()).String(), targetServer.GetConnections()+1))
	targetServer.IncrementConnections()
	// could delete pre-existing X-Forwarded-For header to prevent IP spoofing
	targetServer.Serve(rw, req)
	// fmt.Printf("done processing request at %s\n", targetServer.Address())
	targetServer.DecrementConnections()
}

func (lb *LoadBalancer) ChangeAlgorithm(algo string) bool{
	if algo == "round-robin" || algo == "least-connections" {
		lb.algorithm = algo
		lb.logger.Print(fmt.Sprintf("Algorithm srt to %s \n", algo))
	} else {
		lb.logger.PrintError(fmt.Sprintf("Algorithm %s is not supported\n", algo))
		return false
	}
	return true
}


func (lb *LoadBalancer) AddServer(addr string) bool{
	newServer := NewSimpleServer(NewServerParams(addr))
	if newServer != nil{
		*lb.servers = append(*lb.servers, newServer)
		lb.logger.Print(fmt.Sprintf("New server %s added\n", addr))
		return true
	}
	lb.logger.PrintError(fmt.Sprintf("Failed to add %s\n", addr))
	return newServer != nil
}

func (lb *LoadBalancer) RemoveServer(addr string) bool{
	toRemove := NewSimpleServer(NewServerParams(addr))
	servers := *lb.servers
	for i, server := range servers {
		if server.Address() == toRemove.Address(){
			*lb.servers = append(servers[:i], servers[i+1:]...)
			lb.logger.Print(fmt.Sprintf("Server %s at %d removed\n", addr, i))
			return true
		}
	}
	lb.logger.PrintError(fmt.Sprintf("Failed to remove %s\n", addr))
	return false
}

func (lb *LoadBalancer) ReadFromCLI() {
	for true {
		ReadInput(lb)
	}
}

func (lb *LoadBalancer) ServeForever() {
	portStr := strconv.Itoa(lb.port)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	lb.logger.Print(fmt.Sprintf("serving requests at localhost:%s\n", portStr))
	go http.ListenAndServe(":"+portStr, http.HandlerFunc(handleRedirect))
	lb.ReadFromCLI()
}

func (lb *LoadBalancer) SetLogging(log bool) {
	lb.logger.SetLogging(log)
}

func InitializeLoadBalancer() *LoadBalancer {
	config := Parse()
	lb := NewLoadBalancer(config)
	return lb
}