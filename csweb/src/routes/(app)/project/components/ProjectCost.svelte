<script lang="ts">
	import { Table, TableBody, TableBodyRow, TableBodyCell, TableHead, TableHeadCell} from 'flowbite-svelte';
	import { SectionHeading, DataCard, SectionSubHeading, ResourceList, PieChart } from '$lib/components';
	import { formatCurrency } from '$lib/utils/format';
	import { getDefaultProject} from '$lib/forms/project.validation';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import type { ProjectEnvelope, Resource } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';
	import { ClockOutline, ClockSolid, DollarOutline } from 'flowbite-svelte-icons';
	import { csGroupBy, deepCopy } from '$lib/utils/helpers';
	

	interface Props {
		id: string;
		update: Function;
	}

	interface FlatCostDetail {
		taskID: string;
		taskName: string;
		taskSkill: string;
		milestoneID: string;
		milestoneName: string;
		resourceID: string;
		resourceName: string;
		resource: Resource;
		estimatedHours: number;
		actualizedHours: number;
		resourceCount: number;
		cost: number;
		costPerHour: number;
	}

	let { id }: Props = $props();
	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)
	let costTable:FlatCostDetail[] = $state([] as FlatCostDetail[])

	let resourcePieValues:any[] = $state([] as any[])
	let milestonePieValues:any[] = $state([] as any[])

	let calculatedCostSum:number = $state(0.0);

	const buildFlatCostDetails = (pro:ProjectEnvelope):FlatCostDetail[] => {
		let out = [] as FlatCostDetail[]
		if(!pro.data?.projectMilestones) {
			return out
		}

		for(let i = 0; i < pro.data?.projectMilestones?.length; i++) {
			const milestone = pro.data?.projectMilestones[i]

			if(!milestone || !milestone.tasks) {
				continue
			}

			for (let j = 0; j < milestone.tasks.length; j++) {
				const task = milestone.tasks[j]

				if(task.resources) {
					const resourceCount = task.resources.length
					for (let k = 0; k < resourceCount; k++) {
						const res = task.resources[k]

						let item:FlatCostDetail = {
							milestoneID: milestone.id,
							milestoneName: milestone.phase.name,
							taskID: task.id,
							taskName: task.name,
							taskSkill: task.requiredSkillID,
							resourceID: res.id as string,
							resourceName: res.name as string,
							resource: res,
							resourceCount: resourceCount,
							estimatedHours: (task.hourEstimate / resourceCount),
							actualizedHours: ((task.calculated?.actualizedHoursToComplete || 0) / resourceCount),
							cost: ((task.calculated?.actualizedCost || 0) / resourceCount),
							costPerHour: (task.calculated?.averageHourlyRate || 0),
						}

						calculatedCostSum += item.cost
						out.push(item)
					}
				}
			}
		}

		return out
	}

	const load = async ():Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				project = proj;

				return proj
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectCost): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const loadPage = async () => {
		load().then(p => {
			project = p

			costTable = buildFlatCostDetails(project)

			resourcePieValues = csGroupBy(deepCopy(costTable), "resourceName", "cost")
			milestonePieValues = csGroupBy(costTable, "milestoneName", "cost")
		});
	};
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project} 
		<SectionHeading>
			Estimated Implementation Cost: {project.data?.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}

	<div class="flex mb-8">
		<div class="flex-1 px-r">
			<DataCard dataPoint={formatCurrency.format(project.data?.projectCost.calculated?.initialCost as number)} indicatorClass="text-green-500 dark:text-green-500">
				{#snippet description()}
					Implementation cost
				{/snippet}
				{#snippet indicator()}
					<DollarOutline />
				{/snippet}
			</DataCard>
		</div>

		<div class="flex-1 px-2">
			<DataCard dataPoint={project.data?.projectCost.calculated?.hourEstimate + ""} indicatorClass="text-yellow-500 dark:text-yellow-500">
				{#snippet description()}
					Estimated hours
				{/snippet}
				{#snippet indicator()}
					<ClockOutline />
				{/snippet}
			</DataCard>
		</div>

		<div class="flex-1 pl-2">
			<DataCard dataPoint={project.data?.projectCost.calculated?.hoursActualized+ ""} indicatorClass="text-orange-500 dark:text-orange-500">
				{#snippet description()}
					Adjusted hours
				{/snippet}
				{#snippet indicator()}
					<ClockSolid />
				{/snippet}
			</DataCard>
		</div>

</div>

<div class="flex mb-8">
	<div class="flex-1 px-2">
		<SectionSubHeading>Cost by Resource</SectionSubHeading>
		{#if resourcePieValues}
			<PieChart labels={Object.keys(resourcePieValues)} values={Object.values(resourcePieValues)} format="currency" />
		{/if}
	</div>

	<div class="flex-1 px-2">
		<SectionSubHeading>Cost by Milestone</SectionSubHeading>
		{#if milestonePieValues}
		<PieChart labels={Object.keys(milestonePieValues)} values={Object.values(milestonePieValues)} format="currency" />
		{/if}
	</div>

</div>
<SectionSubHeading>Cost by Task</SectionSubHeading>
		<Table hoverable={true}>
			<TableHead>
			  <TableHeadCell>Milestone</TableHeadCell>
			  <TableHeadCell>Task</TableHeadCell>
			  <TableHeadCell>Resource</TableHeadCell>
			  <TableHeadCell>Estimated hours</TableHeadCell>
			  <TableHeadCell>Actual hours</TableHeadCell>
			  <TableHeadCell>Cost per Hour</TableHeadCell>
			  <TableHeadCell>Cost</TableHeadCell>
			</TableHead>
			<TableBody tableBodyClass="divide-y">
				{#each costTable as line}
			  <TableBodyRow>
				<TableBodyCell>{line.milestoneName}</TableBodyCell>
				<TableBodyCell>
					<div>{line.taskName}</div>
					<small class="clear-both dark:text-gray-400 text-gray-600">{line.taskSkill}</small>
				</TableBodyCell>
				<TableBodyCell>
					<div>
                        <span class="float-left mr-2"><ResourceList resources={[line.resource]} size="sm" maxSize={1} /></span>
						{line.resource.name}
					</div>
				</TableBodyCell>
				<TableBodyCell>{line.estimatedHours}</TableBodyCell>
				<TableBodyCell>{line.actualizedHours}</TableBodyCell>
				<TableBodyCell tdClass="text-right">{formatCurrency.format(line.costPerHour || 0)}</TableBodyCell>
				<TableBodyCell tdClass="text-right">{formatCurrency.format(line.cost || 0)}</TableBodyCell>
			  </TableBodyRow>
			  {/each}
			</TableBody>
			<tfoot>
				<tr>
					<th colspan="3">&nbsp;</th>
					<th>&nbsp;</th>
					<th>&nbsp;</th>
					<th colspan="1">&nbsp;</th>
					<th class="text-right text-lg text-gray-100">{formatCurrency.format(calculatedCostSum)}</th>
				</tr>
			</tfoot>
		  </Table>
{/await}
