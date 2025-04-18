package main

import (
    "net/http"
    "fmt"
    "github.com/FernandoVT10/go-blog/internals/routes"
    "github.com/FernandoVT10/go-blog/internals/db"
)

const PORT = 3000
const UPLOADS_DIR = "./uploads"

func main() {
    router := http.NewServeMux()

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
