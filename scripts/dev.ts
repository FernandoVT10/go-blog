import es from "esbuild";
import path from "path";
import assert from "node:assert";
import ts from "typescript";
import * as sass from "sass";
import fs from "fs";
import lv from "livereload";
import { createErrorSnippet, color } from "./utils.ts";

const ROOT_DIR = path.resolve("./");
const PRODUCTION = false;
const OUT_DIR = path.resolve(ROOT_DIR, "public/build");
const JS_OUT_DIR = path.resolve(OUT_DIR, "js/");
const TSCONFIG_FILE = path.resolve(ROOT_DIR, "src/tsconfig.json");
const SASS_FILE = path.resolve(ROOT_DIR, "src/scss/main.scss");
const SASS_DIR = path.resolve(ROOT_DIR, "src/scss");
const OUT_CSS_FILE = path.resolve(OUT_DIR, "main.css");

function esbuildOnEnd(res: es.BuildResult) {
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

    if(res.errors.length === 0) {
        console.log(color.cyan("Done esbuild - Watching..."));
    }
}

async function esbuild() {
    const inputGlob = path.resolve(ROOT_DIR, "src/ts/*.ts");
    const ctx = await es.context({
        bundle: true,
        splitting: true,
        minify: PRODUCTION,
        format: "esm",
        treeShaking: true,
        outdir: JS_OUT_DIR,
        entryPoints: [inputGlob],
        logLevel: "silent",
        plugins: [
            {
                name: "onBundleDone",
                setup(build: es.PluginBuild) {
                    build.onEnd(esbuildOnEnd);
                },
            },
        ],
    });
    await ctx.watch();
}

function reportDiagnostic(diagnostic: ts.Diagnostic) {
    if(!diagnostic.file) {
        console.log(diagnostic.messageText);
        return;
    }

    if(diagnostic.start && diagnostic.length) {
        const { file } = diagnostic;

        const pos = file.getLineAndCharacterOfPosition(diagnostic.start);
        const lineText = file.getText().split("\n")[pos.line];
        // pos.line starts at 0
        const line = pos.line + 1;
        const column = pos.character;

        const filePath = path.relative(ROOT_DIR, file.fileName);
        const message = typeof(diagnostic.messageText) === "string"
            ? diagnostic.messageText : diagnostic.messageText.messageText;

        const error = createErrorSnippet({
            lineText,
            line,
            column,
            length: diagnostic.length,
            filePath,
            message,
        });
        console.log(error);
    }
}

function tscWatcher(diagnostic: ts.Diagnostic) {
    if(diagnostic.code === 6194) {
        console.log(color.blue("Done tsc - Watching..."));
    }
}

async function tsc() {
    const program = ts.createSemanticDiagnosticsBuilderProgram;
    const host = ts.createWatchCompilerHost(
        TSCONFIG_FILE,
        { noEmit: true },
        ts.sys,
        program,
        reportDiagnostic,
        tscWatcher,
    );
    ts.createWatchProgram(host);
}

async function buildScss() {
    try {
        const res = await sass.compileAsync(SASS_FILE, {
            style: PRODUCTION ? "compressed" : "expanded",
        });
        await fs.promises.mkdir(path.dirname(OUT_CSS_FILE), { recursive: true });
        await fs.promises.writeFile(OUT_CSS_FILE, res.css);
        console.log(color.magenta("Done scss - Watching..."));
    } catch(e) {
        if(e instanceof sass.Exception)
            console.log(e.message);
        console.log(e);
    }
}

async function sassBuilder() {
    await buildScss();

    fs.watch(SASS_DIR, { recursive: true }, async (event) => {
        if(event === "change") {
            await buildScss();
        }
    });
}

function livereload() {
    const server = lv.createServer({
        exts: ["css", "js"],
    }, () => {
        console.log(color.green("Livereload server started!"));
    });
    server.watch(OUT_DIR);
}

async function main() {
    await Promise.allSettled([tsc(), esbuild(), sassBuilder(), livereload()]);
}

main();
