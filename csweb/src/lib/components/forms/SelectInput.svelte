<script lang="ts">
    import { FormErrorMessage}  from "$lib/components/forms";
    import { Label, Select, type SelectOptionType } from "flowbite-svelte";
    import { createEventDispatcher } from "svelte";

    const dispatch = createEventDispatcher();

    function updated() {
        dispatch('updated', {
            text: 'Select input was updated'
        });
    }

    let { fieldName, error, value, options = $bindable() } = $props()
</script>

<div class="mb-6">
<Label>
    {fieldName}
    <Select class="mt-2" items={options as SelectOptionType<any>[]} bind:value={value} on:change={updated} />
</Label>
<FormErrorMessage message={error} />
</div>