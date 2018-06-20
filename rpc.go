package main

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/ethereum/go-ethereum/p2p"
)

// RPCClient defines subset of client that
// can call needed methods to geth's RPC server.
type RPCClient interface {
	AdminPeers(ip string) ([]*Node, error)
	NodeInfo(ip string) (*p2p.NodeInfo, error)
}

// HTTPRPCClient implements RPCClient for
// HTTP transport.
type HTTPRPCClient struct {
}

// NewHTTPRPCClient creates new HTTP RPC client for eth JSON-RPC server.
func NewHTTPRPCClient() *HTTPRPCClient {
	return &HTTPRPCClient{}
}

// AdminPeers executes `admin_peers` RPC call and parses the response.
// Satisfies RPCClient interface.
func (h *HTTPRPCClient) AdminPeers(ip string) ([]*Node, error) {
	data := bytes.NewBufferString(`{"jsonrpc":"2.0","method":"admin_peers","params":[],"id":1}`)
	resp, err := http.Post("http://"+ip, "application/json", data)
	if err != nil {
		return nil, fmt.Errorf("POST RPC request: %s", err)
	}
	defer resp.Body.Close()

	nodes, err := ParsePeersResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get admin peers: %s", err)
	}

	return PeersToNodes(nodes)
}

// NodeInfo executes `admin_nodeInfo` RPC call and parses the response.
// Satisfies RPCClient interface.
func (h *HTTPRPCClient) NodeInfo(ip string) (*p2p.NodeInfo, error) {
	data := bytes.NewBufferString(`{"jsonrpc":"2.0","method":"admin_nodeInfo","params":[],"id":1}`)
	resp, err := http.Post("http://"+ip, "application/json", data)
	if err != nil {
		return nil, fmt.Errorf("POST RPC request: %s", err)
	}
	defer resp.Body.Close()

	nodeInfo, err := ParseNodeInfoResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get node info: %s", err)
	}

	return nodeInfo, err
}
