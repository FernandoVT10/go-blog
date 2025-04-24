package http

import (
    "errors"
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

type ParsedJson map[string]string

func ParseJson(r *http.Request) (ParsedJson, error) {
    if r.Header.Get("Content-Type") != "application/json" {
        return nil, errors.New(`Content-Type should be "application/json"`)
    }

    var data ParsedJson
    err := json.NewDecoder(r.Body).Decode(&data)
    if err != nil {
        return nil, errors.New("Couldn't parse body")
    }

    return data, nil
}
