package p2p

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTCPTransport(t *testing.T) {
	tcpOpts := TCPTransportOpts{
		ListenAddr:    ":4000",
		HandshakeFunc: NOPHandshakeFunc,
	}
	tr := NewTCPTransport(tcpOpts)

	assert.Equal(t, tr.ListenAddr, tcpOpts.ListenAddr)

	// Server
	assert.Nil(t, tr.ListenAndAccept())
}
