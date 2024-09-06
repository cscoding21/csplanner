<script lang="ts">
    import { FormErrorMessage } from "$lib/components/forms";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Number input was updated'
        });
    }

    let { fieldName, error, value = $bindable() } = $props()
</script>

<div class="mb-6">
<label for="money" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<div class="flex">
  <input type="number" id="number" class="rounded-none rounded-e-lg bg-gray-50 border border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="" bind:value={value} onchange={updated} />
</div>
<FormErrorMessage message={error} />
</div>