<script lang="ts">
    import type { Portfolio } from "$lib/graphql/generated/sdk";
	import type { ProjectRow, ScheduleTable } from "$lib/services/portfolio";
	import { getPortfolio, buildPortfolioTable } from "$lib/services/portfolio";
    import { NoResults, CSSection, SectionHeading, ResourceList } from "$lib/components";
    import { formatDate, pluralize, formatPercent } from "$lib/utils/format";
	import { getID } from "$lib/utils/id";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Popover } from "flowbite-svelte";

	let portfolioTable:ScheduleTable = $state({header: [] as string[], body:[] as ProjectRow[] } as ScheduleTable)

	const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		//@ts-ignore
		portfolioTable = buildPortfolioTable(res)

		return res;
	};

	let portfolio = $state({} as Portfolio);
	const loadPage = async () => {
		refresh().then((r) => {
			portfolio = r as Portfolio;
		});
	};
</script>

<div class="p-4">
<CSSection>

	<SectionHeading>Current Roadmap</SectionHeading>

{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	{#if portfolio.schedule != null}
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
			<TableBody tableBodyClass="divide-y">
				{#each portfolioTable.body as row}
					<TableBodyRow>
					<TableBodyCell>
					<a href="/project/detail/{row.project.id}#schedule">
						{row.label}
					</a>
					</TableBodyCell>
					{#each row.weeks as week}
					
					<TableBodyCell tdClass="text-center">
						{#if week.active}
						<Button  size="xs" color="green" pill id={"id_" + getID(row.project.id, formatDate(week.end))}>{week.activities.reduce((acc, curr) => acc + (curr.hoursSpent || 0), 0)}</Button>
						<Popover class="w-64 text-sm font-light " title={"Week ending " + formatDate(week.end)} triggeredBy={"#id_" + getID(row.project.id, formatDate(week.end))}>
							<div class="p-2">
							<ul class="">
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
							</ul>
						</div>
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

	{:else}
		<NoResults title="No Projects" newUrl="">
			No projects have been scheduled.  
		</NoResults>
	{/if}
{/await}

</CSSection>
</div>
