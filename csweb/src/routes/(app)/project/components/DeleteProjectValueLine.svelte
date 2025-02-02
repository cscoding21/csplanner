<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
	import { deleteProjectValueLine } from '$lib/services/project';
	import { addToast } from '$lib/stores/toasts';
    import { formatCurrency } from '$lib/utils/format';
	import type { Snippet } from 'svelte';

	let popupModal = $state(false);

	interface Props {
		projectID: string;
		id: string;
		name: number;
		children: Snippet;
		size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
		update: Function;
	}
	let { projectID, id, name, size, update = $bindable(), children }: Props = $props();

	const delValueLine= async (featureID: string) => {
		deleteProjectValueLine(projectID, featureID).then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project value line deleted successfully',
					dismissible: true,
					type: 'success'
				});

				update();
			} else {
				addToast({
					message: 'Error deleting project value line: ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};
</script>

<Button {size} color="dark" onclick={() => (popupModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={popupModal} size="sm" autoclose>
	<div class="text-center">
		<ExclamationCircleOutline class="mx-auto mb-4 h-12 w-12 text-gray-400 dark:text-gray-200" />
		<h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">
			Are you sure you want to delete this value line <br />(5 year total {formatCurrency.format(name)})?
		</h3>
		<Button color="red" class="mr-2" onclick={() => delValueLine(id)}>Yes, I'm sure</Button>
		<Button color="alternative">No, cancel</Button>
	</div>
</Modal>
