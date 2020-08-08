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

//DialServer calls the net.Dial() function and prints the appropraite log messages
func DialServer(network, address string) (net.Conn, error) {
	connectionWithServer, connectionWithServerErr := net.Dial(network, address)
	if connectionWithServerErr != nil {
		log.Print(connectionWithServerErr)
		return connectionWithServer, connectionWithServerErr
	}
	log.Printf("Establishing connection with server at network address: %v", connectionWithServer.RemoteAddr())
	return connectionWithServer, connectionWithServerErr
}

func writeMessageToServer(connectionWithServer net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		dataForServer, dataForServerError := reader.ReadString('\n')
		if dataForServerError != nil {
			log.Println(dataForServerError)
			wg.Done()
			return
		}
		_, err := fmt.Fprintf(connectionWithServer, dataForServer+"\n")
		if err != nil {
			log.Println(err)
			wg.Done()
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataForServer)) == "STOP CLIENT" {
			os.Exit(0)
		}

	}
}
func acceptMessageFromServer(connectionWithServer net.Conn) {
	for {
		reader := bufio.NewReader(connectionWithServer)
		dataFromServer, dataFromServerError := reader.ReadString('\n')
		if dataFromServerError != nil {
			log.Println(dataFromServerError)
			wg.Done()
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataFromServer)) == "STOP SERVER" {
			log.Println("Stopping client because server has stopped")
			os.Exit(0)
		}
		fmt.Printf("From server-> %s", dataFromServer)
	}
}

//SetupReaderAndWriter runs the reader and write goroutines simultaneously
func SetupReaderAndWriter(connectionWithServer net.Conn) {
	wg.Add(1)
	go writeMessageToServer(connectionWithServer)
	wg.Add(1)
	go acceptMessageFromServer(connectionWithServer)
	log.Println("You can now start communication")
	wg.Wait()

}
