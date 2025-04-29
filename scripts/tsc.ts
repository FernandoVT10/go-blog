import path from "path";
import ts from "typescript";

import { ROOT_DIR, TSCONFIG_FILE } from "./config";
import { createErrorSnippet, color } from "./utils.ts";

export function reportDiagnostic(diagnostic: ts.Diagnostic) {
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

export function tscWatcher(diagnostic: ts.Diagnostic) {
    if(diagnostic.code === 6194) {
        console.log(color.blue("Done tsc - Watching..."));
    }
}
