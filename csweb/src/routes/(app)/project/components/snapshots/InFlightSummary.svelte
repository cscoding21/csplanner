<script lang="ts">
	import { DataCard } from "$lib/components";
	import { formatDate, pluralize } from "$lib/utils/format";
	import { CalendarWeekOutline } from "flowbite-svelte-icons";
	import type { Schedule } from "$lib/graphql/generated/sdk";

    interface Props {
        schedule: Schedule
    }
    let { schedule }:Props = $props()
</script>

{#if schedule}
<div class="flex mb-8">
    <div class="flex-1 px-r mr-2">
        <DataCard dataPoint={formatDate(schedule.end)} indicatorClass="text-green-500 dark:text-green-500">
            {#snippet description()}
                Complete date
            {/snippet}
            {#snippet indicator()}
                <CalendarWeekOutline />
            {/snippet}
        </DataCard>
    </div>

    <div class="flex-1 px-r">
        <DataCard dataPoint={"Fill Me"} indicatorClass="text-green-500 dark:text-green-500">
            {#snippet description()}
                Progress
            {/snippet}
            {#snippet indicator()}
                <CalendarWeekOutline />
            {/snippet}
        </DataCard>
    </div>

    <div class="flex-1 px-2">
        <DataCard dataPoint={schedule.projectActivityWeeks?.length + ""} indicatorClass="text-green-500 dark:text-green-500">
            {#snippet description()}
                {pluralize("Week", schedule.projectActivityWeeks?.length as number)}
            {/snippet}
            {#snippet indicator()}
                <CalendarWeekOutline />
            {/snippet}
        </DataCard>
    </div>
</div>
{/if}