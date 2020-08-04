package main

import (
	"bufio"
	"fmt"
	"testing"
	"websocket/client"
	"websocket/server"
)


func TestMsgFromClientToServer(t *testing.T) {

	message := "Message from client to server checked successfully\n"
	listener := server.SetupListener("tcp", "localhost:49152")

	connectionWithServer := client.DialServer("tcp", "localhost:49152")
	connectionWithClient := server.SetupConnection(listener)

	connectionWithServer.Write([]byte(message))

	buf, _ := bufio.NewReader(connectionWithClient).ReadString('\n')
	fmt.Println(string(buf[:]))
	if msg := string(buf[:]); msg != message {
		t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
	}

	connectionWithServer.Close()
	listener.Close()
	connectionWithClient.Close()

}

func TestMsgFromServerToClient(t *testing.T) {

	message := "Message from  server to client checked successfully\n"

	listener := server.SetupListener("tcp", "localhost:49152")

	connectionWithServer := client.DialServer("tcp", "localhost:49152")
	connectionWithClient := server.SetupConnection(listener)

	connectionWithClient.Write([]byte(message))

	buf, _ := bufio.NewReader(connectionWithServer).ReadString('\n')
	fmt.Println(string(buf[:]))
	if msg := string(buf[:]); msg != message {
		t.Fatalf("Unexpected message:\nGot:\t\t%s\nExpected:\t%s\n", msg, message)
	}

	connectionWithServer.Close()
	listener.Close()
	connectionWithClient.Close()

}
