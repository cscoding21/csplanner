<script lang="ts">
	import { Button, Modal } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';
	import ProjectValueLineForm from './ProjectValueLineForm.svelte';
	import type { ProjectValueLine } from '$lib/graphql/generated/sdk';

	let popupModal = $state(false);

	interface Props {
		projectID: string;
        valueItem: ProjectValueLine | undefined
		children: Snippet;
		size: 'xs' | 'sm' | 'lg' | 'xl' | 'md' | undefined;
		update: Function;
	}
	let { projectID, valueItem, size, update = $bindable(), children }: Props = $props();
</script>

<Button {size} color="dark" onclick={() => (popupModal = true)}>
	{@render children()}
</Button>

<Modal bind:open={popupModal} size="sm" autoclose>
	<ProjectValueLineForm projectID={projectID} valueItem={valueItem} update={update} />
</Modal>
