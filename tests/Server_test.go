package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"
)

func TestMsgFromClientToServer(t *testing.T) {

	message := "Message from client to server checked successfully\n"
	listener, listenerError := net.Listen("tcp", "localhost:49152")
	if listenerError != nil {
		t.Fatal(listenerError)
	}
	connectionWithServer, connectionWithServerError := net.Dial("tcp", "localhost:49152")
	if connectionWithServerError != nil {
		t.Fatal(connectionWithServerError)
	}
	connectionWithClient, connectionWithClientError := listener.Accept()
	if connectionWithClientError != nil {
		return
	}

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

	listener, listenerError := net.Listen("tcp", "localhost:49152")
	if listenerError != nil {
		t.Fatal(listenerError)
	}
	connectionWithServer, connectionWithServerError := net.Dial("tcp", "localhost:49152")
	if connectionWithServerError != nil {
		t.Fatal(connectionWithServerError)
	}
	connectionWithClient, connectionWithClientError := listener.Accept()
	if connectionWithClientError != nil {
		return
	}

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
