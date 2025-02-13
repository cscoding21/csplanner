<script lang="ts">
	import { Card, Heading, Hr, P } from 'flowbite-svelte';
	import type { Project, Resource } from '$lib/graphql/generated/sdk';
	import { formatDate } from '$lib/utils/format';
	import { ResourceList, MoneyDisplay } from '$lib/components';
	import { safeArray } from '$lib/utils/helpers';
	import { BadgeProjectStatus } from '.';

	interface Props {
		project: Project;
	}
	let { project }: Props = $props();
</script>

<Card size="xl">
	<Heading tag="h6">
		<a href="project/detail/{project.id}#snapshot">
			{project.projectBasics?.name}
		</a>
		<BadgeProjectStatus status={project.projectStatusBlock?.status}  />
		{#if project.projectValue?.calculated?.netPresentValue}
			<span class="float-right text-lg">
				<MoneyDisplay amount={project.projectValue?.calculated?.netPresentValue} />
			</span>
		{/if}
	</Heading>

	<P class="line-clamp-2 py-1 text-sm text-slate-600 dark:text-slate-400"
		>{project.projectBasics?.description}</P
	>
	{#if project.projectBasics.startDate}
	<span class="text-xs text-yellow-200">Start date: <b>{formatDate(project.projectBasics?.startDate)}</b></span>
	{/if}
	<Hr classHr="mt-2 mb-4" />
	<div class="px-4">
		<span>
			<ResourceList
				maxSize={4}
				resources={[
					...safeArray(project.projectDaci?.driver as Resource[]),
					...safeArray(project.projectDaci?.approver as Resource[]),
					...safeArray(project.projectDaci?.contributor as Resource[]),
					...safeArray(project.projectDaci?.informed as Resource[])
				]}
				size="sm"
			/>
		</span>
	</div>
</Card>
