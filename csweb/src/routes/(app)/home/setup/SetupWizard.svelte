<script lang="ts">
	import type { Organization } from "$lib/graphql/generated/sdk";
	import { CheckCircleSolid } from "flowbite-svelte-icons";
    
    interface Props {
        org: Organization
        step: number
    }
    let { org, step = $bindable() }:Props = $props()


    let _ms:number = step
    let maxStepAttained = $derived(() => {
        if(step >= _ms) {
            _ms = step
        }

        return _ms
    })

</script>

<ol class="flex items-center justify-center w-full p-3 space-x-2 text-sm font-medium text-center text-gray-500 bg-white border border-gray-200 rounded-lg shadow-xs dark:text-gray-400 sm:text-base dark:bg-gray-800 dark:border-gray-700 sm:p-4 sm:space-x-4 rtl:space-x-reverse">
    {@render stepSnippet(1, "Welcome", true)}
    {@render stepSnippet(2, "Org Settings", true)}
    {@render stepSnippet(3, "Skills", true)}
    {@render stepSnippet(4, "Value Categories", true)}
    {@render stepSnippet(5, "Funding Sources", true)}
    {@render stepSnippet(6, "Roles", true)}
    {@render stepSnippet(7, "Resources", false)}
</ol>



{#snippet stepSnippet(number:number, name:string, hasNext:boolean)}
{#if step == number}
<li class="flex items-center text-blue-600 dark:text-blue-500 text-nowrap">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border  border-blue-600 rounded-full shrink-0 dark:border-blue-500">{number}</span> {name}
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{:else if step > number}
<li class="flex items-center text-green-600 dark:text-green-500 text-nowrap">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">{number}</span> <button onclick={() => { step = number}}>{name}</button>
    <CheckCircleSolid class="ml-2" />
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{:else if maxStepAttained() >= number}
<li class="flex items-center text-green-600 dark:text-green-500  text-nowrap">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">{number}</span> <button onclick={() => { step = number }}>{name}</button>
    <CheckCircleSolid class="ml-2" />
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{:else}
<li class="flex items-center text-gray-600 dark:text-gray-500 text-nowrap">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">{number}</span> {name}
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{/if}
{/snippet}


<!-- <div class="mt-4">
    <ButtonGroup>
        <Button onclick={() => { setStep(--step) }} disabled={step == 1}><ArrowLeftOutline /> Previous</Button>
        <Button onclick={() => { setStep(++step) }} disabled={step == 7}>Next <ArrowRightOutline /></Button>
    </ButtonGroup>
</div> -->