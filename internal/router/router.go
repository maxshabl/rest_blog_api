package router

import (
	"blog/internal/middelware"
	"fmt"
	"net/http"
	"os"
)

type Route struct {
	http.Handler
	Action      http.Handler
	Middlewares []middelware.Middleware
	Path        string
	Method      string
}

func (r *Route) applyMiddelware() {
	for i := len(r.Middlewares) - 1; i >= 0; i-- {
		r.Action = r.Middlewares[i](r.Action)
	}
	r.Action = r.httpMethodMiddelware(r.Action)
}

func (r *Route) httpMethodMiddelware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		if req.Method == r.Method {
			next.ServeHTTP(w, req)
			return
		}
		fmt.Fprintln(os.Stderr, "wrong http method")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Page not found")
	})
}

func ConfigureRoutes(routes []Route) *http.ServeMux {
	mux := http.NewServeMux()
	for _, route := range routes {
		route.applyMiddelware()
		mux.Handle(route.Path, route.Action)
	}
	return mux
}
