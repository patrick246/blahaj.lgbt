package main

import (
	"fmt"
	"github.com/patrick246/blahaj.lgbt/server/internal/config"
	"github.com/patrick246/blahaj.lgbt/server/internal/datasources/prometheus"
	"github.com/patrick246/blahaj.lgbt/server/internal/server"
	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"os"
	"time"

	"go.uber.org/zap"

	// This controls the maxprocs environment variable in container runtimes.
	// see https://martin.baillie.id/wrote/gotchas-in-the-go-network-packages-defaults/#bonus-gomaxprocs-containers-and-the-cfs
	_ "go.uber.org/automaxprocs"
)

func main() {
	cfg, err := config.Get()
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "config read error: %v", err)
		os.Exit(1)
	}

	if err := run(cfg); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "an error occurred: %s\n", err)
		os.Exit(1)
	}
}

func run(cfg config.Config) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	defer logger.Sync()

	sugaredLogger := logger.Sugar()

	promClient, err := api.NewClient(api.Config{
		Address: cfg.PrometheusDatasource.Address,
		RoundTripper: &prometheus.BasicAuthRoundTripper{
			Username: cfg.PrometheusDatasource.Username,
			Password: cfg.PrometheusDatasource.Password,
		},
	})

	datasource := prometheus.NewDatasource(v1.NewAPI(promClient), sugaredLogger.With("package", "prometheus-ds"))
	datasourceCache := cache.New(time.Hour, 30*time.Minute)

	srv := server.NewServer(cfg.Address, datasource, datasourceCache, sugaredLogger.With("package", "server"))

	sugaredLogger.Infow("listening", "addr", cfg.Address)

	err = srv.ListenAndServe()
	if err != nil {
		sugaredLogger.Fatalw("listen error", "error", err)
	}

	return nil
}
