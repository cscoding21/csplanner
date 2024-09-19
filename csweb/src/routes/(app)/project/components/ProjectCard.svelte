<script lang="ts">
	import { Card, Heading, Hr, P, Badge } from 'flowbite-svelte';
	import type { Project } from '$lib/graphql/generated/sdk';
	import { formatDate } from '$lib/utils/format';
	import { UserList, MoneyDisplay } from '$lib/components';

	export let project: Project;
</script>

<Card size="xl">
	<Heading tag="h6">
		<a href="project/workbook/{project.id}">
			{project.projectBasics?.name}
		</a>
		<Badge color="blue">{project.projectBasics?.status}</Badge>
		{#if project.projectValue?.netPresentValue}
			<span class="float-right text-lg">
				<MoneyDisplay amount={project.projectValue?.netPresentValue} />
			</span>
		{/if}
	</Heading>

	<P class="line-clamp-2 py-1 text-sm text-slate-600 dark:text-slate-400"
		>{project.projectBasics?.description}</P
	>
	<span class="float-right">{formatDate(project.projectBasics?.startDate)}</span>
	<Hr classHr="mt-2 mb-4" />
	<div class="px-4">
		<span>
			<UserList
				maxSize={4}
				resourceIDs={[
					...(project.projectDaci?.driver as string[]),
					...(project.projectDaci?.approver as string[]),
					...(project.projectDaci?.contributor as string[]),
					...(project.projectDaci?.informed as string[])
				]}
				size="sm"
			/>
		</span>
	</div>
</Card>
