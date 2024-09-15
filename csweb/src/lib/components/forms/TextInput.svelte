<script lang="ts">
	import FormErrorMessage from './FormErrorMessage.svelte';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function updated() {
		dispatch('updated', {
			text: 'Text input was updated'
		});
	}

	let { fieldName, placeholder, error = $bindable(), value = $bindable()} = $props();
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
		onchange={updated}
	/>

	<FormErrorMessage message={error} />
</div>
