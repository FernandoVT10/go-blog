import { h, VNode } from "snabbdom";

import Cover from "./cover";
import PostEditorController from "./controller";
import ContentEditor from "./contentEditor";

const TITLE_INPUT_ID = "post-editor-title-input";

function titleInput(ctrl: PostEditorController): VNode {
    return h("div.input-group", [
        h(
            "label.input-group__label",
            { props: { htmlFor: TITLE_INPUT_ID }},
            "Title",
        ),
        h(
            `input#${TITLE_INPUT_ID}.input-group__input`,
            {
                props: {
                    type: "text",
                    placeholder: "Write an inspiring title",
                    name: "title",
                    required: true,
                    value: ctrl.title,
                },
                on: {
                    change: (e) => {
                        const v = (e.target as HTMLInputElement).value;
                        ctrl.setTitle(v);
                    },
                },
            },
        ),
    ]);
}

export default function(ctrl: PostEditorController, btnText: string): VNode {
    return h("article.post-editor",[
        h(
            "form",
            {on: {submit: (e) => {
                e.preventDefault();
                ctrl.submitForm(e.target as HTMLFormElement);
            }}},
            [
                Cover(ctrl),
                titleInput(ctrl),
                ContentEditor(ctrl),
                h(
                    "button.button.button--normal.post-editor__button",
                    {props: {
                        type: "submit",
                        disabled: ctrl.loading,
                    }},
                    btnText,
                ),
            ],
        ),
    ]);
}
