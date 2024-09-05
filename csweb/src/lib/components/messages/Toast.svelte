<script lang="ts">
    import { dismissToast, type Toast } from "$lib/stores/toasts";
    import { Avatar } from "flowbite-svelte";
    import { CheckCircleSolid, CloseCircleSolid, ExclamationCircleSolid, ExclamationCircleOutline  } from "flowbite-svelte-icons";

    export let toast: Toast;
</script>

<div id="toast-success" class="flex items-center w-full max-w-xs p-4 mb-4 text-gray-500 bg-white rounded-lg shadow dark:text-gray-400 dark:bg-gray-800" role="alert">
    {#if toast.type === "success"}
    <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
        <CheckCircleSolid class="w-5 h-5" />
        <span class="sr-only">Check icon</span>
    </div>
    {:else if toast.type === "error"}
    <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
        <CloseCircleSolid class="w-5 h-5" />
        <span class="sr-only">Error icon</span>
    </div>
    {:else if toast.type === "warn"}
    <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
        <ExclamationCircleSolid class="w-5 h-5" />
        <span class="sr-only">Warning icon</span>
    </div>
    {:else if toast.type === "info"}    
    <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
        <ExclamationCircleOutline class="w-5 h-5" />
        <span class="sr-only">Info icon</span>
    </div>
    {:else if toast.type === "user"}
    <div class="inline-flex items-center justify-center flex-shrink-0 w-8 h-8 text-green-500 bg-green-100 rounded-lg dark:bg-green-800 dark:text-green-200">
        <Avatar src={toast.image} class="w-5 h-5" />
        <span class="sr-only">User profile image</span>
    </div>
    {/if}
    <div class="ms-3 text-sm font-normal">{toast.message}</div>

    {#if toast.dismissible}
    <button type="button" class="ms-auto -mx-1.5 -my-1.5 bg-white text-gray-400 hover:text-gray-900 rounded-lg focus:ring-2 focus:ring-gray-300 p-1.5 hover:bg-gray-100 inline-flex items-center justify-center h-8 w-8 dark:text-gray-500 dark:hover:text-white dark:bg-gray-800 dark:hover:bg-gray-700" data-dismiss-target="#toast-success" aria-label="Close" on:click={() => dismissToast(toast.id)}>
        <span class="sr-only">Close</span>
        <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
        </svg>
    </button>
    {/if}
</div>