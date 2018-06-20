package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/ethereum/go-ethereum/p2p"
)

// RPCClient defines subset of client that
// can call needed methods to geth's RPC server.
type RPCClient interface {
	AdminPeers(ctx context.Context, ip string) ([]*Node, error)
	NodeInfo(ctx context.Context, ip string) (*p2p.NodeInfo, error)
}

// HTTPRPCClient implements RPCClient for
// HTTP transport.
type HTTPRPCClient struct {
	client *http.Client
}

const HTTPRPCTimeout = 3 * time.Second

// NewHTTPRPCClient creates new HTTP RPC client for eth JSON-RPC server.
func NewHTTPRPCClient() *HTTPRPCClient {
	return &HTTPRPCClient{
		client: &http.Client{
			Timeout: ConsulTimeout,
		},
	}
}

// AdminPeers executes `admin_peers` RPC call and parses the response.
// Satisfies RPCClient interface.
func (h *HTTPRPCClient) AdminPeers(ctx context.Context, ip string) ([]*Node, error) {
	r, err := h.postMethod(ctx, ip, "admin_peers")
	if err != nil {
		return nil, fmt.Errorf("rpc admin_peers: %s", err)
	}
	defer r.Close()

	nodes, err := ParsePeersResponse(r)
	if err != nil {
		return nil, fmt.Errorf("get admin peers: %s", err)
	}

	return PeersToNodes(nodes)
}

// NodeInfo executes `admin_nodeInfo` RPC call and parses the response.
// Satisfies RPCClient interface.
func (h *HTTPRPCClient) NodeInfo(ctx context.Context, ip string) (*p2p.NodeInfo, error) {
	r, err := h.postMethod(ctx, ip, "admin_nodeInfo")
	if err != nil {
		return nil, fmt.Errorf("rpc admin_nodeInfo: %s", err)
	}
	defer r.Close()

	nodeInfo, err := ParseNodeInfoResponse(r)
	if err != nil {
		return nil, fmt.Errorf("get node info: %s", err)
	}

	return nodeInfo, err
}

// postMethod performs POST RPC request for single method RPC calls without params.
// it reads body and return the whole answer.
func (h *HTTPRPCClient) postMethod(ctx context.Context, ip, method string) (io.ReadCloser, error) {
	payload := fmt.Sprintf("{\"jsonrpc\":\"2.0\",\"method\":\"%s\",\"params\":[],\"id\":1}", method)
	data := bytes.NewBufferString(payload)
	url := "http://" + ip
	req, err := http.NewRequest("POST", url, data)
	if err != nil {
		return nil, fmt.Errorf("request: %s", err)
	}
	req = req.WithContext(ctx)

	resp, err := h.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("POST RPC request: %s", err)
	}

	return resp.Body, nil
}
