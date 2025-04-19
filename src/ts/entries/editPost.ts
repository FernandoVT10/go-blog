import { VNode } from "snabbdom";
import { patch } from "../shared/snabbdom";

import PostEditor from "../postEditor";
import PostEditorController from "../postEditor/controller";
import Notify from "../shared/notify";

declare const blogPostJSON: string;
const blogPost: BlogPost = JSON.parse(blogPostJSON);

type BlogPost = {
    Id: string;
    Title: string;
    Cover: string;
    Content: string;
}

async function submitForm(form: HTMLFormElement, ctrl: PostEditorController) {
    const formData = new FormData();

    if(ctrl.title !== blogPost.Title) {
        formData.append("title", ctrl.title);
    }

    if(ctrl.content !== blogPost.Content) {
        formData.append("content", ctrl.content);
    }

    const data = new FormData(form);
    const cover = data.get("cover") as (null | File);
    if(cover && cover.size > 0) {
        formData.append("cover", cover);
    }

    ctrl.loading = true;
    ctrl.reRender();

    const n = Notify.loading("Saving changes...");

    try {
        const res = await fetch(`/api/posts/${blogPost.Id}`, {
            method: "PUT",
            body: formData,
        });

        if(res.status === 200) {
            window.location.href = `/blog/posts/${blogPost.Id}`;
        }
    } catch(e) {
        Notify.error("There was an error trying to save the changes");
        console.error(e);
    } finally {
        ctrl.loading = false;
        ctrl.reRender();
        Notify.remove(n);
    }
}

window.addEventListener("DOMContentLoaded", () => {
    if(!blogPost) {
        console.error("blogPost not defined");
        return;
    }

    const containerId = "edit-post";
    const container = document.getElementById(containerId);

    const ctrl = new PostEditorController(rerender, submitForm);
    ctrl.title = blogPost.Title;
    ctrl.content = blogPost.Content;
    ctrl.imageURL = blogPost.Cover;

    let vnode: VNode;

    function rerender() {
        vnode = patch(vnode, PostEditor(ctrl, "Save Changes"));
    }

    if(container) {
        vnode = patch(container, PostEditor(ctrl, "Save Changes"));
    } else {
        console.error(`There's no html element with id "${containerId}"`);
    }
});
