<script lang="ts">
	import { Avatar } from 'flowbite-svelte';
	import { getInitialsFromName, formatDate } from '$lib/utils/format';
	import type { Resource } from '$lib/graphql/generated/sdk';

	interface UserCardProps {
		resource: Resource;
	}
	let { resource }: UserCardProps = $props();
</script>

<div class="p-2">
	<Avatar src={resource.profileImage as string} class="float-left mr-3 h-16 w-16" rounded
		>{getInitialsFromName(resource.name)}</Avatar
	>
	<div class="text-left">
		<div><a href="/resource/detail/{resource.id}" class="font-semibold">{resource.name}</a></div>
		{#if resource.user}
			<div><p>{resource.user?.email}</p></div>
		{/if}
		<div><p>{formatDate(resource.createdAt)}</p></div>
		<div class="clear-both"></div>
	</div>
</div>
