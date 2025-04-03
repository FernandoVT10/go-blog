package routes

import (
    "net/http"

    "github.com/FernandoVT10/go-blog/internals/db"
    "github.com/FernandoVT10/go-blog/internals/html"

    g "maragu.dev/gomponents"
    ghttp "maragu.dev/gomponents/http"
)

type BlogPostDoc struct {
    db.BlogPost
    Id string
}

func GetRoutes() *http.ServeMux {
    router := http.NewServeMux()

    router.HandleFunc(
        "GET /{$}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost
            // TODO: Add a limit of 3 posts only
            db.BlogPostModel.Find(&blogPosts)

            return html.Home(blogPosts), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/{$}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost
            db.BlogPostModel.Find(&blogPosts)

            return html.Blog(blogPosts), nil
        }),
    )

    return router
}
