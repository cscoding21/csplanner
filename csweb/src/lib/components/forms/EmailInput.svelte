<script lang="ts">
	import { FormErrorMessage } from '.';
	import { createEventDispatcher } from 'svelte';

	const dispatch = createEventDispatcher();

	function updated() {
		dispatch('updated', {
			text: 'Email input was updated'
		});
	}

	let { fieldName, error, value = $bindable() } = $props();
</script>

<div class="mb-6">
	<label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<input
		type="email"
		id="email"
		class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
		placeholder="name@csplanner.io"
		required
		bind:value
		onchange={updated}
	/>

	<FormErrorMessage message={error} />
</div>
