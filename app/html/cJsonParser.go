package html

import (
    "github.com/FernandoVT10/go-blog/app/utils"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func logo() Node {
    return Pre(Class("c-json-parser-header__logo"), Text(`
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

func description() Node {
    return Section(Class("c-json-parser-description"),
        P(Class("c-json-parser-description__text"),
            Text(`
                A little library that parses json into useful C structs.
                There are a lot of helper functions missing,
                that is because this project was meant to be a learning project for me.
                A project to learn how to work with Lexers and Parsers.
            `),
        ),
    )
}

const images_folder = "/images/cJsonParser/";

func feature(image string, title string, description string) Node {
    return Article(Class("c-json-parser-feature"),
        Div(Class("c-json-parser-feature__img-container"),
            Img(
                Class("c-json-parser-feature__img"),
                Src(images_folder + image),
                Alt("Image"),
            ),
        ),
        Div(Class("c-json-parser-feature__content"),
            H3(Class("c-json-parser-feature__title"), Text(title)),
            P(Class("c-json-parser-feature__description"),
                Text(description),
            ),
        ),
    )
}

func features() Node {
    return Section(Class("c-json-parser-section"),
        H2(Class("c-json-parser-subtitle"), Text("Features")),
        Section(Class("c-json-parser-features"),
            feature(
                "Error-Logging.webp",
                "Complex syntax error logging",

                `When an error happens, you know where and why it happened.
                The parser prints the line together with a mark showing exactly where the error is.`,
            ),
            feature(
                "Performance.webp",
                "Simple with great performance",

                `It's a super simple project with great performance, taking 0.14s (in average) to parse a 40MB json file.`,
            ),
        ),
    )
}

func challenges() Node {
    return Section(Class("c-json-parser-section"),
        H2(Class("c-json-parser-subtitle"), Text("Challenges I had to face")),
        Ul(Class("c-json-parser-challenges"),
            Li(Class("c-json-parser-challenge"),
                P(
                    Text("I had to learn new memory management strategies (like "),
                    A(Href("https://fvtblog.com/blog/posts/68260672cdab662d98a4b9ce"), Text("arenas")),
                    Text("), for the purpose of allocating and de-allocating all memory safely."),
                ),
            ),

            Li(Class("c-json-parser-challenge"),
                P(
                    Text("Learning about "),
                    Strong(Text("lexers")),
                    Text(" and "),
                    Strong(Text("parsers")),
                    Text(" was a daunting journey that I had to tackle."),
                ),
            ),

            Li(Class("c-json-parser-challenge"),
                P(
                    Text(`
                        One of the most challenging things I did in this project, was the way to report syntax errors.
                        As you can see in the features, the error logging is somewhat complex.
                        Full of string formatting, edge cases, and many other problems that araised while working on it.
                    `),
                ),
            ),

            Li(Class("c-json-parser-challenge"),
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

func learnings() Node {
    return Section(Class("c-json-parser-section"),
        H2(Class("c-json-parser-subtitle"), Text("What I've learned")),
        P(Class("c-json-parser-learnings"),
            Text("I have learned to manage the memory more efficiently, to think more about better ways to orginize my code to keep it simple and clean."),
            Br(),
            Text("I also have acquired a better way of learning through using notes."),
            Br(),
            Text("Not trying to achieve perfectionism is a greate learning. A phrase that has helped me with this states that \"the cleanest code is the code that has not been written.\""),
            Br(),
            Text("And finally, a learning that will stick with me forever. Reading good code can help you (in this case jlox helped me a lot) to get better at coding."),
        ),
    )
}

func CJsonParser() Node {
    return page(
        "C Json Parser",
        []HeadNodes {
            utils.EsmJs("cJsonParser"),
        },
        Article(Class("c-json-parser"),
            Header(Class("c-json-parser-header"),
                logo(),
            ),
            description(),
            features(),
            challenges(),
            learnings(),
            Div(Class("c-json-parser-bg"), Div(ID("c-json-parser-bg"))),
        ),
    )
}
