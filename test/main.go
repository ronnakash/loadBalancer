package main

import (
	// "context"
	// "flag"
	"fmt"
	"time"

	// "net"
	"net/http"
	// "net/http/httputil"
	"sync"
)



func helloHandler(hostname string, waitTime int) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        fmt.Printf("processing at %s\n", hostname)
        time.Sleep(time.Duration(waitTime)*time.Second)
        fmt.Fprintf(w, "hello from %s", hostname)
        fmt.Printf("done at %s\n", hostname)
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


func main() {
    // var cfg Config

    // flag.StringVar(&cfg.algorithm, "algorithm", "", "Algorithm name")
    // flag.StringVar(&cfg.serverAddr, "server", "", "Server address")
    // flag.IntVar(&cfg.serverPort, "port", 0, "Server port")
    // flag.Parse()

    // flag.Visit(func(f *flag.Flag) {
    //     switch f.Name {
    //     case "algorithm":
    //         fmt.Print("Enter algorithm name: ")
    //         fmt.Scan(&cfg.algorithm)
    //     case "server":
    //         fmt.Print("Enter server address: ")
    //         fmt.Scan(&cfg.serverAddr)
    //     case "port":
    //         fmt.Print("Enter server port: ")
    //         fmt.Scan(&cfg.serverPort)
    //     }
    // })

    // fmt.Println("Using config:")
    // fmt.Println("Algorithm:", cfg.algorithm)
    // fmt.Println("Server:", cfg.serverAddr)
    // fmt.Println("Port:", cfg.serverPort)


    go http.ListenAndServe(":8081", helloHandler("localhost:8081", 3))
    go http.ListenAndServe(":8082", helloHandler("localhost:8082", 2))
    go http.ListenAndServe(":8083", helloHandler("localhost:8083", 1))
    makeRequests(100)
}

