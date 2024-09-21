<script lang="ts">
	import { page } from '$app/stores';
	import { ButtonGroup, Button, Card } from 'flowbite-svelte';
	import { getProject } from '$lib/services/project';
	import { ProjectActionBar, ProjectValueChart, DeleteProject } from '../../components';
	import {
		UserList,
		UserDisplay,
		SectionHeading,
		PageHeading,
		MoneyDisplay,
		QuillEditor
	} from '$lib/components';
	import { formatDate, formatCurrency, formatPercent } from '$lib/utils/format';
	import { CommentList } from "$lib/components";
	import { AdjustmentsHorizontalSolid, TrashBinOutline } from 'flowbite-svelte-icons';
	import { type Project } from '$lib/graphql/generated/sdk';

	const id = $page.params.id;

	let project: Project = $state({} as Project);

	const loadPage = async () => {
		getProject(id).then((p) => {
			project = p;
			return p;
		});
	};

	loadPage();
</script>

{#await loadPage}
	<div>Loading...</div>
{:then promiseData}
	<ProjectActionBar pageDetail={project.projectBasics?.name}>
		<ButtonGroup>
			<Button href="/project/workbook/{id}">
				<AdjustmentsHorizontalSolid class="mr-2 h-3 w-3" />
				Workbook
			</Button>
			<DeleteProject id={id || ''} name={project.projectBasics?.name}>
				<TrashBinOutline class="mr-2 h-3 w-3" />
				Delete
			</DeleteProject>
		</ButtonGroup>
	</ProjectActionBar>

	<PageHeading title={project.projectBasics?.name} />

	<div class="grid grid-cols-6 gap-4">
		<div class="col-span-2">
			<Card size="lg">
				<SectionHeading>Financials</SectionHeading>

				<div class="mb-6 p-4">
					<ProjectValueChart {project}></ProjectValueChart>
				</div>
				<ul class="list mb-6">
					<li>
						<span>Net Present Value</span>
						<span class="float-right flex-auto font-semibold">
							<MoneyDisplay amount={project.projectValue?.netPresentValue || 0} />
						</span>
					</li>

					<li>
						<span>Internal Rate of Return</span>
						<span class="float-right flex-auto font-semibold"
							>{formatPercent.format(project.projectValue?.internalRateOfReturn || 0)}</span
						>
					</li>

					<li>
						<span>Funding Source</span>
						<span class="float-right flex-auto font-semibold"
							>{project.projectValue?.fundingSource}</span
						>
					</li>

					<li>
						<span>Development Cost</span>
						<span class="float-right flex-auto font-semibold"
							>{formatCurrency.format(project.projectCost?.initialCost || 0)}</span
						>
					</li>

					<li>
						<span>Development Hours</span>
						<span class="float-right flex-auto font-semibold"
							>{project.projectCost?.hourEstimate}</span
						>
					</li>

					<li>
						<span>Blended Hourly Rate</span>
						<span class="float-right flex-auto font-semibold"
							>{formatCurrency.format(project.projectCost?.blendedRate || 0)}</span
						>
					</li>
				</ul>

				<SectionHeading>Executive Summary</SectionHeading>
				<p class="mb-6 text-xs">{project.projectBasics?.description}</p>

				<SectionHeading>Details</SectionHeading>
				<ul class="list mb-6">
					<li>
						<span>Owner</span>
						<span class="float-right flex-auto font-semibold">
							{#if project.projectBasics?.ownerID}
								<UserDisplay id={project.projectBasics?.ownerID as string} />
							{/if}
						</span>
					</li>

					<li>
						<span>Created</span>
						<span class="float-right flex-auto font-semibold">{formatDate(project.createdAt)}</span>
					</li>

					<li>
						<span>Status</span>
						<span class="float-right flex-auto font-semibold">{project.projectBasics?.status}</span>
					</li>
				</ul>

				<SectionHeading>Team</SectionHeading>

				<ul class="list">
					{#if project.projectDaci?.driver}
						<li>
							<div>Drivers</div>
							<div class="px-4 pb-2 text-left font-semibold">
								<UserList
									size="md"
									maxSize={4}
									resourceIDs={project.projectDaci?.driver as string[]}
								/>
							</div>
						</li>
					{/if}
					{#if project.projectDaci?.approver}
						<li>
							<div>Approvers</div>
							<div class="px-4 pb-2 text-left font-semibold">
								<UserList
									size="md"
									maxSize={4}
									resourceIDs={project.projectDaci?.approver as string[]}
								/>
							</div>
						</li>
					{/if}
					{#if project.projectDaci?.contributor}
						<li>
							<div>Contributors</div>
							<div class="px-4 pb-2 text-left font-semibold">
								<UserList
									size="md"
									maxSize={4}
									resourceIDs={project.projectDaci?.contributor as string[]}
								/>
							</div>
						</li>
					{/if}
					{#if project.projectDaci?.informed}
						<li>
							<div>Informed</div>
							<div class="px-4 pb-2 text-left font-semibold">
								<UserList
									size="md"
									maxSize={4}
									resourceIDs={project.projectDaci?.informed as string[]}
								/>
							</div>
						</li>
					{/if}
				</ul>
			</Card>
		</div>

		<div class="col-span-4">
			<Card size="xl">
				<SectionHeading>Comments / Activity</SectionHeading>
				<!-- <CommentList projectID={id} /> -->
				 <QuillEditor post={() => {}} contents={""} attachContext={id} />
			</Card>
		</div>
	</div>
{/await}
