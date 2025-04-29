import es from "esbuild";
import fs from "fs";
import lv from "livereload";
import ts from "typescript";
import { reportDiagnostic, tscWatcher } from "./tsc.ts";
import { color, cleanUp } from "./utils.ts";

import { buildScss } from "./scss.ts";
import { getEsbuildConfig, esbuildOnEnd } from "./esbuild.ts";

import * as config from "./config.ts";

async function esbuild() {
    const config = getEsbuildConfig(false, [
        {
            name: "onBundleDone",
            setup(build: es.PluginBuild) {
                build.onEnd(res => {
                    esbuildOnEnd(res);

                    if(res.errors.length === 0) {
                        console.log(color.cyan("Done esbuild - Watching..."));
                    }
                });
            },
        }
    ]);
    const ctx = await es.context(config);
    await ctx.watch();
}

async function sassBuilder() {
    await buildScss(false);
    console.log(color.magenta("Done scss - Watching..."));

    fs.watch(config.SASS_DIR, { recursive: true }, async (event) => {
        if(event === "change") {
            if(await buildScss(false)) {
                console.log(color.magenta("Done scss - Watching..."));
            }
        }
    });
}

function livereload() {
    const server = lv.createServer({
        exts: ["css", "js"],
    }, () => {
        console.log(color.green("Livereload server started!"));
    });
    server.watch(config.BUILD_DIR);
}

async function tsc() {
    const program = ts.createSemanticDiagnosticsBuilderProgram;
    const host = ts.createWatchCompilerHost(
        config.TSCONFIG_FILE,
        { noEmit: true },
        ts.sys,
        program,
        reportDiagnostic,
        tscWatcher,
    );
    ts.createWatchProgram(host);
}

async function main() {
    await cleanUp();
    await Promise.allSettled([tsc(), esbuild(), sassBuilder(), livereload()]);
}

main();
