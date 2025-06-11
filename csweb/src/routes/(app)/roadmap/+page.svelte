<script lang="ts">
    import type { Portfolio } from "$lib/graphql/generated/sdk";
	import type { ProjectRow, ScheduleTable } from "$lib/services/portfolio";
	import { getPortfolio, buildPortfolioTable } from "$lib/services/portfolio";
    import { NoResults, CSSection, DataCard, WeekPopupSummary, SectionSubHeading } from "$lib/components";
    import { formatDate, formatPercent, formatCurrency } from "$lib/utils/format";
	import { getID } from "$lib/utils/id";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Popover } from "flowbite-svelte";
	import { RiskLegend } from "../project/components";
	import { DollarOutline } from "flowbite-svelte-icons";
	import { RoadmapActionBar } from "./components";

	let portfolioTable:ScheduleTable = $state({header: [] as string[], body:[] as ProjectRow[] } as ScheduleTable)

	const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		portfolioTable = buildPortfolioTable(res, undefined, undefined, "")

		return res;
	};

	const getHealth = (data:number) => {
		if(data < 0) {
			return "bad"
		}
		
		return "good"
	}

	const getCellColor = (risks:string[], isPastDue:boolean):"red"|"yellow"|"green"|undefined => {
		if(isPastDue) {
			return "red"
		} else if (risks && risks.length > 0) {
			return "yellow"
		}

		return "green"
	}

	let portfolio = $state({} as Portfolio);
	const loadPage = async () => {
		refresh().then((r) => {
			portfolio = r as Portfolio;
		});
	};
</script>

<RoadmapActionBar pageDetail="" />

<div class="p-4">

{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	<CSSection>

		{#if portfolio && portfolio.schedule && portfolio.schedule.length > 0}
		<div class="flex mb-8">
			<div class="flex-1 pr-2">
				<DataCard health={getHealth(portfolio.calculated?.totalValue)} dataPoint={formatCurrency.format(portfolio.calculated?.totalValue as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Total portfolio value
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>	
			<div class="flex-1  px-2">
				<DataCard health={getHealth(portfolio.calculated?.valueInFlight)} dataPoint={formatCurrency.format(portfolio.calculated?.valueInFlight as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Value in flight
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>
			<div class="flex-1  pl-2">
				<DataCard health={getHealth(portfolio.calculated?.valueScheduled)} dataPoint={formatCurrency.format(portfolio.calculated?.valueScheduled as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Value scheduled
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>
		</div>

		<SectionSubHeading>Portfolio Schedule</SectionSubHeading>
		<Table divClass="h-full">
			<TableHead>
				{#each portfolioTable.header as head, index}
				{#if index === 0}
					<TableHeadCell>{head}</TableHeadCell>
				{:else}
					<TableHeadCell class="text-center">{head}</TableHeadCell>
				{/if}
				{/each}
			</TableHead>
			<TableBody class="divide-y">
				{#each portfolioTable.body as row}
					<TableBodyRow>
					<TableBodyCell>
					<a href="/project/detail/{row.project.id}#schedule">
						{row.label}
					</a>
					</TableBodyCell>
					{#each row.weeks as week}
					
					<TableBodyCell class="text-center">
						{#if week.active}
						{@const cellColor = getCellColor(week.risks, week.isPastDue) }
						{@const popWidth = week.risks.length > 0 ? "w-[600px]" : "w-64" }
						<Button  size="xs" color={cellColor} id={"id_" + getID(row.project.id, formatDate(week.end))}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
						<Popover class=" text-sm font-light {popWidth}" title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
							<WeekPopupSummary activities={week.activities} risks={week.risks} />
						</Popover>
						
						{/if}
					</TableBodyCell>
					{/each}
					</TableBodyRow>
				{/each}
			</TableBody>
			<tfoot>
				<tr class="font-semibold text-gray-900 dark:text-white">
				  <th scope="row" class="py-3 px-6 text-base">Capacity</th>
				  {#each portfolio.weekSummary as week}
				  {#if week}
				  <td class="py-3 px-6 text-center">{formatPercent.format(week.allocatedHours / week.orgCapacity)} <br />
					<small class="whitespace-nowrap">{week.allocatedHours} / {week.orgCapacity}</small>
				  </td>
				  {/if}
				  {/each} 
				</tr>
			  </tfoot>
		</Table>

		<div class="mt-4">
			<RiskLegend />
		</div>
	{:else}
		<NoResults title="No Scheduled Projects" newUrl="">

			<div class="text-center p-8">
			No projects have been scheduled.  
			</div>

		</NoResults>
	{/if}

	</CSSection>
{/await}

</div>
