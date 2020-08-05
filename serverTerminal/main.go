package main

import (
	"websocket/server"
)

func main() {
	listener, listenerErr := server.SetupListener("tcp", "localhost:49152")
	if listenerErr == nil {
		connectionWithClient, connectionWithClientErr := server.SetupConnection(listener)
		if connectionWithClientErr == nil {
			defer connectionWithClient.Close()
			server.SetupReaderAndWriter(connectionWithClient)
		}
	}
}
