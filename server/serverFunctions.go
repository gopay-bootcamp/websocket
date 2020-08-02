package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"sync"
)

var wg = sync.WaitGroup{}

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
func acceptMessageFromClient(connectionWithClient net.Conn) {
	for {
		reader := bufio.NewReader(connectionWithClient)
		dataFromClient, dataFromClientError := reader.ReadString('\n')
		if dataFromClientError != nil {
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataFromClient)) == "STOP" {
			wg.Done()
			return
		}
		fmt.Print("From client -> ", string(dataFromClient))
	}

}
func writeMessageToClient(connectionWithClient net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		dataForClient, dataForClientError := reader.ReadString('\n')
		if dataForClientError != nil {
			wg.Done()
			return
		}
		connectionWithClient.Write([]byte(dataForClient))
		if strings.TrimSpace(string(dataForClient)) == "STOP" {
			wg.Done()
			return
		}

	}
}
func SetupReaderAndWriter(connectionWithClient net.Conn) {
	wg.Add(1)
	go acceptMessageFromClient(connectionWithClient)
	wg.Add(1)
	go writeMessageToClient(connectionWithClient)
	wg.Wait()

}
