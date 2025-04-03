package html

import (
	"fmt"
    "os"
	. "maragu.dev/gomponents"
	. "maragu.dev/gomponents/html"
    . "maragu.dev/gomponents/components"
)

func layout(title string, children ...Node) Node {
    dev := os.Getenv("APP_ENV") != "production"

    return HTML5(HTML5Props {
        Title: title,
        Head: []Node{
            Link(Rel("stylesheet"), Href("/build/main.css")),
            Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
            Link(Rel("icon"), Type("image/svg+xml"), Href("/static/favicon.svg")),
            If(dev, Script(Src("http://localhost:35729/livereload.js"))),
        },
        Body: children,
    })
}

func SVGIcon(iconName string, class string) Node {
    return SVG(If(class != "", Class(class)),
        Raw(fmt.Sprintf(`<use href="/static/icons.svg#%s">`, iconName)),
    )
}


func navbarLink(href, text string) Node {
    return Li(Class("navbar__link-item"),
        A(Href(href), Text(text)),
    )
}

func navbar(isHome bool, title string) Node {
    var class string

    if isHome {
        class = "navbar navbar--home"
    } else {
        class = "navbar"
    }

    return Nav(Class(class),
        Div(Class("navbar__container"),
            A(Class("navbar__link-icon"), Href("/"),
                Img(Src("/static/icon.svg"), Alt("Icon"), Width("30"), Height("30")),
                If(!isHome, H1(Class("navbar__link-title"), Text("Blog"))),
            ),

            Ul(Class("navbar__link-list"),
                navbarLink("/", "Home"),
                navbarLink("/blog", "Blog"),
                navbarLink("/projects", "Projects"),
            ),
        ),
    )
}
