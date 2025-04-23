package main

import (
    "encoding/json"
    "net/http"
    "github.com/FernandoVT10/go-blog/internals/controllers"
    "github.com/FernandoVT10/go-blog/internals/router"
    "github.com/FernandoVT10/go-blog/internals/utils"
    "github.com/FernandoVT10/go-blog/internals/html"
    "github.com/FernandoVT10/go-blog/internals/db"

    g "maragu.dev/gomponents"
    ghttp "maragu.dev/gomponents/http"
    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

const HOME_BLOG_POSTS_LIMIT = 3

func definePages(router *router.Router) {
    router.Get("/", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
        blogPosts := controllers.GetBlogPosts(controllers.GetBlogPostsOpts{
            Limit: HOME_BLOG_POSTS_LIMIT,
            Sort: map[string]int{
                "createdAt": db.DescendingSort,
            },
        })

        return html.Home(blogPosts), nil
    }))

    router.Get("/blog", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
        blogPosts := controllers.GetBlogPosts(controllers.GetBlogPostsOpts{
            Sort: map[string]int{
                "createdAt": db.DescendingSort,
            },
        })
        return html.Blog(blogPosts), nil
    }))

    router.Get("/blog/posts/{id}", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
        blogPost, err := controllers.GetBlogPostByHexId(r.PathValue("id"))

        if err != nil {
            return html.NotFound(), nil
        }

        return html.BlogPost(blogPost), nil
    }))

    router.Get("/blog/create-post", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
        return html.CreatePost(), nil
    }))

    router.Get("/blog/posts/{id}/edit", ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
        blogPost, err := controllers.GetBlogPostByHexId(r.PathValue("id"))
        if err != nil {
            return html.NotFound(), nil
        }

        blogPostJSON, err := json.Marshal(blogPost)
        if err != nil {
            // TODO: return a 500 page
            return nil, httpUtils.ErrorWithStatusCode(http.StatusInternalServerError)
        }

        return html.EditPost(blogPost, string(blogPostJSON)), nil
    }))
}

type StringMap map[string]string

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

func defineApi(router *router.Router) {
    router.Post(
        "/api/posts",
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

    router.Delete("/api/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
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
    })

    router.Put(
        "/api/posts/{id}",
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

    router.Post("/api/render-markdown", func(w http.ResponseWriter, r *http.Request) {
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
    })
}

func GetRoutes() router.Router {
    router := router.NewRouter()

    definePages(&router)
    defineApi(&router)

    return router
}
