<script lang="ts">
	import { dismissToast, type Toast } from '$lib/stores/toasts';
	import { Avatar } from 'flowbite-svelte';
	import {
		CheckCircleSolid,
		CloseCircleSolid,
		ExclamationCircleSolid,
		ExclamationCircleOutline
	} from 'flowbite-svelte-icons';

	interface Props {
		toast: Toast;
	}
	let { toast }: Props = $props();
</script>

<div
	id="toast-success"
	class="mb-4 flex w-full max-w-xs items-center rounded-lg bg-white p-4 text-gray-500 shadow dark:bg-gray-800 dark:text-gray-400"
	role="alert"
>
	{#if toast.type === 'success'}
		<div
			class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-green-100 text-green-500 dark:bg-green-800 dark:text-green-200"
		>
			<CheckCircleSolid class="h-5 w-5" />
			<span class="sr-only">Check icon</span>
		</div>
	{:else if toast.type === 'error'}
		<div
			class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-red-100 text-red-500 dark:bg-red-800 dark:text-red-200"
		>
			<CloseCircleSolid class="h-5 w-5" />
			<span class="sr-only">Error icon</span>
		</div>
	{:else if toast.type === 'warn'}
		<div
			class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-yellow-100 text-yellow-500 dark:bg-yellow-800 dark:text-yellow-200"
		>
			<ExclamationCircleSolid class="h-5 w-5" />
			<span class="sr-only">Warning icon</span>
		</div>
	{:else if toast.type === 'info'}
		<div
			class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-blue-100 text-blue-500 dark:bg-blue-800 dark:text-blue-200"
		>
			<ExclamationCircleOutline class="h-5 w-5" />
			<span class="sr-only">Info icon</span>
		</div>
	{:else if toast.type === 'user'}
		<div
			class="inline-flex h-8 w-8 flex-shrink-0 items-center justify-center rounded-lg bg-gray-100 text-gray-500 dark:bg-gray-800 dark:text-gray-200"
		>
			<Avatar src={toast.image} class="h-5 w-5" />
			<span class="sr-only">User profile image</span>
		</div>
	{/if}
	<div class="ms-3 text-sm font-normal">{toast.message}</div>

	{#if toast.dismissible}
		<button
			type="button"
			class="-mx-1.5 -my-1.5 ms-auto inline-flex h-8 w-8 items-center justify-center rounded-lg bg-white p-1.5 text-gray-400 hover:bg-gray-100 hover:text-gray-900 focus:ring-2 focus:ring-gray-300 dark:bg-gray-800 dark:text-gray-500 dark:hover:bg-gray-700 dark:hover:text-white"
			data-dismiss-target="#toast-success"
			aria-label="Close"
			onclick={() => dismissToast(toast.id as number)}
		>
			<span class="sr-only">Close</span>
			<svg
				class="h-3 w-3"
				aria-hidden="true"
				xmlns="http://www.w3.org/2000/svg"
				fill="none"
				viewBox="0 0 14 14"
			>
				<path
					stroke="currentColor"
					stroke-linecap="round"
					stroke-linejoin="round"
					stroke-width="2"
					d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"
				/>
			</svg>
		</button>
	{/if}
</div>
