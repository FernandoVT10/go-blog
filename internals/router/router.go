package router

import (
    "net/http"
    "strings"
    "regexp"
)

type Router struct {
    routes []Route
}

func NewRouter() Router {
    return Router{
        routes: make([]Route, 0),
    }
}

func (r *Router) DefineRoute(method string, path string, h http.HandlerFunc) *Route {
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

    r.routes = append(r.routes, Route{
        handler: h,
        method: method,
        segments: segments,
    })

    return &r.routes[len(r.routes) - 1]
}

func (r *Router) Get(path string, h http.HandlerFunc) *Route {
    return r.DefineRoute("GET", path, h)
}

func (r *Router) Post(path string, h http.HandlerFunc) *Route {
    return r.DefineRoute("POST", path, h)
}

func (r *Router) Put(path string, h http.HandlerFunc) *Route {
    return r.DefineRoute("PUT", path, h)
}

func (r *Router) Delete(path string, h http.HandlerFunc) *Route {
    return r.DefineRoute("DELETE", path, h)
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

        route.Serve(w, r)
        return true
    }

    return false
}
