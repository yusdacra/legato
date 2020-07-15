package v1

import (
	"time"

	"github.com/harmony-development/legato/server/config"
	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/routing"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/state"
	"github.com/harmony-development/legato/server/storage"
)

// Handlers for ProfileKit
type Handlers struct {
	*Dependencies
}

// Dependencies are the elements that ProfileKit handlers need
type Dependencies struct {
	DB             db.IHarmonyDB
	Config         *config.Config
	StorageManager *storage.Manager
	Logger         logger.ILogger
	State          *state.State
}

// New creates a new set of Handlers
func New(deps *Dependencies) *Handlers {
	return &Handlers{
		Dependencies: deps,
	}
}

// MakeRoutes creates the routes for ProfileKit
func (h Handlers) MakeRoutes() []routing.Route {
	return []routing.Route{
		{
			Path:    "/users/:user_id",
			Handler: h.GetUser,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    2,
			},
			Auth:     true,
			Location: routing.LocationUser,
		},
		{
			Path:    "/users/~/avatar",
			Handler: h.AvatarUpdate,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 10 * time.Second,
				Burst:    2,
			},
			Auth:     true,
			Location: routing.LocationNone,
		},
		{
			Path:    "/users/~/username",
			Handler: h.UsernameUpdate,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    2,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   UsernameUpdateData{},
		},
		{
			Path:    "/users/~/status",
			Handler: h.StatusUpdate,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 5 * time.Second,
				Burst:    4,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   StatusUpdateData{},
		},
		{
			Path:    "/users/~/metadata",
			Handler: h.GetUserMetadata,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 3 * time.Second,
				Burst:    5,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   GetUserMetadataData{},
		},
		{
			Path:    "/users/~/guild/move",
			Handler: h.MoveGuild,
			Method:  routing.PATCH,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   MoveGuildData{},
		},
		{
			Path:    "/users/~/guild/remove",
			Handler: h.RemoveGuild,
			Method:  routing.DELETE,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   RemoveGuildData{},
		},
		{
			Path:    "/users/~/guild/add",
			Handler: h.AddGuild,
			Method:  routing.PUT,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    5,
			},
			Auth:     true,
			Location: routing.LocationNone,
			Schema:   AddGuildData{},
		},
		{
			Path:    "/users/~/guilds",
			Handler: h.GetGuildList,
			Method:  routing.GET,
			RateLimit: &routing.RateLimit{
				Duration: 1 * time.Second,
				Burst:    3,
			},
			Auth:     true,
			Location: routing.LocationNone,
		},
	}
}
