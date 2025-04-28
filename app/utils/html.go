package utils

import (
    "fmt"
    "path"

    . "maragu.dev/gomponents"
    . "maragu.dev/gomponents/html"
)

const BUILD_JS = "/build/js/";

// returns a script (module) pointing to a compiled js file
func EsmJs(jsKey string) Node {
    fileName := fmt.Sprintf("%s.js", jsKey)
    return Script(Src(path.Join(BUILD_JS, fileName)), Type("module"))
}
