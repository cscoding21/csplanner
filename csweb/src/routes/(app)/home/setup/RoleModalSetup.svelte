<script lang="ts">
	import { Modal } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';
	import RoleFormSetup from './RoleFormSetup.svelte';
	import type { Role } from '$lib/graphql/generated/sdk';
	import { callIf } from '$lib/utils/helpers';

	let popupModal = $state(false);

	interface Props {
        role?: Role | undefined
		children: Snippet;
		update?: Function;
	}
	let { role = $bindable(), update = $bindable(), children }: Props = $props();

	let formUpdated = () => {
		popupModal = false

		callIf(update)
	}
</script>

<button onclick={() => (popupModal = true)}>
	{@render children()}
</button>

<Modal bind:open={popupModal}>
	<RoleFormSetup bind:role={role} update={formUpdated} />
</Modal>
