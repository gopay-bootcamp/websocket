package server

import (
	"log"
	"net"
)

func SetupListener(network, address string) net.Listener {
	listener, listenerErr := net.Listen(network, address)
	if listenerErr != nil {
		log.Fatal(listenerErr)
	}
	log.Printf("Server is listening at network address: %v\n", listener.Addr())
	return listener
}

func SetupConnection(clientListener net.Listener) net.Conn {
	log.Println("Waiting for client to dial...")
	connectionWithClient, connectionWithClientErr := clientListener.Accept()
	if connectionWithClientErr != nil {
		log.Fatal(connectionWithClientErr)
	}
	log.Printf("Establishing connection with client at network address: %v", connectionWithClient.RemoteAddr())
	return connectionWithClient
}
