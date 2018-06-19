package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/divan/graph-experiments/graph"
	"github.com/ethereum/go-ethereum/p2p"
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
		_ = peer
		//AddPeer(g, "", node.PI)
	}
	for _, link := range links {
		_ = link
		//AddPeer(g, "", node.PI)
	}

	fmt.Printf("Graph has %d nodes and %d links\n", len(g.Nodes()), len(g.Links()))
}

func AddPeer(g *graph.Graph, fromID string, to *p2p.PeerInfo) {
	toID := to.ID
	addNode(g, fromID, false)
	//addNode(g, toID, isClient(to.Name))
	addNode(g, toID, false)

	if g.LinkExistsByID(fromID, toID) {
		return
	}
	if to.Network.Inbound == false {
		g.AddLinkByIDs(fromID, toID)
	} else {
		g.AddLinkByIDs(toID, fromID)
	}
}

func addNode(g *graph.Graph, id string, client bool) {
	if _, err := g.NodeByID(id); err == nil {
		// already exists
		return
	}
	//node := NewNode(id)
	//g.AddNode(node)
}
