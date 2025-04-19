import {
    h,
    init,
    classModule,
    propsModule,
    styleModule,
    attributesModule,
    eventListenersModule,
    VNode,
} from "snabbdom";

import Cover from "./cover";
import Editor from "./editor";
import CreatePostController from "./controller";

let vnode: VNode;

const patch = init([
    classModule,
    propsModule,
    styleModule,
    attributesModule,
    eventListenersModule,
]);

function titleInput(): VNode {
    const inputId = "create-post-title-input";
    return h("div.input-group", [
        h(
            "label.input-group__label",
            { props: { htmlFor: inputId }},
            "Title",
        ),
        h(
            `input#${inputId}.input-group__input`,
            {props: {
                type: "text",
                placeholder: "Write an inspiring title",
                name: "title",
                required: true,
            }},
        ),
    ]);
}

function view(ctrl: CreatePostController): VNode {
    return h("article.create-post",[
        h(
            "form",
            {on: {submit: (e) => {
                e.preventDefault();
                ctrl.submitForm(e.target as HTMLFormElement);
            }}},
            [
                Cover(ctrl),
                titleInput(),
                Editor(ctrl),
                h(
                    "button.create-post__button",
                    {props: {
                        type: "submit",
                        disabled: ctrl.creatingPost,
                    }},
                    "Create Post",
                ),
            ],
        ),
    ]);
}

window.addEventListener("DOMContentLoaded", () => {
    const containerId = "create-post";
    const container = document.getElementById(containerId);

    const ctrl = new CreatePostController(rerender);

    function rerender() {
        vnode = patch(vnode, view(ctrl));
    }

    if(container) {
        vnode = patch(container, view(ctrl));
    } else {
        console.error(`There's no html element with id "${containerId}"`);
    }
});
