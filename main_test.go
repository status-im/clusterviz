package main

import (
	"context"
	"testing"
)

func TestGraphCreate(t *testing.T) {
	cluster := NewMockConsulSource()
	rpc := NewMockRPCClient()

	f := NewFetcher(cluster, rpc)

	ctx := context.Background()
	nodes, err := f.Nodes(ctx, "", "eth.beta")
	if err != nil {
		t.Fatal(err)
	}

	got := len(nodes)
	expected := 15
	if got != expected {
		t.Fatalf("Expected %d nodes, got %d", expected, got)
	}

	peers, links, err := f.NodePeers(ctx, nodes)
	if err != nil {
		t.Fatal(err)
	}

	got = len(peers)
	expected = 49
	if got != expected {
		t.Fatalf("Expected %d nodes, got %d", expected, got)
	}

	got = len(links)
	expected = 200
	if got != expected {
		t.Fatalf("Expected %d links, got %d", expected, got)
	}
}
