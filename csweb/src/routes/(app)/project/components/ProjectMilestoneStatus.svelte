<script lang="ts">
	import type { ProjectMilestone } from "$lib/graphql/generated/sdk";
	import { normalizeGUID } from "$lib/utils/id";
    import { Progressbar, Popover } from "flowbite-svelte";

    interface Props {
        milestone: ProjectMilestone
    }
    let { 
        milestone = $bindable() 
    }:Props = $props()

    let percentDone = $derived(Math.round(milestone.calculated?.percentDone as number * 100.0))

    const getID = (prefix:string):string => {
        return prefix + normalizeGUID(milestone.id)
    }
</script>

<div class="p2">
<div class="mb-2">
    <b>{milestone.phase.name}: </b>
    {milestone.phase.description}
</div>

{#if milestone.calculated?.isComplete}
<Progressbar progress={percentDone} id={getID("mil_")} size="h-4" color="green" labelOutside="Complete" />
{:else if milestone.calculated?.isInFlight}
<Progressbar progress={percentDone} id={getID("mil_")} size="h-4" labelOutside={"Remaining Hours: " + milestone.calculated?.hoursRemaining} />
{:else}
<Progressbar progress={percentDone} id={getID("mil_")} size="h-4" color="yellow" labelOutside="Not Started" />
{/if}

<Popover triggeredBy={"#" + getID("mil_")}>
    <div class="p-4">
    <ul>
        <li class="flex items-center">
            Number of Tasks: <span class="text-white ml-2"> {milestone.calculated?.totalTasks}</span>
        </li>
        <li class="flex items-center">
            Completed Tasks: <span class="text-white ml-2"> {milestone.calculated?.completedTasks}</span>
        </li>

        <li class="flex items-center mb-1">
            Total Hours: <span class="text-white ml-2"> {milestone.calculated?.totalHours}</span>
        </li>
        <li class="flex items-center mb-1">
            Remaining Hours: <span class="text-white ml-2"> {milestone.calculated?.hoursRemaining}</span>
        </li>
      </ul>
    </div>
</Popover>
</div>

