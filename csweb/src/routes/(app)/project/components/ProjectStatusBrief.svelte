<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
	import { getID } from "$lib/utils/id";
	import { Progressbar } from "flowbite-svelte";
	import { ExclamationCircleOutline, QuestionCircleOutline } from "flowbite-svelte-icons";


    interface Props {
        project: Project
    }
    let { project }:Props = $props()

    let isAtRisk = $state(project.calculated?.unhealthyTasks as number > 0)
    let isPastDue = $state(false)

    const getBarColor = (ar:boolean, pd:boolean):"red"|"yellow"|"green" => {
        if(ar) 
            return "red"
        else if (pd)
            return "yellow"

        return "green"
    }

    let barColor = $derived(getBarColor(isAtRisk, isPastDue))


</script>

<div class="flex w-72">
<div class="w-full">
<Progressbar progress={(project.calculated?.projectPercentComplete as number) * 100.0} id={getID("mil_")} size="h-4" color={barColor} labelInside={true} />
</div>

{#if isPastDue}
    <div class="ml-2" title="Has past-due tasks"><ExclamationCircleOutline color="red" size="md"  /></div>
{/if}
{#if isAtRisk}
    <div class="ml-2" title="Has at-risk tasks"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}
</div>

