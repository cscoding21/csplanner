<script lang="ts">
	import type { Project } from '$lib/graphql/generated/sdk';
	import { Hr, Tabs, TabItem } from 'flowbite-svelte';
	import {
		ChevronRightOutline,
	} from 'flowbite-svelte-icons';
	import { SectionHeading } from '$lib/components';
	import { findAllProjectTemplates } from '$lib/services/project';
	import { ProjectTaskForm, ProjectTaskDisplay, ProjectTemplateSelector, ProjectMilestoneStatus, BadgeProjectStatus, ShowIfStatus } from '.';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import { getDefaultProject } from '$lib/forms/project.validation';
	import { callIf } from '$lib/utils/helpers';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let modalState: boolean[] = $state([]);
	let hasMilestones: boolean = $state(true);
	let project = $state(getDefaultProject() as Project);
	let editTask = $state("")

	const load = async (id: string) => {
		await getProject(id)
			.then((proj) => {
				project = proj;
				hasMilestones = (proj.projectMilestones && proj.projectMilestones.length > 0) as boolean;
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
			Implementation Plan & Milestones: {project.projectBasics.name}
			<span class="float-right"><BadgeProjectStatus status={project.projectStatusBlock?.status} /></span>
		</SectionHeading>
	{/if}

	{#if project.projectMilestones && project.projectMilestones.length > 0}
		<div class="">
			<div class="">
				<Tabs tabStyle="underline">
					{#each project.projectMilestones as milestone, index}
					<TabItem open={index === 0}>
						<div slot="title" class="flex items-center gap-2">
						  <ChevronRightOutline size="md" />
						  {milestone.phase.name}
						</div>
						<div class="text-sm text-gray-500 dark:text-gray-400">
							<ProjectMilestoneStatus milestone={project.projectMilestones[index]} />
							<Hr />
							{#if milestone.tasks}
								{#each milestone.tasks as task, tindex}
									{#if editTask === task.id} 
									<div class="my-3">
									<ProjectTaskForm {task} milestoneID={milestone.id} projectID={id} update={() => editTaskComplete()} />
									</div>
										<Hr />
									{:else}
									<ProjectTaskDisplay projectID={id} milestoneID={milestone.id} editClick={() => editTask = task.id} update={() => editTaskComplete()} task={project.projectMilestones[index].tasks[tindex]} />
									{/if}
								{/each}
							{/if}
							<ShowIfStatus scope={["new", "draft"]} status={project.projectStatusBlock?.status}>
								<Hr />
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
