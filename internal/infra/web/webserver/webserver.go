package webserver

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"net/http"
)

type Handler struct {
	fn     http.HandlerFunc
	method string
	path   string
}

type WebServer struct {
	Router        chi.Router
	Handlers      []Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      make([]Handler, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
	fmt.Println("Adding handler", path, method)
	s.Handlers = append(s.Handlers, Handler{
		fn:     handler,
		method: method,
		path:   path,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	fmt.Println(s.Handlers)
	for _, handler := range s.Handlers {
		switch handler.method {
		case "GET":
			s.Router.Get(handler.path, handler.fn)
		case "POST":
			s.Router.Post(handler.path, handler.fn)
		default:
			s.Router.NotFound(http.NotFound)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
