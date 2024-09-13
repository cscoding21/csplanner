<script lang="ts">
	import { FormErrorMessage } from '$lib/components/forms';
	import { createEventDispatcher } from 'svelte';
	import { SalePercentOutline } from 'flowbite-svelte-icons';

	const dispatch = createEventDispatcher();

	function updated() {
		dispatch('updated', {
			text: 'Percent input was updated'
		});
	}

	let { fieldName, error, value = $bindable() } = $props();
</script>

<div class="mb-6">
	<label for="percent" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<div class="flex">
		<span
			class="inline-flex items-center rounded-s-md border border-e-0 border-gray-300 bg-gray-200 px-3 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-600 dark:text-gray-400"
		>
			<SalePercentOutline size="sm" />
		</span>
		<input
			type="percent"
			id="money"
			class="block w-full min-w-0 flex-1 rounded-none rounded-e-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
			placeholder=""
			bind:value
			onchange={updated}
		/>
	</div>
	<FormErrorMessage message={error} />
</div>
