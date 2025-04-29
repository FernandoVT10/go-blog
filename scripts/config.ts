import path from "path";

export const ROOT_DIR = path.resolve("./");
export const BUILD_DIR = path.resolve(ROOT_DIR, "public/build");
export const JS_OUT_DIR = path.resolve(BUILD_DIR, "js/");
export const TSCONFIG_FILE = path.resolve(ROOT_DIR, "src/tsconfig.json");
export const SASS_FILE = path.resolve(ROOT_DIR, "src/scss/main.scss");
export const SASS_DIR = path.resolve(ROOT_DIR, "src/scss");
export const OUT_CSS_FILE = path.resolve(BUILD_DIR, "main.css");
