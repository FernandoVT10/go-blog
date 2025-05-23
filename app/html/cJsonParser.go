package html

import (
    "github.com/FernandoVT10/go-blog/app/utils"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func logo() Node {
    return Pre(Class("c-json-parser__logo"), Text(`
 ▄████▄      ▄▄▄██▀▀▀  ██████  ▒█████   ███▄    █        ██▓███   ▄▄▄       ██▀███    ██████ ▓█████  ██▀███
▒██▀ ▀█        ▒██   ▒██    ▒ ▒██▒  ██▒ ██ ▀█   █       ▓██░  ██▒▒████▄    ▓██ ▒ ██▒▒██    ▒ ▓█   ▀ ▓██ ▒ ██▒
▒▓█    ▄       ░██   ░ ▓██▄   ▒██░  ██▒▓██  ▀█ ██▒      ▓██░ ██▓▒▒██  ▀█▄  ▓██ ░▄█ ▒░ ▓██▄   ▒███   ▓██ ░▄█ ▒
▒▓▓▄ ▄██▒   ▓██▄██▓    ▒   ██▒▒██   ██░▓██▒  ▐▌██▒      ▒██▄█▓▒ ▒░██▄▄▄▄██ ▒██▀▀█▄    ▒   ██▒▒▓█  ▄ ▒██▀▀█▄
▒ ▓███▀ ░    ▓███▒   ▒██████▒▒░ ████▓▒░▒██░   ▓██░      ▒██▒ ░  ░ ▓█   ▓██▒░██▓ ▒██▒▒██████▒▒░▒████▒░██▓ ▒██▒
░ ░▒ ▒  ░    ▒▓▒▒░   ▒ ▒▓▒ ▒ ░░ ▒░▒░▒░ ░ ▒░   ▒ ▒       ▒▓▒░ ░  ░ ▒▒   ▓▒█░░ ▒▓ ░▒▓░▒ ▒▓▒ ▒ ░░░ ▒░ ░░ ▒▓ ░▒▓░
  ░  ▒       ▒ ░▒░   ░ ░▒  ░ ░  ░ ▒ ▒░ ░ ░░   ░ ▒░      ░▒ ░       ▒   ▒▒ ░  ░▒ ░ ▒░░ ░▒  ░ ░ ░ ░  ░  ░▒ ░ ▒░
░            ░ ░ ░   ░  ░  ░  ░ ░ ░ ▒     ░   ░ ░       ░░         ░   ▒     ░░   ░ ░  ░  ░     ░     ░░   ░
░ ░          ░   ░         ░      ░ ░           ░                      ░  ░   ░           ░     ░  ░   ░
░
`))
}


func challenges() Node {
    return Section(Class("c-json-parser__section-challenges"),
        H2(Class("c-json-parser__subtitle"), Text("Challenges I had to face")),
        Ul(Class("c-json-parser__challenges"),
            Li(Class("c-json-parser__challenge"),
                P(
                    Text("I had to learn new memory management strategies (like "),
                    A(Href("https://fvtblog.com/blog/posts/68260672cdab662d98a4b9ce"), Text("arenas")),
                    Text("), for the purpose of allocating and de-allocating all memory safely."),
                ),
            ),

            Li(Class("c-json-parser__challenge"),
                P(
                    Text("Learning about "),
                    Strong(Text("lexers")),
                    Text(" and "),
                    Strong(Text("parsers")),
                    Text(" was a daunting journey that I had to tackle."),
                ),
            ),

            Li(Class("c-json-parser__challenge"),
                P(
                    Text(`
                        One of the most challenging things I did in this project, was the way to report syntax errors.
                        As you can see in the features, the error logging is somewhat complex.
                        Full of string formatting, edge cases, and many other problems that araised while working on it.
                    `),
                ),
            ),

            Li(Class("c-json-parser__challenge"),
                P(
                    Text("How can I parse json into "),
                    Strong(Text("C")),
                    Text(`? Which data types should I use? How will I structure objects to be easily traverse?
                        Those are some of the questions I had to answer. They took a while, but the result was so cool!
                    `),
                ),
            ),
        ),
    )
}

const images_folder = "/images/cJsonParser/";

func features() Node {
    return Section(Class("c-json-parser__features-section"),
        H2(Class("c-json-parser__subtitle"), Text("Features")),
        Section(Class("c-json-parser__features"),
            Article(Class("c-json-parser__feature"),
                Img(
                    Class("c-json-parser__feature-img"),
                    Src(images_folder + "Error-Logging.webp"),
                    Alt("Performance Image"),
                ),
                Div(Class("c-json-parser__feature-content"),
                    H3(Class("c-json-parser__feature-title"), Text("Complex syntax error logging")),
                    P(Class("c-json-parser__feature-description"),
                        Text(`
                            When an error happens, you know where and why it happened.
                            The parser prints the line together with a mark showing exactly where the error is.
                        `),
                    ),
                ),
            ),

            Article(Class("c-json-parser__feature"),
                Img(
                    Class("c-json-parser__feature-img"),
                    Src(images_folder + "Performance.webp"),
                    Alt("Performance Image"),
                ),
                Div(Class("c-json-parser__feature-content"),
                    H3(Class("c-json-parser__feature-title"), Text("Simple with great performance")),
                    P(Class("c-json-parser__feature-description"),
                        Text(`
                            It's a super simple project with great performance, taking 0.14s (in average) to parse a 40MB json file.
                        `),
                    ),
                ),
            ),
        ),
    )
}

func learnings() Node {
    return Section(Class("c-json-parser__learnings"))
}

func CJsonParser() Node {
    return page(
        "C Json Parser",
        []HeadNodes {
            utils.EsmJs("cJsonParser"),
        },
        Article(Class("c-json-parser"),
            Header(Class("c-json-parser__header"),
                logo(),
            ),
            Section(Class("c-json-parser__description"),
                P(Class("c-json-parser__description-p"),
                    Text(`
                        A little library that parses json into useful C structs.
                        There are a lot of helper functions missing,
                        that is because this project was meant to be a learning project for me.
                        A project to learn how to work with Lexers and Parsers.
                    `),
                ),
            ),
            features(),
            challenges(),
            learnings(),
            Div(Class("c-json-parser__bg"), Div(ID("c-json-parser-bg"))),
        ),
    )
}
