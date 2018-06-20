package main

import "testing"

func TestIsClient(t *testing.T) {
	node := NewNode("id", "Statusd/v0.9.9-8141657a/linux-amd64/go1.10.2")
	got := node.IsClient()
	if got != false {
		t.Fatalf("Expect IsClient to be false")
	}

	node = NewNode("id", "StatusIM/v0.9.9-a339d7e/android-arm/go1.10.1")
	got = node.IsClient()
	if got != true {
		t.Fatalf("Expect IsClient to be true")
	}
}
