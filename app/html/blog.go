package html

import (
    "fmt"

    "github.com/FernandoVT10/go-blog/app/db"
    "github.com/FernandoVT10/go-blog/app/utils"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func blogPostItem(blogPost db.BlogPost) Node {
    link := fmt.Sprintf("/blog/posts/%s", blogPost.Id.Hex())

    return A(Href(link),
        Article(Class("blog-post-large-card"),
            Img(
                Class("blog-post-large-card__cover"),
                Src(blogPost.Cover),
                Alt("Blog Post Cover"),
            ),

            Div(Class("blog-post-large-card__details"),
                H3(Class("blog-post-large-card__title"), Text(blogPost.Title)),

                Div(Class("blog-post-large-card__date"),
                    SVGIcon("clock", ""),
                    Span(Text(utils.GetTimeAgo(blogPost.CreatedAt))),
                ),
            ),
        ),
    )
}

func Blog(blogPosts []db.BlogPost, pageData PageData) Node {
    posts := make([]Node, 0)

    for _, blogPost := range blogPosts {
        posts = append(posts, blogPostItem(blogPost))
    }

    return page(
        "Blog",
        nil,
        baseNavbar(false, "Blog", pageData.IsAuthenticated),
        Section(Class("page-wrapper"),
            Group(posts),
        ),
    )
}
