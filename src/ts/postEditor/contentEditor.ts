import { h, VNode } from "snabbdom";

import PostEditorController from "./controller";

const MAX_CONTENT_LENGTH = 5000;
const CONTENT_TEXTAREA_ID = "content-textarea";

function Options(ctrl: PostEditorController): VNode {
    const showPreview = ctrl.mdPreview.show;
    const activeOptCls = ".post-editor__content-editor-opt--active";

    return h("div.post-editor__content-editor-opts", [
        h(
            `button.post-editor__content-editor-opt${!showPreview ? activeOptCls : ""}`,
            {
                props: { type: "button" },
                on: { click: () => ctrl.setMDPreviewStatus(false) },
            },
            "Editor"
        ),
        h(
            `button.post-editor__content-editor-opt${showPreview ? activeOptCls : ""}`,
            {
                props: { type: "button" },
                on: { click: () => ctrl.setMDPreviewStatus(true) },
            },
            "Preview"
        ),
    ]);
}

function Preview(ctrl: PostEditorController): VNode {
    if(ctrl.mdPreview.loading) {
        return h("div.post-editor__preview", [
            h("div.post-editor__preview-loader", [
                h("span.post-editor__preview-spinner"),
                h("p.post-editor__preview-text", "Rendering markdown..."),
            ]),
        ]);
    }

    return h("div.post-editor__preview.markdown-container", {
        props: { innerHTML: ctrl.mdPreview.rawHtml },
    });
}

function Textarea(ctrl: PostEditorController): VNode {
    return h(
        `textarea#${CONTENT_TEXTAREA_ID}.post-editor__textarea`,
        {
            props: {
                maxLength: MAX_CONTENT_LENGTH,
                required: true,
                name: "content",
                value: ctrl.content,
            },
            on: {
                change: (e) => {
                    e.preventDefault();
                    const textarea = e.target as HTMLTextAreaElement;
                    ctrl.setContent(textarea.value);
                },
            },
        },
    );
}

export default function(ctrl: PostEditorController): VNode {
    return h(
        `div.post-editor__content-editor`,
        [
            Options(ctrl),
            ctrl.mdPreview.show ? Preview(ctrl) : Textarea(ctrl),
        ],
    );
}
