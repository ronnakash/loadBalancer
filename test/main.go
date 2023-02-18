package main

import (
	// "context"
	"flag"
	"fmt"
	"time"

	// "net"
	"net/http"
	// "net/http/httputil"
	"sync"
)

type TestServer struct {
    hostname            string
    helloHandler        http.Handler
    waitTime            time.Duration
    connections         int
    mutexCount 		    sync.Mutex
    mutexProcess 		sync.Mutex

}

func NewTestServer(port int, waitTime int) *TestServer {
    hostname := fmt.Sprintf("localhost:%d", port)
    testServer := &TestServer{
        hostname:       hostname,
        waitTime:       time.Duration(waitTime)*time.Second,
        connections:    0,
    }
    testServer.helloHandler = helloHandler(testServer)
    return testServer
}

func helloHandler(testServer *TestServer) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("processing at %s\n", testServer.hostname)
        testServer.mutexCount.Lock()
        testServer.connections++
        testServer.mutexCount.Unlock()
        testServer.mutexProcess.Lock()
        testServer.mutexCount.Lock()
        fmt.Printf("requests at %s : %d\n", testServer.hostname, testServer.connections)
        testServer.mutexCount.Unlock()
        time.Sleep(testServer.waitTime)
        fmt.Fprintf(w, "hello from %s", testServer.hostname)
        testServer.mutexProcess.Unlock()
        testServer.mutexCount.Lock()
        testServer.connections--
        testServer.mutexCount.Unlock()
        fmt.Printf("done at %s\n", testServer.hostname)
    })
}

func makeRequests(requests int) {
    var wg sync.WaitGroup
    for i := 0; i < requests; i++ {
        time.Sleep(time.Second/4)
        fmt.Println("making request")
        wg.Add(1)
        go func() {
            defer wg.Done()
            _, err := http.Get("http://localhost:8080")
            if err != nil {
                // handle error
                fmt.Println("request failed!")
            }
        }()
    }
    wg.Wait()
}

type Config struct {
    algorithm string
    serverAddr string
    serverPort int
}


func TestConfig() {
    var cfg Config

    flag.StringVar(&cfg.algorithm, "algorithm", "", "Algorithm name")
    flag.StringVar(&cfg.serverAddr, "server", "", "Server address")
    flag.IntVar(&cfg.serverPort, "port", 0, "Server port")
    flag.Parse()

    flag.Visit(func(f *flag.Flag) {
        switch f.Name {
        case "algorithm":
            fmt.Print("Enter algorithm name: ")
            fmt.Scan(&cfg.algorithm)
        case "server":
            fmt.Print("Enter server address: ")
            fmt.Scan(&cfg.serverAddr)
        case "port":
            fmt.Print("Enter server port: ")
            fmt.Scan(&cfg.serverPort)
        }
    })

    fmt.Println("Using config:")
    fmt.Println("Algorithm:", cfg.algorithm)
    fmt.Println("Server:", cfg.serverAddr)
    fmt.Println("Port:", cfg.serverPort)
}

func main() {
    server1 := NewTestServer(8081, 100)
    server2 := NewTestServer(8082, 10)
    server3 := NewTestServer(8083, 1)
    go http.ListenAndServe(":8081", server1.helloHandler)
    go http.ListenAndServe(":8082", server2.helloHandler)
    go http.ListenAndServe(":8083", server3.helloHandler)
    makeRequests(100)
}

