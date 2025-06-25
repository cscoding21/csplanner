<script lang="ts">
	import { Heading, P } from 'flowbite-svelte';
	import type { ProjectEnvelope, Resource } from '$lib/graphql/generated/sdk';
	import { formatDate } from '$lib/utils/format';
	import { ResourceList, MoneyDisplay, CSHR } from '$lib/components';
	import { safeArray } from '$lib/utils/helpers';
	import { BadgeProjectStatus } from '.';

	interface Props {
		project: ProjectEnvelope;
	}
	let { project }: Props = $props();
</script>

<a href="project/detail/{project.meta?.id}#snapshot" class="p-4 m-1 rounded-lg dark:bg-gray-800 bg-gray-100">
	<Heading tag="h6">
			{project.data?.projectBasics?.name}
		<BadgeProjectStatus status={project.data?.projectStatusBlock?.status}  />
		{#if project.data?.projectValue?.calculated?.netPresentValue}
			<span class="float-right text-lg">
				<MoneyDisplay amount={project.data?.projectValue?.calculated?.netPresentValue} />
			</span>
		{/if}
	</Heading>

	<P class="line-clamp-2 py-1 text-sm text-slate-600 dark:text-slate-400"
		>{project.data?.projectBasics?.description}</P
	>
	{#if project.data?.projectBasics.startDate}
	<span class="text-xs text-yellow-200">Kickoff date: <b>{formatDate(project.data?.projectBasics?.startDate)}</b></span>
	{/if}
	 <CSHR />
	<div class="px-4">
		<span>
			<ResourceList
				maxSize={4}
				resources={safeArray(project.data.calculated?.teamMembers as Resource[])}
				size="sm"
			/>
		</span>
	</div>
</a>
