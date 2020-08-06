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

//SetupListener calls the net.listen function and prints the appropriate log messages
func SetupListener(network, address string) (net.Listener, error) {
	listener, listenerErr := net.Listen(network, address)
	if listenerErr != nil {
		log.Print(listenerErr)
		return listener, listenerErr
	}
	log.Printf("Server is listening at network address: %v\n", listener.Addr())
	return listener, listenerErr
}

//SetupConnection calls the net.listener.Accept() function and prints the appropraite log messages
func SetupConnection(clientListener net.Listener) (net.Conn, error) {
	log.Println("Waiting for client to dial...")
	connectionWithClient, connectionWithClientErr := clientListener.Accept()
	if connectionWithClientErr != nil {
		log.Print(connectionWithClientErr)
		return connectionWithClient, connectionWithClientErr
	}
	log.Printf("Establishing connection with client at network address: %v", connectionWithClient.RemoteAddr())
	return connectionWithClient, connectionWithClientErr
}
func acceptMessageFromClient(connectionWithClient net.Conn) {
	for {
		reader := bufio.NewReader(connectionWithClient)
		dataFromClient, dataFromClientError := reader.ReadString('\n')
		if dataFromClientError != nil {
			log.Println(dataFromClientError)
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataFromClient)) == "STOP" {
			log.Println("Server cannot accept messages from client now")
			wg.Done()
			return
		}
		fmt.Printf("From client-> %s", string(dataFromClient))
	}

}
func writeMessageToClient(connectionWithClient net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		dataForClient, dataForClientError := reader.ReadString('\n')
		if dataForClientError != nil {
			log.Println(dataForClientError)
			wg.Done()
			return
		}
		_, err := connectionWithClient.Write([]byte(dataForClient))
		if err != nil {
			log.Println(err)
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataForClient)) == "STOP" {
			log.Println("Server cannot send messages to client now")
			wg.Done()
			return
		}

	}
}

//SetupReaderAndWriter runs the reader and write goroutines simultaneously
func SetupReaderAndWriter(connectionWithClient net.Conn) {
	wg.Add(1)
	go acceptMessageFromClient(connectionWithClient)
	wg.Add(1)
	go writeMessageToClient(connectionWithClient)
	log.Println("You can now start communication")
	wg.Wait()

}
