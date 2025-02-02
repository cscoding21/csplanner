<script lang="ts">
	import { Button, Table, TableBody, TableBodyRow, TableBodyCell, TableHead, TableHeadCell, ButtonGroup, Toggle  } from 'flowbite-svelte';
	import { EditOutline, TrashBinOutline } from 'flowbite-svelte-icons';
	import type { SelectOptionType } from 'flowbite-svelte';
	import { MoneyInput, PercentInput, SectionHeading, SelectInput, SectionSubHeading } from '$lib/components';
	import { getDefaultProject, valueSchema, valueDefaultForm } from '$lib/forms/project.validation';
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
	import { BadgeProjectStatus, DeleteProjectValueLine, ProjectValueChart, ProjectValueCategoryDistributionChart } from '.';
	import { formatCurrency } from '$lib/utils/format';

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
								message: 'Error updating project value: ' + res.status?.message,
								dismissible: true,
								type: 'error'
							});
						}
					})
					.catch((err) => {
						addToast({
							message: 'Error updating project value: ' + err,
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
		load().then(p => {
			project = p
		});
	};

	let valueForm = $state(valueDefaultForm());
	let isCapitalized = $state(false)
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

	{#if project.projectValue.projectValueLines}
	<div class="flex mb-8">
		<div class="flex-1 px-2">
			{#if valueForm}
			<SectionSubHeading>Parameters</SectionSubHeading>
		<PercentInput
			bind:value={valueForm.discountRate}
			fieldName="Discount rate percentage"
			error={errors.discountRate}
			update={() => callIf(update)}
		/>

		<div class="pb-4 mb-2">
			<Toggle class="mt-3" bind:checked={valueForm.isCapitalized}>Capitalized</Toggle>
		</div>

		<Button
			size="xs"
			onclick={() => updateValue()}>
			Update
		</Button>
		{/if}

		</div>

		<div class="flex-1 px-2">
			<SectionSubHeading>Five Year Outlook</SectionSubHeading>
			<ProjectValueChart {project}></ProjectValueChart>
		</div>

		<div class="flex-1 px-2">
			<SectionSubHeading>Category Distribution</SectionSubHeading>
			<ProjectValueCategoryDistributionChart {project}></ProjectValueCategoryDistributionChart>
		</div>
		
	</div>

	<SectionSubHeading>
		Revenue Items
	</SectionSubHeading>
	<Table>
		<TableHead>
		  <TableHeadCell>Funding Source</TableHeadCell>
		  <TableHeadCell>Category</TableHeadCell>
		  <TableHeadCell>Year One</TableHeadCell>
		  <TableHeadCell>Year Two</TableHeadCell>
		  <TableHeadCell>Year Three</TableHeadCell>
		  <TableHeadCell>Year Four</TableHeadCell>
		  <TableHeadCell>Year Two</TableHeadCell>
		  <TableHeadCell>Actions</TableHeadCell>
		</TableHead>
		<TableBody tableBodyClass="divide-y">
			{#each project.projectValue.projectValueLines as line}
		  <TableBodyRow>
			<TableBodyCell>{line.fundingSource}</TableBodyCell>
			<TableBodyCell>{line.valueCategory}</TableBodyCell>
			<TableBodyCell>{formatCurrency.format(line.yearOneValue || 0)}</TableBodyCell>
			<TableBodyCell>{formatCurrency.format(line.yearTwoValue || 0)}</TableBodyCell>
			<TableBodyCell>{formatCurrency.format(line.yearThreeValue || 0)}</TableBodyCell>
			<TableBodyCell>{formatCurrency.format(line.yearFourValue || 0)}</TableBodyCell>
			<TableBodyCell>{formatCurrency.format(line.yearFiveValue || 0)}</TableBodyCell>
			<TableBodyCell>
				<ButtonGroup>
				<Button
					color="dark"
					class=""
					onclick={() => {
						console.log('open a modal');
					}}
				>
					<EditOutline size="sm" />
				</Button>
				<DeleteProjectValueLine
					id={line.id || ''}
					name={0 + (line.yearOneValue || 0) + (line.yearTwoValue || 0) + (line.yearThreeValue || 0) + (line.yearFourValue || 0) + (line.yearFiveValue || 0)}
					projectID={id}
					size="md"
					update={() => callIf(update)}
				>
					<TrashBinOutline size="sm" class="" />
				</DeleteProjectValueLine>
				</ButtonGroup>
			</TableBodyCell>
		  </TableBodyRow>
		  {/each}
		</TableBody>
	  </Table>
	  {:else}
	  	No Value here
	  {/if}
{/await}
