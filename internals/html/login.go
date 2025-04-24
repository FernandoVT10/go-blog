package html

import (
    "github.com/FernandoVT10/go-blog/internals/utils"
    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

func Login() Node {
    return page(
        "Login",
        []HeadNodes{
            utils.EsmJs("login"),
        },
        navbar(false, ""),
        Article(Class("login"),
            Form(ID("login-form"),
                H1(Class("login__title"), Text("Login")),
                Div(Class("input-group"),
                    Input(
                        Type("password"),
                        Class("input-group__input login__input"),
                        Placeholder("Password..."),
                        Name("password"),
                        Required(),
                    ),
                ),
                Button(Class("button button--normal login__button"), Type("Submit"), Text("Login")),
            ),
        ),
    )
}
