package html

import (
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func ServerError(pageData PageData) Node {
    return page("500", nil,
        navbar(pageData.IsAuthenticated),
        Div(Class("error-message"),
            H1(Text("500")),
            H2(Text("Internal Server Error")),
        ),
    )
}
