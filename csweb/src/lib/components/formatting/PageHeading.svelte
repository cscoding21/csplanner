<script lang="ts">
	import { Heading } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';

	interface PageHeadingProps {
		title: string;
		children?: Snippet;
	}
	let { title, children = $bindable() }: PageHeadingProps = $props();
</script>

<Heading tag="h3" class="mb-4 px-2">
	{title}

	<span class="float-right">
		{#if children}
			{@render children()}
		{/if}
	</span>
</Heading>
