import { h, VNode } from "snabbdom";

import CreatePostController from "./controller";

const INPUT_ID = "cover-selector-input";

function getImageOrLabel(ctrl: CreatePostController): VNode {
    if(ctrl.imageURL) {
        return h("div", [
            h(
                "img.cover-selector__image",
                {
                    props: { alt: "Cover", src: ctrl.imageURL },
                },
            ),
            h(
                "label.cover-selector__button",
                {
                    props: { htmlFor: INPUT_ID },
                },
                [
                    h("svg", h("use", { attrs: { href: "/icons.svg#image" } })),
                    "Change Cover",
                ],
            ),
        ]);
    }

    return h(
        "label.cover-selector__label",
        {
            props: { htmlFor: INPUT_ID },
        },
        "Click to upload an image",
    );
}

export default function(ctrl: CreatePostController): VNode {
    return h("div.cover-selector", [
        h(
            `input#${INPUT_ID}.cover-selector__input`,
            {
                props: { type: "file", accept: ".jpg,.jpeg,.png", name: "cover" },
                on: {
                    change: (e) => ctrl.setImagePreview(e.target as HTMLInputElement)
                },
            },
        ),
        getImageOrLabel(ctrl),
    ]);
}
