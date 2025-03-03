<script lang="ts">
	import {
		Avatar,
		Card,
		Alert,
		Table,
		TableBody,
		TableBodyRow,
		TableHead,
		TableHeadCell,
		TableBodyCell,
		Button,
		ButtonGroup,
		Range
	} from 'flowbite-svelte';
	import { TrashBinOutline, PenOutline } from 'flowbite-svelte-icons';
	import { page } from '$app/state';
	import {
		ResourceActionBar,
		AddSkill,
		DeleteResource,
		UpdateResourceModal
	} from '../../components';
	import { getResource, deleteResourceSkill, decodeProficiency } from '$lib/services/resource';
	import { formatDate, formatCurrency, titleCase } from '$lib/utils/format';
	import { getInitialsFromName } from '$lib/utils/format';
	import type { Resource, Portfolio, Skill } from '$lib/graphql/generated/sdk';
	import { addToast } from '$lib/stores/toasts';
	import { findScheduledWorkForResource, flattenPortfolio, type FlatPortfolioItem } from '$lib/services/portfolio';
	import { SectionSubHeading, CSSection , PieChart, CSHR, PortfolioTable } from '$lib/components';
	import { csGroupBy, deepCopy } from '$lib/utils/helpers';
	import SkillsTable from '../../components/SkillsTable.svelte';

	const id = page.params.id;

	const startDate = new Date()    
    const endDate = new Date(new Date().setDate(new Date().getDate() + 7 * 12));

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

				refresh().then((r) => {
					resourcePromise = r as Resource;
				})
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
			resourceSkills = r.skills as Skill[] 

			return r
		}).then(p => {
			 refreshWork().then(p => {
				portfolio = p

				flatData = flattenPortfolio(p, startDate, endDate)
			 })
		});
	}

	let flatData = $state([] as FlatPortfolioItem[])
	let resourceSkills = $state([] as Skill[])
	let resourcePromise:Resource = $state({} as Resource);
	let projectPieData = $derived(csGroupBy(deepCopy(flatData), "projectName", "actualizedHours"))
	let skillPieData = $derived(csGroupBy(deepCopy(flatData), "requiredSkill", "actualizedHours"))


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
				<UpdateResourceModal {id} update={updateResource}
								><PenOutline class="mr-2 h-3 w-3" /> Edit</UpdateResourceModal
							>
				<DeleteResource id={resourcePromise.id || ''} name={resourcePromise.name}>
					<TrashBinOutline class="mr-2 h-3 w-3" />
					Delete
				</DeleteResource>
			</ButtonGroup>
		</ResourceActionBar>

		<div class="grid grid-cols-3 w-full">
			<div class="mr-4 col-span-1">
				<Card padding="sm" size="xl">
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
					<div class="mb-6">
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
								<span>Hours per Week</span>
								<span class="float-right flex-auto font-semibold">{resourcePromise.availableHoursPerWeek}</span>
							</li>
							<li>
								<span>Onboarding Cost</span>
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

							<CSHR />

							<li>
								<span>Hourly Rate</span>
								<span class="float-right flex-auto font-semibold"
									>{formatCurrency.format(resourcePromise.calculated?.hourlyCost as number)}</span
								>
							</li>
							<li>
								<span>Rate Determination Method</span>
								<span class="float-right flex-auto font-semibold"
									>{resourcePromise.calculated?.hourlyCostMethod}</span
								>
							</li>
						</ul>
					</div>

					{#if resourcePromise.type === "human"}
					<SectionSubHeading>User skills</SectionSubHeading>
					<SkillsTable resourceID={id} skills={resourceSkills} update={() => updateResource()} allowAdd={true} />
					{/if}
				</Card>
			</div>

			<div class="col-span-2">
				<CSSection>
					<SectionSubHeading>Project Allocation: {formatDate(startDate)} - {formatDate(endDate)}</SectionSubHeading>

					{#if portfolio && portfolio.schedule}
					<div class="mb-2">
						<PortfolioTable {portfolio} {startDate} {endDate}></PortfolioTable>
					</div>
					{/if}

				<div class="mt-8 grid grid-cols-2">
					<div class="p-4">
						<SectionSubHeading>Project Distribution</SectionSubHeading>
						{#if projectPieData}
							<PieChart labels={Object.keys(projectPieData)} values={Object.values(projectPieData)} />
						{/if}
					</div>


					<div class="p-4">
						<SectionSubHeading>Skills Distribution</SectionSubHeading>
						{#if projectPieData}
							<PieChart labels={Object.keys(skillPieData)} values={Object.values(skillPieData)} />
						{/if}
					</div>
				</div>
				</CSSection>
			</div>
		</div>
	{/if}
{/await}

</div>
