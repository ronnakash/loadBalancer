package main


func main() {
	lb:= InitializeLoadBalancer()
	lb.ServeForever()
}
