<script lang="ts">
	import { Table, TableBody, TableBodyRow, TableBodyCell, TableHead, TableHeadCell} from 'flowbite-svelte';
	import { SectionHeading, DataCard, SectionSubHeading } from '$lib/components';
	import { formatCurrency } from '$lib/utils/format';
	import { getDefaultProject, costSchema } from '$lib/forms/project.validation';
	import { updateProjectCost, getProject } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { coalesceToType } from '$lib/forms/helpers';
	import type { UpdateProjectCost, Project, Resource } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';
	import { ClockOutline, ClockSolid, DollarOutline } from 'flowbite-svelte-icons';
	import ResourceList from '$lib/components/widgets/ResourceList.svelte';

	interface Props {
		id: string;
		update: Function;
	}

	interface FlatCostDetail {
		taskID: string;
		taskName: string;
		milestoneID: string;
		milestoneName: string;
		resource: Resource;
		estimatedHours: number;
		actualizedHours: number;
		resourceCount: number;
		cost: number;
		costPerHour: number;
	}

	let { id, update }: Props = $props();
	let project:Project = $state(getDefaultProject() as Project)
	let costTable:FlatCostDetail[] = $state([] as FlatCostDetail[])
	let calculatedCostSum:number = $state(0.0);

	let errors: any = $state({ ongoing: '', blendedRate: '' });

	const buildFlatCostDetails = (pro:Project):FlatCostDetail[] => {
		let out = [] as FlatCostDetail[]
		if(!pro.projectMilestones) {
			return out
		}

		for(let i = 0; i < pro.projectMilestones?.length; i++) {
			const milestone = pro.projectMilestones[i]

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
							taskName: task.name + " (" + task.requiredSkillID + ")",
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

	const load = async ():Promise<Project> => {
		return await getProject(id)
			.then((proj) => {
				project = proj;
				costForm = coalesceToType<UpdateProjectCost>(proj.projectCost, costSchema);

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

	const updateCost = async () => {
		errors = {};

		const projectCostParsed = costSchema.cast(costForm);
		costSchema
			.validate(projectCostParsed, { abortEarly: false })
			.then(async () => {
				await updateProjectCost(id, projectCostParsed)
					.then((res) => {
						if (res.status?.success) {
							load().then((p) => {
								addToast({
									message: 'Project cost updated successfully',
									dismissible: true,
									type: 'success'
								});

								update();
							});
						} else {
							addToast({
								message: 'Error updating project basics: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project basics: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.log(errors);
			});
	};

	const loadPage = async () => {
		load().then(p => {
			project = p

			costTable = buildFlatCostDetails(project)
		});
	};

	let costForm = $state(getDefaultProject().projectCost);
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project} 
		<SectionHeading>
			Estimated Implementation Cost: {project.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}
	{#if costForm}
	<div class="flex mb-8">
		<div class="flex-1 px-2">
	<DataCard dataPoint={formatCurrency.format(project.projectCost.calculated?.initialCost as number)} indicatorClass="text-green-500 dark:text-green-500">
		{#snippet description()}
			Implementation costs
		{/snippet}
		{#snippet indicator()}
			<DollarOutline />
		{/snippet}
	</DataCard>
</div>

	<div class="flex-1 px-2">
<DataCard dataPoint={project.projectCost.calculated?.hourEstimate + ""} indicatorClass="text-yellow-500 dark:text-yellow-500">
	{#snippet description()}
		Estimated hours
	{/snippet}
	{#snippet indicator()}
		<ClockOutline />
	{/snippet}
</DataCard>
</div>

	<div class="flex-1 px-2">
<DataCard dataPoint={project.projectCost.calculated?.hoursActualized+ ""} indicatorClass="text-orange-500 dark:text-orange-500">
	{#snippet description()}
		Adjusted hours
	{/snippet}
	{#snippet indicator()}
		<ClockSolid />
	{/snippet}
</DataCard>
</div>

</div>



<SectionSubHeading>Cost Table</SectionSubHeading>
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
				<TableBodyCell>{line.taskName}</TableBodyCell>
				<TableBodyCell>
					<div>
                        <span class="float-left mr-2"><ResourceList resources={[line.resource]} size="sm" maxSize={1} /></span>
						{line.resource.name}
					</div>
				</TableBodyCell>
				<TableBodyCell>{line.estimatedHours}</TableBodyCell>
				<TableBodyCell>{line.actualizedHours}</TableBodyCell>
				<TableBodyCell>{formatCurrency.format(line.costPerHour || 0)}</TableBodyCell>
				<TableBodyCell>{formatCurrency.format(line.cost || 0)}</TableBodyCell>
			  </TableBodyRow>
			  {/each}
			</TableBody>
			<tfoot>
				<tr>
					<th colspan="3">&nbsp;</th>
					<th>&nbsp;</th>
					<th>&nbsp;</th>
					<th colspan="1">&nbsp;</th>
					<th>{formatCurrency.format(calculatedCostSum)}</th>
				</tr>
			</tfoot>
		  </Table>
	{/if}
{/await}
