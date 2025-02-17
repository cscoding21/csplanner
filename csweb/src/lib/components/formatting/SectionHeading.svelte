<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import { CSHR } from '..';
	import type { Snippet } from 'svelte';

	interface Props {
		children?: Snippet;
	}
	let { children }: Props = $props();
</script>

<Heading tag="h6">
	{#if children}
		{@render children()}
	{/if}
</Heading>
<CSHR />
