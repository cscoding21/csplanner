<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { callIf } from '$lib/utils/helpers';
	import type { Snippet } from 'svelte';

	interface Props {
		fieldName: string;
		error: string;
		placeholder?: string;
		children?: Snippet;
		value: string;
		rows: number;
		update?: Function;
	}
	let {
		fieldName,
		error,
		placeholder,
		children,
		update,
		value = $bindable(),
		rows = $bindable()
	}: Props = $props();
</script>

<div class="mb-6">
	<label for="input" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<textarea
		class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
		rows={rows as number}
		bind:value
		placeholder={placeholder as string}
		onchange={() => callIf(update)}
	></textarea>

	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
