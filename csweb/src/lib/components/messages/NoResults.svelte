<script lang="ts">
	import { Button } from 'flowbite-svelte';
	import { ArrowRightOutline } from 'flowbite-svelte-icons';
	import type { Snippet } from 'svelte';
	import { CSSection, SectionHeading } from '..';

	interface Props {
		title: string;
		newUrl?: string;
		children: Snippet;
	}
	let { title, newUrl, children = $bindable() }: Props = $props();
</script>

<CSSection>
	<SectionHeading>{title}</SectionHeading>
	<p class="mb-3 font-normal leading-tight text-gray-700 dark:text-gray-400">
		{@render children()}
	</p>
	{#if newUrl}
		<Button class="w-fit" href={newUrl as string}>
			Create new <ArrowRightOutline class="ml-2 h-3.5 w-3.5 text-white" />
		</Button>
	{/if}
</CSSection>
