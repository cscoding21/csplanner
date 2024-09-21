<script lang="ts">
    import { Card } from "flowbite-svelte";
	import type { Snippet } from "svelte";
    
    interface Props {
        dataPoint: string
        indicatorClass: string
        description: Snippet
        indicator: Snippet
    }
    let { 
        dataPoint, 
        indicatorClass,
        description,
        indicator
    }:Props = $props()
</script>


<Card size="xl" class="mr-2">
<div class="flex justify-between">
<div>
    <h5 class="leading-none text-3xl font-bold text-gray-900 dark:text-white pb-2">{dataPoint}</h5>
    <p class="text-base font-normal text-gray-500 dark:text-gray-400">{@render description()}</p>
</div>
<div class="flex items-center px-2.5 py-0.5 text-base font-semibold {indicatorClass} text-center">
    {@render indicator()}
</div>
</div>
</Card>