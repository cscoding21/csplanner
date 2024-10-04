<script lang="ts">
	import { ButtonGroup, Button, CardPlaceholder } from 'flowbite-svelte';
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

<div class="p-4">
{#await projects}
	<CardPlaceholder />
	<CardPlaceholder />
	<CardPlaceholder />
{:then promiseData}
	{#if promiseData.results != null}
		<div class="grid grid-cols-3 gap-3">
			{#each promiseData.results as p}
				<ProjectCard project={p} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Projects" newUrl="/resource/new">No projects...create one now.</NoResults>
	{/if}
{/await}
</div>
