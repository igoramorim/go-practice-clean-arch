package webserver

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func New(port string) *WebServer {
	r := chi.NewRouter()

	// TODO: Add middlewares

	return &WebServer{
		router:   r,
		handlers: make(map[string]http.HandlerFunc),
		port:     port,
	}
}

type WebServer struct {
	router   chi.Router
	handlers map[string]http.HandlerFunc
	port     string
}

func (s *WebServer) AddHandler(method, path string, handler http.HandlerFunc) {
	s.router.MethodFunc(method, path, handler)
}

func (s *WebServer) Run() error {
	return http.ListenAndServe(fmt.Sprintf(":%s", s.port), s.router)
}
