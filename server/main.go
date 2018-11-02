package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/nakabonne/sxds/config"
	"github.com/nakabonne/sxds/domain"
	"github.com/nakabonne/sxds/server/cacher"
	"github.com/nakabonne/sxds/server/xds"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()

	conf, err := config.New()
	if err != nil {
		panic(err)
	}

	logger, err := newLogger(conf.IsProduction())
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	// FIXME: ADSモードをフラグで制御。下記参照
	// https://github.com/envoyproxy/data-plane-api/blob/master/XDS_PROTOCOL.md#aggregated-discovery-services-ads
	isAds := false
	snapshotCache := cache.NewSnapshotCache(isAds, domain.Hasher{}, &snapshotLogger{})

	xdsServer := xds.NewServer(ctx, snapshotCache, &conf.Xds, logger)
	grpcServer, err := xdsServer.Run()
	if err != nil {
		panic(err)
	}

	cacherServer := cacher.NewServer(ctx, snapshotCache, &conf.Cacher, logger)
	cacherServer.Run()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	logger.Info("signal.received", zap.String("signal", fmt.Sprintf("%d", <-quit)))

	grpcServer.GracefulStop()
	logger.Info("init.grpc.server.config",
		zap.String("status", "stop"),
	)

	os.Exit(0)
}

func newLogger(isProduction bool) (logger *zap.Logger, err error) {
	if isProduction {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}
	return
}

// TODO: Use "go.uber.org/zap"
type snapshotLogger struct{}

// Infof logs a formatted informational message.
func (l *snapshotLogger) Infof(format string, args ...interface{}) {
	fmt.Printf("SnapshotCache: "+format+"\n", args...)
}

// Errorf logs a formatted error message.
func (l *snapshotLogger) Errorf(format string, args ...interface{}) {
	fmt.Printf("SnapshotCache: "+format+"\n", args...)
}
