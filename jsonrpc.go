package main

import (
	"encoding/json"
	"io"

	"github.com/ethereum/go-ethereum/p2p"
)

// PeersResponse represents JSON-RPC response for `admin_peers` command from
// geth instance.
type PeersResponse struct {
	Version string          `json:"jsonrpc"`
	Id      interface{}     `json:"id,omitempty"`
	Result  []*p2p.PeerInfo `json:"result"`
}

// ParsePeersResponse parses JSON-RPC 'admin_peers' response from reader r.
func ParsePeersResponse(r io.Reader) ([]*p2p.PeerInfo, error) {
	lr := io.LimitReader(r, 10e6) // 1MB should be more than enough
	var resp PeersResponse
	err := json.NewDecoder(lr).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}

// NodeInfoResponse represents JSON-RPC response for `admin_nodeInfo` command from
// geth instance.
type NodeInfoResponse struct {
	Version string        `json:"jsonrpc"`
	Id      interface{}   `json:"id,omitempty"`
	Result  *p2p.NodeInfo `json:"result"`
}

// ParseNodeInfoResponse parses JSON-RPC 'admin_nodeInfo' response from reader r.
func ParseNodeInfoResponse(r io.Reader) (*p2p.NodeInfo, error) {
	lr := io.LimitReader(r, 10e6) // 1MB should be more than enough
	var resp NodeInfoResponse
	err := json.NewDecoder(lr).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
