<script lang="ts">
	import type { ProjectActivity } from "$lib/graphql/generated/sdk";
	import { formatDate, pluralize } from "$lib/utils/format";
	import { Alert } from "flowbite-svelte";
	import { InfoCircleSolid } from "flowbite-svelte-icons";
	import { ResourceList } from "..";

    interface Props {
		activities: ProjectActivity[];
        risks: string[];
	}
	let { activities, risks }: Props = $props();

    
</script>

<div class="p-2 flex">
    <div class="flex-none">
    <ul class="">
        {#each activities as activity}
        <li class="text-left ml-2 p-2">
            <span class="float-left mr-2">
                <ResourceList resources={[activity.resource]} size="sm" maxSize={1} />
            </span>
            {activity.taskName}<br /> 
            <small class="text-gray-800 dark:text-gray-100">{activity.hoursSpent + " " + pluralize("hour", activity.hoursSpent || 0)}</small><br />
            <!-- <small class="text-gray-800 dark:text-gray-100">Task completes on <b>{formatDate(activity.taskEndDate)}</b></small> -->
            <br class="clear-both" />
        </li>
        {/each}
    </ul>
    </div>
    {#if risks && risks.length > 0}
    <div class="flex-1 ml-8">
    <Alert class="items-start! text-left" border>
       {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
        <ul class="mt-1.5 ms-4 list-disc list-inside text-wrap">
            {#each risks as risk}
                <li>{risk}</li>
            {/each}
        </ul>
    </Alert> 
    </div>
    <br class="clear-both" />
    {/if}
</div>