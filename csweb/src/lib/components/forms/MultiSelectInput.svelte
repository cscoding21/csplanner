<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { callIf } from '$lib/utils/helpers';
	import { Label, MultiSelect, type SelectOptionType } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';

	interface Props {
		fieldName: string;
		error: string;
		value: string[];
		update?: Function;
		children?: Snippet;
		options: SelectOptionType<string | number>[];
	}
	let {
		fieldName,
		children,
		update,
		error = $bindable(),
		value = $bindable(),
		options = $bindable()
	}: Props = $props();
</script>

<div class="mb-6">
	<Label>
		{fieldName}
		<MultiSelect
			class="mt-2"
			items={options}
			bind:value
			on:change={() => {
				callIf(update)
			}}
		/>
	</Label>
	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>
