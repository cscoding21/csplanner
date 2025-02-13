<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { SectionHeading, DateInput } from '$lib/components';
	import { updateProjectBasics } from '$lib/services/project';
	import { basicSchema } from '$lib/forms/project.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import type { Project, UpdateProjectBasics } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import { coalesceToType } from '$lib/forms/helpers';

	interface Props {
		project: Project;
		update?: Function;
	}
	let { project = $bindable(), update }: Props = $props();

    let startDate = $state(project.projectBasics.startDate)

	const updateStartDate = async () => {
		errors = {};

		let bf = coalesceToType<UpdateProjectBasics>(project.projectBasics, basicSchema);
		bf.startDate = new Date(startDate)

		const projectBasicsParsed = basicSchema.cast(bf);
		basicSchema
			.validate(projectBasicsParsed, { abortEarly: false })
			.then(async () => {
				if (!project.id) {
					addToast({
						message: 'Error updating project start date: project id not available',
						dismissible: true,
						type: 'error'
					});

					return
				}

				updateProjectBasics(project.id, projectBasicsParsed)
					.then((res) => {
						if (res.status?.success) {
							addToast({
								message: 'Project start date updated successfully',
								dismissible: true,
								type: 'success'
							});

							callIf(update);
						} else {
							addToast({
								message: 'Error updating project start date: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project start date: ' + err,
							dismissible: true,
							type: 'error'
						});
					});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));
			});
	};

	let errors: any = $state({ name: '', ownerID: '', description: '' });
</script>

	<SectionHeading>
		Set Project Start Date: {project.projectBasics.name}
	</SectionHeading>

<div class="flex">
	<div class="flex-1">
		<DateInput bind:value={startDate} fieldName={""} error="" />
		
		</div>
		<div class="flex-1">
		<span class="float-right">
			<Button onclick={updateStartDate}>Set start date</Button>
		</span>
		<br class="clear-both" />
	</div>
</div>

