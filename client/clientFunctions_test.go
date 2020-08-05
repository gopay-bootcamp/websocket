package client

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"websocket/server"
)

func TestDialServer(t *testing.T) {
	listener,_:= server.SetupListener("tcp", "localhost:49153")
	_,connectionWithServerErr:=DialServer("tcp", "localhost:49153")
	assert.Equal(t,nil,connectionWithServerErr)
	listener.Close()

	listener,_= server.SetupListener("tcp", "localhost:49153")
	_,connectionWithServerErr=DialServer("tcp", "localhost:9153")
	assert.NotEqual(t,nil,connectionWithServerErr)
	listener.Close()
}

