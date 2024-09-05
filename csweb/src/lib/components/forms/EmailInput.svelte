<script lang="ts">
    import { FormErrorMessage } from ".";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Email input was updated'
        });
    }

    export let fieldName:string = "";
    export let error:string = "";
    export let value:any;

    $: value
    $: error

</script>

<div class="mb-6">
<label for="email" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<input type="email" id="email" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="name@analyzer.net" required bind:value={value} on:change={updated} />

<FormErrorMessage message={error} />
</div>