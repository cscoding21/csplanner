<script lang="ts">
	import {
		Avatar,
		Rating,
		Card,
		Alert,
		Table,
		TableBody,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		TableBodyCell,
		Button,
		ButtonGroup
	} from 'flowbite-svelte';
	import { TrashBinOutline, PenOutline } from 'flowbite-svelte-icons';
	import { page } from '$app/stores';
	import {
		ResourceActionBar,
		AddSkill,
		DeleteResource,
		UpdateResourceModal
	} from '../../components';
	import { getResource, deleteResourceSkill } from '$lib/services/resource';
	import { formatDate, formatCurrency, titleCase, pluralize } from '$lib/utils/format';
	import { getInitialsFromName } from '$lib/utils/format';
	import type { Resource, Portfolio } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { findScheduledWorkForResource } from '$lib/services/portfolio';

	const id = $page.params.id;

	const refresh = async ():Promise<Resource> => {
		const res = getResource(id);

		return res;
	};

	const refreshWork = async ():Promise<Portfolio> => {
		return findScheduledWorkForResource(id).then(p => {
			portfolio = p

			return p
		})
	}

	const deleteSkill = async (skillID: string) => {
		deleteResourceSkill(id, skillID).then((res) => {
			if (res.success) {
				addToast({
					message: 'Resource skill deleted successfully',
					dismissible: true,
					type: 'success'
				});

				//resourcePromise = refresh();
			} else {
				addToast({
					message: 'Error updating resource: ' + res.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};

	const updateResource = () => {
		refresh().then((r) => {
			resourcePromise = r as Resource;
		}).then(p => {
			 refreshWork().then(p => {
				portfolio = p
			 })
		});
	}


	let resourcePromise:Resource = $state({} as Resource);
	let portfolio:Portfolio = $state({} as Portfolio);
	const loadPage = async () => {
		updateResource()
	}
</script>

<div class="p-4">
{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	{#if resourcePromise}
		<ResourceActionBar pageDetail={resourcePromise.name}>
			<ButtonGroup>
				<DeleteResource id={resourcePromise.id || ''} name={resourcePromise.name}>
					<TrashBinOutline class="mr-2 h-3 w-3" />
					Delete
				</DeleteResource>
			</ButtonGroup>
		</ResourceActionBar>

		<div class="grid grid-cols-2">
			<div class="mr-4">
				<Card padding="sm" size="xl">
					<div class="flex justify-end">
						<ButtonGroup>
							<UpdateResourceModal {id} update={updateResource}
								><PenOutline /></UpdateResourceModal
							>
						</ButtonGroup>
					</div>
					<div class="flex flex-col items-center pb-4">
						<Avatar size="lg" src={resourcePromise.profileImage as string} rounded
							>{getInitialsFromName(resourcePromise.name)}</Avatar
						>
						<h5 class="mb-1 mt-2 text-xl font-medium text-gray-900 dark:text-white">
							{resourcePromise.name}
						</h5>
						<span class="text-sm text-gray-500 dark:text-gray-400">{resourcePromise.role?.name}</span>
					</div>
					<hr class="mb-4 mt-2" />
					<div>
						<ul class="list">
							<li>
								<span>Type</span>
								<span class="float-right flex-auto font-semibold">{titleCase(resourcePromise.type)}</span>
							</li>
							<li>
								<span>Status</span>
								<span class="float-right flex-auto font-semibold">{titleCase(resourcePromise.status)}</span>
							</li>
							<li>
								<span>Role</span>
								<span class="float-right flex-auto font-semibold">{resourcePromise.role?.name}</span>
							</li>
							<li>
								<span>User</span>
								{#if resourcePromise.user}
									<span class="float-right flex-auto font-semibold">{resourcePromise.user?.email}</span>
								{:else}
									<span class="float-right flex-auto font-semibold">No User Account</span>
								{/if}
							</li>
							<li>
								<span>Initial Cost</span>
								<span class="float-right flex-auto font-semibold">{formatCurrency.format(resourcePromise.initialCost as number)}</span>
							</li>
							<li>
								<span>Annualized Cost</span>
								<span class="float-right flex-auto font-semibold">{formatCurrency.format(resourcePromise.annualizedCost as number)}</span>
							</li>
							<li>
								<span>Created Date</span>
								<span class="float-right flex-auto font-semibold"
									>{formatDate(resourcePromise.createdAt)}</span
								>
							</li>
						</ul>
					</div>

					{#if resourcePromise.type === "human"}
					<caption
						class="mt-4 bg-white text-left text-lg font-semibold text-gray-900 dark:bg-gray-800 dark:text-white"
					>
						User skills
						<p class="mb-4 mt-1 text-sm font-normal text-gray-500 dark:text-gray-400">
							Identifying user skills allows <b>csPlanner</b> to properly estimate and propose project
							resourcing strategies.
						</p>
					</caption>

					{#if resourcePromise.skills && resourcePromise.skills.length > 0}
						<Table>
							<TableHead>
								<TableHeadCell>Skill</TableHeadCell>
								<TableHeadCell>Proficienty</TableHeadCell>
								<TableHeadCell>
									<span class="sr-only">Action</span>
								</TableHeadCell>
							</TableHead>
							<TableBody>
								{#each resourcePromise.skills as s (s)}
									<TableBodyRow>
										<TableBodyCell>{s.name}</TableBodyCell>
										<TableBodyCell
											><Rating rating={s.proficiency?.valueOf()} total={3} /></TableBodyCell
										>
										<TableBodyCell tdClass="float-right pt-2">
											<Button outline color="dark" on:click={() => deleteSkill(s.id)}>
												<TrashBinOutline size="xs" />
											</Button>
										</TableBodyCell>
									</TableBodyRow>
								{/each}
							</TableBody>
						</Table>
					{:else}
						<Alert>No skills have been added for this resource.</Alert>
					{/if}

					<hr class="my-4" />
					<AddSkill resourceID={id} update={updateResource} />
					{/if}
				</Card>
			</div>

			<div>
				<Card size="lg">
					{#if portfolio && portfolio.schedule}
					<ul>
					{#each portfolio.schedule as schedule}
						<li>
							<div class="text-lg">{schedule.project.projectBasics.name}</div>
							{#if schedule.projectActivityWeeks}
							{#each schedule.projectActivityWeeks as week} 
								<div class="text-sm text-red-500">{formatDate(week.end)}</div>
								{#if week.activities}
									<ul class="list-disc">
										{#each week.activities as activity}
											<li class="text-yellow-200 ml-3 text-xs">{activity.taskName} ({activity.hoursSpent} {pluralize("hour", activity.hoursSpent || 0)})</li>
										{/each}
									</ul>
								{/if}
							{/each}
							{/if}
						</li>
					{/each}
				</ul>
					{/if}
				</Card>
			</div>
		</div>
	{/if}
{/await}

</div>
