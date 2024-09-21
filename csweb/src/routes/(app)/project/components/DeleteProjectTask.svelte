<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { deleteProjectTask } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import { callIf } from '$lib/utils/helpers';
	import type { Snippet } from 'svelte';

	let popupModal = $state(false);

	interface Props {
		id: string;
		projectID: string;
		milestoneID: string;
		name: string;
		size: 'xs' | 'sm' | 'md' | 'lg' | 'xl' | undefined;
		update?: Function;
		children: Snippet;
	}
	let { id, projectID, milestoneID, name, size, update, children }: Props = $props();

	const delTask = async (taskID: string) => {
		deleteProjectTask(projectID, milestoneID, taskID).then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project feature deleted successfully',
					dismissible: true,
					type: 'success'
				});

				callIf(update);
			} else {
				addToast({
					message: 'Error deleting project: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};
</script>

<Button {size} color="dark" on:click={() => (popupModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={popupModal} size="xs" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Are you sure you want to delete the task, "{name}"?
		</h3>
		<Button color="red" class="mr-2" onclick={() => delTask(id)}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
