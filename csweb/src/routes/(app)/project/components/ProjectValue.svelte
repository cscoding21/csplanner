<script lang="ts">
	import { Heading, Button, type SelectOptionType } from 'flowbite-svelte';
	import { MoneyInput, PercentInput, SectionHeading, SelectInput } from '$lib/components';
	import { getDefaultProject, valueSchema } from '$lib/forms/project.validation';
	import { getProject, updateProjectValue } from '$lib/services/project';
	import {
		parseErrors,
		mergeErrors,
		coalesceToType,
		findSelectOptsFromList
	} from '$lib/forms/helpers';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import type { UpdateProjectValue, Project } from '$lib/graphql/generated/sdk';
	import { getList } from '$lib/services/list';
	import { BadgeProjectStatus } from '.';

	let errors: any = $state({});

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let project:Project = $state(getDefaultProject() as Project)

	const load = async ():Promise<Project> => {
		return await getProject(id)
			.then((proj) => {
				valueForm = coalesceToType<UpdateProjectValue>(proj.projectValue, valueSchema);

				return proj
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project (ProjectValue): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

	const updateValue = async () => {
		errors = {};

		const projectValueParsed = valueSchema.cast(valueForm);
		valueSchema
			.validate(projectValueParsed, { abortEarly: false })
			.then(async () => {
				await updateProjectValue(id, projectValueParsed)
					.then((res) => {
						if (res.status?.success) {
							console.log('update success ', res.status);
							load().then((p) => {
								addToast({
									message: 'Project value updated successfully',
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

				console.error(errors);
			});
	};

	const loadPage = async () => {
		getList('Funding Source')
			.then((l) => {
				fundingSourceOpts = findSelectOptsFromList(l);
			})
			.then(() => {
				load().then(p => {
					project = p
				});
			});
	};

	let valueForm = $state(getDefaultProject().projectValue);
	let fundingSourceOpts = $state([] as SelectOptionType<string>[]);
</script>


{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project}
	<SectionHeading>
		Value Proposition: {project.projectBasics.name}
		<span class="float-right"><BadgeProjectStatus status={project.projectStatusBlock?.status} /></span>
	</SectionHeading>
	{/if}

	{#if valueForm}
		<SelectInput
			bind:value={valueForm.fundingSource as string}
			fieldName="Funding Source"
			error={errors.fundingSource}
			options={fundingSourceOpts}
			update={() => callIf(update)}
		/>

		<PercentInput
			bind:value={valueForm.discountRate as number}
			fieldName="Discount Rate"
			error={errors.discountRate}
			update={() => callIf(update)}
		/>

		<Heading tag="h6">Five Year Forecast</Heading>

		<MoneyInput
			bind:value={valueForm.yearOneValue as number}
			fieldName="Estimated Year One Returns"
			error={errors.yearOneValue}
			update={() => callIf(update)}
		/>

		<MoneyInput
			bind:value={valueForm.yearTwoValue as number}
			fieldName="Estimated Year Two Returns"
			error={errors.yearTwoValue}
			update={() => callIf(update)}
		/>

		<MoneyInput
			bind:value={valueForm.yearThreeValue as number}
			fieldName="Estimated Year Tnree Returns"
			error={errors.yearThreeValue}
			update={() => callIf(update)}
		/>

		<MoneyInput
			bind:value={valueForm.yearFourValue as number}
			fieldName="Estimated Year Four Returns"
			error={errors.yearFourValue}
			update={() => callIf(update)}
		/>

		<MoneyInput
			bind:value={valueForm.yearFiveValue as number}
			fieldName="Estimated Year Five Returns"
			error={errors.yearFiveValue}
			update={() => callIf(update)}
		/>

		<div class="col-span-4">
			<span class="float-right">
				<Button on:click={updateValue}>Update Value</Button>
			</span>
			<br class="clear-both" />
		</div>
	{/if}
{/await}
