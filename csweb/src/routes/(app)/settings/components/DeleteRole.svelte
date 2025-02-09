<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { deleteRole } from '$lib/services/resource';
	import { goto } from '$app/navigation';
	import { addToast } from '$lib/stores/toasts';
	import type { Snippet } from 'svelte';
	import { callIf } from '$lib/utils/helpers';

	let popupModal = $state(false);

	interface Props {
		id: string;
		name: string;
		children: Snippet;
        update?: Function;
        size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
	}

	let { id, name, update, size, children }: Props = $props();

	const deleteItem = async () => {
		if (!id) {
			return;
		}

		const res = await deleteRole(id);
		if (res.success) {
			addToast({
				message: 'Role deleted successfully',
				dismissible: true,
				type: 'success'
			});

			callIf(update);
		} else {
			addToast({
				message: 'Error deleting role: ' + res.message,
				dismissible: true,
				type: 'error'
			});
		}
	};
</script>

<Button on:click={() => (popupModal = true)}>{@render children()}</Button>

<Modal bind:open={popupModal} size={size} autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Are you sure you want to delete the role, {name}?
		</h3>
		<Button color="red" class="mr-2" on:click={deleteItem}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
