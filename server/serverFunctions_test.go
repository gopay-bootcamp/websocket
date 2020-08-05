package server

import (
	"testing"
	"websocket/client"

	"github.com/stretchr/testify/assert"
)

func TestSetupListener(t *testing.T) {
	listener, listenerErr := SetupListener("tcp", "localhost:49152")
	assert.Equal(t, nil, listenerErr)
	listener.Close()

	_, listenerErr = SetupListener("tcp", "localhost:49152403293")
	assert.NotEqual(t, nil, listenerErr)

}

func TestSetupConnection(t *testing.T) {
	listener, _ := SetupListener("tcp", "localhost:49153")
	connectionWithServer, _ := client.DialServer("tcp", "localhost:49153")
	_, connectionWithClientErr := SetupConnection(listener)
	assert.Equal(t, nil, connectionWithClientErr)
	listener.Close()
	connectionWithServer.Close()

}
