package main

import (
	"flag"
	"fmt"
	"server"
)

func main() {
	config := new(server.Config)
	flag.StringVar(&config.Port, "port", "8888", "Port of service")
	flag.StringVar(&config.Ip, "ip", "0.0.0.0", "Ip of service")
	fmt.Print(config)
	server.Server(config)
}
