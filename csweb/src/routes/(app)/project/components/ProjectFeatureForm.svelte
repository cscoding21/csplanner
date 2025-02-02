<script lang="ts">
	import { SectionHeading, TextInput, SelectInput } from '$lib/components';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import { featureSchema, getNewFeature } from '$lib/forms/project.validation';
	import { mergeErrors, parseErrors } from '$lib/forms/helpers';
	import { onMount } from 'svelte';
	import { updateProjectFeature } from '$lib/services/project';
	import { is } from '$lib/utils/check';
	import type { ProjectFeature } from '$lib/graphql/generated/sdk';
	import { TextAreaInput } from '$lib/components';
	import { addToast } from '$lib/stores/toasts';
	import { callIf, deepCopy } from '$lib/utils/helpers';

	onMount(async () => {
		if (feature && feature.id) {
			featureForm = deepCopy(feature);
			featureForm.projectID = projectID;
		} else {
			featureForm = getNewFeature(projectID);
		}
	});

	interface Props {
		id?: string;
		feature: ProjectFeature | undefined;
		projectID: string;
		update?: Function;
	}
	let { 
		id, 
		feature = $bindable(), 
		projectID, 
		update 
	}: Props = $props();

	let errors: any = $state({});

	const priorityOpts: SelectOptionType<string>[] = [
		{ value: 'veryhigh', name: 'Must Have' },
		{ value: 'high', name: 'High' },
		{ value: 'medium', name: 'Medium' },
		{ value: 'low', name: 'Low' }
	];

	const submitForm = () => {
		errors = {};

		const featureFormParsed = featureSchema.cast(featureForm);
		featureSchema
			.validate(featureFormParsed, { abortEarly: false })
			.then(() => {
				updateProjectFeature(featureFormParsed).then((res) => {
					if (res && res.status?.success) {
						addToast({
							message: 'Project features updated successfully',
							dismissible: true,
							type: 'success'
						});

						featureForm = getNewFeature(projectID);
						callIf(update);
					} else {
						addToast({
							message: 'Error updating project features: ' + res.status?.message,
							dismissible: true,
							type: 'error'
						});
					}
				});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.error(errors);
			});
	};

	let featureForm = $state(getNewFeature(projectID));
</script>

<SectionHeading>
	{#if is(feature?.id)}
		Update feature
	{:else}
		Add a New feature
	{/if}
</SectionHeading>
<form>
	<div class="grid grid-cols-4 gap-4">
		<div class="col-span-3">
			<TextInput
				bind:value={featureForm.name}
				placeholder="Feature name"
				fieldName="Feature Name"
				error={errors.name}
			/>
		</div>
		<div class="col-span-1">
			<SelectInput
				fieldName="Priority"
				bind:value={featureForm.priority}
				error={errors.priority}
				options={priorityOpts}
			/>
		</div>
		<div class="col-span-4">
			<TextAreaInput
				bind:value={featureForm.description}
				rows={4}
				fieldName="Feature Description"
				placeholder="Feature description"
				error={errors.description}
			/>
		</div>
		<div class="col-span-4">
			<span class="float-right">
				<Button onclick={submitForm}>
					{#if featureForm.id}
						Update feature
					{:else}
						Add feature
					{/if}
				</Button>
			</span>
		</div>
	</div>
</form>
