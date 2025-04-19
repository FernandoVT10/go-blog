import Notify from "../shared/notify";

const SUPPORTED_IMAGE_TYPES = [
    "image/jpeg", "image/png", "image/jpg",
];

type OnSubmitForm = (form: HTMLFormElement, ctrl: PostEditorController) => void;

export default class PostEditorController {
    reRender: () => void;
    imageURL = "";
    loading = false;
    mdPreview = {
        show: false,
        loading: false,
        rawHtml: "",
    };
    content = "";
    title = "";
    onSubmitForm: OnSubmitForm;

    constructor(reRender: () => void, onSubmitForm: OnSubmitForm) {
        this.reRender = reRender;
        this.onSubmitForm = onSubmitForm;
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

    submitForm(form: HTMLFormElement) {
        this.onSubmitForm(form, this);
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

    setTitle(title: string) {
        this.title = title;
        this.reRender();
    }
}
