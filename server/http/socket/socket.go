package socket

import (
	"net/http"
	"sync"

	"github.com/harmony-development/legato/server/db"
	"github.com/harmony-development/legato/server/http/socket/client"
	"github.com/harmony-development/legato/server/logger"
	"github.com/harmony-development/legato/server/state"

	"github.com/gorilla/websocket"
)

// Handler is an instance of the socket handler
type Handler struct {
	Upgrader *websocket.Upgrader
	DB       db.IHarmonyDB
	Logger   logger.ILogger
	Bus      client.Bus
	State    *state.State
}

type Dependencies struct {
	DB     db.IHarmonyDB
	Logger logger.ILogger
	State  *state.State
}

// NewHandler creates a new socket handler
func NewHandler(deps *Dependencies) *Handler {
	var bus = make(client.Bus)
	h := &Handler{
		Upgrader: &websocket.Upgrader{
			ReadBufferSize:  2048,
			WriteBufferSize: 2048,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
			EnableCompression: true,
		},
		Bus:    bus,
		State:  deps.State,
		DB:     deps.DB,
		Logger: deps.Logger,
	}
	h.Setup()
	return h
}

// Handle upgrades an HTTP request to a Client
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) *client.Client {
	conn, err := h.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		h.Logger.Exception(err)
		return nil
	}
	c := &client.Client{
		Conn:       conn,
		Bus:        h.Bus,
		Out:        make(chan []byte),
		Deregister: h.Deregister,
		RWMutex:    &sync.RWMutex{},
	}
	go c.Reader()
	go c.Writer()
	return c
}
