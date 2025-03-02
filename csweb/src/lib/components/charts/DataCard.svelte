<script lang="ts">
    import { Card } from "flowbite-svelte";
	import type { Snippet } from "svelte";
    
    interface Props {
        dataPoint: string
        indicatorClass: string
        description: Snippet
        indicator: Snippet
        health?: "good"|"normal"|"bad"|undefined
    }
    let { 
        dataPoint, 
        indicatorClass,
        description,
        indicator,
        health
    }:Props = $props()


    const getLabelColor = (h:"good"|"normal"|"bad"|undefined):string => {
        switch(h) {
            case("good"):
                return "text-green-500 dark:text-green-500"
            case("bad"):
                return "text-red-500 dark:text-red-500"
            default:
                return "text-gray-900 dark:text-white"
        }
    }

    let labelColor = $derived(getLabelColor(health))
</script>


<Card size="xl" class="mr-2">
<div class="flex justify-between">
<div>
    <h5 class="leading-none text-3xl font-bold  pb-2 {labelColor}">{dataPoint}</h5>
    <p class="text-base font-normal text-gray-500 dark:text-gray-400">{@render description()}</p>
</div>
<div class="flex items-center px-2.5 py-0.5 text-base font-semibold {indicatorClass} text-center">
    {@render indicator()}
</div>
</div>
</Card>