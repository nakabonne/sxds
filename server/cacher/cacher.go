package cacher

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/domain"
)

type cacher struct {
	snapshotCache cache.SnapshotCache
}

func (c *cacher) setSnapshot(nodeType domain.NodeType, resources *domain.Resources) error {
	snapshot := c.newSnapshot(nodeType, resources)
	err := c.snapshotCache.SetSnapshot(nodeType, *snapshot)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacher) newSnapshot(nodeType domain.NodeType, resources *domain.Resources) *cache.Snapshot {
	endpoints := make([]cache.Resource, 0, len(resources.Endpoints))
	for _, e := range resources.Endpoints {
		endpoints = append(endpoints, e)
	}
	clusters := make([]cache.Resource, 0, len(resources.Clusters))
	for _, c := range resources.Clusters {
		clusters = append(clusters, c)
	}
	routes := make([]cache.Resource, 0, len(resources.Routes))
	for _, r := range resources.Routes {
		routes = append(routes, r)
	}
	listeners := make([]cache.Resource, 0, len(resources.Listeners))
	for _, l := range resources.Listeners {
		listeners = append(listeners, l)
	}

	snapshot := cache.NewSnapshot(resources.Version, endpoints, clusters, routes, listeners)
	return &snapshot
}
