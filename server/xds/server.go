package xds

import (
	"context"
	"fmt"
	"net"

	api "github.com/envoyproxy/go-control-plane/envoy/api/v2"
	discovery "github.com/envoyproxy/go-control-plane/envoy/service/discovery/v2"
	"github.com/envoyproxy/go-control-plane/pkg/cache"
	xds "github.com/envoyproxy/go-control-plane/pkg/server"
	"github.com/nakabonne/sxds/config"
	"go.uber.org/zap"

	"google.golang.org/grpc"
)

// Server is xDS Server
type Server struct {
	ctx           context.Context
	snapshotCache cache.SnapshotCache
	conf          *config.Xds
	logger        *zap.Logger
}

// NewServer creates Server
func NewServer(ctx context.Context, cache cache.SnapshotCache, conf *config.Xds, logger *zap.Logger) *Server {
	return &Server{
		ctx:           ctx,
		snapshotCache: cache,
		conf:          conf,
		logger:        logger,
	}
}

// Run runs xDS gRPC server
func (s *Server) Run() (*grpc.Server, error) {
	grpcServer := grpc.NewServer()
	server := xds.NewServer(s.snapshotCache, nil)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.conf.Port))
	if err != nil {
		return nil, err
	}

	discovery.RegisterAggregatedDiscoveryServiceServer(grpcServer, server)
	api.RegisterEndpointDiscoveryServiceServer(grpcServer, server)
	api.RegisterClusterDiscoveryServiceServer(grpcServer, server)
	api.RegisterRouteDiscoveryServiceServer(grpcServer, server)
	api.RegisterListenerDiscoveryServiceServer(grpcServer, server)

	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			s.logger.Error("xDS server closed", zap.Error(err))
		}
	}()
	s.logger.Info("xDS server is listening",
		zap.Any("conf", s.conf),
	)

	return grpcServer, nil
}
