<script lang="ts">
	import { ButtonGroup } from 'flowbite-svelte';
	import { TableHeader } from 'flowbite-svelte-blocks'
	import { UserAddOutline } from 'flowbite-svelte-icons';
	import { findResources } from '$lib/services/resource';
	import { ResourceActionBar, ResourceCard, UpdateResourceModal, ResourceSearchFilters } from './components';
	import { CSPaging, NoResults } from '$lib/components';
	import type { ResourceResults, PageAndFilter, InputFilters, Pagination, Filters } from '$lib/graphql/generated/sdk';

	let filters:InputFilters = $state({}) as InputFilters;
	let paging:Pagination = $state({ pageNumber: 1, resultsPerPage: 10 }) as Pagination;

	const refresh = async (): Promise<ResourceResults> => {
		const res = await findResources(getFilters());

		return res;
	};

	const getFilters = ():PageAndFilter => {
		let pageAndFilter: PageAndFilter = {
			paging: paging,
		};

		pageAndFilter.filters = filters

		return pageAndFilter
	} 

	const filterChange = (f:any) => {
		filters = f

		refresh().then(r => {
			resources = r as ResourceResults;
			pagingDisplay = r.paging as Pagination
			filterDisplay = r.filters
		});
	}

	const pagingChange = (np:Pagination) => {
		paging = np

		refresh().then((r) => {
			resources = r as ResourceResults;
			pagingDisplay = r.paging as Pagination
			filterDisplay = r.filters
		});
	}

	let resources = $state({} as ResourceResults);
	const loadPage = async () => {
		refresh().then((r) => {
			resources = r as ResourceResults;
			pagingDisplay = r.paging as Pagination
			filterDisplay = r.filters
		});
	};

	let pagingDisplay:Pagination = $state({}) as Pagination
	let filterDisplay:Filters = $state({}) as Filters
</script>

<ResourceActionBar pageDetail="">
	<ButtonGroup>
		<UpdateResourceModal update={() => refresh()}>
			<UserAddOutline class="mr-2 h-4 w-4" />
			Add Resource
		</UpdateResourceModal>
	</ButtonGroup>
</ResourceActionBar>

<div class="px-4">
	<TableHeader headerType="search">
		<ResourceSearchFilters change={filterChange} />
	</TableHeader>
</div>

<div class="p-4">
{#await loadPage()}
	<div>Loading...</div>
{:then promiseData}
	{#if resources.results != null && resources.results.length}
		<CSPaging paging={pagingDisplay} change={pagingChange} />

		<div class="grid grid-cols-3 gap-3">
			{#each resources.results as r(r.id)}
				<ResourceCard resource={r} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Resources" newUrl="/resource/new">
			No resources...create one now.
		</NoResults>
	{/if}
{/await}
</div>
