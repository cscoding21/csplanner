<script lang="ts">
	import { findAllResources } from '$lib/services/resource';
	import { MultiSelectInput, SectionHeading } from '$lib/components';
	import { daciSchema } from '$lib/forms/project.validation';
	import { getProject, updateProjectDaci } from '$lib/services/project';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, safeArray } from '$lib/utils/helpers';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();
	let errors: any = $state({});
	let df = $state({
		driverIDs: [] as string[],
		approverIDs: [] as string[],
		contributorIDs: [] as string[],
		informedIDs: [] as string[]
	});

	const load = async () => {
		await getProject(id)
			.then((proj) => {
				// daciForm = coalesceToType<UpdateProjectDaci>(proj.projectDaci, daciSchema);
				df = {
					driverIDs: safeArray(proj.projectDaci?.driver?.map((d) => d?.id) as string[]),
					approverIDs: safeArray(proj.projectDaci?.approver?.map((d) => d?.id) as string[]),
					contributorIDs: safeArray(proj.projectDaci?.contributor?.map((d) => d?.id) as string[]),
					informedIDs: safeArray(proj.projectDaci?.informed?.map((d) => d?.id) as string[])
				};
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectDACI): ' + err,
					dismissible: true,
					type: 'error'
				});
			});
	};

	const updateDACI = async () => {
		errors = {};

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
				load();
			});
	};

	let resourceOpts = $state([] as SelectOptionType<string>[]);
</script>

<SectionHeading>DACI</SectionHeading>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if df}
		<MultiSelectInput
			fieldName="Drivers"
			error={errors.driver}
			options={resourceOpts}
			bind:value={df.driverIDs as string[]}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Approvers"
			error={errors.approver}
			options={resourceOpts}
			bind:value={df.approverIDs as string[]}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Contributors"
			error={errors.contributor}
			options={resourceOpts}
			bind:value={df.contributorIDs as string[]}
			update={() => callIf(update)}
		/>

		<MultiSelectInput
			fieldName="Informed"
			error={errors.informed}
			options={resourceOpts}
			bind:value={df.informedIDs as string[]}
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
