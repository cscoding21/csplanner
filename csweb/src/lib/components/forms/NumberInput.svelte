<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { callIf } from '$lib/utils/helpers';

	interface NumberInputProps {
		fieldName: string;
		error: string;
		value: number;
		min?: number;
		update?: Function;
	}
	let { 
		fieldName,
		min,
		update, 
		error = $bindable(), 
		value = $bindable() 
	}:NumberInputProps = $props();
</script>

<div class="mb-6">
	<label for="money" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
		>{fieldName}</label
	>
	<div class="flex">
		<input
			type="number"
			id="number"
			min={min}
			class="block w-full min-w-0 flex-1 rounded-none rounded-e-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
			placeholder=""
			bind:value
			onchange={() => { callIf(update) }}
		/>
	</div>
	<FormErrorMessage message={error} />
</div>
