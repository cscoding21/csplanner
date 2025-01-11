<script module lang="ts">
    export type Week = {
        showTasks: boolean
        weekEnding: string
        activities: ProjectActivity[]
    }
</script>

<script lang="ts">
	import type { ProjectActivity } from "$lib/graphql/generated/sdk";
    import { Popover, Button } from "flowbite-svelte";
    import { normalizeGUID } from "$lib/utils/id";
	import { pluralize } from "$lib/utils/format";

    interface Props {
        week: Week
    }
    let { week }: Props = $props();

    let uuid = normalizeGUID(crypto.randomUUID());
</script>


{#if week.activities.length > 0}
<Button  size="xs" color="green" pill id={"id_" + uuid}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
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
        <li class="list-disc text-left ml-2">
            {activity.resource?.name} ({activity.hoursSpent + " " + pluralize("hour", activity.hoursSpent || 0)} out of {activity.resource?.availableHoursPerWeek})
        </li>
        {/each}
        {/if}
    </ul>
</div>
</Popover>
{/if}