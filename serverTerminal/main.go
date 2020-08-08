package main

import (
	"log"
	"websocket/server"
)

func main() {
	listener, listenerErr := server.SetupListener("tcp", "localhost:49152")
	if listenerErr == nil {
		connectionWithClient, connectionWithClientErr := server.SetupConnection(listener)
		if connectionWithClientErr != nil {
			log.Println(connectionWithClientErr)
			return
		}
		server.SetupReaderAndWriter(connectionWithClient)
	}
}
