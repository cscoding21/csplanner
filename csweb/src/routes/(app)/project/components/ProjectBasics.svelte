<script lang="ts">
	import { Button, Toggle, type SelectOptionType } from 'flowbite-svelte';
	import { SelectInput, TextInput, TextAreaInput, SectionHeading } from '$lib/components';
	import { getProject, ProjectStatusApproved, ProjectStatusBacklogged, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusNew, ProjectStatusProposed, ProjectStatusScheduled, updateProjectBasics } from '$lib/services/project';
	import { basicSchema, getDefaultProject } from '$lib/forms/project.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import type { UpdateProjectBasics, ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { coalesceToType } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import { findAllUsers } from '$lib/services/user';
	import { BadgeProjectStatus, ShowIfStatus } from '.';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

	const load = async (): Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
				basicsForm = coalesceToType<UpdateProjectBasics>(proj.data?.projectBasics, basicSchema);

				return proj;
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectBasics): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err;
			});
	};

	const updateBasics = async () => {
		errors = {};

		const projectBasicsParsed = basicSchema.cast(basicsForm);
		basicSchema
			.validate(projectBasicsParsed, { abortEarly: false })
			.then(async () => {
				updateProjectBasics(id, projectBasicsParsed)
					.then((res) => {
						if (res.status?.success) {
							load().then(() => {
								addToast({
									message: 'Project basics updated successfully',
									dismissible: true,
									type: 'success'
								});

								callIf(update);
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
			});
	};

	const loadPage = async () => {
		findAllUsers()
			.then((r) => {
				userOpts = r.results?.map((r) => ({
					name: r.firstName + ' ' + r.lastName,
					value: r.email as string
				})) as SelectOptionType<string>[];

				return r;
			})
			.then(() => {
				load().then(p => {
					project = p
				});
			});
	};

	let errors: any = $state({ name: '', ownerID: '', description: '' });
	let userOpts = $state([] as SelectOptionType<string>[]);
	let basicsForm = $state(getDefaultProject().projectBasics);
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project}
		<SectionHeading>
			Basics: {project.data?.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}

	{#if basicsForm}	
		<TextInput
			bind:value={basicsForm.name}
			fieldName="Project Name"
			placeholder="Project name"
			error={errors.name}
		/>

		<SelectInput
			bind:value={basicsForm.ownerID as string}
			fieldName="Owner"
			error={errors.ownerID}
			options={userOpts}
		/>

		<TextAreaInput
			bind:value={basicsForm.description as string}
			rows={8}
			error={errors.description}
			placeholder="Executive summary"
			fieldName="Executive Summary"
		/>

		<div class="pb-4 mb-2">
			<Toggle class="mt-3" bind:checked={basicsForm.isCapitalized}>Capitalized</Toggle>
		</div>

		<ShowIfStatus scope={[ProjectStatusNew, ProjectStatusDraft, ProjectStatusProposed, ProjectStatusApproved, ProjectStatusBacklogged, ProjectStatusScheduled, ProjectStatusInflight]} status={project.data?.projectStatusBlock?.status}>
		<div class="">
			<span class="float-right">
				<Button onclick={updateBasics}>Update Basics</Button>
			</span>
			<br class="clear-both" />
		</div>
		</ShowIfStatus>
	{/if}
{/await}
