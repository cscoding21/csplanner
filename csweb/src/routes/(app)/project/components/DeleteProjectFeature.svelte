<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { deleteProjectFeature } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
	import type { Snippet } from 'svelte';

	let popupModal = $state(false);

	interface Props {
		projectID: string;
		id: string;
		name: string;
		children: Snippet;
		size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
		update: Function;
	}
	let { projectID, id, name, size, update = $bindable(), children }: Props = $props();

	const delFeature = async (featureID: string) => {
		deleteProjectFeature(projectID, featureID).then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project feature deleted successfully',
					dismissible: true,
					type: 'success'
				});

				update();
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
			Are you sure you want to delete the feature, "{name}"?
		</h3>
		<Button color="red" class="mr-2" on:click={() => delFeature(id)}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
