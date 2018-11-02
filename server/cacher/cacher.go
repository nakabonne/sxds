package cacher

import (
	"io"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/domain"
)

type cacher struct {
	snapshotCache cache.SnapshotCache
}

func (c *cacher) setSnapshot(nodeType domain.NodeType, resources io.Reader) error {
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

func (s *cacher) newSnapshot(nodeType domain.NodeType, resources io.Reader) (*cache.Snapshot, error) {
	// FIXME: resourcesのバリデーション
	return nil, nil
}
