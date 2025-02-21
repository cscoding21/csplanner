<script lang="ts">
    import { Heading, Popover } from "flowbite-svelte";
	import type { Portfolio } from "$lib/graphql/generated/sdk";
    import { buildPortfolioTable, type ScheduleTable } from "$lib/services/portfolio";
    import { getID } from "$lib/utils/id";
    import { pluralize, formatDate, formatPercent } from "$lib/utils/format";

    interface Props {
		portfolio: Portfolio;
        startDate?: Date;
        endDate?: Date;
	}
	let { portfolio, startDate, endDate }: Props = $props();

    let portfolioTable:ScheduleTable = $state(buildPortfolioTable(portfolio, startDate, endDate))
</script>

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
                {@const cellColor = week.risks.length > 0 ? "bg-yellow-400" : "bg-green-600 " }
                {#if week.active}
                <td class="text-center text-xs p-1 text-gray-100 {cellColor}">
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
        <tfoot>
            <tr class="font-semibold text-sm text-gray-900 dark:text-white">
            <th scope="row" class="py-3 px-6 text-sm">Capacity</th>
            {#each portfolioTable.footer as sum}
            {#if sum && sum.orgCapacity > 0}
            <td class="py-3 px-6 text-center" title={sum.allocatedHours + " / " + sum.orgCapacity}>
                {formatPercent.format(sum.allocatedHours / sum.orgCapacity)}
            </td>
            {:else}
                <td class="py-3 px-6 text-center text-gray-500" title={sum.allocatedHours + " / " + sum.orgCapacity}>n/a</td>
            {/if}
            {/each} 
            </tr>
        </tfoot>
    </table>
{/if}
