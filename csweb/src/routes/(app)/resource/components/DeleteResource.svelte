<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { goto } from '$app/navigation';
	import { deleteResource } from '$lib/services/resource';
	import { addToast } from '$lib/stores/toasts';
	import type { Snippet } from 'svelte';

	let deleteResourceModal = $state(false);

	interface Props {
		id: string;
		name: string|undefined;
		children: Snippet;
	}
	let { id, name, children }: Props = $props();

	const deleteItem = async () => {
		if (!id) {
			return;
		}

		deleteResource(id)
			.then((res) => {
				if (res.success) {
					addToast({
						message: 'Resource deleted successfully',
						dismissible: true,
						type: 'success'
					});

					goto('/resource');
				} else {
					addToast({
						message: 'Error deleting resource: ' + res.message,
						dismissible: true,
						type: 'error'
					});
				}
			})
			.catch((err) => {
				addToast({
					message: 'Error deleting resource: ' + err,
					dismissible: true,
					type: 'error'
				});
			});
	};
</script>

<Button onclick={() => (deleteResourceModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={deleteResourceModal} size="xs" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Are you sure you want to delete the resource, {name}?
		</h3>
		<Button color="red" class="mr-2" onclick={deleteItem}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
