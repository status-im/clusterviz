package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/divan/graph-experiments/graph"
)

func main() {
	var (
		consulAddr = flag.String("consul", "localhost:8500", "Host:port for consul address to query")
		testMode   = flag.Bool("test", false, "Test mode (use local test data)")
	)
	flag.Parse()

	var rpc RPCClient
	rpc = NewHTTPRPCClient()
	if *testMode {
		rpc = NewMockRPCClient()
	}
	var cluster ClusterSource
	cluster = NewConsul(*consulAddr)
	if *testMode {
		cluster = NewMockConsulSource()
	}

	fetcher := NewFetcher(cluster, rpc)
	nodes, err := fetcher.Nodes("", "eth.beta")
	if err != nil {
		log.Fatalf("Getting list of ips: %s", err)
	}

	peers, links, err := fetcher.NodePeers(nodes)
	if err != nil {
		log.Fatalf("Getting list of ips: %s", err)
	}

	g := graph.NewGraph()
	for _, peer := range peers {
		if _, err := g.NodeByID(peer.ID()); err == nil {
			// already exists
			continue
		}
		g.AddNode(peer)
	}
	for _, link := range links {
		AddLink(g, link.FromID, link.ToID)
	}

	fmt.Printf("Graph has %d nodes and %d links\n", len(g.Nodes()), len(g.Links()))
}

// AddLink is a wrapper around adding link to graph with proper checking for duplicates.
func AddLink(g *graph.Graph, fromID, toID string) {
	if g.LinkExistsByID(fromID, toID) {
		return
	}
	g.AddLinkByIDs(fromID, toID)
}
