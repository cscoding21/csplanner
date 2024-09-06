<script lang="ts">
    import FormErrorMessage from "./FormErrorMessage.svelte";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Textarea input was updated'
        });
    }

    let { fieldName, error, value, placeholder, rows = $bindable() } = $props()
</script>


<div class="mb-6">
<label for="input" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<textarea class="block p-2.5 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" rows={rows as number} bind:value={value} placeholder={placeholder as string} onchange={updated}></textarea>   

<FormErrorMessage message={error} />
</div>