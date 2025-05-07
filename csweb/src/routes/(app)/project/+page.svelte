<script lang="ts">
	import { ButtonGroup, Button, CardPlaceholder } from 'flowbite-svelte';
	import { NewspaperOutline } from 'flowbite-svelte-icons';
	import { findProjects } from '$lib/services/project';
	import { ProjectActionBar, ProjectCard, ProjectSearchFilters } from './components';
	import { NoResults, CSPaging, CSSection } from '$lib/components';
	import type { PageAndFilter, ProjectResults, InputFilters, Pagination, Filters } from '$lib/graphql/generated/sdk';

	let filters:InputFilters = $state({}) as InputFilters;
	let paging:Pagination = $state({ pageNumber: 1, resultsPerPage: 10 }) as Pagination;

	const getFilters = ():PageAndFilter => {
		let pageAndFilter: PageAndFilter = {
			paging: paging,
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
			pagingDisplay = p.paging as Pagination
			filterDisplay = p.filters
		});
	}

	const pagingChange = (np:Pagination) => {
		paging = np

		refresh().then((p) => {
			projects = p as ProjectResults;
			pagingDisplay = p.paging as Pagination
			filterDisplay = p.filters
		});
	}

	let projects = $state({} as ProjectResults);
	const loadPage = async () => {
		refresh().then((p) => {
			projects = p as ProjectResults;
			pagingDisplay = p.paging as Pagination
			filterDisplay = p.filters
		});
	};

	let pagingDisplay:Pagination = $state({}) as Pagination
	let filterDisplay:Filters = $state({}) as Filters
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
<CSSection>
	<ProjectSearchFilters change={filterChange} />
</CSSection>
</div>

<div class="px-4 mt-2">
{#await loadPage()}
	<CardPlaceholder />
	<CardPlaceholder />
	<CardPlaceholder />
{:then promiseData}
	{#if projects.results != null && projects.results.length > 0}
		<CSPaging paging={pagingDisplay} change={pagingChange} />

		<div class="grid grid-cols-3 gap-3">
			{#each projects.results as p}
				<ProjectCard project={p} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Projects" newUrl="/project/new">No projects...create one now.</NoResults>
	{/if}
{/await}
</div>
