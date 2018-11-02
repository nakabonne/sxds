package cacher

import (
	"context"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/config"
	"go.uber.org/zap"
)

// Server is server that set cache
type Server struct {
	ctx           context.Context
	snapshotCache cache.SnapshotCache
	conf          *config.Cacher
	logger        *zap.Logger
}

// NewServer generates a Server
func NewServer(ctx context.Context, sc cache.SnapshotCache, conf *config.Cacher, l *zap.Logger) *Server {
	return &Server{
		ctx:           ctx,
		snapshotCache: sc,
		conf:          conf,
		logger:        l,
	}
}
