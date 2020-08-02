package main

import (
	"websocket/client"
)

func main() {
	connectionWithServer := client.DialServer("tcp", "localhost:49152")
	client.SetupReaderAndWriter(connectionWithServer)
}
