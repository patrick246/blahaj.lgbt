package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/patrick246/blahaj.lgbt/server/internal/datasources"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/patrickmn/go-cache"
)

const (
	internalServerError = `Internal Server Error`
	globalCacheKey      = "global"
	storeCacheKey       = "stores"
)

var (
	httpRequestDurationSeconds = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:        "http_request_duration_seconds",
		Help:        "Histogram for the time that the server took to handle the request",
		ConstLabels: nil,
		Buckets:     prometheus.DefBuckets,
	}, []string{"path", "code", "method"})
)

func init() {
	prometheus.MustRegister(httpRequestDurationSeconds)
}

type Server struct {
	server     http.Server
	datasource datasources.Datasource
	cache      *cache.Cache
	log        *zap.SugaredLogger
}

func NewServer(
	addr string,
	datasource datasources.Datasource,
	cache *cache.Cache,
	log *zap.SugaredLogger,
) *Server {
	m := mux.NewRouter()
	srv := &Server{
		datasource: datasource,
		cache:      cache,
		log:        log,
		server: http.Server{
			Addr:              addr,
			Handler:           m,
			ReadTimeout:       10 * time.Second,
			ReadHeaderTimeout: 10 * time.Second,
		},
	}

	m.Handle("/api/availability/countries", promhttp.InstrumentHandlerDuration(
		httpRequestDurationSeconds.MustCurryWith(prometheus.Labels{
			"path": "/api/availability/countries",
		}),
		http.HandlerFunc(srv.handleGlobalAvailabilityRequest)),
	)

	m.Handle("/api/availability/stores", promhttp.InstrumentHandlerDuration(
		httpRequestDurationSeconds.MustCurryWith(prometheus.Labels{
			"path": "/api/availability/stores",
		}),
		http.HandlerFunc(srv.handleAvailabilityRequest)),
	)

	m.Handle("/metrics", promhttp.Handler())

	m.HandleFunc("/ready", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(200)
		_, _ = writer.Write([]byte("ready"))
	})

	return srv
}

func (s *Server) handleGlobalAvailabilityRequest(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if cacheEntry, ok := s.cache.Get(globalCacheKey); ok {
		err := json.NewEncoder(writer).Encode(cacheEntry)
		if err != nil {
			s.log.Errorw("json marshalling error", "error", err)
			http.Error(writer, internalServerError, http.StatusInternalServerError)
			return
		}

		return
	}

	globalAvailability, err := s.datasource.GlobalAvailability(req.Context())
	if err != nil {
		s.log.Errorw("datasource error", "error", err)
		http.Error(writer, internalServerError, http.StatusInternalServerError)
		return
	}

	s.cache.SetDefault(globalCacheKey, globalAvailability)

	err = json.NewEncoder(writer).Encode(globalAvailability)
	if err != nil {
		s.log.Errorw("json marshalling error", "error", err)
		http.Error(writer, internalServerError, http.StatusInternalServerError)
	}
}

func (s *Server) handleAvailabilityRequest(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	if cacheEntry, ok := s.cache.Get(storeCacheKey); ok {
		err := json.NewEncoder(writer).Encode(cacheEntry)
		if err != nil {
			s.log.Errorw("json marshalling error", "error", err)
			http.Error(writer, internalServerError, http.StatusInternalServerError)
			return
		}
		return
	}

	availability, err := s.datasource.Availability(req.Context())
	if err != nil {
		s.log.Errorw("datasource error", "error", err)
		http.Error(writer, internalServerError, http.StatusInternalServerError)
		return
	}

	s.cache.SetDefault(storeCacheKey, availability)

	err = json.NewEncoder(writer).Encode(availability)
	if err != nil {
		s.log.Errorw("json marshalling error", "error", err)
		http.Error(writer, internalServerError, http.StatusInternalServerError)
	}
}

func (s *Server) ListenAndServe() error {
	return s.server.ListenAndServe()
}
