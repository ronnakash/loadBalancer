# LoadBalancer

The load balancer is a reverse proxy load balancer implemented in Golang, which is a programming language designed for building high-performance and scalable systems. The load balancer is designed to distribute incoming HTTP requests to multiple servers, effectively handling the load and improving the performance of the system. It uses multiple algorithms to determine the best server to handle each incoming request, ensuring that the load is distributed as evenly as possible.

The load balancer uses Reverse Proxy, which is a technique for directing network traffic to different servers, based on the request's URL or other information. This allows the load balancer to redirect requests to the appropriate server, improving the performance and scalability of the system.


The load balancer can be configured using a .yaml configuration file, or by passing arguments to the program when it is launched. This allows for flexibility in how the load balancer is set up and configured. Additionally, the load balancer can be interacted with using the command line, providing a user-friendly interface for managing and monitoring the load balancer.

# How To Use

The load balancer configuration is initialized using a config.yaml file in the following format:

    algorithm: (round-robin/least-connections)

    port: 8080

    logging: false

    servers:

    - address: localhost

    port: 8081

    - address: localhost

    port: 8082

    - address: localhost

    port: 8083

If the algorithm or port are not specified, the load balancer will request the user to input them using the cli, with default values of port 8080 and round-robin algorithm.

The load balancer configuration can be changed while it is running using the cli and the following commands:

    configuration commands:
        - as/add-server [addr:port]:
            adds a server to the cluster with the specified address and port
        - rs/remove-server [addr:port]:
            removes a server from the cluster with the specified address and port if it exists, will print error otherwise
        - a/algo [round-robin/least-connections]:
            sets the load balancer algorithm accordingly
        - l/log [true/false]:
            enable or disable logging

    additinal commands:
        -ll/list:
            list all working servers addresses, ports and open connections count

