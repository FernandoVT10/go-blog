package http

import (
    "errors"
    "net/http"
    "encoding/json"
    "github.com/FernandoVT10/go-blog/app/html"
    "github.com/FernandoVT10/go-blog/app/controllers"

    g "maragu.dev/gomponents"
)

func SendJson(w http.ResponseWriter, statusCode int, data any) {
    w.Header().Add("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    err := json.NewEncoder(w).Encode(data)

    if err != nil {
        http.Error(w, "Unexpected server error", 500)
    }
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

func SendNode(w http.ResponseWriter, r *http.Request, node g.Node) {
    err := node.Render(w)
    if err != nil {
        http.Error(w, "error rendering node: "+err.Error(), http.StatusInternalServerError)
    }
}

// returns data that is gonna be used for all pages
func GetPageData(r *http.Request) html.PageData {
    return html.PageData{
        IsAuthenticated: controllers.IsAuthenticated(r),
    }
}

func Send404Page(w http.ResponseWriter, r *http.Request) {
    page := html.NotFound(GetPageData(r))
    SendNode(w, r, page)
}

func Send500Page(w http.ResponseWriter, r *http.Request) {
    page := html.ServerError(GetPageData(r))
    SendNode(w, r, page)
}
