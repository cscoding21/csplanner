<script lang="ts">
    import FormErrorMessage from "./FormErrorMessage.svelte";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Text input was updated'
        });
    }

    let { fieldName, error, value, placeholder = $bindable() } = $props()

</script>


<div class="mb-6">
<label for="input" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<input type="text" id="input" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder={placeholder as string} required bind:value={value} onchange={updated} />

<FormErrorMessage message={error} />
</div>