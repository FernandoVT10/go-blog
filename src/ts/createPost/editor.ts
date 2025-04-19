import { h, VNode } from "snabbdom";

import CreatePostController from "./controller";

const MAX_CONTENT_LENGTH = 5000;
const CONTENT_TEXTAREA_ID = "content-textarea";

function Options(ctrl: CreatePostController): VNode {
    const showPreview = ctrl.mdPreview.show;
    const activeOptCls = ".create-post__editor-opt--active";

    return h("div.create-post__editor-opts", [
        h(
            `button.create-post__editor-opt${!showPreview ? activeOptCls : ""}`,
            {
                props: { type: "button" },
                on: { click: () => ctrl.setMDPreviewStatus(false) },
            },
            "Editor"
        ),
        h(
            `button.create-post__editor-opt${showPreview ? activeOptCls : ""}`,
            {
                props: { type: "button" },
                on: { click: () => ctrl.setMDPreviewStatus(true) },
            },
            "Preview"
        ),
    ]);
}

function Preview(ctrl: CreatePostController): VNode {
    if(ctrl.mdPreview.loading) {
        return h("div.create-post__preview", [
            h("div.create-post__preview-loader", [
                h("span.create-post__preview-spinner"),
                h("p.create-post__preview-text", "Rendering markdown..."),
            ]),
        ]);
    }

    return h("div#create-post-preview.create-post__preview.markdown-container", {
        props: { innerHTML: ctrl.mdPreview.rawHtml },
    });
}

function Textarea(ctrl: CreatePostController): VNode {
    return h(
        `textarea#${CONTENT_TEXTAREA_ID}.create-post__textarea`,
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

export default function(ctrl: CreatePostController): VNode {
    return h(
        `div.create-post__editor`,
        [
            Options(ctrl),
            ctrl.mdPreview.show ? Preview(ctrl) : Textarea(ctrl),
        ],
    );
}
