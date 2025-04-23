package router

import (
    "net/http"
    "strings"
    "regexp"
)

type segment struct {
    literal string
    wild bool
}

type route struct {
    handler http.HandlerFunc
    method string
    segments []segment
}

type Router struct {
    routes []route
}

func NewRouter() Router {
    return Router{
        routes: make([]route, 0),
    }
}

func (r *Router) DefineRoute(method string, path string, h http.HandlerFunc) {
    if !strings.HasPrefix(path, "/") {
        path = "/" + path
    }

    splittedStr := strings.Split(path, "/")
    segments := make([]segment, 0)

    for index, str := range splittedStr {
        if index == 0 {
            segments = append(segments, segment{
                literal: "/",
                wild: false,
            })
            continue
        }
        r, err := regexp.Compile("^{.*}$")
        if err != nil {
            continue
        }

        if r.MatchString(str) {
            l := len(str) - 1
            literal := str[1:l]

            segments = append(segments, segment{
                literal: literal,
                wild: true,
            })
        } else {
            segments = append(segments, segment{
                literal: str,
                wild: false,
            })
        }
    }

    r.routes = append(r.routes, route{
        handler: h,
        method: method,
        segments: segments,
    })
}

func (r *Router) Get(path string, h http.HandlerFunc) {
    r.DefineRoute("GET", path, h)
}

func (r *Router) Post(path string, h http.HandlerFunc) {
    r.DefineRoute("POST", path, h)
}

func (r *Router) Put(path string, h http.HandlerFunc) {
    r.DefineRoute("PUT", path, h)
}

func (r *Router) Delete(path string, h http.HandlerFunc) {
    r.DefineRoute("DELETE", path, h)
}

// returns true when path is found
func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) bool {
    rPath := r.URL.Path

    if !strings.HasPrefix(rPath, "/") {
        rPath = "/" + rPath
    }

    pathSegments := strings.Split(rPath, "/")
    pathSegments[0] = "/"

    Loop:
    for _, route := range router.routes {
        if route.method != r.Method || len(route.segments) != len(pathSegments) {
            continue
        }

        for i, segment := range route.segments {
            if segment.wild {
                val := pathSegments[i]
                r.SetPathValue(segment.literal, val)
            } else if pathSegments[i] != segment.literal {
                continue Loop
            }
        }

        route.handler.ServeHTTP(w, r)
        return true
    }

    return false
}
