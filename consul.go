package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// Consul implements ClusterSource for Consul.
type Consul struct {
	hostport string
	client   *http.Client
}

const ConsulTimeout = 5 * time.Second // it's localhost, should be in ms actually

// NewConsul creates new Consul source. It doesn't attempt
// to connect or verify if address is correct.
func NewConsul(hostport string) *Consul {
	return &Consul{
		hostport: hostport,
		client: &http.Client{
			Timeout: ConsulTimeout,
		},
	}
}

// IPs returns the list of IPs for the given datacenter and tag from Consul.
func (c *Consul) IPs(ctx context.Context, dc, tag string) ([]string, error) {
	url := fmt.Sprintf("http://%s/v1/catalog/service/statusd-rpc?tag=%s", c.hostport, tag)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("request: %s", err)
	}
	req = req.WithContext(ctx)

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("http call: %s", err)
	}
	defer resp.Body.Close()

	ips, err := ParseConsulResponse(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("get nodes list: %s", err)
	}
	return ips, err
}

// ConsulResponse describes response structure from Consul.
type ConsulResponse []*ConsulNodeInfo

// ConsulNodeInfo describes single node as reported by Consul.
type ConsulNodeInfo struct {
	ServiceAddress string
	ServicePort    int
}

// ToIP converts ConsulNodeInfo fields into hostport representation of IP.
func (c *ConsulNodeInfo) ToIP() string {
	return fmt.Sprintf("%s:%d", c.ServiceAddress, c.ServicePort)
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
