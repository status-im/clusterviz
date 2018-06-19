package main

import (
	"encoding/json"
	"io"

	"github.com/ethereum/go-ethereum/p2p"
)

// JSONRPCResponse represents JSON-RPC response for `admin_peers` command from
// geth instance.
type JSONRPCResponse struct {
	Version string          `json:"jsonrpc"`
	Id      interface{}     `json:"id,omitempty"`
	Result  []*p2p.PeerInfo `json:"result"`
}

// ParseResponse parses JSON-RPC 'admin_peers' response from reader r.
func ParseResponse(r io.Reader) ([]*p2p.PeerInfo, error) {
	var resp JSONRPCResponse
	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return nil, err
	}

	return resp.Result, nil
}
