<script lang="ts">
	import { Card, Heading, Hr, P } from 'flowbite-svelte';
	import type { ProjectEnvelope, Resource } from '$lib/graphql/generated/sdk';
	import { formatDate } from '$lib/utils/format';
	import { ResourceList, MoneyDisplay } from '$lib/components';
	import { safeArray } from '$lib/utils/helpers';
	import { BadgeProjectStatus } from '.';

	interface Props {
		project: ProjectEnvelope;
	}
	let { project }: Props = $props();
</script>

<Card size="xl">
	<Heading tag="h6">
		<a href="project/detail/{project.meta?.id}#snapshot">
			{project.data?.projectBasics?.name}
		</a>
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
	<span class="text-xs text-yellow-200">Start date: <b>{formatDate(project.data?.projectBasics?.startDate)}</b></span>
	{/if}
	<Hr classHr="mt-2 mb-4" />
	<div class="px-4">
		<span>
			<ResourceList
				maxSize={4}
				resources={[
					...safeArray(project.data?.projectDaci?.driver as Resource[]),
					...safeArray(project.data?.projectDaci?.approver as Resource[]),
					...safeArray(project.data?.projectDaci?.contributor as Resource[]),
					...safeArray(project.data?.projectDaci?.informed as Resource[])
				]}
				size="sm"
			/>
		</span>
	</div>
</Card>
