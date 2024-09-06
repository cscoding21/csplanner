<script lang="ts">
    import { FormErrorMessage } from "$lib/components/forms";
    import { DollarOutline } from "flowbite-svelte-icons";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Money input was updated'
        });
    }

    let { fieldName, error, value = $bindable() } = $props()
</script>

<div class="mb-6">
<label for="money" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<div class="flex">
  <span class="inline-flex items-center px-3 text-sm text-gray-900 bg-gray-200 border border-e-0 border-gray-300 rounded-s-md dark:bg-gray-600 dark:text-gray-400 dark:border-gray-600">
    <DollarOutline size="sm" />
  </span>
  <input type="number" id="money" class="rounded-none rounded-e-lg bg-gray-50 border border-gray-300 text-gray-900 focus:ring-blue-500 focus:border-blue-500 block flex-1 min-w-0 w-full text-sm p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" placeholder="" bind:value={value} onchange={updated} />
</div>
<FormErrorMessage message={error} />
</div>