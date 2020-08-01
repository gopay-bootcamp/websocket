package server

import (
	"log"
	"net"
)

func SetupListener(network, address string) net.Listener {
	clientListener, clientListenerErr := net.Listen(network, address)
	if clientListenerErr != nil {
		log.Fatal(clientListenerErr)
	}
	log.Printf("Server is listening at network address: %v\n", clientListener.Addr())
	return clientListener
}
