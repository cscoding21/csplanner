<script lang="ts">
    import type { SelectOptionType } from "flowbite-svelte";
    import { SelectInput, MoneyInput } from "$lib/components"
    import { Heading, Button } from "flowbite-svelte"

    let valueForm = $state({})
    let fundingSourceOpts = $state([] as SelectOptionType<string>[]);
</script>


{#if valueForm}
<SelectInput
    bind:value={valueForm.fundingSource as string}
    fieldName="Funding Source"
    error={errors.fundingSource}
    options={fundingSourceOpts}
    update={() => callIf(update)}
/>

<Heading tag="h6">Five Year Forecast</Heading>

<MoneyInput
    bind:value={valueForm.yearOneValue as number}
    fieldName="Estimated Year One Returns"
    error={errors.yearOneValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueForm.yearTwoValue as number}
    fieldName="Estimated Year Two Returns"
    error={errors.yearTwoValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueForm.yearThreeValue as number}
    fieldName="Estimated Year Tnree Returns"
    error={errors.yearThreeValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueForm.yearFourValue as number}
    fieldName="Estimated Year Four Returns"
    error={errors.yearFourValue}
    update={() => callIf(update)}
/>

<MoneyInput
    bind:value={valueForm.yearFiveValue as number}
    fieldName="Estimated Year Five Returns"
    error={errors.yearFiveValue}
    update={() => callIf(update)}
/>

<div class="col-span-4">
    <span class="float-right">
        <Button on:click={updateValue}>Update Value</Button>
    </span>
    <br class="clear-both" />
</div>

{/if}