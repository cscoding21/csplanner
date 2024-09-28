<script lang="ts">
	import { Avatar, Popover } from 'flowbite-svelte';
	import { ResourceCard } from '..';
	import { normalizeID } from '$lib/utils/id';
	import type { Resource } from '$lib/graphql/generated/sdk';
	import { getInitialsFromName } from '$lib/utils/format';

	interface Props {
		resources: Resource[];
		size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | undefined;
		maxSize: number;
	}
	let { resources, size, maxSize }: Props = $props();

	const deDuped = resources.filter((value, index) => {
		const _value = JSON.stringify(value);
		return index === resources.findIndex(obj => {
			return JSON.stringify(obj) === _value;
		});
	});

	const loopLength = maxSize > deDuped.length ? deDuped.length : maxSize;
	const more = resources.length - maxSize;

	const translateResource = (): Resource[] => {
		let toShow:Resource[] = [];
		for (let i = 0; i < loopLength; i++) {
			if (deDuped[i] === null) {
				continue
			}

			toShow.push(deDuped[i]);
		}

		return toShow;
	};

	const rss = translateResource();
</script>

{#if rss != null && rss.length > 0 }
<div class="flex">
	{#each rss as res (res.id)}
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
			<ResourceCard resource={res} />
		</Popover>
	{/each}

	{#if more > 0}
		<span class="ml-1 p-1">+{more}</span>
	{/if}
</div>
{:else}
	<span></span>
{/if}

