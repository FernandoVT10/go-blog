package html

import (
    "github.com/FernandoVT10/go-blog/internals/db"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func BlogPost(blogPost db.BlogPost) Node {
    return layout(blogPost.Title,
        navbar(false, ""),
        Article(Class("blog-post"),
            Section(Class("blog-post__cover-container"),
                Img(
                    Src(blogPost.Cover),
                    Alt("Blog Post Cover"),
                    Class("blog-post__cover"),
                ),
                Div(Class("blog-post__date"),
                    SVGIcon("clock", ""),
                    // TODO: get actual date
                    Span(Text("19 days ago")),
                ),
            ),
            Section(Class("blog-post__content-container"),
                H1(Class("blog-post__title"), Text(blogPost.Title)),

                // TODO: blogPost.Content is a markdown text, render it
                Div(Class("blog-post__content"), Text(blogPost.Content)),
            ),
        ),
    )
}
