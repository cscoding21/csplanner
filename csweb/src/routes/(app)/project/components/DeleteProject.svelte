<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { deleteProject } from '$lib/services/project';
	import { goto } from '$app/navigation';
	import { addToast } from '$lib/stores/toasts';
	import type { Snippet } from 'svelte';

	let popupModal = $state(false);

	interface Props {
		id: string;
		name: string;
		children: Snippet;
	}

	let { id, name, children }: Props = $props();

	const deleteItem = async () => {
		if (!id) {
			return;
		}

		const res = await deleteProject(id);
		if (res.success) {
			addToast({
				message: 'Project deleted successfully',
				dismissible: true,
				type: 'success'
			});

			goto('/project');
		} else {
			addToast({
				message: 'Error deleting project: ' + res.message,
				dismissible: true,
				type: 'error'
			});
		}
	};
</script>

<Button on:click={() => (popupModal = true)}>{@render children()}</Button>

<Modal bind:open={popupModal} size="xs" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Are you sure you want to delete the project, {name}?
		</h3>
		<Button color="red" class="mr-2" on:click={deleteItem}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
