<script lang="ts">
    import { Popover } from "flowbite-svelte";
	import type { Portfolio } from "$lib/graphql/generated/sdk";
    import { buildPortfolioTable, type ScheduleTable } from "$lib/services/portfolio";
    import { getID } from "$lib/utils/id";
    import { formatDate, formatPercent } from "$lib/utils/format";
	import { RiskLegend } from "../../../routes/(app)/project/components";
	import WeekPopupSummary from "./WeekPopupSummary.svelte";

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
                {@const popWidth = week.risks.length > 0 ? "w-[600px]" : "w-64" }
                {#if week.active}
                <td class="text-center text-xs p-1 text-gray-100 {cellColor}">
                    <button class="text-xs" id={"id_" + getID(row.project.id, formatDate(week.end))}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}
                        <Popover class="text-sm font-light {popWidth}" title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
                            <WeekPopupSummary activities={week.activities} risks={week.risks} />
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

    <div class="mt-4">
        <RiskLegend />
    </div>
{/if}
