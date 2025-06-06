<script lang="ts">
	import { Heading, Hr } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';
	import { CSHR } from '..';

	interface Props {
		children?: Snippet;
		cssClass?: string;
	}
	let { children, cssClass }: Props = $props();
</script>

<Heading tag="h6" class="text-md dark:text-gray-300 text-gray-600 {cssClass}">
	{#if children}
		{@render children()}
	{/if}
</Heading>
<CSHR />