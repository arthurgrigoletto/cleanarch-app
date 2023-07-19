package webserver

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServerRouter struct {
	Path    string
	Handler http.Handler
	Method  string
}

type WebServer struct {
	Router        chi.Router
	Handlers      []WebServerRouter
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []WebServerRouter{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, method string, handler http.HandlerFunc) {
	s.Handlers = append(s.Handlers, WebServerRouter{
		Path:    path,
		Handler: handler,
		Method:  method,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		fmt.Println(handler.Method, handler.Path)
		s.Router.Method(handler.Method, handler.Path, handler.Handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
