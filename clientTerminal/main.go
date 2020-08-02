package main

import (
	"fmt"
	"websocket/client"
)

func main() {
	connectionWithServer := client.DialServer("tcp", "localhost:49152")
	client.SetupReaderAndWriter(connectionWithServer)
	fmt.Printf("Local address of connection: %v of type: %T\n", connectionWithServer.LocalAddr(), connectionWithServer.LocalAddr())
	fmt.Printf("Remote address of connection: %v of type: %T\n", connectionWithServer.RemoteAddr(), connectionWithServer.RemoteAddr())
}
