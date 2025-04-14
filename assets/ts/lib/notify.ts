declare global {
    interface Window {
        Notify: {
            success: (message: string) => number;
            loading: (message: string) => number;
            error: (message: string) => number;
            remove: (id: number) => void;
        },
    }
}

enum NotificationType {
    Success,
    Error,
    Loading,
}

const SUCCESS_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="currentColor"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M17 3.34a10 10 0 1 1 -14.995 8.984l-.005 -.324l.005 -.324a10 10 0 0 1 14.995 -8.336zm-1.293 5.953a1 1 0 0 0 -1.32 -.083l-.094 .083l-3.293 3.292l-1.293 -1.292l-.094 -.083a1 1 0 0 0 -1.403 1.403l.083 .094l2 2l.094 .083a1 1 0 0 0 1.226 0l.094 -.083l4 -4l.083 -.094a1 1 0 0 0 -.083 -1.32z" /></svg>`;

const CLOSE_ICON = `<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 18" fill="none" stroke="currentColor" stroke-width="4px" stroke-linecap="round" stroke-linejoin="round"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M18 6l-12 12" /><path d="M6 6l12 12" /></svg>`;

const ERROR_ICON = `<svg xmlns="http://www.w3.org/2000/svg"width="20"height="20"viewBox="0 0 24 24"fill="currentColor"><path stroke="none" d="M0 0h24v24H0z" fill="none"/><path d="M12 2c5.523 0 10 4.477 10 10s-4.477 10 -10 10s-10 -4.477 -10 -10s4.477 -10 10 -10m3.6 5.2a1 1 0 0 0 -1.4 .2l-2.2 2.933l-2.2 -2.933a1 1 0 1 0 -1.6 1.2l2.55 3.4l-2.55 3.4a1 1 0 1 0 1.6 1.2l2.2 -2.933l2.2 2.933a1 1 0 0 0 1.6 -1.2l-2.55 -3.4l2.55 -3.4a1 1 0 0 0 -.2 -1.4" /></svg>`;

let notifications: Notification[] = [];
let notificationId = 0;

function getIcon(type: NotificationType): HTMLDivElement {
    const div = document.createElement("div");
    div.classList.add("notification__icon");

    switch(type) {
        case NotificationType.Loading:
            div.innerHTML = `<span class="notification__spinner"></span>`;
            break;
        case NotificationType.Error:
            div.innerHTML = ERROR_ICON;
            break;
        case NotificationType.Success:
        default:
            div.innerHTML = SUCCESS_ICON;
    }

    return div;
}

function getButton(closeFn: () => void): HTMLButtonElement {
    const btn = document.createElement("button");
    btn.classList.add("notification__close-btn");
    btn.innerHTML = CLOSE_ICON;
    btn.type = "button";

    btn.addEventListener("click", closeFn);
    return btn;
}

function getNotification(type: NotificationType, message: string, closeFn: () => void): HTMLDivElement {
    const div = document.createElement("div");

    let modifier: string;
    switch(type) {
        case NotificationType.Loading:
            modifier = "notification--loading";
            break;
        case NotificationType.Error:
            modifier = "notification--error";
            break;
        case NotificationType.Success:
        default:
            modifier = "notification--success";
    }

    div.classList.add("notification", modifier, "fadein");

    const messageEl = document.createElement("p");
    messageEl.classList.add("notification__message");
    messageEl.textContent = message;

    div.append(getIcon(type), messageEl, getButton(closeFn));

    return div;
}

const NOTIFICATION_DELAY = 5000;

class Notification {
    public id: number;
    private type: NotificationType;
    private messsage: string;

    public el: HTMLDivElement;

    constructor(id: number, type: NotificationType, message: string) {
        this.id = id;
        this.type = type;
        this.messsage = message;

        this.el = getNotification(this.type, this.messsage, () => {
            this.startRemoving();
        });

        this.el.addEventListener("animationend", () => {
            if(!this.el.classList.contains("fadeout")) {
                this.el.classList.remove("fadein");

                if(this.type !== NotificationType.Loading) {
                    setTimeout(() => {
                        this.el.classList.add("fadeout");
                    }, NOTIFICATION_DELAY);
                }
            } else {
                this.remove();
            }
        });
    }

    // starts fadeout animation that when ends removes the notification
    public startRemoving() {
        this.el.classList.add("fadeout");
    }

    private remove() {
        this.el.remove();
        notifications = notifications.filter(n => n.id !== this.id);
    }
}

function pushNotification(type: NotificationType, message: string): number {
    const id = notificationId++;

    const n = new Notification(id, type, message);
    notifications.push(n);

    container.appendChild(n.el);

    return id;
}

window.Notify = {
    success: (message: string): number => {
        return pushNotification(NotificationType.Success, message);
    },
    loading: (message: string): number => {
        return pushNotification(NotificationType.Loading, message);
    },
    error: (message: string): number => {
        return pushNotification(NotificationType.Error, message);
    },
    remove: (id: number): void => {
        const notification = notifications.find(n => n.id === id);
        if(!notification) return;

        notification.startRemoving();
    },
};

let container = document.createElement("div");

window.addEventListener("DOMContentLoaded", () => {
    container.classList.add("notify-container");
    document.body.appendChild(container);
});

export {}
