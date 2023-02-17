package main

import (
	"fmt"
	"net/http"
	"strconv"
)


func main() {
	
	config := Parse()
	// fmt.Printf("number of servers:%d'\n", len(servers))
	lb := NewLoadBalancer(config)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}
	portStr := strconv.Itoa(lb.port)
	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at localhost:%s\n", portStr)
	http.ListenAndServe(":" + portStr, nil)
}
