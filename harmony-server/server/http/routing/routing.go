package routing

import (
	"time"

	"github.com/labstack/echo/v4"

	"harmony-server/server/http/hm"
)

type RequestMethod string

const (
	GET    RequestMethod = "GET"
	POST   RequestMethod = "POST"
	PUT    RequestMethod = "PUT"
	DELETE RequestMethod = "DELETE"
	PATCH  RequestMethod = "PATCH"
	ANY    RequestMethod = "ANY"
)

type RateLimit struct {
	Duration time.Duration
	Burst    int
}

type Route struct {
	*echo.Group
	Path      string
	Handler   echo.HandlerFunc
	RateLimit *RateLimit
	Auth      bool
	Method    RequestMethod
}

type Router struct {
	Middlewares *hm.Middlewares
}

func (r Router) BindRoute(g *echo.Group, endPoint Route) {
	var middleware []echo.MiddlewareFunc
	if endPoint.Auth {
		middleware = append(middleware, r.Middlewares.WithAuth)
	}
	if endPoint.RateLimit != nil {
		middleware = append(middleware, r.Middlewares.RateLimit(endPoint.RateLimit.Duration, endPoint.RateLimit.Burst))
	}
	switch endPoint.Method {
	case GET:
		{
			g.GET(endPoint.Path, endPoint.Handler, middleware...)
		}
	case POST:
		{
			g.POST(endPoint.Path, endPoint.Handler, middleware...)
		}
	case PUT:
		{
			g.PUT(endPoint.Path, endPoint.Handler, middleware...)
		}
	case DELETE:
		{
			g.DELETE(endPoint.Path, endPoint.Handler, middleware...)
		}
	case PATCH:
		{
			g.PATCH(endPoint.Path, endPoint.Handler, middleware...)
		}
	case ANY:
		{
			g.Any(endPoint.Path, endPoint.Handler, middleware...)
		}
	}
}

func (r Router) BindRoutes(g *echo.Group, endPoints []Route) {
	for _, endPoint := range endPoints {
		r.BindRoute(g, endPoint)
	}
}