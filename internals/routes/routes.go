package routes

import (
    "net/http"
    "fmt"

    "github.com/FernandoVT10/go-blog/internals/db"
    "github.com/FernandoVT10/go-blog/internals/html"

    g "maragu.dev/gomponents"
    ghttp "maragu.dev/gomponents/http"
)

const POSTS_UPLOADS_URL = "http://localhost:3000/uploads/posts"

type BlogPostDoc struct {
    db.BlogPost
    Id string
}

func ConvertCoverToUrl(cover string) string {
    return fmt.Sprintf("%s/%s", POSTS_UPLOADS_URL, cover)
}

// converts all covers names (19238...2.webp) into an url that can be send to the frontend
func ConvertCoversToUrl(blogPosts []db.BlogPost) {
    for i := range len(blogPosts) {
        blogPosts[i].Cover = ConvertCoverToUrl(blogPosts[i].Cover)
    }
}

func GetRoutes() *http.ServeMux {
    router := http.NewServeMux()

    router.Handle("/api/", http.StripPrefix("/api", getApiRoutes()));

    router.HandleFunc(
        "GET /{$}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost

            // TODO: Add a limit of 3 posts only
            db.BlogPostModel.Find(&blogPosts)
            ConvertCoversToUrl(blogPosts)

            return html.Home(blogPosts), nil
        }),
    )

    router.HandleFunc(
        "GET /blog",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost

            db.BlogPostModel.Find(&blogPosts)
            ConvertCoversToUrl(blogPosts)

            return html.Blog(blogPosts), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/posts/{id}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPost db.BlogPost
            // TODO: make this work correctly
            db.BlogPostModel.FindOne(&blogPost)
            return html.BlogPost(blogPost), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/create-post",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            return html.CreatePost(), nil
        }),
    )

    return router
}
