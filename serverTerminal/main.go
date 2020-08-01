package main

import (
	"websocket/server"
)

func main() {
	server.SetupListener("tcp", "localhost:49152")
}
