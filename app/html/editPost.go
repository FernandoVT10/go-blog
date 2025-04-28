package html

import (
    "fmt"
    "strconv"
    "github.com/FernandoVT10/go-blog/app/db"
    "github.com/FernandoVT10/go-blog/app/utils"

    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

// blogPostJSON is the blogPost encoded as JSON
func EditPost(blogPost db.BlogPost, blogPostJSON string, pageData PageData) Node {
    title := fmt.Sprintf("Editing - %s", blogPost.Title)
    scriptData := fmt.Sprintf(
        "const blogPostJSON = %s;", strconv.Quote(blogPostJSON),
    )

    return page(
        title,
        []HeadNodes{
            Script(Raw(scriptData)),
            utils.EsmJs("editPost"),
        },
        navbar(pageData.IsAuthenticated),
        Div(ID("edit-post")),
    )
}
