package main

import (
	"github.com/ethereum/go-ethereum/p2p"
)

// Node represents single node information to be used in Graph.
type Node struct {
	ID_    string `json:"id"`
	Group_ int    `json:"group"`

	meta *Metadata
}

// NewNode creates new Node object for the given peerinfo.
func NewNode(id, name string) *Node {
	meta := NewMetadata(name)
	return &Node{
		ID_:    id,
		Group_: clientGroup(meta),

		meta: meta,
	}
}

// IsClient returns true if node is identified as a mobile client, rather then server.
func (n *Node) IsClient() bool {
	return n.meta.Name == "StatusIM"
}

// clientGroup returns group id based in server type.
func clientGroup(meta *Metadata) int {
	switch meta.Name {
	case "StatusIM":
		return 1
	case "Statusd":
		return 2
	default:
		return 3

	}
	return 3
}

// ClusterNode represents single cluster node information.
type ClusterNode struct {
	IP   string
	ID   string
	Type string // name field in JSON (statusd or statusIM)
}

// NewClusterNode creates new Node object for the given peerinfo.
func NewClusterNode(ip string, peer *p2p.NodeInfo) *ClusterNode {
	return &ClusterNode{
		IP:   ip,
		ID:   peer.ID,
		Type: peer.Name,
	}
}

// PeersToNodes converts PeerInfo to Nodes.
func PeersToNodes(peers []*p2p.PeerInfo) ([]*Node, error) {
	ret := make([]*Node, len(peers))
	for i, peer := range peers {
		ret[i] = NewNode(peer.ID, peer.Name)
	}
	return ret, nil
}

// ID returns ID of the node. Satisfies graph.Node interface.
func (n *Node) ID() string {
	return n.ID_
}

// Group returns group of the node. Satisfies graph.Node interface.
func (n *Node) Group() int {
	return n.Group_
}

// Link represents link between two nodes.
type Link struct {
	FromID, ToID string
}

// NewLinks creates link for the given IDs.
func NewLink(fromID, toID string) *Link {
	return &Link{fromID, toID}
}
