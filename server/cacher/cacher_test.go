package cacher

import (
	"testing"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/domain"
	"github.com/stretchr/testify/assert"
)

var (
	allNodeTypes = []domain.NodeType{"sidecar", "router", "ingress"}

	simpleVersion      = "1.0"
	simpleClusterName  = "cluster1"
	simpleRouteName    = "route1"
	simpleListenerName = "listener1"

	simpleCluster  = &api.Cluster{Name: simpleClusterName}
	simpleEndpoint = &api.ClusterLoadAssignment{ClusterName: simpleClusterName}
	simpleRoute    = &api.RouteConfiguration{Name: simpleRouteName}
	simpleListener = &api.Listener{Name: simpleListenerName}
)

func TestNewSnapshot(t *testing.T) {
	nodeType := allNodeTypes[0]

	resources := &domain.Resources{
		Clusters:  []*api.Cluster{simpleCluster},
		Endpoints: []*api.ClusterLoadAssignment{simpleEndpoint},
		Routes:    []*api.RouteConfiguration{simpleRoute},
		Listeners: []*api.Listener{simpleListener},
	}
	snapshotCache := cache.NewSnapshotCache(false, domain.Hasher{}, nil)
	c := &cacher{snapshotCache: snapshotCache}
	snapshot := c.newSnapshot(nodeType, resources)

	clusters := snapshot.GetResources(cache.ClusterType)
	endpoints := snapshot.GetResources(cache.EndpointType)
	routes := snapshot.GetResources(cache.RouteType)
	listeners := snapshot.GetResources(cache.ListenerType)

	assert.Equal(t, clusters[simpleClusterName], simpleCluster)
	assert.Equal(t, endpoints[simpleClusterName], simpleEndpoint)
	assert.Equal(t, routes[simpleRouteName], simpleRoute)
	assert.Equal(t, listeners[simpleListenerName], simpleListener)
}
