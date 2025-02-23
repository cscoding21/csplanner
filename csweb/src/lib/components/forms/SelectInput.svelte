<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { Label, Select, type SelectOptionType } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';

	interface Props {
		fieldName: string;
		error?: string;
		value: string;
		children?: Snippet;
		update?: Function;
		options: SelectOptionType<string | number>[];
	}
	let {
		fieldName,
		update,
		children,
		error = $bindable(),
		value = $bindable(),
		options = $bindable()
	}: Props = $props();
</script>

<div class="mb-6">
	<Label>
		{fieldName}
		<Select
			class="mt-2"
			items={options}
			bind:value
			on:change={() => {
				if (update) {
					update();
				}
			}}
		/>
	</Label>
	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
