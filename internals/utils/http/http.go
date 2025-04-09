package http

import (
    "net/http"
    "encoding/json"
)

func SendJson(w http.ResponseWriter, statusCode int, data any) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    err := json.NewEncoder(w).Encode(data)

    if err != nil {
        http.Error(w, "Unexpected server error", 500)
    }
}
