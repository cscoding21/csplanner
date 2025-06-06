<script lang="ts">
	import type { ProjectEnvelope } from '$lib/graphql/generated/sdk';
	import { Hr, Tabs, TabItem } from 'flowbite-svelte';
	import {
		ChevronRightOutline,
	} from 'flowbite-svelte-icons';
	import { SectionHeading } from '$lib/components';
	import { findAllProjectTemplates } from '$lib/services/template';
	import { ProjectTaskForm, ProjectTaskDisplay, ProjectTemplateSelector, ProjectMilestoneStatus, BadgeProjectStatus, ShowIfStatus } from '.';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let modalState: boolean[] = $state([]);
	let hasMilestones: boolean = $state(true);
	let project = $state({} as ProjectEnvelope);
	let editTask = $state("")

	const load = async (id: string) => {
		await getProject(id)
			.then((proj) => {
				project = proj;
				hasMilestones = (proj.data?.projectMilestones && proj.data?.projectMilestones.length > 0) as boolean;
			})
			.catch((err) => {
				console.log('Error on get ', err);
				addToast({
					message: 'Error loading project (ProjectMilestones): ' + err,
					dismissible: true,
					type: 'error'
				});
			});
	};

	function refresh() {
		load(id);

		modalState = [];
	}

	const editTaskComplete = () => {
		editTask = ""
		console.log("editTask", editTask)

		callIf(update)

		refresh()
	}

	const loadPage = async () => {
		findAllProjectTemplates()
			.then((r) => r)
			.then((r) => {
				load(id);

				return r;
			})
			.then((r) => r);
	};
</script>

{#await loadPage()}
	Loading...
{:then promiseData}
	{#if project}
		<SectionHeading>
			Implementation Plan & Milestones: {project.data?.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.data?.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}

	{#if project.data?.projectMilestones && project.data?.projectMilestones.length > 0}
		<div class="">
			<div class="">
				<Tabs tabStyle="underline">
					{#each project.data?.projectMilestones as milestone, index}
					<TabItem open={index === 0}>
						{#snippet titleSlot()}
						<div class="flex items-center gap-2">
						  <ChevronRightOutline size="md" />
						  {milestone.phase.name}
						  </div>
						{/snippet}
						<div class="text-sm text-gray-500 dark:text-gray-400">
							<ProjectMilestoneStatus milestone={project.data?.projectMilestones[index]} />
							<Hr />
							{#if milestone.tasks}
								{#each milestone.tasks as task, tindex}
									{#if editTask === task.id} 
									<div class="my-4 rounded-lg bg-white p-4 shadow-sm dark:bg-gray-900 sm:p-6">
									<ProjectTaskForm {task} milestoneID={milestone.id} projectID={id} update={() => editTaskComplete()} />
									</div>
									{:else}
									<ProjectTaskDisplay projectStatus={project.data?.projectStatusBlock.status} projectID={id} milestoneID={milestone.id} editClick={() => editTask = task.id} update={() => editTaskComplete()} task={project.data?.projectMilestones[index].tasks[tindex]} />
									{/if}
								{/each}
							{/if}
							<ShowIfStatus scope={["new", "draft"]} status={project.data?.projectStatusBlock?.status}>
								<SectionHeading>Add New Task to Milestone {milestone.phase.name}</SectionHeading>
								<ProjectTaskForm milestoneID={milestone.id} projectID={id} update={() => refresh()} />
							</ShowIfStatus>
						</div>
					  </TabItem>
					{/each}
					</Tabs>	
			</div>
		</div>	
	{:else}
		<ProjectTemplateSelector {id} />
	{/if}
{/await}
