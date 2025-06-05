<script lang="ts">
	import { Avatar, Heading } from 'flowbite-svelte';
	import { getInitialsFromName } from '$lib/utils/format';
	import { ResourceSkillsList } from '$lib/components';
	import type { ResourceEnvelope, Skill } from '$lib/graphql/generated/sdk';
	import { BadgeResourceStatus } from '.'

	interface Props {
		resource: ResourceEnvelope;
	}
	let { resource = $bindable() }: Props = $props();
</script>

<div
	class="items-center rounded-lg bg-gray-50 shadow dark:border-gray-700 dark:bg-gray-800 sm:flex"
>
	<div class="mx-4 my-2 justify-start">
		<a href="/resource/detail/{resource.meta.id}">
			<Avatar size="lg" src={resource.data.profileImage || ''}>{getInitialsFromName(resource.data?.name || '')}</Avatar>
		</a>
	</div>
	<div class="p-5">
		<Heading tag="h6">
			<a href="/resource/detail/{resource.meta.id}">{resource.data?.name}</a>

			<BadgeResourceStatus status={resource.data.status || ''}  />
		</Heading>
		{#if resource.data?.role}
		<span class="text-gray-500 dark:text-gray-400">{resource.data.role.name}</span>
		{/if}
		<p class="mb-3 text-xs font-light text-gray-500 dark:text-gray-400">
			<ResourceSkillsList skills={resource.data.skills as Skill[]} />
		</p>
	</div>
</div>
