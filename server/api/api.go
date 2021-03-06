package api

import (
	"net/http"
	"time"

	authv1 "github.com/harmony-development/legato/gen/auth/v1"
	chatv1 "github.com/harmony-development/legato/gen/chat/v1"
	"github.com/harmony-development/legato/server/api/authsvc"
	"github.com/harmony-development/legato/server/api/chat"
	"github.com/harmony-development/legato/server/api/chat/v1/permissions"
	"github.com/harmony-development/legato/server/api/middleware"
	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/attachments/backend"
	"github.com/harmony-development/legato/server/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sony/sonyflake"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Dependencies struct {
	Logger         logger.ILogger
	DB             db.IHarmonyDB
	Sonyflake      *sonyflake.Sonyflake
	AuthManager    *auth.Manager
	Config         *config.Config
	Permissions    *permissions.Manager
	StorageBackend backend.AttachmentBackend
}

// API contains the component of the server responsible for APIs
type API struct {
	Dependencies
	GrpcServer       *grpc.Server
	GrpcWebServer    *grpcweb.WrappedGrpcServer
	PrometheusServer *http.Server
	ChatSvc          *chat.Service
}

// New creates a new API instance
func New(deps Dependencies) *API {
	api := &API{
		Dependencies: deps,
	}
	m := middleware.New(middleware.Dependencies{
		Logger: deps.Logger,
		DB:     deps.DB,
		Perms:  api.Permissions,
	})
	api.GrpcServer = grpc.NewServer(
		grpc_middleware.WithUnaryServerChain(
			m.HarmonyContextInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(m.RecoveryFunc)),
			grpc_prometheus.UnaryServerInterceptor,
			m.ErrorInterceptor,
			m.RateLimitInterceptor,
			m.ValidatorInterceptor,
			m.AuthInterceptor,
			m.LocationInterceptor,
			m.GuildPermissionInterceptor,
			m.LoggingInterceptor,
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(m.RecoveryFunc)),
			grpc_prometheus.StreamServerInterceptor,
			m.HarmonyContextInterceptorStream,
			m.ErrorInterceptorStream,
			m.RateLimitStreamInterceptorStream,
		))
	api.GrpcWebServer = grpcweb.WrapServer(api.GrpcServer, grpcweb.WithOriginFunc(func(_ string) bool {
		return true
	}), grpcweb.WithWebsockets(true), grpcweb.WithWebsocketOriginFunc(func(req *http.Request) bool {
		return true
	}), grpcweb.WithWebsocketPingInterval(10*time.Second))
	prometheusMux := http.NewServeMux()
	prometheusMux.Handle("/metrics", promhttp.Handler())
	api.PrometheusServer = &http.Server{
		Handler: prometheusMux,
	}

	chatv1.RegisterChatServiceServer(api.GrpcServer, chat.New(&chat.Dependencies{
		DB:             api.DB,
		Logger:         api.Logger,
		Sonyflake:      api.Sonyflake,
		Perms:          api.Permissions,
		Config:         deps.Config,
		StorageBackend: deps.StorageBackend,
	}).V1)
	authv1.RegisterAuthServiceServer(api.GrpcServer, authsvc.New(&authsvc.Dependencies{
		DB:          api.DB,
		Logger:      api.Logger,
		Sonyflake:   api.Sonyflake,
		AuthManager: api.AuthManager,
		Config:      api.Config,
	}))
	reflection.Register(api.GrpcServer)
	grpc_prometheus.Register(api.GrpcServer)
	grpc_prometheus.EnableHandlingTimeHistogram()

	return api
}
