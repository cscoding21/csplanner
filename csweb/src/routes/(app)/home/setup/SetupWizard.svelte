<script lang="ts">
	import type { Organization } from "$lib/graphql/generated/sdk";
    
    interface Props {
        org: Organization
    }
    let { org }:Props = $props()

    let step = $state(1)
    
    const getClass = (thisStep:number):string => {
        if (thisStep == step)
            return "text-blue-600 dark:text-blue-500"
        else if (thisStep < step)
            return "text-gray-600 dark:text-gray-500"

        return ""
    }
</script>

<!-- <textarea cols="100" rows="26">{JSON.stringify(org, null, 3)}</textarea> -->


<ol class="flex items-center w-full p-3 space-x-2 text-sm font-medium text-center text-gray-500 bg-white border border-gray-200 rounded-lg shadow-xs dark:text-gray-400 sm:text-base dark:bg-gray-800 dark:border-gray-700 sm:p-4 sm:space-x-4 rtl:space-x-reverse">
    {@render stepSnippet(1, "Welcome", true)}
    {@render stepSnippet(2, "Org Settings", true)}
    {@render stepSnippet(3, "Skills", true)}
    {@render stepSnippet(4, "Value Categories", true)}
    {@render stepSnippet(5, "Funding Sources", true)}
    {@render stepSnippet(6, "Roles", true)}
    {@render stepSnippet(7, "Review", false)}
</ol>

<div class="mt-4">
    <button onclick={() => step--}>Previous</button>
    <button onclick={() => step++}>Next</button>
</div>


{#snippet stepSnippet(number:number, name:string, hasNext:boolean)}
{#if step == number}
<li class="flex items-center text-blue-600 dark:text-blue-500">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border  border-blue-600 rounded-full shrink-0 dark:border-blue-500">{number}</span> {name}
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{:else if step > number}
<li class="flex items-center text-gray-600 dark:text-gray-500">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">{number}</span> {name}
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{:else}
<li class="flex items-center text-gray-600 dark:text-gray-500">
    <span class="flex items-center justify-center w-5 h-5 me-2 text-xs border border-gray-500 rounded-full shrink-0 dark:border-gray-400">{number}</span> {name}
    {#if hasNext}
    <svg class="w-3 h-3 ms-2 sm:ms-4 rtl:rotate-180" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 12 10">
        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 9 4-4-4-4M1 9l4-4-4-4"/>
    </svg>
    {/if}
</li>
{/if}
{/snippet}