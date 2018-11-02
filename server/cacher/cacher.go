package cacher

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/domain"
)

type cacher struct {
	snapshotCache cache.SnapshotCache
}

func (c *cacher) setSnapshot(nodeType domain.NodeType, resources *domain.Resources) error {
	snapshot, err := c.newSnapshot(nodeType, resources)
	if err != nil {
		return err
	}
	err = c.snapshotCache.SetSnapshot(nodeType, *snapshot)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacher) newSnapshot(nodeType domain.NodeType, resources *domain.Resources) (*cache.Snapshot, error) {
	// FIXME: 実装する
	return &cache.Snapshot{}, nil
}
