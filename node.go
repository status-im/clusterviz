package main

import "github.com/ethereum/go-ethereum/p2p"

// Node represents single node information.
type Node struct {
	*p2p.PeerInfo
}

// NewNode creates new Node object for the given peerinfo.
func NewNode(peer *p2p.PeerInfo) *Node {
	return &Node{
		PeerInfo: peer,
	}
}

// PeersToNodes converts PeerInfo to Nodes.
func PeersToNodes(peers []*p2p.PeerInfo) ([]*Node, error) {
	ret := make([]*Node, len(peers))
	for i := range peers {
		ret[i] = NewNode(peers[i])
	}
	return ret, nil
}
