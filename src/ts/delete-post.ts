import Notify from "./shared/notify";

const BUTTON_ID = "delete-post-btn";
const ESCAPE_CODE = "Escape";

declare const postId: string;

let modal: Modal;

class Modal {
    private htmlEl: HTMLDivElement;

    constructor(container: HTMLElement) {
        this.htmlEl = this.getHTMLEl();
        container.appendChild(this.htmlEl);

        document.addEventListener("keydown", (e: KeyboardEvent) => {
            if(e.code === ESCAPE_CODE) {
                e.preventDefault();

                this.close();
            }
        });
    }

    private cancelBtn(): HTMLButtonElement {
        const btn = document.createElement("button");
        btn.type = "button";
        btn.classList.add("button", "button--secondary");
        btn.textContent = "Cancel";
        btn.addEventListener("click", () => this.close());
        return btn;
    }

    private confirmBtn(): HTMLButtonElement {
        const btn = document.createElement("button");
        btn.type = "button";
        btn.classList.add("button", "button--danger");
        btn.textContent = "Delete Post";
        btn.addEventListener("click", () => this.deletePost());
        return btn;
    }

    private getHTMLEl(): HTMLDivElement {
        const div = document.createElement("div");
        div.classList.add("modal");
        div.role = "dialog";

        const content = document.createElement("div");
        content.classList.add("modal__content");
        div.appendChild(content);

        const h2 = document.createElement("h2");
        h2.textContent = "Do you really want to delete this post?";
        h2.classList.add("modal__title");
        content.appendChild(h2);

        const btnContainer = document.createElement("div");
        btnContainer.classList.add("modal__btn-container");
        btnContainer.appendChild(this.cancelBtn());
        btnContainer.appendChild(this.confirmBtn());
        content.appendChild(btnContainer);

        return div;
    }

    public open() {
        this.htmlEl.classList.add("modal--open");
        document.body.style.overflow = "hidden";
    }

    public close() {
        this.htmlEl.classList.remove("modal--open");
        document.body.style.overflow = "auto";
    }

    private async deletePost() {
        const n = Notify.loading("Deleting post...");
        this.close();

        try {
            const res = await fetch(`/api/posts/${postId}`, {
                method: "DELETE",
            });

            if(res.status === 200) {
                Notify.success("Post deleted successfully!");
                window.setTimeout(() => {
                    window.location.href = "/";
                }, 1000);
            } else {
                const data = await res.json();
                Notify.error(data.error);
            }
        } catch(e) {
            console.error(e);
        }

        Notify.remove(n);
    }
}

window.addEventListener("DOMContentLoaded", () => {
    const btn = document.getElementById(BUTTON_ID);
    if(btn === null) {
        console.error(`There's no button with id "${BUTTON_ID}"`);
        return;
    }

    if(postId === undefined) {
        console.error(`"postId" variable is needed`);
        return;
    }

    modal = new Modal(document.body);

    btn.addEventListener("click", () => modal.open());
});
