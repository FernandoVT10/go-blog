package utils

import (
    "io"

    "github.com/gomarkdown/markdown"
    "github.com/gomarkdown/markdown/ast"
    "github.com/alecthomas/chroma/v2"
    "github.com/alecthomas/chroma/v2/lexers"
    "github.com/alecthomas/chroma/v2/formatters/html"

    mdHtml "github.com/gomarkdown/markdown/html"
)

var customFormatter *html.Formatter
var customStyle *chroma.Style

func getCustomStyle() (*chroma.Style, error) {
    if customStyle != nil {
        return customStyle, nil
    }

    white := "#cdd6f4"
    peach := "#fab285"
    green := "#a6e3a1"
    // sapphire := "#74c6ec"
    red := "#f38ca9"
    mauve := "#caa6f7"
    blue := "#89b5fa"
    // lavander := "#b4befe"
    gray := "#a6adc9"

    entries := chroma.StyleEntries{
        chroma.Background: white,
        chroma.Error: white,
        chroma.NameAttribute: green,
        chroma.NameBuiltin: red,
        chroma.NameBuiltinPseudo: red,

        chroma.NameFunction: blue,

        chroma.NameTag: peach,

        chroma.String: green,
        chroma.StringDouble: green,
        chroma.LiteralStringAffix: "#f00",

        chroma.Punctuation: red,
        chroma.KeywordType: mauve,
        chroma.Keyword: mauve,

        chroma.Comment: gray,
        chroma.Operator: blue,

        chroma.LiteralNumber: peach,
    }

    style, err := chroma.NewStyle("custom", entries)
    if err != nil {
        return nil, err
    }

    customStyle = style
    return style, err
}

func getFormatter() *html.Formatter {
    if customFormatter != nil {
        return customFormatter
    }

    return html.New()
}

func mdRenderHook(w io.Writer, node ast.Node, entering bool) (ast.WalkStatus, bool) {
    if code, ok := node.(*ast.CodeBlock); ok {
        lang := string(code.Info)

        if lang == "" {
            return ast.GoToNext, false
        }

        lexer := lexers.Get(lang)
        if lexer == nil {
            return ast.GoToNext, false
        }

        iterator, err := lexer.Tokenise(nil, string(code.Literal))
        if err != nil {
            return ast.GoToNext, false
        }

        style, err := getCustomStyle()
        if err != nil {
            return ast.GoToNext, false
        }

        formatter := getFormatter()
        err = formatter.Format(w, style, iterator)
        if err != nil {
            return ast.GoToNext, false
        }

        return ast.GoToNext, true
    }

    return ast.GoToNext, false
}

func MarkdownToHTML(md string) string {
    opts := mdHtml.RendererOptions{
        Flags: mdHtml.CommonFlags,
        RenderNodeHook: mdRenderHook,
    }
    renderer := mdHtml.NewRenderer(opts)
    html := markdown.Render(
        markdown.Parse([]byte(md), nil),
        renderer,
    )
    return string(html)
}
