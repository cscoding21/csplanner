<script lang="ts">
    import type { ProjectActivity, ProjectActivityWeek, Portfolio, Project } from "$lib/graphql/generated/sdk";
	import { getPortfolio } from "$lib/services/portfolio";
    import { NoResults, CSSection, SectionHeading, ResourceList } from "$lib/components";
    import { formatDate, formatDateNoYear, pluralize, formatPercent } from "$lib/utils/format";
	import { dateCompare } from "$lib/utils/check";
	import { getID } from "$lib/utils/id";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell, Button, Popover } from "flowbite-svelte";

	let portfolioTable = $state({header: [] as string[], body:[] as ProjectRow[] })

	interface ProjectRow {
        label: string
        project: Project
        weeks: ProjectRowCell[]
    }
	interface ProjectRowCell {
		active: boolean
		end: Date
		orgCapacity: number
		activities: ProjectActivity[]
	}

	let orgCapacity = $state(0)

	const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		portfolioTable = buildPortfolioTable(res)

		return res;
	};

	const buildPortfolioTable = (res:Portfolio):any => {
		//weeks = getCSWeeks(new Date(res.begin), new Date(res.end))
		portfolioTable.header = ["Project", ...res.weekSummary.map(w => formatDateNoYear(w?.end))]

		for(let i = 0; i < res.schedule.length; i++) {
			let schedule = res.schedule[i]
			let row = {weeks: [] as ProjectRowCell[]} as ProjectRow

			row.label = schedule.project.projectBasics.name
			row.project = schedule.project

			for (let j = 0; j < res.weekSummary.length; j++) {
				const thisWeek = res.weekSummary[j] 
				let cell = {active:false, activities:[], end: thisWeek?.end, orgCapacity: thisWeek?.orgCapacity } as ProjectRowCell

				const paw = getWeekActivities(schedule.projectActivityWeeks as ProjectActivityWeek[], new Date(thisWeek?.end))
				cell.orgCapacity = paw.orgCapacity
				orgCapacity = paw.orgCapacity

				if (paw.activities && paw.activities.length > 0) {
					cell.active = true
					cell.activities = paw.activities as ProjectActivity[]
				}

				row.weeks.push(cell)
			}

			portfolioTable.body.push(row)
		}

		return portfolioTable
	}

	const getWeekActivities = (paWeeks:ProjectActivityWeek[], week:Date):ProjectActivityWeek  => {
		if (paWeeks && paWeeks.length > 0) {
			for (let i = 0; i < paWeeks.length; i++) {
				const paw = paWeeks[i]

				if(dateCompare(new Date(paw.end), week)) {
					return paw as ProjectActivityWeek
				}
			}
		}

		return {} as ProjectActivityWeek
	}

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
