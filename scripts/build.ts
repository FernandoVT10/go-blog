import path from "path";
import es from "esbuild";
import ts from "typescript";

import { buildScss } from "./scss.ts";
import { getEsbuildConfig, esbuildOnEnd } from "./esbuild.ts";

import { color, cleanUp } from "./utils.ts";
import { TSCONFIG_FILE } from "./config.ts";
import { reportDiagnostic } from "./tsc.ts";

async function esbuild() {
    const config = getEsbuildConfig(true, []);
    const res = await es.build(config);
    esbuildOnEnd(res);

    if(res.errors.length === 0) {
        console.log(color.cyan("Done esbuild"));
    }
}

function tsc(): boolean {
    const configFile = ts.readConfigFile(TSCONFIG_FILE, ts.sys.readFile);
    const parsedConfig = ts.parseJsonConfigFileContent(
        configFile.config, ts.sys, path.dirname(TSCONFIG_FILE)
    );
    const program = ts.createProgram(parsedConfig.fileNames, { noEmit: true });
    const diagnostics = ts.getPreEmitDiagnostics(program);

    for(const diagnostic of diagnostics) {
        reportDiagnostic(diagnostic);
    }

    return diagnostics.length === 0;
}

async function main() {
    await cleanUp();

    if(await buildScss(true)) {
        console.log(color.magenta("Done scss"));
    }

    if(!tsc()) {
        return;
    }
    console.log(color.blue("Done tsc --noEmit"));

    await esbuild();
}

main();
