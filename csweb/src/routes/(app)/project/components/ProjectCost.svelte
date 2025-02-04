<script lang="ts">
	import { Table, TableBody, TableBodyRow, TableBodyCell, TableHead, TableHeadCell} from 'flowbite-svelte';
	import { SectionHeading, DataCard, SectionSubHeading } from '$lib/components';
	import { formatCurrency } from '$lib/utils/format';
	import { getDefaultProject, costSchema } from '$lib/forms/project.validation';
	import { updateProjectCost, getProject } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { coalesceToType } from '$lib/forms/helpers';
	import type { UpdateProjectCost, Project } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';
	import { ClockOutline, ClockSolid, DollarOutline } from 'flowbite-svelte-icons';

	interface Props {
		id: string;
		update: Function;
	}

	interface FlatCostDetail {
		taskID: string;
		taskName: string;
		milestoneID: string;
		milestoneName: string;
		resourceID: string;
		resourceName: string;
		estimatedHours: number;
		actualizedHours: number;
		contention: number;
		cost: number;
		costPerHour: number;
	}

	let { id, update }: Props = $props();
	let project:Project = $state(getDefaultProject() as Project)
	let costTable:FlatCostDetail[] = $state([] as FlatCostDetail[])
	//let costByResource

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
							taskName: task.name,
							resourceID: res.id as string,
							resourceName: res.name,
							contention: resourceCount,
							estimatedHours: (task.hourEstimate / resourceCount),
							actualizedHours: ((task.calculated?.actualizedHoursToComplete || 0) / resourceCount),
							cost: ((task.calculated?.actualizedCost || 0) / resourceCount),
							costPerHour: (res.annualizedCost || 0) / 2000,
						}

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
			  <TableHeadCell>Contention</TableHeadCell>
			  <TableHeadCell>Cost</TableHeadCell>
			</TableHead>
			<TableBody tableBodyClass="divide-y">
				{#each costTable as line}
			  <TableBodyRow>
				<TableBodyCell>{line.milestoneName}</TableBodyCell>
				<TableBodyCell>{line.taskName}</TableBodyCell>
				<TableBodyCell>{line.resourceName}</TableBodyCell>
				<TableBodyCell>{line.estimatedHours}</TableBodyCell>
				<TableBodyCell>{line.actualizedHours}</TableBodyCell>
				<TableBodyCell>{line.contention}</TableBodyCell>
				<TableBodyCell>{formatCurrency.format(line.cost || 0)}</TableBodyCell>
			  </TableBodyRow>
			  {/each}
			</TableBody>
		  </Table>

		<!-- <MoneyInput
			bind:value={costForm.ongoing as number}
			fieldName="Ongoing Cost of Ownership (annually)"
			error={errors.ongoing}
			update={() => update()}
		/>

		<MoneyInput
			bind:value={costForm.blendedRate as number}
			fieldName="Blended Hourly Rate"
			error={errors.blendedRate}
			update={() => update()}
		/>

		<div class="col-span-4">
			<span class="float-right">
				<Button on:click={updateCost}>Update Costs</Button>
			</span>
			<br class="clear-both" />
		</div> -->
	{/if}
{/await}
