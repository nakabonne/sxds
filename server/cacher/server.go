package cacher

import (
	"context"
	"fmt"
	"net/http"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/julienschmidt/httprouter"
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

func (s *Server) Run() {
	router := httprouter.New()
	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	defer s.logger.Sync()

	router.PUT("/resources/:node_type", putResources)
	go func(p int, r *httprouter.Router) {
		if err := http.ListenAndServe(fmt.Sprintf(":%d", p), r); err != nil {
			// TODO: error handling
			s.logger.Error("cacher server closed", zap.Error(err))
		}
	}(conf.Cacher.Port, router)

	s.logger.Info("cacher server is listening",
		zap.Any("conf", conf.Cacher),
	)
	return
}

func putResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
