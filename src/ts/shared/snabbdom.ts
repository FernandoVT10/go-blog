import {
    init,
    classModule,
    propsModule,
    styleModule,
    attributesModule,
    eventListenersModule,
} from "snabbdom";

export const patch = init([
    classModule,
    propsModule,
    styleModule,
    attributesModule,
    eventListenersModule,
]);
