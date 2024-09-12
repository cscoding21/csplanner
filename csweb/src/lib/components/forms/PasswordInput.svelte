<script lang="ts">
    import { FormErrorMessage } from ".";
    import { createEventDispatcher } from "svelte";
    import { Input, ButtonGroup, InputAddon } from 'flowbite-svelte';
    import { EyeOutline, EyeSlashOutline } from 'flowbite-svelte-icons';

    let show = $state(false);

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Password Input was updated'
        });
    }

    let { fieldName, error, value = $bindable() } = $props()
</script>

<div class="mb-6">
<label for="password" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">{fieldName}</label>
<ButtonGroup class="w-full">
<InputAddon>
    <button onclick={() => { (show = !show); }}>
    {#if show}
        <EyeOutline class="w-6 h-6" />
    {:else}
        <EyeSlashOutline class="w-6 h-6" />
    {/if}
    </button>
</InputAddon>
<Input type={show ? 'text' : 'password'} id="password" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500" required bind:value={value} onchange={updated} />
</ButtonGroup>

<FormErrorMessage message={error} />
</div>
