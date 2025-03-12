<script lang="ts">
	import { findAllResources } from '$lib/services/resource';
	import { MultiSelectInput, SectionHeading } from '$lib/components';
	import { daciSchema, getDefaultProject } from '$lib/forms/project.validation';
	import { getProject, updateProjectDaci } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, safeArray } from '$lib/utils/helpers';
	import type { UpdateProjectDaci, ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus, ShowIfStatus } from '.';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)
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

	const load = async (): Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				// daciForm = coalesceToType<UpdateProjectDaci>(proj.projectDaci, daciSchema);
				dids = safeArray(proj.data?.projectDaci?.driver?.map((d) => d?.id) as string[])
				aids = safeArray(proj.data?.projectDaci?.approver?.map((d) => d?.id) as string[])
				cids = safeArray(proj.data?.projectDaci?.contributor?.map((d) => d?.id) as string[])
				iids = safeArray(proj.data?.projectDaci?.informed?.map((d) => d?.id) as string[])

				df = {
					driverIDs: safeArray(proj.data?.projectDaci?.driver?.map((d) => d?.id) as string[]),
					approverIDs: safeArray(proj.data?.projectDaci?.approver?.map((d) => d?.id) as string[]),
					contributorIDs: safeArray(proj.data?.projectDaci?.contributor?.map((d) => d?.id) as string[]),
					informedIDs: safeArray(proj.data?.projectDaci?.informed?.map((d) => d?.id) as string[])
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

		console.log("dids", dids)

		const dfParsed = daciSchema.cast(df);
		daciSchema
			.validate(dfParsed, { abortEarly: false })
			.then(async () => {

				console.log("dfParsed", dfParsed)
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
						console.log("DACI:" + err)
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
				resourceOpts = r.results?.filter(r => r.data?.type === "human" || true).map((r) => ({
					name: r.data?.name,
					value: r.meta?.id as string
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
			DACI: {project.data?.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
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

		<ShowIfStatus scope={["new", "draft", "proposed", "approved", "backlogged", "scheduled", "inflight"]} status={project.data?.projectStatusBlock?.status}>
		<div class="col-span-4">
			<span class="float-right">
				<Button on:click={updateDACI}>Update Team</Button>
			</span>
			<br class="clear-both" />
		</div>
		</ShowIfStatus>
	{/if}
{/await}
