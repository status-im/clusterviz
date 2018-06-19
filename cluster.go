package main

import (
	"log"
	"sort"
)

// ClusterSource represents knowledge source of
// cluster configuration.
type ClusterSource interface {
	IPs(dc, tag string) ([]string, error)
}

// Fetched implements data fetching from multiple sources
// to get the needed peers data.
type Fetcher struct {
	cluster ClusterSource
	rpc     RPCClient
}

// NewFetcher creates new Fetcher.
func NewFetcher(cluster ClusterSource, rpc RPCClient) *Fetcher {
	return &Fetcher{
		cluster: cluster,
		rpc:     rpc,
	}
}

// Nodes returns the list of nodes for the given datacentre 'dc' and tag.
func (f *Fetcher) Nodes(dc, tag string) ([]*ClusterNode, error) {
	ips, err := f.cluster.IPs(dc, tag)
	if err != nil {
		return nil, err
	}

	var ret []*ClusterNode
	for _, ip := range ips {
		nodeInfo, err := f.rpc.NodeInfo(ip)
		if err != nil {
			return nil, err
		}
		node := NewClusterNode(ip, nodeInfo)
		ret = append(ret, node)
	}

	return ret, nil
}

// NodePeers runs `admin_peers` command for each node.
func (f *Fetcher) NodePeers(nodes []*ClusterNode) ([]*Node, []*Link, error) {
	m := make(map[string]*Node)
	var links []*Link
	for _, node := range nodes {
		// TODO: run concurrently
		peers, err := f.rpc.AdminPeers(node.IP)
		if err != nil {
			log.Printf("[ERROR] Failed to get peers from %s\n", node.IP)
			continue
		}

		for _, peer := range peers {
			m[peer.ID()] = peer

			link := NewLink(node.ID, peer.ID())
			links = append(links, link)
		}
	}

	var ret []*Node
	for _, node := range m {
		ret = append(ret, node)
	}
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID() < ret[j].ID()
	})
	return ret, links, nil
}
