<script lang="ts">
    import { Popover } from "flowbite-svelte";
	import type { Portfolio } from "$lib/graphql/generated/sdk";
    import { buildPortfolioTable, getCellColor, type ScheduleTable } from "$lib/services/portfolio";
    import { getID } from "$lib/utils/id";
    import { formatDate, formatPercent } from "$lib/utils/format";
	import { RiskLegend } from "../../../routes/(app)/project/components";
	import WeekPopupSummary from "./WeekPopupSummary.svelte";

    interface Props {
		portfolio: Portfolio;
        startDate?: Date;
        endDate?: Date;
        highlightProjectID?: string
	}
	let { portfolio, startDate, endDate, highlightProjectID }: Props = $props();

    let portfolioTable:ScheduleTable = $derived(buildPortfolioTable(portfolio, startDate, endDate, highlightProjectID))
</script>

{#if portfolioTable.body.length > 0}
<table class="w-full">
    <thead class="">
        <tr>
        {#each portfolioTable.header as head, index}
        
        {#if index === 0}
            <th class="text-xs text-left py-2">{head}</th>
        {:else}
            <th class="text-center text-xs  py-2">{head}</th>
        {/if}
        {/each}
    </tr>
    </thead>
    <tbody class="">
        {#each portfolioTable.body as row}
        {@const highlightClass = row.highlight ? "text-gray-100 text-sm" : "text-gray-500" }
        <tr class="border-separate">
            <td class="text-xs whitespace-nowrap">
                <div class="pr-2 py-1">
                <a href="/project/detail/{row.project.id}#schedule" class={highlightClass}>{row.label}</a>
                </div>
            </td>

            {#each row.weeks as week}
                {@const cellColor = getCellColor(week.risks, week.isPastDue)}
                {@const popWidth = week.risks.length > 0 ? "w-[600px]" : "w-64" }
                {#if week.active}
                <td class="text-center text-xs text-gray-100 py-1">
                    <div class="px-0 bg-{cellColor}-500">
                    <button class="text-xs cursor-pointer" id={"id_" + getID(row.project.id, formatDate(week.end))}>
                        <!-- {week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)} -->
                         &nbsp;&nbsp;&nbsp;&nbsp;
                        <Popover class="text-sm font-light {popWidth}" trigger="click" title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
                            <WeekPopupSummary activities={week.activities} risks={week.risks} />
                        </Popover>
                    </button> 
                    </div>
                    
                </td>
            {:else}
                <td></td>
            {/if}
            
            {/each}
        </tr>
        {/each}
        </tbody>

        <tfoot>
            <tr class="font-semibold text-sm text-gray-900 dark:text-gray-300">
            <th scope="row" class="py-3  pr-4 text-xs text-right">Capacity</th>
            {#each portfolioTable.footer as sum}
            {#if sum && sum.orgCapacity > 0}
            <td class="py-3 text-center" title={sum.allocatedHours + " / " + sum.orgCapacity}>
                {formatPercent.format(sum.allocatedHours / sum.orgCapacity)}
            </td>
            {:else}
                <td class="py-3 text-center text-gray-700" title={sum.allocatedHours + " / " + sum.orgCapacity}>n/a</td>
            {/if}
            {/each} 
            </tr>
        </tfoot>

    </table>

    <div class="mt-4">
        <RiskLegend />
    </div>
{/if}
