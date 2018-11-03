package domain

import (
	"testing"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

var (
	testNodeType1 = "sidecar"
	testNodeType2 = "router"
	testNodeType3 = "ingress"
	testNode1     = &core.Node{Id: testNodeType1 + "-app1"}
	testNode2     = &core.Node{Id: testNodeType2 + "-app1"}
	testNode3     = &core.Node{Id: testNodeType3 + "-app1"}
)

func TestID(t *testing.T) {
	hasher := &Hasher{}

	if got, want := hasher.ID(testNode1), testNodeType1; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got, want := hasher.ID(testNode2), testNodeType2; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
	if got, want := hasher.ID(testNode3), testNodeType3; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}
}
