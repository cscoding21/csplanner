<script lang="ts">
	import { Popover } from 'flowbite-svelte';
	import { getResource } from '$lib/services/resource';
	import { UserCard } from '..';
	import { normalizeID } from '$lib/utils/id';
	import type { Resource } from '$lib/graphql/generated/sdk';

	interface Props {
		id: string;
	}
	let { id }: Props = $props();

	let resource: Resource = $state({} as Resource);

	const loadPage = async () => {
		return getResource(id)
			.then((r) => {
				resource = r;
				return resource;
			})
			.then((r) => r)
			.catch((err) => {
				return Promise.reject(err);
			});
	};
</script>

{#await loadPage()}
	NoID
{:then promiseData}
	{#if resource.name}
		<button class="bg-initial" id={normalizeID(resource.id as string)}>{resource.name}</button>

		<Popover
			class="w-64 text-left text-sm font-light"
			title={resource.name}
			triggeredBy="#{normalizeID(resource.id as string)}"
			trigger="click"
		>
			<UserCard {resource} />
		</Popover>
	{/if}
{:catch error}
	Error: {error}
{/await}
