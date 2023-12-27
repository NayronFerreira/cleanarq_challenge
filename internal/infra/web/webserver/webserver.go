package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Handler struct {
	Func   http.HandlerFunc
	Method string
}

type WebServer struct {
	Router        chi.Router
	Handlers      map[string]Handler
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(JSONMiddleware)
	return &WebServer{
		Router:        r,
		Handlers:      make(map[string]Handler),
		WebServerPort: serverPort,
	}
}

func (ws *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
	ws.Router.MethodFunc(method, path, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		handler(w, r)
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	for path, handler := range s.Handlers {
		s.Router.Method(handler.Method, path, handler.Func)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}

func JSONMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
