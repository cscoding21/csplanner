import { writable } from "svelte/store";

export const toasts = writable([]);

export type Toast = {
    id: number;
    message: string;
    type: "info" | "success" | "error" | "warn" | "user";
    dismissible: boolean;
    timeout?: number;
    image?: string;
}

export const addToast = (toast:Toast) => {
    // Create a unique ID so we can easily find/remove it
    // if it is dismissible/has a timeout.
    const id = Math.floor(Math.random() * 10000);

    // Setup some sensible defaults for a toast.
    const defaults = {
        id: id,
        type: "info",
        dismissible: true,
        timeout: 4000,
    };

    toast = { ...defaults, ...toast }

    // Push the toast to the top of the list of toasts
    // @ts-ignore
    toasts.update((all) => [toast, ...all]);

    // If toast is dismissible, dismiss it after "timeout" amount of time.
    if (toast.timeout) {
        setTimeout(() => dismissToast(id), toast.timeout);
    }
};

export const dismissToast = (id:number) => {
    // @ts-ignore
    toasts.update((all) => all.filter(t => t.id !== id));
};