<script lang="ts">
	import { Heading, Button } from 'flowbite-svelte';
	import { MoneyInput, SectionHeading } from '$lib/components';
	import { formatCurrency } from '$lib/utils/format';
	import { getDefaultProject, costSchema } from '$lib/forms/project.validation';
	import { updateProjectCost, getProject } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { coalesceToType } from '$lib/forms/helpers';
	import type { UpdateProjectCost, Project } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';

	interface Props {
		id: string;
		update: Function;
	}

	let { id, update }: Props = $props();
	let project:Project = $state(getDefaultProject() as Project)

	let errors: any = $state({ ongoing: '', blendedRate: '' });

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
		<Heading tag="h6">Project Implementation Cost</Heading>
		<p class="text-gray-500 dark:text-gray-400">
			Implementation cost is determined by multiplying the hourly estimates for the milestone tasks
			with the blended hourly rate.
		</p>
		<div class="mb-4 text-gray-50">
			{formatCurrency.format(project.projectCost.calculated?.initialCost as number)}
		</div>

		<Heading tag="h6">Project Hours</Heading>
		<div class="mb-4 text-gray-50">{project.projectCost.calculated?.hourEstimate}</div>

		<MoneyInput
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
		</div>
	{/if}
{/await}
