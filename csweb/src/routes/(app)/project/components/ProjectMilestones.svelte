<script lang="ts">
	import type { ProjectMilestone, Skill, Project } from '$lib/graphql/generated/sdk';
	import { P, Button, Hr, Modal } from 'flowbite-svelte';
	import {
		ChevronDoubleRightOutline,
		ChevronRightOutline,
		EditOutline,
		TrashBinOutline
	} from 'flowbite-svelte-icons';
	import { SectionHeading, ResourceList } from '$lib/components';
	import { findAllProjectTemplates } from '$lib/services/project';
	import { ProjectTaskForm, DeleteProjectTask, ProjectTemplateSelector } from '.';
	import { formatToCommaSepList } from '$lib/utils/format';
	import { getProject } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';

	interface Props {
		id: string;
		update?: Function;
	}
	let { id, update }: Props = $props();

	let modalState: boolean[] = $state([]);
	//let modalState:Map<string, boolean> = new Map<string, boolean>()
	let hasMilestones: boolean = $state(true);
	let project = $state({} as Project);
	let currentMilestone = $state({} as ProjectMilestone);

	const load = async (id: string, milestoneID: string) => {
		await getProject(id)
			.then((proj) => {
				project = proj;
				hasMilestones = (proj.projectMilestones && proj.projectMilestones.length > 0) as boolean;

				setCurrentMilestone(milestoneID);
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
		load(id, currentMilestone.id);

		modalState = [];
	}

	const getRequiredSkillsDisplay = (task: any): string => {
		if (task.skills == null || task.skills.length === 0) {
			return 'Not set';
		}

		return formatToCommaSepList(task.skills.map((s: Skill) => s.name));
	};

	const setCurrentMilestone = (milestoneID: any): ProjectMilestone => {
		if (!hasMilestones) {
			//---TODO: handle this
			return {} as ProjectMilestone;
		}

		if (
			project == null ||
			project.projectMilestones == null ||
			project.projectMilestones.length == 0
		) {
			return {} as ProjectMilestone;
		}

		if (!milestoneID) {
			currentMilestone = project.projectMilestones[0] as ProjectMilestone;
		}

		for (let i = 0; i < project.projectMilestones.length; i++) {
			const m = project.projectMilestones[i] as ProjectMilestone;

			if (m.id == milestoneID) {
				currentMilestone = m;
				break;
			}
		}

		return currentMilestone;
	};

	const loadPage = async () => {
		findAllProjectTemplates()
			.then((r) => r)
			.then((r) => {
				load(id, '');

				return r;
			})
			.then((r) => (currentMilestone = setCurrentMilestone('')));
	};

	loadPage();
</script>

<SectionHeading>Implementation Plan & Milestones</SectionHeading>

{#await loadPage}
	Loading...
{:then promiseData}
	{#if project.projectMilestones && project.projectMilestones.length > 0}
		<div class="grid grid-cols-4 gap-8">
			<div class="col-span-1">
				<ol
					class="relative border-s border-gray-200 text-gray-500 dark:border-gray-700 dark:text-gray-400"
				>
					{#each project.projectMilestones as milestone (milestone)}
						<li class="mb-10 ms-6">
							{#if currentMilestone.id === milestone?.id}
								<span
									class="absolute -start-4 flex h-8 w-8 items-center justify-center rounded-full bg-green-200 ring-4 ring-white dark:bg-green-700 dark:text-gray-50 dark:ring-gray-900"
									><ChevronDoubleRightOutline size="xs" /></span
								>
							{:else}
								<span
									class="absolute -start-4 flex h-8 w-8 items-center justify-center rounded-full bg-gray-100 ring-4 ring-white dark:bg-gray-700 dark:ring-gray-900"
									><ChevronRightOutline size="xs" /></span
								>
							{/if}
							<button onclick={(e) => setCurrentMilestone(milestone?.id)}>
								<h3 class="font-medium leading-tight">{milestone?.phase.name}</h3>
							</button>
							<p class="text-xs">{milestone?.phase.description}</p>
						</li>
					{/each}
				</ol>
			</div>
			<div class="col-span-3">
				{#if currentMilestone.tasks}
					{#each currentMilestone.tasks as task (task.id)}
						{@const taskID = task.id}
						<div>
							<span class="text-gray-100">{task.name}</span>
							<P class="mb-3" weight="light" color="text-gray-500 dark:text-gray-100 float-right"
								>{task.hourEstimate} hours</P
							>
						</div>
						<div>
							<div class="text-sm text-gray-400">{task.description}</div>
							<div class="my-2 text-sm text-gray-100">
								Skills Required: {getRequiredSkillsDisplay(task)}
							</div>

							<div class="pl-4">
								<ResourceList maxSize={4} size="xs" resources={task.resources || []} />
							</div>
						</div>
						<div class="mt-2">
							<Button
								size="xs"
								color="dark"
								on:click={() => {
									modalState[taskID] = true;
								}}
							>
								<EditOutline size="xs" class="mr-2" />
								Edit
							</Button>
							<DeleteProjectTask
								name={task.name}
								size="xs"
								projectID={id}
								milestoneID={currentMilestone.id}
								on:updated={refresh}
								id={task.id}
							>
								<TrashBinOutline size="xs" class="mr-2" />
								Delete
							</DeleteProjectTask>
						</div>

						<Modal bind:open={modalState[taskID]} size="xl" autoclose={false}>
							<ProjectTaskForm
								{task}
								milestoneID={currentMilestone.id}
								projectID={id}
								update={refresh}
							/>
						</Modal>
						<br class="clear-both" />
					{/each}
				{/if}
				<Hr />
				<ProjectTaskForm milestoneID={currentMilestone.id} projectID={id} on:updated={refresh} />
			</div>
		</div>
	{:else}
		<ProjectTemplateSelector {id} />
	{/if}
{/await}
