package main

import (
    "net/http"
    "fmt"
    "github.com/FernandoVT10/go-blog/internals/routes"
    "github.com/FernandoVT10/go-blog/internals/db"
)

const PORT = 3000
const STATIC_DIR = "./static"
const BUILD_DIR = "./build"
const UPLOADS_DIR = "./uploads"

func main() {
    router := http.NewServeMux()

    staticFs := http.FileServer(http.Dir(STATIC_DIR))
    router.Handle("/static/", http.StripPrefix("/static/", staticFs))

    buildFs := http.FileServer(http.Dir(BUILD_DIR))
    router.Handle("/build/", http.StripPrefix("/build/", buildFs))

    uploadsFs := http.FileServer(http.Dir(UPLOADS_DIR))
    router.Handle("/uploads/", http.StripPrefix("/uploads/", uploadsFs))

    router.Handle("/", routes.GetRoutes())

    db.Connect()

    fmt.Println("[INFO] Server listening on port", PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), router)

    if err != nil {
        panic(err)
    }
}
