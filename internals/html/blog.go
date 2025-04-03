package html

import (
    "fmt"

    "github.com/FernandoVT10/go-blog/internals/db"
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
                    // TODO: Add real date
                    Span(Text("19 days ago")),
                ),
            ),
        ),
    )
}

func Blog(blogPosts []db.BlogPost) Node {
    posts := make([]Node, 0)

    for _, blogPost := range blogPosts {
        posts = append(posts, blogPostItem(blogPost))
    }

    return layout("Blog",
        navbar(false, "Blog"),
        Section(Class("page-wrapper"),
            Group(posts),
        ),
    )
}
