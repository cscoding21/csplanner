<script lang="ts">
	import { ButtonGroup, Button } from 'flowbite-svelte';
	import { NewspaperOutline } from 'flowbite-svelte-icons';
	import { findProjects } from '$lib/services/project';
	import { ProjectActionBar, ProjectCard } from './components';
	import { NoResults, PageHeading } from '$lib/components';
	import { onMount } from 'svelte';
	import type { PageAndFilter, ProjectResults } from '$lib/graphql/generated/sdk';

	// let pageAndFilter: PageAndFilter = {
	// 	paging: { pageNumber: 1, resultsPerPage: 20 },
	// 	filters: { filters: [{ key: 'basics.status', value: 'approved', operation: 'eq' }] }
	// };

	let pageAndFilter: PageAndFilter = {
		paging: { pageNumber: 1, resultsPerPage: 20 }
	};

	const refresh = async (): Promise<ProjectResults> => {
		const res = await findProjects(pageAndFilter);

		return res;
	};

	onMount(() => {
		projects = refresh();
	});

	let projects = $state(refresh());
</script>

<ProjectActionBar pageDetail="">
	<ButtonGroup>
		<Button href="/project/new">
			<NewspaperOutline class="mr-2 h-3 w-3" />
			Add Project
		</Button>
	</ButtonGroup>
</ProjectActionBar>

<PageHeading title="Projects Portfolio" />

{#await projects}
	<div>Loading...</div>
{:then promiseData}
	{#if promiseData.results != null}
		<div class="grid grid-cols-3 gap-3">
			{#each promiseData.results as p}
				<ProjectCard project={p} on:updated={refresh} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Projects" newUrl="/resource/new">No projects...create one now.</NoResults>
	{/if}
{/await}
