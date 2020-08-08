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
var ConnectionsMap = make(map[net.Addr]net.Conn)

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
	ConnectionsMap[connectionWithClient.RemoteAddr()] = connectionWithClient
	fmt.Println("Conncetions map:", ConnectionsMap)
	log.Printf("Establishing connection with client at network address: %v", connectionWithClient.RemoteAddr())
	return connectionWithClient, connectionWithClientErr
}
func acceptMessageFromClient(connectionWithClient net.Conn) {
	for {
		reader := bufio.NewReader(connectionWithClient)
		dataFromClient, dataFromClientError := reader.ReadString('\n')
		if dataFromClientError != nil {
			log.Printf("Deleting client@%v from map because of %v", connectionWithClient.RemoteAddr(), dataFromClientError)
			wg.Done()
			return
		}
		if strings.TrimSpace(string(dataFromClient)) == "STOP CLIENT" {
			delete(ConnectionsMap, connectionWithClient.RemoteAddr())
			log.Println("Received STOP CLIENT command from", connectionWithClient.RemoteAddr())
			log.Printf("Deleting client@%v from connections map", connectionWithClient.RemoteAddr())
			wg.Done()
			return
		}
		fmt.Printf("From client@%v-> %s", connectionWithClient.RemoteAddr(), string(dataFromClient))
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
		for clientAddress := range ConnectionsMap {
			//fmt.Println("Capital of",country,"is",countryCapitalMap[country])
			connectionWithClient := ConnectionsMap[clientAddress]
			_, err := connectionWithClient.Write([]byte(dataForClient))
			if err != nil {
				delete(ConnectionsMap, connectionWithClient.RemoteAddr())
				log.Printf("Deleting client@%v from map because of %v", clientAddress, err)
				wg.Done()
				return
			}
		}
		if strings.TrimSpace(string(dataForClient)) == "STOP SERVER" {
			log.Println("Stopping server...")
			os.Exit(0)
		}

	}
}

//SetupReaderAndWriter runs the reader and write goroutines simultaneously
func SetupReaderAndWriter(connectionWithClient net.Conn) {
	wg.Add(1)
	go acceptMessageFromClient(connectionWithClient)
	wg.Add(1)
	go writeMessageToClient(connectionWithClient)
	log.Printf("You can now start communication with client at network address: %v", connectionWithClient.RemoteAddr())
	wg.Wait()

}
