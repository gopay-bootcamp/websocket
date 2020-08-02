package client

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

func DialServer(network, address string) net.Conn {
	connectionWithServer, connectionWithServerErr := net.Dial(network, address)
	if connectionWithServerErr != nil {
		log.Fatal(connectionWithServerErr)
	}
	log.Printf("Establishing connection with server at network address: %v", connectionWithServer.RemoteAddr())
	return connectionWithServer
}

func writeMessageToServer(connectionWithServer net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		dataForServer, dataForServerError := reader.ReadString('\n')
		if dataForServerError != nil {
			log.Fatal(dataForServerError)
		}
		fmt.Fprintf(connectionWithServer, dataForServer+"\n")
		if strings.TrimSpace(string(dataForServer)) == "STOP" {
			log.Println("Client cannot send messages to server now")
			wg.Done()
			return
		}

	}
}
func acceptMessageFromServer(connectionWithServer net.Conn) {
	for {
		reader := bufio.NewReader(connectionWithServer)
		dataFromServer, dataFromServerError := reader.ReadString('\n')
		if dataFromServerError != nil {
			log.Fatal(dataFromServerError)
		}
		if strings.TrimSpace(string(dataFromServer)) == "STOP" {
			log.Println("Client cannot receive messages from server now")
			wg.Done()
			return
		}
		fmt.Print("From server-> " + dataFromServer)
	}
}
func SetupReaderAndWriter(connectionWithServer net.Conn) {
	wg.Add(1)
	go writeMessageToServer(connectionWithServer)
	wg.Add(1)
	go acceptMessageFromServer(connectionWithServer)
	wg.Wait()

}
