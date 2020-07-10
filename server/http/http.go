package http

import (
	"github.com/harmony-development/legato/server/http/profile"
	"github.com/harmony-development/legato/server/http/protocol"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/sony/sonyflake"

	"github.com/harmony-development/legato/server/auth"
	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/core"
	"github.com/harmony-development/legato/server/http/hm"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/http/socket"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/state"
	"github.com/harmony-development/legato/server/storage"
)

// Server is an instance of the HTTP server
type Server struct {
	*echo.Echo
	Router     *routing.Router
	CoreAPI    *core.API
	ProfileAPI *profile.API
	Socket     *socket.Handler
	Deps       *Dependencies
}

// Dependencies are elements that a Server needs
type Dependencies struct {
	DB             db.IHarmonyDB
	AuthManager    *auth.Manager
	StorageManager *storage.Manager
	Logger         logger.ILogger
	State          *state.State
	Config         *config.Config
	Sonyflake      *sonyflake.Sonyflake
}

// New creates a new HTTP server instance
func New(deps *Dependencies) *Server {
	s := &Server{
		Echo: echo.New(),
		Socket: socket.NewHandler(&socket.Dependencies{
			DB:     deps.DB,
			Logger: deps.Logger,
			State:  deps.State,
		}),
		Deps: deps,
	}
	s.Pre(middleware.RemoveTrailingSlash())
	s.Use(middleware.CORS())
	s.Use(middleware.RecoverWithConfig(middleware.RecoverConfig{
		StackSize:       1 << 10,
		DisableStackAll: true,
	}))
	s.Use(sentryecho.New(sentryecho.Options{
		Repanic:         true,
		WaitForDelivery: false,
	}))
	s.Validator = &HarmonyValidator{
		Validator: validator.New(),
	}
	m := hm.New(deps.DB, deps.Logger)
	s.Router = &routing.Router{Middlewares: m}
	api := s.Group("/api")
	api.Use(m.WithHarmony)
	s.CoreAPI = core.New(&core.Dependencies{
		Router:         s.Router,
		APIGroup:       api,
		DB:             s.Deps.DB,
		Config:         s.Deps.Config,
		AuthManager:    s.Deps.AuthManager,
		StorageManager: s.Deps.StorageManager,
		Logger:         s.Deps.Logger,
		State:          s.Deps.State,
		Sonyflake:      s.Deps.Sonyflake,
	})
	s.ProfileAPI = profile.New(&profile.Dependencies{
		Router:         s.Router,
		APIGroup:       api,
		DB:             s.Deps.DB,
		Config:         s.Deps.Config,
		StorageManager: s.Deps.StorageManager,
		Logger:         s.Deps.Logger,
		State:          s.Deps.State,
	})
	protocol.New(&protocol.Dependencies{
		Router:      s.Router,
		APIGroup:    api,
		DB:          s.Deps.DB,
		Logger:      s.Deps.Logger,
		AuthManager: s.Deps.AuthManager,
		Sonyflake:   s.Deps.Sonyflake,
		Config:      s.Deps.Config,
		Socket:      s.Socket,
	})
	return s
}
