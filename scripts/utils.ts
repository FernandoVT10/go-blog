import fs from "fs";

import { BUILD_DIR } from "./config";

function getColoredText(codeColor: number, text: string): string {
    return `\x1b[${codeColor}m${text}\x1b[0m`;
}

export const color = {
    red(text: string): string {
        return getColoredText(31, text);
    },
    cyan(text: string): string {
        return getColoredText(36, text);
    },
    blue(text: string): string {
        return getColoredText(34, text);
    },
    magenta(text: string): string {
        return getColoredText(35, text);
    },
    green(text: string): string {
        return getColoredText(32, text);
    },
};

export function spaces(n: number): string {
    return "".padStart(n, " ");
}

type SnippetError = {
    lineText: string;
    line: number;
    column: number;
    length: number;
    message: string;
    filePath: string;
};

export function createErrorSnippet(err: SnippetError): string {
    let error = "";

    error += color.red(`${err.filePath}:${err.line}:${err.column}`) + "\n";
    error +=`${color.red("Error:")} ${err.message}\n`;

    const leftCode = err.lineText.slice(0, err.column);
    const highligtedCode = color.red(err.lineText.slice(err.column, err.column + err.length));
    const rightCode = err.lineText.slice(err.column + err.length);
    const codeLine = `${spaces(4)}${err.line}`;
    error += `${codeLine} |${leftCode}${highligtedCode}${rightCode}\n`;

    const errorMark = color.red("^".padEnd(err.length, "~"));
    error += `${spaces(4 + err.line.toString().length)} |${spaces(err.column)}${errorMark}`;

    return error;
}

export async function cleanUp() {
    console.log("Cleaning up build folder...")
    await fs.promises.rm(BUILD_DIR, { recursive: true, force: true });
}
