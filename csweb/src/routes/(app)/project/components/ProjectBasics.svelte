<script lang="ts">
	import { onMount } from 'svelte';
	import { Button } from 'flowbite-svelte';
	import { resourceStore, refreshResourceStore } from '$lib/stores/resource';
	import { SelectInput, TextInput, TextAreaInput, SectionHeading } from '$lib/components';
	import { createEventDispatcher } from 'svelte';
	import { getProject, updateProjectBasics } from '$lib/services/project';
	import { getDefaultProject, basicSchema } from '$lib/forms/project.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import type { UpdateProjectBasics } from '$lib/graphql/generated/sdk';
	import { coalesceToType } from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';

	const dispatch = createEventDispatcher();

	function updated() {
		dispatch('updated', {
			text: 'Project basics was updated'
		});
	}

	export let id: string;

	let errors: any = {};
	let resourceOpts: any[] = [];

	onMount(async () => {
		await refreshResourceStore();

		load(id);
	});

	const load = async (id: string) => {
		await getProject(id)
			.then((proj) => {
				basicsForm = coalesceToType<UpdateProjectBasics>(proj.projectBasics, basicSchema);
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectBasics): ' + err,
					dismissible: true,
					type: 'error'
				});
			});
	};

	const update = async () => {
		errors = {};

		const projectBasicsParsed = basicSchema.cast(basicsForm);
		basicSchema
			.validate(projectBasicsParsed, { abortEarly: false })
			.then(async () => {
				updateProjectBasics(id, projectBasicsParsed)
					.then((res) => {
						if (res.status?.success) {
							load(id).then(() => {
								addToast({
									message: 'Project basics updated successfully',
									dismissible: true,
									type: 'success'
								});

								updated();
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

	$: resourceOpts = $resourceStore.map((r) => ({ name: r.name, value: r.id }));
	$: basicsForm = getDefaultProject().projectBasics;
	$: project = getDefaultProject();
</script>

<SectionHeading>Basics</SectionHeading>

{#await project}
	Loading...
{:then promiseData}
	{#if basicsForm}
		<TextInput
			bind:value={basicsForm.name}
			fieldName="Project Name"
			placeholder="Project name"
			error={errors.name}
		/>

		<SelectInput
			bind:value={basicsForm.ownerID}
			fieldName="Owner"
			error={errors.ownerID}
			options={resourceOpts}
		/>

		<TextAreaInput
			bind:value={basicsForm.description}
			rows={8}
			error={errors.description}
			placeholder="Executive summary"
			fieldName="Executive Summary"
		/>

		<div class="col-span-4">
			<span class="float-right">
				<Button on:click={update}>Update Basics</Button>
			</span>
			<br class="clear-both" />
		</div>
	{/if}
{/await}
