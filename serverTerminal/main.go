package main

import (
	"websocket/server"
)

func main() {
	listener := server.SetupListener("tcp", "localhost:49152")
	connectionWithClient := server.SetupConnection(listener)
	defer connectionWithClient.Close()
	server.SetupReaderAndWriter(connectionWithClient)
}
