package html

import (
    "fmt"
    "os"

    "github.com/FernandoVT10/go-blog/app/config"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/components"
    . "maragu.dev/gomponents/html"
)

type HeadNodes = Node

type PageData struct {
    IsAuthenticated bool
}

func page(title string, headNodes []HeadNodes, children ...Node) Node {
    dev := os.Getenv("APP_ENV") != "production"

    commonHeadNodes := []Node {
        Link(Rel("stylesheet"), Href("/build/main.css")),
        Meta(Name("viewport"), Content("width=device-width, initial-scale=1.0")),
        Link(Rel("icon"), Type("image/svg+xml"), Href("/favicon.svg")),
        If(!config.GetEnv().Production, Script(Src("http://localhost:35729/livereload.js"))),
    }

    return HTML5(HTML5Props{
        Title: title,
        Head: append(commonHeadNodes, headNodes...),
        Body: children,
    })
}

func SVGIcon(iconName string, class string) Node {
    return SVG(If(class != "", Class(class)),
        Raw(fmt.Sprintf(`<use href="/icons.svg#%s">`, iconName)),
    )
}

func navbarLink(href, text string) Node {
    return Li(Class("navbar__link-item"),
        A(Href(href), Text(text)),
    )
}

func baseNavbar(isHome bool, title string, isAuthenticated bool) Node {
    var class string

    if isHome {
        class = "navbar navbar--home"
    } else {
        class = "navbar"
    }

    return Nav(Class(class),
        Div(Class("navbar__container"),
            A(Class("navbar__link-icon"), Href("/"),
                Img(Src("/icon.svg"), Alt("Icon"), Width("30"), Height("30")),
                If(title != "", H1(Class("navbar__link-title"), Text("Blog"))),
            ),

            Ul(Class("navbar__link-list"),
                navbarLink("/", "Home"),
                navbarLink("/blog", "Blog"),
                If(isAuthenticated, navbarLink("/blog/create-post", "Create Post")),
            ),
        ),
    )
}

// general and most used navbar
func navbar(isAuthenticated bool) Node {
    return baseNavbar(false, "", isAuthenticated)
}
