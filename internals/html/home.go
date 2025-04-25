package html

import (
    "fmt"

    "github.com/FernandoVT10/go-blog/internals/db"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func homeSocialMediaLink(link, iconName string) Node {
    return Li(Class("home-header__media-item"), 
        A(Target("_blank"), Href(link), Class("home-header__media-link"),
            SVGIcon(iconName, ""),
        ),
    )
}

func homeHeader() Node {
    description := "Hi, I'm a Javascript full stack developer. I like video games, anime, music, and trying to find the best way to write the cleanest and best code."

    return Header(Class("home-header"),
        H1(Class("home-header__full-name"), Text("Fernando Vaca Tamayo")),
        Hr(Class("home-header__separator")),

        P(Class("home-header__description"), Text(description)),

        Ul(Class("home-header__media-list"),
            homeSocialMediaLink("https://twitter.com/FernandoVT10", "twitter"),
            homeSocialMediaLink("https://github.com/FernandoVT10", "github"),
            homeSocialMediaLink("https://github.com/FernandoVT10", "linkedin"),
        ),
    )
}

func homeLanguageList() []Node {
    languages := [...]string{
        "React.js", "Node.js", "Typescript", "HTML/CSS",
        "SQL/NoSQL", "Linux", "Python", "C", "Go",
    }
    nodes := make([]Node, 0)

    for i, language := range languages {
        if i > 0 {
            nodes = append(nodes, Text(", "))
        }
        nodes = append(nodes, Span(Text(language)))
    }

    return nodes
}

func homePresentationCard(iconName, title string, children ...Node) Node {
    return Article(Class("presentation-card"),
        SVGIcon(iconName, "presentation-card__icon"),
        H2(Class("presentation-card__title"), Text(title)),
        P(Class("presentation-card__description"),
            Group(children),
        ),
    )
}

func homePresentationSection() Node {
    return Section(Class("presentation-cards-section"),
        homePresentationCard("code", "Technologies",
            Text("I can use "), Group(homeLanguageList()), Text(" and more."),
        ),
        homePresentationCard("world", "Languages",
            Text("Spanish is my native language, and I have a B1 level of English."),
        ),
        homePresentationCard("stack", "Projects",
            Text("You can click "),
            A(Href("/projects"), Class("presentation-card__link"), Text("here")),
            Text(" to see my projects."),
        ),
    )
}

func homeBlogPost(blogPost db.BlogPost) Node {
    link := fmt.Sprintf("/blog/posts/%s", blogPost.Id.Hex())

    return Article(Class("blog-post-card"),
        A(Href(link),
            Div(Class("blog-post-card__cover"),
                Img(
                    Class("blog-post-card__cover-image"),
                    Src(blogPost.Cover),
                    Alt("Blog Post Cover"),
                ),
            ),
            H3(Class("blog-post-card__title"), Text(blogPost.Title)),
        ),
    )
}

func homeBlogPosts(blogPosts []db.BlogPost) Node {
    posts := make([]Node, 0)

    for _, blogPost := range blogPosts {
        posts = append(posts, homeBlogPost(blogPost))
    }

    return Section(Class("blog-posts-cards"),
        Group(posts),
    )
}

func Home(blogPosts []db.BlogPost, pageData PageData) Node {
    return page(
        "Fernando Vaca Tamayo",
        nil,
        baseNavbar(true, "", pageData.IsAuthenticated),
        homeHeader(),
        Section(Class("page-wrapper"),
            homePresentationSection(),
            H2(Class("section-title"), Text("Posts from my blog")),
            homeBlogPosts(blogPosts),
        ),
    )
}
