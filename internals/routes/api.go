package routes

import (
    "net/http"

    "github.com/FernandoVT10/go-blog/internals/controllers"

    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

const MAX_TITLE_LENGTH = 100
const MAX_CONTENT_LENGTH = 5000

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
    );

    return router
}
