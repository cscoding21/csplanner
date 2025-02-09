<script lang="ts">
	import { callIf } from '$lib/utils/helpers';
	import type { Snippet } from 'svelte';
	import FormErrorMessage from './FormErrorMessage.svelte';

	interface Props {
		fieldName: string;
		placeholder?: string;
		children?: Snippet;
		error: string;
		value: string;
		update?: Function;
	}
	let {
		fieldName,
		placeholder,
		children,
		update,
		error = $bindable(),
		value = $bindable()
	}: Props = $props();
</script>

<div class="mb-6">
	<label for="input" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<input
		type="text"
		id="input"
		class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
		placeholder={placeholder as string}
		required
		bind:value
		onchange={() => callIf(update)}
	/>

	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
