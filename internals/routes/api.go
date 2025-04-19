package routes

import (
    "net/http"
    "encoding/json"

    "github.com/FernandoVT10/go-blog/internals/controllers"
    "github.com/FernandoVT10/go-blog/internals/utils"

    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

const MAX_TITLE_LENGTH = 100
const MAX_CONTENT_LENGTH = 5000


type StringMap map[string]string

var createPostValidator = []httpUtils.Validator{
    httpUtils.StringValidator{
        Required: true,
        MaxLength: MAX_TITLE_LENGTH,
        Key: "title",
    },
    httpUtils.StringValidator{
        Required: true,
        MaxLength: MAX_CONTENT_LENGTH,
        Key: "content",
    },
    httpUtils.ImageValidator{
        Required: true,
        Key: "cover",
    },
}

var editPostValidator = []httpUtils.Validator{
    httpUtils.StringValidator{
        Required: false,
        MaxLength: MAX_TITLE_LENGTH,
        Key: "title",
    },
    httpUtils.StringValidator{
        Required: false,
        MaxLength: MAX_CONTENT_LENGTH,
        Key: "content",
    },
    httpUtils.ImageValidator{
        Required: false,
        Key: "cover",
    },
}

func getApiRoutes() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc(
        "POST /posts",
        httpUtils.ValidateReq(
            httpUtils.Multipart,
            createPostValidator,
            func(w http.ResponseWriter, r *http.Request) {
                title := r.FormValue("title")
                content := r.FormValue("content")

                cover, _, err := r.FormFile("cover")
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }
                defer cover.Close()

                err, postId := controllers.CreateBlogPost(title, content, cover)
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                httpUtils.SendJson(w, http.StatusOK, map[string]string{"postId": postId})
            },
        ),
    )

    router.HandleFunc(
        "DELETE /posts/{id}",
        func(w http.ResponseWriter, r *http.Request) {
            id := r.PathValue("id")
            if id == "" {
                httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": "Id is required"})
                return
            }

            err := controllers.DeleteBlogPost(id)
            if err != nil {
                httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": err.Error()})
                return
            }

            w.WriteHeader(http.StatusOK);
        },
    )

    router.HandleFunc(
        "PUT /posts/{id}",
        httpUtils.ValidateReq(
            httpUtils.Multipart,
            editPostValidator,
            func(w http.ResponseWriter, r *http.Request) {
                id := r.PathValue("id")
                if id == "" {
                    httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": "Id is required"})
                    return
                }

                title := r.FormValue("title")
                content := r.FormValue("content")

                cover, _, err := r.FormFile("cover")
                if err == nil {
                    defer cover.Close()
                }

                err = controllers.UpdateBlogPost(id, controllers.UpdateBlogPostData{
                    Title: title,
                    Content: content,
                    Cover: cover,
                })
                if err != nil {
                    httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": err.Error()})
                    return
                }

                w.WriteHeader(http.StatusOK);
            },
        ),
    )

    router.HandleFunc(
        "POST /render-markdown",
        func(w http.ResponseWriter, r *http.Request) {
            if r.Header.Get("Content-Type") != "application/json" {
                httpUtils.SendJson(w, http.StatusBadRequest, StringMap{
                    "error": `Content-Type should be "application/json"`,
                })
                return
            }

            var data map[string]string
            err := json.NewDecoder(r.Body).Decode(&data)
            if err != nil {
                httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": "Body couldn't be parsed"})
                return
            }

            md := data["markdown"]
            html := utils.MarkdownToHTML(md)

            httpUtils.SendJson(w, http.StatusOK, StringMap{"rawHtml": html})
        },
    )

    return router
}
