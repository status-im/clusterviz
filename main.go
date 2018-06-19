package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/divan/graph-experiments/graph"
)

func main() {
	var consulAddr = flag.String("consul", "localhost:8500", "Host:port for consul address to query")
	flag.Parse()

	cluster := NewConsulSource(*consulAddr)
	nodes, err := cluster.Nodes("", "eth.beta")
	if err != nil {
		log.Fatalf("Getting list of nodes: %s", err)
	}

	g := graph.NewGraph()
	for _, node := range nodes {
		_ = node
		//g.AddNode(node)
	}

	fmt.Printf("Graph has %d nodes and %d links\n", len(g.Nodes()), len(g.Links()))
}
