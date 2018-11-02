package cacher

import (
	"context"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"go.uber.org/zap"
)

// Server is server that set cache
type Server struct {
	ctx           context.Context
	snapshotCache cache.SnapshotCache
	logger        *zap.Logger
}

// NewServer generates a Server
func NewServer(ctx context.Context, sc cache.SnapshotCache, l *zap.Logger) *Server {
	return &Server{
		ctx:           ctx,
		snapshotCache: sc,
		logger:        l,
	}
}
