package p2p

// Message holds any arbitrary data that is being set over the
// each transport between two nodes in the network

type Message struct {
	Payload []byte
}
