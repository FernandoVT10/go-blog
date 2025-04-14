package html

import (
    "fmt"

    "github.com/FernandoVT10/go-blog/internals/db"
    "github.com/FernandoVT10/go-blog/internals/utils"

    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func BlogPost(blogPost db.BlogPost) Node {
    contentHtml := utils.MarkdownToHTML(blogPost.Content)
    postIdJs := fmt.Sprintf(`const postId = "%s";`, blogPost.Id.Hex())

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
                    Time(
                        Text(utils.GetTimeAgo(blogPost.CreatedAt)),
                        Title(utils.FormatTime(blogPost.CreatedAt)),
                        DateTime(blogPost.CreatedAt.String()),
                    ),
                ),
            ),
            Section(Class("blog-post__content-container"),
                H1(Class("blog-post__title"), Text(blogPost.Title)),
                Div(Class("blog-post__content"), Raw(contentHtml)),
                Button(
                    ID("delete-post-btn"),
                    Class("button button--danger"),
                    Text("Delete Post"),
                ),
            ),
        ),
        Script(Raw(postIdJs)),
        Script(Src("/build/js/lib/notify.js")),
        Script(Src("/build/js/delete-post.js")),
    )
}
