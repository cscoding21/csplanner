<script lang="ts">
	import {
		SectionHeading,
		TextInput,
		SelectInput,
		NumberInput,
		MultiSelectInput,
		TextAreaInput
	} from '$lib/components';
	import { Button, type SelectOptionType } from 'flowbite-svelte';
	import type {
		UpdateProjectMilestoneTask,
		ProjectMilestoneTask,
		Resource
	} from '$lib/graphql/generated/sdk';
	import { taskSchema } from '$lib/forms/project.validation';
	import {
		coalesceToType,
		mergeErrors,
		parseErrors,
		findSelectOptsFromList
	} from '$lib/forms/helpers';
	import { updateProjectTask } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import { findAllResources } from '$lib/services/resource';
	import { getList } from '$lib/services/list';

	let errors: any = $state({});

	interface Props {
		milestoneID: string;
		projectID: string;
		task?: ProjectMilestoneTask;
		update?: Function;
	}
	let { milestoneID, projectID, task, update }: Props = $props();

	const statusOpts = [
		{ value: 'new', name: 'New' },
		{ value: 'accepted', name: 'Accepted' },
		{ value: 'removed', name: 'Removed' },
		{ value: 'done', name: 'Done' }
	];

	const resetNewTask = () => {
		(taskForm.projectID = projectID), (taskForm.milestoneID = milestoneID), (taskForm.id = '');
		taskForm.name = '';
		taskForm.hourEstimate = 0;
		taskForm.description = '';
		taskForm.status = 'new';
		taskForm.requiredSkillIDs = [];
		taskForm.resourceIDs = [];
		taskForm.startDate = null;
		taskForm.endDate = null;
	};

	const getTaskForm = (): UpdateProjectMilestoneTask => {
		if (task && task.id) {
			let outTask = coalesceToType<UpdateProjectMilestoneTask>(task, taskSchema);
			outTask.milestoneID = milestoneID;
			outTask.projectID = projectID;

			return outTask;
		}

		return {
			projectID: projectID,
			milestoneID: milestoneID,
			id: '',
			name: '',
			description: '',
			hourEstimate: 0,
			status: '',
			requiredSkillIDs: [],
			resourceIDs: [],
			startDate: null,
			endDate: null
		};
	};

	const submitForm = () => {
		errors = {};

		const taskFormParsed = taskSchema.cast(taskForm);
		taskSchema
			.validate(taskFormParsed, { abortEarly: false })
			.then(() => {
				updateProjectTask(taskFormParsed).then((res) => {
					if (res.status?.success) {
						addToast({
							message: 'Project milestone tasks updated successfully',
							dismissible: true,
							type: 'success'
						});

						resetNewTask();
						console.log(update)
						callIf(update);
						console.log("end of updateProjectTask")
					} else {
						addToast({
							message: 'Error updating project milestone tasks: ' + res.status?.message,
							dismissible: true,
							type: 'error'
						});
					}
				});
			})
			.catch((err) => {
				errors = mergeErrors(errors, parseErrors(err));

				console.log(errors);
			});
	};

	const getElibigleResources = (skillsList: string[]): any[] => {
		//---TODO: figure out the reactivity and required logic to male
		//   list limited to resoures that have skills required by the task
		if (skillsList && skillsList.length > 0) {
			return resourceRoster
				.filter((r) => {
					if (r.skills && r.skills.length > 0) {
						const val = r.skills?.map((rs) => rs.id).some((so) => skillsList.includes(so));
						console.log('val: ', val);
						return val;
					} else {
						return false;
					}
				})
				.map((r) => ({ name: r.name, value: r.id }));
		}

		return resourceRoster.map((r) => ({ name: r.name, value: r.id }));
	};

	const loadPage = async () => {
		getList('Skills')
			.then((skillList: any) => {
				skillsOpts = findSelectOptsFromList(skillList);
			})
			.then(() => {
				findAllResources()
					.then((r) => {
						resourceRoster = r.results as Resource[];

						return r;
					})
					.then((ro) => {
						resourceOpts = ro.results?.map((t: any) => ({
							name: t.name,
							value: t.id as string
						})) as SelectOptionType<string>[];
					});
			});
	};

	let taskForm = $state(getTaskForm());
	let resourceRoster = [] as Resource[];
	let resourceOpts = $state(getElibigleResources(taskForm.requiredSkillIDs as string[]));
	let skillsOpts = $state([] as SelectOptionType<string>[]);

	loadPage();
</script>

{#if resourceOpts && skillsOpts && taskForm}
	<form>
		<div class="grid grid-cols-4 gap-4">
			<div class="col-span-2">
				<TextInput bind:value={taskForm.name} fieldName="Task Name" error={errors.name} />
			</div>
			<div class="col-span-1">
				<NumberInput
					bind:value={taskForm.hourEstimate}
					fieldName="Estimated Hours"
					min={0}
					error={errors.hourEstimate}
				/>
			</div>
			<div class="col-span-1">
				<SelectInput
					fieldName="Status"
					bind:value={taskForm.status}
					error={errors.status}
					options={statusOpts}
				/>
			</div>
		</div>

		<div>
			<TextAreaInput
				bind:value={taskForm.description as string}
				rows={2}
				fieldName="Task Description"
				error={errors.description}
			/>
		</div>

		<div class="grid grid-cols-4 gap-4">
			<div class="col-span-2">
				<MultiSelectInput
					bind:value={taskForm.requiredSkillIDs as string[]}
					fieldName="Required Skills"
					options={skillsOpts}
					error={errors.requiredSkills}
				/>
			</div>
			<div class="col-span-2">
				<MultiSelectInput
					bind:value={taskForm.resourceIDs as string[]}
					fieldName="Resources"
					options={resourceOpts}
					error={errors.resourceIDs}
				/>
			</div>
		</div>

		<div class="col-span-4">
			<span class="float-right">
				{#if taskForm && taskForm.id}
					<Button on:click={submitForm}>Edit Task</Button>
				{:else}
					<Button on:click={submitForm}>Add Task</Button>
				{/if}
			</span>
			<br class="clear-both" />
		</div>
	</form>
{/if}
