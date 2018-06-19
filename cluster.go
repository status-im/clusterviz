package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

// ClusterSource represents knowledge source of
// cluster configuration.
type ClusterSource interface {
	Nodes(dc, tag string) ([]*Node, error)
}

// ConsulSource implements Consul clients that fetches
// actual information about hosts in cluster.
type ConsulSource struct {
	hostport string
}

// NewConsulSource creates new Consul source. It doesn't attempt
// to connect or verify if address is correct.
func NewConsulSource(hostport string) ClusterSource {
	return &ConsulSource{
		hostport: hostport,
	}
}

// Node returns the list of nodes for the given datacentre 'dc' and tag.
// Satisfies ClusterSource interface.
func (c *ConsulSource) Nodes(dc, tag string) ([]*Node, error) {
	url := fmt.Sprintf("http://%s/v1/catalog/service/statusd-rpc?tag=%s", c.hostport, tag)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http call failed: %s", err)
	}
	defer resp.Body.Close()

	ips, err := ParseConsulResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get nodes list: %s", err)
	}

	var ret []*Node
	for _, ip := range ips {
		// TODO: run concurrently
		rpc := NewHTTPRPCClient(ip)
		nodes, err := rpc.AdminPeers()
		if err != nil {
			log.Println("[ERROR] Failed to get peers from %s", ip)
			continue
		}
		ret = append(ret, nodes...)
	}

	return ret, nil

	return nil, errors.New("TBD")
}

// ConsulResponse describes response structure from Consul.
type ConsulResponse []*ConsulNodeInfo

// ConsulNodeInfo describes single node as reported by Consul.
type ConsulNodeInfo struct {
	ServiceAddress string
	ServicePort    string
}

// ToIP converts ConsulNodeInfo fields into hostport representation of IP.
func (c *ConsulNodeInfo) ToIP() string {
	return fmt.Sprintf("%s:%s", c.ServiceAddress, c.ServicePort)
}

// ParseConsulResponse parses JSON output from Consul response with
// the list of service and extracts IP addresses.
func ParseConsulResponse(r io.Reader) ([]string, error) {
	var resp ConsulResponse
	err := json.NewDecoder(r).Decode(&resp)
	if err != nil {
		return nil, fmt.Errorf("unmarshal Consul JSON response: %s", err)
	}

	ret := make([]string, len(resp))
	for i := range resp {
		ret[i] = resp[i].ToIP()
	}
	return ret, nil
}
