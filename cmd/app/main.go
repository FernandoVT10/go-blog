package main

import (
    "net/http"
    "path"
    "fmt"
    "os"
    "github.com/FernandoVT10/go-blog/internals/html"
    "github.com/FernandoVT10/go-blog/internals/db"
)

const PUBLIC_DIR = "./public"
const PORT = 3000
const UPLOADS_DIR = "./uploads"

func main() {
    mux := http.NewServeMux()

    uploadsFs := http.FileServer(http.Dir(UPLOADS_DIR))
    mux.Handle("/uploads/", http.StripPrefix("/uploads/", uploadsFs))

    router := GetRoutes()

    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        rPath := r.URL.Path

        if(router.ServeHTTP(w, r)) {
            return
        }

        filepath := path.Join(PUBLIC_DIR, rPath)
        if _, err := os.Stat(filepath); err == nil {
            http.ServeFile(w, r, filepath)
            return
        }

        notFoundNode := html.NotFound()
        notFoundNode.Render(w)
    })

    db.Connect()

    fmt.Println("[INFO] Server listening on port", PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)

    if err != nil {
        panic(err)
    }
}
