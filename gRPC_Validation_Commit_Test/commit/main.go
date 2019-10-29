package main

import (
	"daad/commit/server"
	"flag"
)

var (
	host = flag.String("host", "", "hostname")
	port = flag.Int("port", 28000, "port")
)

func main() {
	server.Main(*host, *port)
}
