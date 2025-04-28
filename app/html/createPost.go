package html

import (
    "github.com/FernandoVT10/go-blog/app/utils"

    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func CreatePost(pageData PageData) Node {
    return page(
        "Create Blog Post",
        []HeadNodes {
            utils.EsmJs("createPost"),
        },
        navbar(pageData.IsAuthenticated),
        Div(ID("create-post")),
    )
}
