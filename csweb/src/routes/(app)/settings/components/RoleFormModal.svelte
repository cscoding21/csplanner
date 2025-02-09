<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';
	import { RoleForm } from '.';
	import type { Role } from '$lib/graphql/generated/sdk';

	let popupModal = $state(false);

	interface Props {
        role?: Role | undefined
		children: Snippet;
		size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
		update?: Function;
	}
	let { role, size, update = $bindable(), children }: Props = $props();
</script>

<Button {size} color="dark" onclick={() => (popupModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={popupModal} size="sm" autoclose>
	<RoleForm role={role} update={update} />
</Modal>
