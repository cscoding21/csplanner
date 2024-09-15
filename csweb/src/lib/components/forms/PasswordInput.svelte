<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { createEventDispatcher } from 'svelte';
	import { Input, ButtonGroup, InputAddon } from 'flowbite-svelte';
	import { EyeOutline, EyeSlashOutline } from 'flowbite-svelte-icons';

	let show = $state(false);

	const dispatch = createEventDispatcher();

	function updated() {
		dispatch('updated', {
			text: 'Password Input was updated'
		});
	}

	let { fieldName, error, value = $bindable() } = $props();
</script>

<div class="mb-6">
	<label for="password" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<ButtonGroup class="w-full">
		<InputAddon>
			<button
				onclick={() => {
					show = !show;
				}}
			>
				{#if show}
					<EyeOutline class="h-6 w-6" />
				{:else}
					<EyeSlashOutline class="h-6 w-6" />
				{/if}
			</button>
		</InputAddon>
		<Input
			type={show ? 'text' : 'password'}
			id="password"
			class="block w-full rounded-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
			required
			bind:value
			onchange={updated}
		/>
	</ButtonGroup>

	<FormErrorMessage message={error} />
</div>
