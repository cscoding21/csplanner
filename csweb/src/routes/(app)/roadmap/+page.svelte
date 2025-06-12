<script lang="ts">
    import type { Portfolio, ProjectActivity } from "$lib/graphql/generated/sdk";
	import type { ProjectRow, ScheduleTable } from "$lib/services/portfolio";
	import { getPortfolio, buildPortfolioTable, getCellColor, findOpenPastDueTasksInPortfolio } from "$lib/services/portfolio";
    import { NoResults, CSSection, DataCard, WeekPopupSummary, SectionSubHeading } from "$lib/components";
    import { formatDate, formatPercent, formatCurrency, pluralize } from "$lib/utils/format";
	import { getID } from "$lib/utils/id";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Popover, Alert } from "flowbite-svelte";
	import { RiskLegend } from "../project/components";
	import { DollarOutline, InfoCircleSolid } from "flowbite-svelte-icons";
	import { RoadmapActionBar } from "./components";

	let portfolioTable:ScheduleTable = $state({header: [] as string[], body:[] as ProjectRow[] } as ScheduleTable)
	let pastDueTasks:ProjectActivity[] = $state([])

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

	const getPastDueMessage = ():string => {
		if (!pastDueTasks || pastDueTasks.length == 0) {
			return "None"
		}

		const hours =  pastDueTasks.reduce( function(a, b){
			return a + b.hoursSpent;
		}, 0);

		return pastDueTasks.length + " (" + hours +  pluralize(" hour", hours) + ")"
	}

	let portfolio = $state({} as Portfolio);
	const loadPage = async () => {
		refresh().then((r) => {
			portfolio = r as Portfolio;

			pastDueTasks = findOpenPastDueTasksInPortfolio(portfolio)
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
			<div class="flex-1  pl-2">
				<DataCard health={pastDueTasks.length === 0 ? "good" : "bad"} dataPoint={getPastDueMessage()} indicatorClass="text-red-500 dark:text-red-500">
					{#snippet description()}
						Past Due Tasks
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>
		</div>

		{#if pastDueTasks && pastDueTasks.length > 0}
			<div class="my-4">
			<Alert class="items-start!" border>
				{#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
				<p class="font-medium">The following tasks are past due:</p>
				<ul class="mt-1.5 ms-4 list-disc list-inside">
					{#each pastDueTasks as t}
						<li>{t.project?.projectBasics.name}: {t.taskName} ({t.hoursSpent} {pluralize("hour", t.hoursSpent)}) was due on {formatDate(t.taskEndDate)}</li>
					{/each}
				</ul>
			</Alert> 
			</div>
		{/if}

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
						<Button  size="xs"  class="cursor-pointer" color={cellColor} id={"id_" + getID(row.project.id, formatDate(week.end))}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
						<Popover class=" text-sm font-light {popWidth}" title={"Week ending " + formatDate(week.end)} trigger="click" triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
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
