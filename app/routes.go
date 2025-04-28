package main

import (
    "encoding/json"
    "net/http"
    "github.com/FernandoVT10/go-blog/app/controllers"
    "github.com/FernandoVT10/go-blog/app/middlewares"
    "github.com/FernandoVT10/go-blog/app/config"
    "github.com/FernandoVT10/go-blog/app/router"
    "github.com/FernandoVT10/go-blog/app/utils"
    "github.com/FernandoVT10/go-blog/app/html"
    "github.com/FernandoVT10/go-blog/app/db"

    httpUtils "github.com/FernandoVT10/go-blog/app/utils/http"
)

const HOME_BLOG_POSTS_LIMIT = 3

func definePages(router *router.Router) {
    router.Get("/", func(w http.ResponseWriter, r *http.Request) {
        blogPosts := controllers.GetBlogPosts(controllers.GetBlogPostsOpts{
            Limit: HOME_BLOG_POSTS_LIMIT,
            Sort: map[string]int{
                "createdAt": db.DescendingSort,
            },
        })

        page := html.Home(blogPosts, httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    })

    router.Get("/blog", func(w http.ResponseWriter, r *http.Request) {
        blogPosts := controllers.GetBlogPosts(controllers.GetBlogPostsOpts{
            Sort: map[string]int{
                "createdAt": db.DescendingSort,
            },
        })

        page := html.Blog(blogPosts, httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    })

    router.Get("/blog/posts/{id}", func(w http.ResponseWriter, r *http.Request) {
        blogPost, err := controllers.GetBlogPostByHexId(r.PathValue("id"))

        if err != nil {
            httpUtils.Send404Page(w, r)
            return
        }

        page := html.BlogPost(blogPost, httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    })

    router.Get("/blog/create-post", func(w http.ResponseWriter, r *http.Request) {
        page := html.CreatePost(httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    }).Use(middlewares.AuthPage())

    router.Get("/blog/posts/{id}/edit", func(w http.ResponseWriter, r *http.Request) {
        blogPost, err := controllers.GetBlogPostByHexId(r.PathValue("id"))
        if err != nil {
            httpUtils.Send404Page(w, r)
            return
        }

        blogPostJSON, err := json.Marshal(blogPost)
        if err != nil {
            httpUtils.Send500Page(w, r)
            return
        }

        page := html.EditPost(blogPost, string(blogPostJSON), httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    }).Use(middlewares.AuthPage())

    router.Get("/login", func(w http.ResponseWriter, r *http.Request) {
        page := html.Login(httpUtils.GetPageData(r))
        httpUtils.SendNode(w, r, page)
    })
}

type StringMap map[string]string

const MAX_TITLE_LENGTH = 100
const MAX_CONTENT_LENGTH = 5000
const COOKIE_MAX_AGE = 60 * 60 * 24 * 30 // 30 days in seconds

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
    ).Use(middlewares.AuthApi())

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
    }).Use(middlewares.AuthApi())

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
    ).Use(middlewares.AuthApi())

    router.Post("/api/render-markdown", func(w http.ResponseWriter, r *http.Request) {
        data, err := httpUtils.ParseJson(r)
        if err != nil {
            httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": err.Error()})
            return
        }

        md := data["markdown"]
        html := utils.MarkdownToHTML(md)

        httpUtils.SendJson(w, http.StatusOK, StringMap{"rawHtml": html})
    })

    router.Post("/api/login", func(w http.ResponseWriter, r *http.Request) {
        data, err := httpUtils.ParseJson(r)
        if err != nil {
            httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": err.Error()})
            return
        }

        password := data["password"]

        signedStr, err := controllers.Login(password)
        if err != nil {
            httpUtils.SendJson(w, http.StatusBadRequest, StringMap{"error": err.Error()})
            return
        }

        cookie := http.Cookie{
            Name: "token",
            Value: signedStr,
            Path: "/",
            MaxAge: COOKIE_MAX_AGE,
            HttpOnly: true,
            Secure: config.GetEnv().Production,
            SameSite: http.SameSiteStrictMode,
        }

        http.SetCookie(w, &cookie)

        w.WriteHeader(http.StatusOK)
    })
}

func GetRoutes() router.Router {
    router := router.NewRouter()

    definePages(&router)
    defineApi(&router)

    return router
}
