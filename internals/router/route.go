package router

import (
    "net/http"
)

type segment struct {
    literal string
    wild bool
}

type NextFunction func()

type Middleware func(w http.ResponseWriter, r *http.Request, next NextFunction)

type Route struct {
    handler http.HandlerFunc
    method string
    segments []segment
    middlewares []Middleware
}

func (r *Route) Use(middlewares ...Middleware) {
    r.middlewares = middlewares
}

func (route Route) Serve(w http.ResponseWriter, r *http.Request) {
    for _, middleware := range route.middlewares {
        isNextCalled := false

        middleware(w, r, func() {
            isNextCalled = true
        })

        if !isNextCalled {
            return
        }
    }

    route.handler.ServeHTTP(w, r)
}
