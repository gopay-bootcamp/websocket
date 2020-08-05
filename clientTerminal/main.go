package main

import (
	"websocket/client"
)

func main() {
	connectionWithServer, connectionWithServerErr := client.DialServer("tcp", "localhost:49152")
	if connectionWithServerErr == nil {
		defer connectionWithServer.Close()
		client.SetupReaderAndWriter(connectionWithServer)
	}
}
