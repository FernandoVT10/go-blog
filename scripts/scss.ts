import fs from "fs";
import sass from "sass";
import path from "path";

import { SASS_FILE, OUT_CSS_FILE } from "./config";

export async function buildScss(production: boolean): Promise<boolean> {
    try {
        const res = await sass.compileAsync(SASS_FILE, {
            style: production ? "compressed" : "expanded",
        });
        await fs.promises.mkdir(path.dirname(OUT_CSS_FILE), { recursive: true });
        await fs.promises.writeFile(OUT_CSS_FILE, res.css);

        return true;
    } catch(e) {
        if(e instanceof sass.Exception)
            console.log(e.message);
        console.log(e);

        return false;
    }
}
