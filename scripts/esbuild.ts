import path from "path";
import es from "esbuild";
import assert from "node:assert";

import { createErrorSnippet } from "./utils.ts";

import { JS_OUT_DIR, ROOT_DIR } from "./config";

export function getEsbuildConfig(production: boolean, plugins: es.Plugin[]): es.BuildOptions {
    const inputGlob = path.resolve(ROOT_DIR, "src/ts/entries/*.ts");

    return {
        bundle: true,
        splitting: true,
        minify: production,
        format: "esm",
        treeShaking: true,
        outdir: JS_OUT_DIR,
        entryPoints: [inputGlob],
        logLevel: "silent",
        plugins,
    }
}

export function esbuildOnEnd(res: es.BuildResult) {
    for(const error of res.errors) {
        if(!error.location) continue;

        const { lineText, line, column, length, file } = error.location;

        const errorMsg = createErrorSnippet({
            lineText,
            line,
            column,
            length,
            message: error.text,
            filePath: file,
        });
        console.log(errorMsg);
    }

    assert(res.warnings.length === 0, "TODO: handle esbuild warnings!");
}
