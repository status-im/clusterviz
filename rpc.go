package main

import (
	"bytes"
	"fmt"
	"net/http"
)

// RPCClient defines subset of client that
// can call needed methods to geth's RPC server.
type RPCClient interface {
	AdminPeers() ([]*Node, error)
}

// HTTPRPCClient implements RPCClient for
// HTTP transport.
type HTTPRPCClient struct {
	IP string
}

// NewHTTPRPCClient creates new HTTP RPC client for eth JSON-RPC server.
func NewHTTPRPCClient(ip string) *HTTPRPCClient {
	return &HTTPRPCClient{
		IP: ip,
	}
}

// AdminPeers executes `admin_peers` RPC call and parses the response.
// Satisfies RPCClient interface.
func (h *HTTPRPCClient) AdminPeers() ([]*Node, error) {
	data := bytes.NewBufferString(`{"jsonrpc":"2.0","method":"admin_peers","params":[],"id":1}`)
	resp, err := http.Post("https://"+h.IP, "application/json", data)
	if err != nil {
		return nil, fmt.Errorf("POST RPC request: %s", err)
	}
	defer resp.Body.Close()

	nodes, err := ParseResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get admin peers: %s", err)
	}

	return PeersToNodes(nodes)
}
