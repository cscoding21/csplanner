<script lang="ts">
    import { Popover } from "flowbite-svelte";
	import type { Portfolio } from "$lib/graphql/generated/sdk";
    import { buildPortfolioTable, type ScheduleTable } from "$lib/services/portfolio";
    import { getID } from "$lib/utils/id";
    import { pluralize, formatDate } from "$lib/utils/format";
	import { NoResults } from "$lib/components";

    interface Props {
		portfolio: Portfolio;
	}
	let { portfolio }: Props = $props();

    const startDate = new Date()    
    const endDate = new Date(new Date().setDate(new Date().getDate() + 7 * 12));

    //@ts-ignore
    let portfolioTable:ScheduleTable = $state(buildPortfolioTable(portfolio, startDate, endDate))
</script>


<h3 class="mb-2">{formatDate(portfolioTable.startDate)} - {formatDate(portfolioTable.endDate)}</h3>

{#if portfolioTable.body.length > 0}
<table class="w-full">
    <thead class="">
        <tr>
        {#each portfolioTable.header as head, index}
        
        {#if index === 0}
            <th class="text-xs text-left py-2 px-1">{head}</th>
        {:else}
            <th class="text-center text-xs  py-2 px-1">{head}</th>
        {/if}
        {/each}
    </tr>
    </thead>
    <tbody class="">
        {#each portfolioTable.body as row}
        <tr class="border-separate">
            <td class="text-xs whitespace-nowrap py-2"><a href="/project/detail/{row.project.id}#schedule">{row.label}</a></td>

            {#each row.weeks as week}
            
                {#if week.active}
                <td class="text-center text-xs bg-green-600 p-1 text-gray-100">
                <button class="text-xs" id={"id_" + getID(row.project.id, formatDate(week.end))}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}
                <Popover class="w-64 text-sm font-light " title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
                    <div class="">
                    <ul class="">
                        {#each week.activities as activity}
                        <li class="text-left ml-2 p-1">
                            {activity.taskName}<br /> 
                            <small class="text-gray-100">{activity.hoursSpent + " " + pluralize("hour", activity.hoursSpent || 0)}</small>
                            <br class="clear-both" />
                        </li>
                        {/each}
                    </ul>
                    </div>
                </Popover>
            </button> 
                
            </td>
            {:else}
                <td></td>
            {/if}
            
            {/each}
        </tr>
        {/each}
        </tbody>
    </table>
{:else}
    <NoResults title="Resource not allocated">This resource is not allocated to any projects</NoResults>
{/if}

<!--
<table border=1>
    <thead>
        <tr>
            {#each portfolio.weekSummary as week, x}
            <th class="bg-blue-300">{formatDateNoYear(week?.end)}</th>
            {/each}
        </tr>
    </thead>
    <tbody>
        {#each portfolio.schedule as sch, y}
            {#if sch.projectActivityWeeks}
            <tr>
                {#each sch.projectActivityWeeks as week, x}
                    {@const paw = getWeekActivities(sch.projectActivityWeeks, week.end)}
                    {#if paw.activities}
                    <td><div class="bg-green-200 p-1"></div></td>
                    {:else}
                    <td><div class="bg-yellow-200 p-1">{x},{y}</div></td>
                    {/if}
                {/each}
            </tr>
            {/if}
        {/each}
        
    </tbody>
    <tfoot>
        <tr>
            {#each Array.from({ length: cols }) as _, x}
            <th></th>
            {/each}
        </tr>
    </tfoot>
</table>
-->