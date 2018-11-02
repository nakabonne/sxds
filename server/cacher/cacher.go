package cacher

import (
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/domain"
)

type cacher struct {
	snapshotCache cache.SnapshotCache
}

func (c *cacher) setSnapshot(nodeType domain.NodeType) error {
	snapshot, err := c.newSnapshot(nodeType)
	if err != nil {
		return err
	}
	err = c.snapshotCache.SetSnapshot(nodeType, *snapshot)
	if err != nil {
		return err
	}
	return nil
}

func (s *cacher) newSnapshot(nodeType domain.NodeType) (*cache.Snapshot, error) {
	return nil, nil
}
