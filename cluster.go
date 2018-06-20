package main

import (
	"context"
	"log"
	"sort"
	"sync"
)

// ClusterSource represents knowledge source of
// cluster configuration.
type ClusterSource interface {
	IPs(ctx context.Context, tag string) ([]string, error)
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
func (f *Fetcher) Nodes(ctx context.Context, tag string) ([]*ClusterNode, error) {
	ips, err := f.cluster.IPs(ctx, tag)
	if err != nil {
		return nil, err
	}

	var ret []*ClusterNode
	for _, ip := range ips {
		nodeInfo, err := f.rpc.NodeInfo(ctx, ip)
		if err != nil {
			return nil, err
		}
		node := NewClusterNode(ip, nodeInfo)
		ret = append(ret, node)
	}

	return ret, nil
}

// NodePeers runs `admin_peers` command for each node.
func (f *Fetcher) NodePeers(ctx context.Context, nodes []*ClusterNode) ([]*Node, []*Link, error) {
	var (
		m       sync.Map
		wg      sync.WaitGroup
		linksMu sync.Mutex
		links   []*Link
	)
	for _, node := range nodes {
		wg.Add(1)
		go func(node *ClusterNode) {
			defer wg.Done()
			peers, err := f.rpc.AdminPeers(ctx, node.IP)
			if err != nil {
				log.Printf("[ERROR] Failed to get peers from %s\n", node.IP)
				return
			}

			for _, peer := range peers {
				m.Store(peer.ID(), peer)

				link := NewLink(node.ID, peer.ID())
				linksMu.Lock()
				links = append(links, link)
				linksMu.Unlock()
			}
		}(node)
	}
	wg.Wait()

	var ret []*Node
	m.Range(func(k, v interface{}) bool {
		ret = append(ret, v.(*Node))
		return true
	})
	sort.Slice(ret, func(i, j int) bool {
		return ret[i].ID() < ret[j].ID()
	})
	return ret, links, nil
}
