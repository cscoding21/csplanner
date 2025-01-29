<script lang="ts">
	import { findAllResources } from '$lib/services/resource';
	import { MultiSelectInput, SectionHeading } from '$lib/components';
	import { daciSchema, getDefaultProject } from '$lib/forms/project.validation';
	import { getProject, updateProjectDaci } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, safeArray } from '$lib/utils/helpers';
	import type { UpdateProjectDaci, Project } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus } from '.';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let project:Project = $state(getDefaultProject() as Project)
	let errors: any = $state({});
	let df:UpdateProjectDaci = {
		driverIDs: [] as string[],
		approverIDs: [] as string[],
		contributorIDs: [] as string[],
		informedIDs: [] as string[]
	};

	let dids:string[] = $state([])
	let aids:string[] = $state([])
	let cids:string[] = $state([])
	let iids:string[] = $state([])

	const load = async (): Promise<Project> => {
		return await getProject(id)
			.then((proj) => {
				// daciForm = coalesceToType<UpdateProjectDaci>(proj.projectDaci, daciSchema);
				dids = safeArray(proj.projectDaci?.driver?.map((d) => d?.id) as string[])
				aids = safeArray(proj.projectDaci?.approver?.map((d) => d?.id) as string[])
				cids = safeArray(proj.projectDaci?.contributor?.map((d) => d?.id) as string[])
				iids = safeArray(proj.projectDaci?.informed?.map((d) => d?.id) as string[])

				df = {
					driverIDs: safeArray(proj.projectDaci?.driver?.map((d) => d?.id) as string[]),
					approverIDs: safeArray(proj.projectDaci?.approver?.map((d) => d?.id) as string[]),
					contributorIDs: safeArray(proj.projectDaci?.contributor?.map((d) => d?.id) as string[]),
					informedIDs: safeArray(proj.projectDaci?.informed?.map((d) => d?.id) as string[])
				};

				return proj
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectDACI): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const updateDACI = async () => {
		errors = {};

		df.driverIDs = dids
		df.approverIDs = aids
		df.contributorIDs = cids
		df.informedIDs = iids

		const dfParsed = daciSchema.cast(df);
		daciSchema
			.validate(dfParsed, { abortEarly: false })
			.then(async () => {

				await updateProjectDaci(id, dfParsed)
					.then((res) => {
						if (res.status?.success) {
							load().then((p) => {
								addToast({
									message: 'Project DACI updated successfully',
									dismissible: true,
									type: 'success'
								});

								callIf(update);
							});
						} else {
							addToast({
								message: 'Error updating project DACI: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project DACI: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.error(errors);
			});
	};

	const loadPage = async () => {
		findAllResources()
			.then((r) => r)
			.then((r) => {
				resourceOpts = r.results?.map((r) => ({
					name: r.name,
					value: r.id as string
				})) as SelectOptionType<string>[];
			})
			.then(() => {
				load().then(p => {
					project = p
				});
			});
	};

	let resourceOpts = $state([] as SelectOptionType<string>[]);
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project}
		<SectionHeading>
			DACI: {project.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}

	{#if df}
		<MultiSelectInput
			fieldName="Drivers"
			error={errors.driver}
			options={resourceOpts}
			bind:value={dids}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Approvers"
			error={errors.approver}
			options={resourceOpts}
			bind:value={aids}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Contributors"
			error={errors.contributor}
			options={resourceOpts}
			bind:value={cids}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Informed"
			error={errors.informed}
			options={resourceOpts}
			bind:value={iids}
			update={() => callIf(update)}
		/>

		<div class="col-span-4">
			<span class="float-right">
				<Button on:click={updateDACI}>Update Team</Button>
			</span>
			<br class="clear-both" />
		</div>
	{/if}
{/await}
