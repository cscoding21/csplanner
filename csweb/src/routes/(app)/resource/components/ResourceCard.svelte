<script lang="ts">
	import { Avatar, Heading } from 'flowbite-svelte';
	import { formatToCommaSepList, getInitialsFromName } from '$lib/utils/format';
	import { ResourceSkillsList } from '$lib/components';
	import type { Resource, Skill } from '$lib/graphql/generated/sdk';

	interface Props {
		resource: Resource;
	}
	let { resource = $bindable() }: Props = $props();
</script>

<div
	class="items-center rounded-lg bg-gray-50 shadow dark:border-gray-700 dark:bg-gray-800 sm:flex"
>
	<div class="mx-4 my-2 justify-start">
		<a href="/resource/detail/{resource.id}">
			<Avatar size="lg" src={resource.profileImage || ''} rounded
				>{getInitialsFromName(resource.name)}</Avatar
			>
		</a>
	</div>
	<div class="p-5">
		<Heading tag="h6">
			<a href="/resource/detail/{resource.id}">{resource.name}</a>
		</Heading>
		<span class="text-gray-500 dark:text-gray-400">{resource.role}</span>
		<p class="mb-3 text-xs font-light text-gray-500 dark:text-gray-400">
			<ResourceSkillsList skills={resource.skills as Skill[]} />
		</p>
	</div>
</div>
