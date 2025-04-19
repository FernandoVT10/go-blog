import Notify from "../shared/notify";

const SUPPORTED_IMAGE_TYPES = [
    "image/jpeg", "image/png", "image/jpg",
];

export default class CreatePostController {
    reRender: () => void;
    imageURL = "";
    creatingPost = false;
    mdPreview = {
        show: false,
        loading: false,
        rawHtml: "",
    };
    content = "";

    constructor(reRender: () => void) {
        this.reRender = reRender;
    }

    private async readImageAsURL(image: File): Promise<string> {
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

    setImagePreview(input: HTMLInputElement) {
        if(!input.files || !input.files[0]) return;
        const file = input.files[0];

        const isImageValid = SUPPORTED_IMAGE_TYPES.find(v => v === file.type) !== undefined;
        if(!isImageValid) {
            Notify.error("You must select a valid image");
            return;
        }

        this.readImageAsURL(file)
            .then(url => {
                this.imageURL = url;
                this.reRender();
            })
            .catch(e => console.error(e));
    }

    async submitForm(form: HTMLFormElement) {
        const formData = new FormData(form);

        const cover = formData.get("cover") as File;
        if(cover.size === 0) {
            Notify.error("Cover is required");
            return;
        }

        this.creatingPost = true;
        this.reRender();

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
            this.creatingPost = false;
            this.reRender();
        }
    }

    async setMDPreviewStatus(status: boolean) {
        this.mdPreview.show = status;
        this.reRender();

        if(this.mdPreview.show
            && !this.mdPreview.loading
            && this.content.length > 0
        ) {
            this.mdPreview.loading = true;
            try {
                const res = await fetch("/api/render-markdown", {
                    method: "POST",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({ markdown: this.content }),
                });
                const json = await res.json();
                this.mdPreview.rawHtml = json.rawHtml;
                this.reRender();
            } catch(e) {
                console.error(e);
            } finally {
                this.mdPreview.loading = false;
                this.reRender();
            }
        }
    }

    setContent(content: string) {
        this.content = content;
        this.reRender();
    }
}
