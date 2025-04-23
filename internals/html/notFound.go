package html

import (
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func NotFound() Node {
    return page("404", nil,
        navbar(false, ""),
        Div(Class("not-found"),
            H1(Text("404")),
            H2(Text("Page Not Found")),
        ),
    )
}
