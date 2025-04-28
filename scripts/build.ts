import sass from "sass";
import path from "path";
import fs from "fs";

import { color, cleanUp } from "./utils.ts";

import * as config from "./config.ts";

async function buildScss() {
    try {
        const res = await sass.compileAsync(config.SASS_FILE, {
            style: "expanded",
        });
        await fs.promises.mkdir(path.dirname(config.OUT_CSS_FILE), { recursive: true });
        await fs.promises.writeFile(config.OUT_CSS_FILE, res.css);
        console.log(color.magenta("Done scss"));
    } catch(e) {
        if(e instanceof sass.Exception)
            console.log(e.message);
        console.log(e);
    }
}

async function main() {
    await cleanUp();
    await buildScss();
}

main();
