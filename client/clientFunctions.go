package client

import (
	"log"
	"net"
)

func DialServer(network, address string) net.Conn {
	connectionWithServer, connectionWithServerErr := net.Dial(network, address)
	if connectionWithServerErr != nil {
		log.Fatal(connectionWithServerErr)
	}
	log.Printf("Establishing connection with server at network address: %v", connectionWithServer.RemoteAddr())
	return connectionWithServer
}
