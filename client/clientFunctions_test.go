package client

import (
	"github.com/stretchr/testify/assert"
	"net"
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

func TestSetupReaderAndWriter(t *testing.T) {
	type args struct {
		connectionWithServer net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_acceptMessageFromServer(t *testing.T) {
	type args struct {
		connectionWithServer net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}

func Test_writeMessageToServer(t *testing.T) {
	type args struct {
		connectionWithServer net.Conn
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}