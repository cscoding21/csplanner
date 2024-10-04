<script lang="ts">
	import { ButtonGroup } from 'flowbite-svelte';
	import { UserAddOutline } from 'flowbite-svelte-icons';
	import { findAllResources } from '$lib/services/resource';
	import { ResourceActionBar, ResourceCard, UpdateResourceModal } from './components';
	import { NoResults, PageHeading } from '$lib/components';
	import type { ResourceResults } from '$lib/graphql/generated/sdk';

	const refresh = async (): Promise<ResourceResults> => {
		const res = await findAllResources();

		return res;
	};

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
			<UserAddOutline class="mr-2 h-3 w-3" />
			Add Resource
		</UpdateResourceModal>
	</ButtonGroup>
</ResourceActionBar>

<PageHeading title="Resource Roster" />

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
