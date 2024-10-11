<script lang="ts">
	import { ButtonGroup, Button, CardPlaceholder } from 'flowbite-svelte';
	import { TableHeader } from 'flowbite-svelte-blocks'
	import { NewspaperOutline } from 'flowbite-svelte-icons';
	import { findProjects } from '$lib/services/project';
	import { ProjectActionBar, ProjectCard, ProjectSearchFilters } from './components';
	import { NoResults } from '$lib/components';
	import type { PageAndFilter, ProjectResults, InputFilters } from '$lib/graphql/generated/sdk';

	let filters:InputFilters = $state({}) as InputFilters;

	const getFilters = ():PageAndFilter => {
		let pageAndFilter: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 20 },
		};

		pageAndFilter.filters = filters

		return pageAndFilter
	} 

	const refresh = async (): Promise<ProjectResults> => {
		const res = await findProjects(getFilters());

		return res;
	};

	const filterChange = (f:any) => {
		filters = f

		refresh().then((p) => {
			projects = p as ProjectResults;
		});
	}

	let projects = $state({} as ProjectResults);
	const loadPage = async () => {
		refresh().then((p) => {
			projects = p as ProjectResults;
		});
	};
</script>

<ProjectActionBar pageDetail="">
	<ButtonGroup>
		<Button href="/project/new">
			<NewspaperOutline class="mr-2 h-4 w-4" />
			Add Project
		</Button>
	</ButtonGroup>
</ProjectActionBar>

<div class="px-4">
<TableHeader headerType="search">
	<ProjectSearchFilters change={filterChange} />
</TableHeader>
</div>

<div class="p-4">
{#await loadPage()}
	<CardPlaceholder />
	<CardPlaceholder />
	<CardPlaceholder />
{:then promiseData}
	{#if projects.results != null}
		<div class="grid grid-cols-3 gap-3">
			{#each projects.results as p}
				<ProjectCard project={p} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Projects" newUrl="/resource/new">No projects...create one now.</NoResults>
	{/if}
{/await}
</div>
