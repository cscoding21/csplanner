<script lang="ts">
	import { Breadcrumb, BreadcrumbItem } from 'flowbite-svelte';
	import type { Snippet } from 'svelte';

	interface Props {
		pageDetail: string|undefined;
		children: Snippet;
	}
	let { pageDetail, children }: Props = $props();
</script>

<div class="grid w-full grid-cols-4">
	<div class="col-span-2 justify-start p-2">
		<Breadcrumb aria-label="Solid background breadcrumb example">
			<BreadcrumbItem href="/" home>Home</BreadcrumbItem>
			<BreadcrumbItem href="/resource">Resources</BreadcrumbItem>
			{#if pageDetail}
				<BreadcrumbItem>{pageDetail}</BreadcrumbItem>
			{/if}
		</Breadcrumb>
	</div>
	<div>&nbsp;</div>
	<div class="justify-self-end p-2">
		{@render children()}
	</div>
</div>
