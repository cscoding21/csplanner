<script lang="ts">
	import { ButtonGroup, Button } from 'flowbite-svelte';
	import { UserAddOutline } from 'flowbite-svelte-icons';
	import { findAllResources } from '$lib/services/resource';
	import { ResourceActionBar, ResourceCard } from './components';
	import { NoResults } from '$lib/components/messages';
	import { PageHeading } from '$lib/components/formatting';
	import { is } from '$lib/utils/check';
	import { onMount } from 'svelte';
	import type { ApolloQueryResult } from '@apollo/client';

	const refresh = async (): Promise<ApolloQueryResult<any>> => {
		console.log('resource refresh');
		const res = await findAllResources();

		return res;
	};

	onMount(() => {
		resources = refresh();
	});

	let resources = $state(refresh());

	console.log('resource page.svelte');
</script>

<ResourceActionBar pageDetail="">
	<ButtonGroup>
		<Button href="/resource/new">
			<UserAddOutline class="mr-2 h-3 w-3" />
			Add Resource
		</Button>
	</ButtonGroup>
</ResourceActionBar>

<PageHeading title="Resource Roster"><span></span></PageHeading>

{#await resources}
	<div>Loading...</div>
{:then promiseData}
	{#if is(promiseData.data.findAllResources.results)}
		<div class="grid grid-cols-3 gap-3">
			{#each promiseData.data.findAllResources.results as r}
				<ResourceCard resource={r} on:updated={refresh} />
			{/each}
		</div>
	{:else}
		<NoResults title="No Resources" newUrl="/resource/new">
			No resources...create one now.
		</NoResults>
	{/if}
{/await}
