<script lang="ts">
	import { Table, TableBody, TableBodyRow, TableBodyCell, TableHead, TableHeadCell, ButtonGroup, Hr  } from 'flowbite-svelte';
	import { CheckCircleSolid, CircleMinusOutline, EditOutline, TrashBinOutline } from 'flowbite-svelte-icons';
	import { SectionHeading, SectionSubHeading, MoneyDisplay } from '$lib/components';
	import { getDefaultProject } from '$lib/forms/project.validation';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import type { ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { BadgeProjectStatus, DeleteProjectValueLine, ProjectValueChart, ProjectValueCategoryDistributionChart } from '.';
	import { formatCurrency, formatPercent } from '$lib/utils/format';
	import { ProjectValueLineFormModal } from '.';
	import ProjectValueLineForm from './ProjectValueLineForm.svelte';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let project:ProjectEnvelope = $state(getDefaultProject() as ProjectEnvelope)

	const load = async ():Promise<ProjectEnvelope> => {
		return await getProject(id)
			.then((proj) => {
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

	function refresh() {
		load().then(p => {
			project = p

			callIf(update);
		})
	}

	const loadPage = async () => {
		load().then(p => {
			project = p
		});
	};
</script>


{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project}
	<SectionHeading>
		Value Proposition: {project.data?.projectBasics.name}
		<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
	</SectionHeading>
	

	{#if project.data?.projectValue.projectValueLines && project.data?.projectValue.projectValueLines.length > 0}
	<div class="flex mb-8">
		<div class="flex-1 px-4">
			<ul class="list mb-6 col-span-2 p-2 text-sm space-y-3">
				<li>
					<span>Five Year Gross</span>
					<span class="float-right flex-auto font-semibold">
						<MoneyDisplay amount={project.data?.projectValue?.calculated?.fiveYearGross || 0} />
					</span>
				</li>


				<li>
					<span>Net Present Value</span>
					<span class="float-right flex-auto font-semibold">
						<MoneyDisplay amount={project.data?.projectValue?.calculated?.netPresentValue || 0} />
					</span>
				</li>
		
				<li>
					<span>Internal Rate of Return</span>
					<span class="float-right flex-auto font-semibold"
						>{formatPercent.format(project.data?.projectValue?.calculated?.internalRateOfReturn || 0)}</span
					>
				</li>

				<li>
					<span>Is Capitalized</span>
					{#if project.data?.projectBasics.isCapitalized}
					<span class="float-right flex-auto font-semibold text-green-400"><CheckCircleSolid /></span>	
					{:else}
					<span class="float-right flex-auto font-semibold text-gray-400"><CircleMinusOutline /></span>	
					{/if}
				</li>
			</ul>
		</div>

		<div class="flex-1 px-2">
			<SectionSubHeading>Five Year Outlook</SectionSubHeading>
			<ProjectValueChart project={project.data}></ProjectValueChart>
		</div>

		<div class="flex-1 px-2">
			<SectionSubHeading>Category Distribution</SectionSubHeading>
			<ProjectValueCategoryDistributionChart {project}></ProjectValueCategoryDistributionChart>
		</div>
		
	</div>

	<SectionSubHeading>
		Revenue Items

		<span class="float-right">
			<ProjectValueLineFormModal
				valueItem={undefined}
				projectID={id}
				size="md"
				update={() => refresh()}>
				Add Value Line
			</ProjectValueLineFormModal>
		</span>
	</SectionSubHeading>
	<Table hoverable={true}>
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
			{#each project.data?.projectValue.projectValueLines as line}
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
					<ProjectValueLineFormModal
						valueItem={line}
						projectID={id}
						size="md"
						update={() => refresh()}>
						<EditOutline size="sm" />
					</ProjectValueLineFormModal>
				<DeleteProjectValueLine
					id={line.id || ''}
					name={0 + (line.yearOneValue || 0) + (line.yearTwoValue || 0) + (line.yearThreeValue || 0) + (line.yearFourValue || 0) + (line.yearFiveValue || 0)}
					projectID={id}
					size="md"
					update={() => refresh()}
				>
					<TrashBinOutline size="sm" class="" />
				</DeleteProjectValueLine>
				</ButtonGroup>
			</TableBodyCell>
		  </TableBodyRow>
		  {/each}
		</TableBody>
		<tfoot>
			<tr class="font-semibold text-gray-900 dark:text-white">
			  <th scope="row" class="py-3 px-6 text-base">Totals</th>
			  <th scope="row" class="py-3 px-6 text-base">&nbsp;</th>
			  <td class="py-3 px-6">{formatCurrency.format(project.data?.projectValue.calculated?.yearOneValue || 0)}</td>
			  <td class="py-3 px-6">{formatCurrency.format(project.data?.projectValue.calculated?.yearTwoValue || 0)}</td>
			  <td class="py-3 px-6">{formatCurrency.format(project.data?.projectValue.calculated?.yearThreeValue || 0)}</td>
			  <td class="py-3 px-6">{formatCurrency.format(project.data?.projectValue.calculated?.yearFourValue || 0)}</td>
			  <td class="py-3 px-6">{formatCurrency.format(project.data?.projectValue.calculated?.yearFiveValue || 0)}</td>
			  <th scope="row" class="py-3 px-6 text-base">&nbsp;</th>
			</tr>
		  </tfoot>
	  </Table>

	  {:else}
	  	<ProjectValueLineForm projectID={id} valueItem={undefined} update={refresh} />
	  {/if}
	  {/if}
{/await}



