import * as snabbdom from "snabbdom";

declare global {
    interface Window {
        snabbdom: typeof snabbdom;
    }
}

window.snabbdom = snabbdom;
