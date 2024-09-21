<script lang="ts">
	import { Avatar, Popover } from 'flowbite-svelte';
	import { getResource } from '$lib/services/resource';
	import { UserCard } from '..';
	import { normalizeID } from '$lib/utils/id';
	import { deDupeArray } from '$lib/utils/helpers';
	import type { Resource } from '$lib/graphql/generated/sdk';
	import { getInitialsFromName } from '$lib/utils/format';

	interface Props {
		resourceIDs: string[];
		size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | undefined;
		maxSize: number;
	}
	let { resourceIDs, size, maxSize }: Props = $props();

	let resources: Resource[] = $state([]);

	const deDuped = deDupeArray(resourceIDs);
	const loopLength = maxSize > deDuped.length ? deDuped.length : maxSize;
	const more = resourceIDs.length - maxSize;

	const translateResource = async (): Promise<Resource[]> => {
		for (let i = 0; i < loopLength; i++) {
			const r = await getResource(deDuped[i]);

			if (resources.some((rs) => rs.id === r.id)) {
				continue;
			}

			resources.push(r);
		}

		return resources;
	};
</script>

{#await translateResource()}
	<span>Loading...</span>
{:then promiseData}
	<div class="flex">
		{#each promiseData as res (res.id)}
			<Avatar
				id={'pop' + normalizeID(res.id || '')}
				src={res.profileImage || ''}
				stacked
				class="cursor-pointer"
				{size}>{getInitialsFromName(res.name)}</Avatar
			>
			<Popover
				class="w-64 text-left text-sm font-light"
				title={res.name}
				triggeredBy={'#pop' + normalizeID(res.id || '')}
				trigger="click"
			>
				<UserCard resource={res} />
			</Popover>
		{/each}

		{#if more > 0}
			<span class="ml-1 p-1">+{more}</span>
		{/if}
	</div>
{/await}
