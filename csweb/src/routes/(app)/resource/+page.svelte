<script lang="ts">
	import { ButtonGroup } from 'flowbite-svelte';
	import { TableHeader } from 'flowbite-svelte-blocks'
	import { UserAddOutline } from 'flowbite-svelte-icons';
	import { findResources } from '$lib/services/resource';
	import { ResourceActionBar, ResourceCard, UpdateResourceModal, ResourceSearchFilters } from './components';
	import { NoResults } from '$lib/components';
	import type { ResourceResults, PageAndFilter, InputFilters } from '$lib/graphql/generated/sdk';

	let filters:InputFilters = $state({}) as InputFilters;
	const refresh = async (): Promise<ResourceResults> => {
		const res = await findResources(getFilters());

		return res;
	};

	const getFilters = ():PageAndFilter => {
		let pageAndFilter: PageAndFilter = {
			paging: { pageNumber: 1, resultsPerPage: 20 },
		};

		pageAndFilter.filters = filters

		return pageAndFilter
	} 

	const filterChange = (f:any) => {
		filters = f

		refresh().then(r => {
			resources = r as ResourceResults;
		});
	}

	let resources = $state({} as ResourceResults);
	const loadPage = async () => {
		refresh().then((r) => {
			resources = r as ResourceResults;
		});
	};
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
	{#if resources.results != null}
		<div class="grid grid-cols-3 gap-3">
			{#each resources.results as r}
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
