package http

import (
    "net/http"
    "encoding/json"
)

type errorWithStatusCode struct {
    Code int
    error
}

func (e errorWithStatusCode) StatusCode() int {
    return e.Code
}

func SendJson(w http.ResponseWriter, statusCode int, data any) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    err := json.NewEncoder(w).Encode(data)

    if err != nil {
        http.Error(w, "Unexpected server error", 500)
    }
}

func ErrorWithStatusCode(code int) error {
    return errorWithStatusCode{code, nil}
}
