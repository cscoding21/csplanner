<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { Label, MultiSelect, type SelectOptionType } from 'flowbite-svelte';

	interface Props {
		fieldName: string;
		error: string;
		value: string[];
		update?: Function;
		options: SelectOptionType<string | number>[];
	}
	let {
		fieldName,
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
				if (update) {
					update();
				}
			}}
		/>
	</Label>
	<FormErrorMessage message={error} />
</div>
