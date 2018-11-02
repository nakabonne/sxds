package domain

import (
	"strings"

	"github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
)

// Hasher is a calculator that calculates the node id.
type Hasher struct{}

// ID generates a key for cache
// Snapshot is cached per node type.
// ID reads the node type from the node id.
// Also, envoy id needs to be as follows. "sidecar-app1"
func (h Hasher) ID(node *core.Node) string {
	if node == nil {
		return "unknown"
	}
	nodeType := strings.Split(node.Id, "-")[0]
	return nodeType
}
