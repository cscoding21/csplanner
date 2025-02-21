<script module lang="ts">
    export type Week = {
        showTasks: boolean
        weekEnding: string
        activities: ProjectActivity[]
        risks: string[]
    }
</script>

<script lang="ts">
	import type { ProjectActivity } from "$lib/graphql/generated/sdk";
    import { Popover, Button } from "flowbite-svelte";
    import { normalizeGUID } from "$lib/utils/id";
	import { pluralize } from "$lib/utils/format";
    import { ResourceList } from "$lib/components";

    interface Props {
        week: Week
    }
    let { week }: Props = $props();

    let uuid = normalizeGUID(crypto.randomUUID());
</script>


{#if week.activities.length > 0}
{@const cellColor = week.risks.length > 0 ? "yellow" : "green" }
<Button  size="xs" color={cellColor} pill id={"id_" + uuid}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
<Popover class="w-64 text-sm font-light " title={"Week ending " + week.weekEnding} triggeredBy={"#id_" + uuid}>
    <div class="p-2">
    <ul class="">
        {#if week.showTasks}
        {#each week.activities as activity}
        <li class="list-disc text-left ml-2">
            {activity.taskName} ({activity.hoursSpent + " " + pluralize("hour", activity.hoursSpent || 0)})
        </li>
        {/each}
        {:else}
        {#each week.activities as activity}
        <li class="text-left ml-2 p-2">
            <span class="float-left mr-2">
                <ResourceList resources={[activity.resource]} size="sm" maxSize={1} />
            </span>
            {activity.taskName}<br /> 
            <small class="text-gray-100">{activity.hoursSpent + " " + pluralize("hour", activity.hoursSpent || 0)}</small>
            <br class="clear-both" />
        </li>
        {/each}
        {/if}
    </ul>
</div>
</Popover>
{/if}