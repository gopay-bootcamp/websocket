package main

import (
	"fmt"
	"websocket/server"
)

func main() {
	listener := server.SetupListener("tcp", "localhost:49152")
	connectionWithClient := server.SetupConnection(listener)
	fmt.Println(connectionWithClient)
}
