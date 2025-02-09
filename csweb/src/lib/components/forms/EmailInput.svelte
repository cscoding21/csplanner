<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import type { Snippet } from 'svelte';

	interface Props {
		fieldName: string;
		error: string;
		value: string;
		placeholder?: string;
		update: Function;
		children?: Snippet;
	}
	let {
		fieldName,
		error = $bindable(),
		value = $bindable(),
		placeholder,
		children,
		update
	}: Props = $props();
</script>

<div class="mb-6">
	<label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<input
		type="email"
		id="email"
		class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
		{placeholder}
		required
		bind:value
		onchange={() => update()}
	/>

	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
