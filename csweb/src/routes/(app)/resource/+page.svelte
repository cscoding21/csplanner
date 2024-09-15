<script lang="ts">
	import { ButtonGroup, Button } from 'flowbite-svelte';
	import { UserAddOutline } from 'flowbite-svelte-icons';
	import { findAllResources } from '$lib/services/resource';
	import { ResourceActionBar, ResourceCard } from './components';
	import { NoResults, PageHeading } from '$lib/components';
	import { onMount } from 'svelte';
	import type { ResourceResults } from '$lib/graphql/generated/sdk';

	const refresh = async (): Promise<ResourceResults> => {
		const res = await findAllResources();
		console.log(res)

		return res;
	};

	onMount(() => {
		resources = refresh();
	});

	let resources = $state(refresh());
</script>

<ResourceActionBar pageDetail="">
	<ButtonGroup>
		<Button href="/resource/new">
			<UserAddOutline class="mr-2 h-3 w-3" />
			Add Resource
		</Button>
	</ButtonGroup>
</ResourceActionBar>

<PageHeading title="Resource Roster" />

{#await resources}
	<div>Loading...</div>
{:then promiseData}
	{#if promiseData.results != null}
		<div class="grid grid-cols-3 gap-3">
			{#each promiseData.results as r}
				<ResourceCard resource={r} on:updated={refresh} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Resources" newUrl="/resource/new">
			No resources...create one now.
		</NoResults>
	{/if}
{/await}
