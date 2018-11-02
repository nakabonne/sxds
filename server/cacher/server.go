package cacher

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/julienschmidt/httprouter"
	"github.com/nakabonne/sxds/config"
	"go.uber.org/zap"
)

type (
	// Server is server that set cache
	Server struct {
		ctx           context.Context
		snapshotCache cache.SnapshotCache
		conf          *config.Cacher
		logger        *zap.Logger
	}

	exception struct {
		Message string `json:"message"`
	}
)

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

	router.PUT("/resources/:node_type", s.putResources)
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

func (s *Server) putResources(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	nodeType := ps.ByName("node_type")
	if nodeType == "" {
		msg := "Node type does not exist in path"
		s.logger.Error(msg)
		w.WriteHeader(400)
		writeJSON(w, exception{Message: msg})
		return
	}

	c := cacher{snapshotCache: s.snapshotCache}
	if err := c.setSnapshot(nodeType, r.Body); err != nil {
		msg := "Faild to cache resources"
		s.logger.Error(msg, zap.Error(err), zap.Any("node_type", nodeType))
		w.WriteHeader(500)
		writeJSON(w, exception{Message: msg})
		return
	}
	w.WriteHeader(200)
	w.Write([]byte("true"))
}

func writeJSON(w http.ResponseWriter, res interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	r, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		return err
	}
	w.Write(r)
	return nil
}
