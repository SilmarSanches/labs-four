package webserver

import (
	"labs-four/config"
	"labs-four/internal/infra/ratelimit"
	"net/http"

	midd_ratelimit "labs-four/internal/infra/web/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Route struct {
	Method  string
	Handler http.HandlerFunc
}

type WebServer struct {
	Router   *chi.Mux
	conf     *config.AppSettings
	Handlers map[string][]Route
	rateLimiter ratelimit.RateLimitInterface
}

func NewWebServer(conf *config.AppSettings, rateLimit ratelimit.RateLimitInterface) *WebServer {
	return &WebServer{
		Router:   chi.NewRouter(),
		Handlers: make(map[string][]Route),
		conf:     conf,
		rateLimiter: rateLimit,
	}
}

func (w *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	w.Handlers[path] = append(w.Handlers[path], Route{Method: method, Handler: handler})
}

func (w *WebServer) Start() {
	w.Router.Use(middleware.Logger)

	if w.rateLimiter != nil {
		w.Router.Use(midd_ratelimit.RateLimitMiddleware(w.rateLimiter))
	}

	for path, routes := range w.Handlers {
		for _, route := range routes {
			w.Router.MethodFunc(route.Method, path, route.Handler)
		}
	}

	http.ListenAndServe(":"+w.conf.Port, w.Router)
}
