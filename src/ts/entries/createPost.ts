import { VNode } from "snabbdom";
import { patch } from "../shared/snabbdom";

import PostEditor from "../postEditor";
import PostEditorController from "../postEditor/controller";
import Notify from "../shared/notify";

async function submitForm(form: HTMLFormElement, ctrl: PostEditorController) {
    const formData = new FormData(form);

    const cover = formData.get("cover") as File;
    if(cover.size === 0) {
        Notify.error("Cover is required");
        return;
    }

    ctrl.loading = true;
    ctrl.reRender();

    const n = Notify.loading("Creating post...");

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
        Notify.error("There was an error trying to create the post");
        console.error(e);
    } finally {
        ctrl.loading = false;
        ctrl.reRender();
        Notify.remove(n);
    }
}

window.addEventListener("DOMContentLoaded", () => {
    const containerId = "create-post";
    const container = document.getElementById(containerId);

    const ctrl = new PostEditorController(rerender, submitForm);
    let vnode: VNode;

    function rerender() {
        vnode = patch(vnode, PostEditor(ctrl, "Create Post"));
    }

    if(container) {
        vnode = patch(container, PostEditor(ctrl, "Create Post"));
    } else {
        console.error(`There's no html element with id "${containerId}"`);
    }
});
