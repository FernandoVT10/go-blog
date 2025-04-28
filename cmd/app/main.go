package main

import (
    "net/http"
    "path"
    "fmt"
    "os"

    "github.com/FernandoVT10/go-blog/internals/db"
    "github.com/joho/godotenv"

    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

const PUBLIC_DIR = "./public"
const PORT = 3001
const UPLOADS_DIR = "./uploads"

func main() {
    if err := godotenv.Load(); err != nil {
        panic(err)
    }

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

        httpUtils.Send404Page(w, r)
    })

    db.Connect()

    fmt.Println("[INFO] Server listening on port", PORT)
    err := http.ListenAndServe(fmt.Sprintf(":%d", PORT), mux)

    if err != nil {
        panic(err)
    }
}
