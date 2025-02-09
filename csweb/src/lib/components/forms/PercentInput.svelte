<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { SalePercentOutline } from 'flowbite-svelte-icons';
	import { callIf } from '$lib/utils/helpers';
	import type { Snippet } from 'svelte';

	interface Props {
		fieldName: string;
		error: string;
		children?: Snippet;
		value: number;
		update?: Function;
	}
	let { fieldName, children, error, value = $bindable(), update }: Props = $props();
</script>

<div class="mb-6">
	<label for="percent" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<div class="flex">
		<span
			class="inline-flex items-center rounded-s-md border border-e-0 border-gray-300 bg-gray-200 px-3 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-600 dark:text-gray-400"
		>
			<SalePercentOutline size="lg" />
		</span>
		<input
			type="percent"
			id="money"
			class="block w-full min-w-0 flex-1 rounded-none rounded-e-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
			placeholder=""
			bind:value
			onchange={() => callIf(update)}
		/>
	</div>
	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
