package main

import (
	"fmt"
	"net/http"
)




func main() {
	
	config := parse()
	// fmt.Printf("number of servers:%d'\n", len(servers))
	lb := NewLoadBalancer(config)
	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		lb.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)
	fmt.Printf("serving requests at localhost:%s\n", lb.port)
	http.ListenAndServe(":"+lb.port, nil)
}
