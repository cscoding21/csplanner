<script lang="ts">
	import { Heading, Hr } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';

	interface Props {
		children?: Snippet;
	}
	let { 
		children
	 }: Props = $props();
</script>

<Heading tag="h6">
	{#if children}
		{@render children()}
	{/if}
</Heading>
<Hr hrClass="h-px mt-2 mb-6 bg-gray-200 border-0 dark:bg-gray-700" />
