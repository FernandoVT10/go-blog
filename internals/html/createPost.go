package html

import (
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func CreatePost() Node {
    return layout("Create Blog Post",
        navbar(false, ""),
        Div(ID("create-post")),
        Script(Src("/build/js/lib/notify.js")),
        Script(Src("/build/js/lib/snabbdom.js")),
        Script(Src("/build/js/create-post.js")),
    )
}
