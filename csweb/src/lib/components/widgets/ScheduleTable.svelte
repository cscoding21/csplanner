<script lang="ts">
	import type { Schedule } from "$lib/graphql/generated/sdk";
	import { Table, TableHead, TableHeadCell, TableBody, TableBodyRow, TableBodyCell } from "flowbite-svelte";
	import { ResourceList } from "..";
	import { ProjectScheduleCell, RiskLegend } from "../../../routes/(app)/project/components";
	import { getScheduleTableByResource, getScheduleTableByTask } from "$lib/services/schedule";


    interface Props {
        schedule: Schedule
        view?: "task"|"resource"
    }
    let { schedule, view }:Props = $props()

    let scheduleTable = $derived(view === "task" ? getScheduleTableByTask(schedule) : getScheduleTableByResource(schedule))
</script>

{#if scheduleTable}
<Table class="w-full">
    <TableHead>
        {#each scheduleTable.header as head, index}
        {#if index === 0}
            <TableHeadCell>{head}</TableHeadCell>
        {:else}
            <TableHeadCell class="text-center">{head}</TableHeadCell>
        {/if}
        {/each}
    </TableHead>
    <TableBody>
        {#each scheduleTable.body as row}
            <TableBodyRow>
            <TableBodyCell>
                <div>
                {#if row.resource}
                <span class="float-left mr-2">
                        <ResourceList resources={[row.resource]} size="sm" maxSize={1} />
                </span>      
                <a href="/resource/detail/{row.resource.id}">{row.label}</a>
                {:else}
                    {row.label} 
                {/if}
            </div>
            </TableBodyCell>
            {#each row.weeks as week}
            <TableBodyCell class="text-center">
                <ProjectScheduleCell week={week} />
            </TableBodyCell>
            {/each}
            </TableBodyRow>
        {/each}
    </TableBody>
</Table>

<div class="mt-4">
    <RiskLegend />
</div>
{/if}


