import Notify from "../shared/notify";

const FORM_ID = "login-form";

async function login(password: string) {
    const n = Notify.loading("Authenticating...");

    try {
        const res = await fetch("/api/login", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ password }),
        });

        if(res.status === 200) {
            window.location.href = "/";
        } else {
            Notify.error("Incorrect password");
        }
    } catch(e) {
        console.error(e);
        Notify.error("There was an error trying to login");
    } finally {
        Notify.remove(n);
    }
}

window.addEventListener("DOMContentLoaded", () => {
    const form = document.getElementById(FORM_ID) as HTMLFormElement;

    if(!form) {
        console.error(`There is not a form with id "${FORM_ID}"`);
        return;
    }

    form.addEventListener("submit", e => {
        e.preventDefault();
        const formData = new FormData(form);
        const password = formData.get("password");

        if(!password) return;

        login(password?.toString());
    });
});
