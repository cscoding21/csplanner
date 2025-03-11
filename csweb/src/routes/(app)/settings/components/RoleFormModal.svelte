<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';
	import { RoleForm } from '.';
	import type { RoleEnvelope} from '$lib/graphql/generated/sdk';
	import { callIf } from '$lib/utils/helpers';

	let popupModal = $state(false);

	interface Props {
        role?: RoleEnvelope | undefined
		children: Snippet;
		size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
		update?: Function;
	}
	let { role, size, update = $bindable(), children }: Props = $props();

	let formUpdated = () => {
		popupModal = false

		callIf(update)
	}
</script>

<Button {size} color="dark" onclick={() => (popupModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={popupModal} size="sm">
	<RoleForm role={role} update={formUpdated} />
</Modal>
