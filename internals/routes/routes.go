package routes

import (
    "net/http"
    "encoding/json"
    "fmt"

    "github.com/FernandoVT10/go-blog/internals/db"
    "github.com/FernandoVT10/go-blog/internals/html"
    "go.mongodb.org/mongo-driver/v2/mongo/options"
    "go.mongodb.org/mongo-driver/v2/bson"

    g "maragu.dev/gomponents"
    ghttp "maragu.dev/gomponents/http"

    httpUtils "github.com/FernandoVT10/go-blog/internals/utils/http"
)

type ErrorWithStatusCode struct {
    Code int
    error
}

func (e ErrorWithStatusCode) StatusCode() int {
    return e.Code;
}

const PUBLIC_DIR = "./public"
const POSTS_UPLOADS_URL = "http://localhost:3000/uploads/posts"
const HOME_BLOG_POSTS_LIMIT = 3

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

    publicFs := http.FileServer(http.Dir(PUBLIC_DIR))
    router.Handle("/", publicFs)

    router.Handle("/api/", http.StripPrefix("/api", getApiRoutes()));

    router.HandleFunc(
        "GET /{$}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost

            opts := options.Find().SetLimit(HOME_BLOG_POSTS_LIMIT)
            db.BlogPostModel.Find(&blogPosts, bson.D{}, opts)
            ConvertCoversToUrl(blogPosts)

            return html.Home(blogPosts), nil
        }),
    )

    router.HandleFunc(
        "GET /blog",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPosts []db.BlogPost

            opts := options.Find().SetSort(bson.D{{Key: "createdAt", Value: db.DescendingSort}})
            // TODO: handle this error
            db.BlogPostModel.Find(&blogPosts, bson.D{}, opts)
            ConvertCoversToUrl(blogPosts)

            return html.Blog(blogPosts), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/posts/{id}",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPost db.BlogPost

            id, err := bson.ObjectIDFromHex(r.PathValue("id"))
            if err != nil {
                // TODO: return a 404 page
                return nil, httpUtils.ErrorWithStatusCode(http.StatusBadRequest)
            }

            db.BlogPostModel.FindById(&blogPost, id)
            blogPost.Cover = ConvertCoverToUrl(blogPost.Cover)

            if (blogPost == db.BlogPost{}) {
                // TODO: return a 404 page
                return nil, httpUtils.ErrorWithStatusCode(http.StatusNotFound)
            }

            return html.BlogPost(blogPost), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/create-post",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            return html.CreatePost(), nil
        }),
    )

    router.HandleFunc(
        "GET /blog/posts/{id}/edit",
        ghttp.Adapt(func(w http.ResponseWriter, r *http.Request) (g.Node, error) {
            var blogPost db.BlogPost

            id, err := bson.ObjectIDFromHex(r.PathValue("id"))
            if err != nil {
                // TODO: return a 404 page
                return nil, httpUtils.ErrorWithStatusCode(http.StatusBadRequest)
            }

            db.BlogPostModel.FindById(&blogPost, id)

            if (blogPost == db.BlogPost{}) {
                // TODO: return a 404 page
                return nil, httpUtils.ErrorWithStatusCode(http.StatusNotFound)
            }

            blogPost.Cover = ConvertCoverToUrl(blogPost.Cover)

            blogPostJSON, err := json.Marshal(blogPost)
            if err != nil {
                // TODO: return a 500 page
                return nil, httpUtils.ErrorWithStatusCode(http.StatusInternalServerError)
            }

            return html.EditPost(blogPost, string(blogPostJSON)), nil
        }),
    )

    return router
}
