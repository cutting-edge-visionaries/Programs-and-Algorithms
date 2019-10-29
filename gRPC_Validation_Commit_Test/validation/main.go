package main

import (
	"daad/validation/server"
	"flag"
)

var (
	host = flag.String("host", "", "hostname")
	port = flag.Int("port", 29000, "port")
)

func main() {
	server.Main(*host, *port)
}
