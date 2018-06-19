package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/divan/graph-experiments/graph"
)

func main() {
	var (
		consulAddr     = flag.String("consul", "localhost:8500", "Host:port for consul address to query")
		testMode       = flag.Bool("test", false, "Test mode (use local test data)")
		port           = flag.String("port", ":20002", "Port to bind server to")
		updateInterval = flag.Duration("i", 10*time.Second, "Update interval")
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

	ws := NewWSServer(fetcher, *updateInterval)
	ws.refresh()

	log.Printf("Starting web server...")
	startWeb(ws, *port)
}

// BuildGraph performs new cycle of updating data from
// fetcher source and populating graph object.
func BuildGraph(fetcher *Fetcher) (*graph.Graph, error) {
	nodes, err := fetcher.Nodes("", "eth.beta")
	if err != nil {
		return nil, fmt.Errorf("list of ips: %s", err)
	}

	peers, links, err := fetcher.NodePeers(nodes)
	if err != nil {
		return nil, fmt.Errorf("get node peers: %s", err)
	}

	g := graph.NewGraph()
	for _, peer := range peers {
		AddNode(g, peer)
	}
	for _, link := range links {
		AddLink(g, link.FromID, link.ToID)
	}

	fmt.Printf("Graph has %d nodes and %d links\n", len(g.Nodes()), len(g.Links()))
	return g, nil
}

// AddNode is a wrapper around adding node to graph with proper checking for duplicates.
func AddNode(g *graph.Graph, node *Node) {
	if _, err := g.NodeByID(node.ID()); err == nil {
		// already exists
		return
	}
	g.AddNode(node)
}

// AddLink is a wrapper around adding link to graph with proper checking for duplicates.
func AddLink(g *graph.Graph, fromID, toID string) {
	if g.LinkExistsByID(fromID, toID) {
		return
	}
	g.AddLinkByIDs(fromID, toID)
}
