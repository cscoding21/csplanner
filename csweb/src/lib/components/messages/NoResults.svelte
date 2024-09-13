<script lang="ts">
	import { Card, Button } from 'flowbite-svelte';
	import { ArrowRightOutline } from 'flowbite-svelte-icons';
	import type { Snippet } from 'svelte';

	interface NoResultsProps {
		title: string;
		newUrl: string;
		children: Snippet;
	}
	let { title, newUrl, children = $bindable() }: NoResultsProps = $props();
</script>

<Card class="w-full" size="xl">
	<h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 dark:text-white">{title}</h5>
	<p class="mb-3 font-normal leading-tight text-gray-700 dark:text-gray-400">
		{@render children()}
	</p>
	{#if newUrl}
		<Button class="w-fit" href={newUrl as string}>
			Create new <ArrowRightOutline class="ml-2 h-3.5 w-3.5 text-white" />
		</Button>
	{/if}
</Card>
