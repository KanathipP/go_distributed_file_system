package p2p

// Peer is an interface that represents the remote node
type Peer interface {
	Close() error
}

// transport is anything that handles the communication
// between the node in the network. This can be of the
// form ( TCP, UDP, websockets, ...)
type Transport interface {
	ListenAndAccept() error
	Consume() <-chan RPC
}
