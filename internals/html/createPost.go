package html

import (
    "github.com/FernandoVT10/go-blog/internals/utils"

    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func CreatePost() Node {
    return page(
        "Create Blog Post",
        []HeadNodes {
            utils.EsmJs("createPost"),
        },
        navbar(false, ""),
        Div(ID("create-post")),
    )
}
