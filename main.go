package main

import (
	"fmt"
	"flag"
	"server"
	"context"
)

func main() {
	flag.StringVar(&context.Config.Server.Port, "port", "8888", "Port of service")
	flag.StringVar(&context.Config.Server.Ip, "ip", "0.0.0.0", "Ip of service")
	flag.Parse()
	server.Server()
}
