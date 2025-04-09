import type { VNode } from "snabbdom";

const snabbdom = window.snabbdom;
const { h } = window.snabbdom;
const inputId = "cover-selector-input";

const MAX_CONTENT_LENGTH = 5000;

let vnode: VNode;

const patch = snabbdom.init([
    snabbdom.classModule,
    snabbdom.propsModule,
    snabbdom.styleModule,
    snabbdom.attributesModule,
    snabbdom.eventListenersModule,
]);

const data = {
    imageURL: "",
    title: "",
    creatingPost: false,
};

function readImageAsURL(image: File): Promise<string> {
  const reader = new FileReader();

  return new Promise((resolve, reject) => {
    reader.onload = () => {
      resolve(reader.result as string);
    };

    reader.onerror = () => {
      reject();
    };

    reader.readAsDataURL(image);
  });
}

const SUPPORTED_IMAGE_TYPES = [
    "image/jpeg", "image/png", "image/jpg",
];

async function handleChangeInput(e: Event) {
    const input = e.target as HTMLInputElement;

    if(!input.files || !input.files[0]) return;
    const file = input.files[0];

    const isImageValid = SUPPORTED_IMAGE_TYPES.find(v => v === file.type) !== undefined;
    if(!isImageValid) {
        window.Notify.error("You must select a valid image");
        return;
    }

    try {
        data.imageURL = await readImageAsURL(file);
        render();
    } catch (e) {
        console.error(e);
    }
}

function getImageOrLabel(): VNode {
    if(data.imageURL) {
        return h("div", [
            h(
                "img.cover-selector__image",
                {
                    props: { alt: "Cover", src: data.imageURL },
                },
            ),
            h(
                "label.cover-selector__button",
                {
                    props: { htmlFor: inputId },
                },
                [
                    h("svg", h("use", { attrs: { href: "/static/icons.svg#image" } })),
                    "Change Cover",
                ],
            ),
        ]);
    }

    return h(
        "label.cover-selector__label",
        {
            props: { htmlFor: inputId },
        },
        "Click to upload an image",
    );
}

function coverSelector(): VNode {
    return h("div.cover-selector", [
        h(
            `input#${inputId}.cover-selector__input`,
            {
                props: { type: "file", accept: ".jpg,.jpeg,.png", name: "cover" },
                on: { change: handleChangeInput }
            },
        ),
        getImageOrLabel(),
    ]);
}

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

function contentTextarea(): VNode {
    return h(
        "textarea#content-textarea.create-post__textarea",
        {props: {
            maxLength: MAX_CONTENT_LENGTH,
            required: true,
            name: "content",
        }}
    );
}

async function handleForm(e: Event) {
    e.preventDefault();
    const form = e.target as HTMLFormElement;
    const formData = new FormData(form);

    const cover = formData.get("cover") as File;
    if(cover.size === 0) {
        window.Notify.error("Cover is required");
        return;
    }

    data.creatingPost = true;
    render();

    try {
        const res = await fetch("/api/posts", {
            method: "POST",
            body: formData,
        });

        if(res.status === 200) {
            const json = await res.json();
            window.location.href = `/blog/posts/${json.postId}`;
        }
    } catch(e) {
        window.Notify.error("There was an error trying to create the post");
        console.error(e);
    }

    data.creatingPost = false;
    render();
}

function view(): VNode {
    return h("article.create-post",[
        h(
            "form",
            {on: {submit: handleForm}},
            [
                coverSelector(),
                titleInput(),
                contentTextarea(),
                h(
                    "button.create-post__button",
                    {props: {
                        type: "submit",
                        disabled: data.creatingPost,
                    }},
                    "Create Post",
                ),
            ],
        ),
    ]);
}

function render() {
    if(vnode) {
        vnode = patch(vnode, view());
    }
}

window.addEventListener("DOMContentLoaded", () => {
    const containerId = "create-post";
    const container = document.getElementById(containerId);

    if(container) {
        vnode = patch(container, view());
    } else {
        console.error(`There's no html element with id "${containerId}"`);
    }
});
