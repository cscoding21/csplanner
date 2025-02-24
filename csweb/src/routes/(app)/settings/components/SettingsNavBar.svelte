<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/state';
	import { CSPageSidebar, CSSidebarItem } from '$lib/components';

	import {
		UserOutline,
		AdjustmentsVerticalOutline,
		ListOutline,
		RectangleListOutline,
		BookOpenOutline

	} from 'flowbite-svelte-icons';
	import { scrollTop } from 'svelte-scrolling';

	let hash = $state(page.url.hash);

	afterNavigate((navigation) => {
		scrollTop()
	});
</script>

<svelte:window
	on:hashchange={() => {
		hash = page.url.hash
	}}
/>

<CSPageSidebar>
	<CSSidebarItem href="#my-info" label="My Info" active={hash === "#my-info"}> 
		{#snippet icon()}
			<UserOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#org-settings" label="Org Info" active={hash === "#org-settings"}> 
		{#snippet icon()}
			<AdjustmentsVerticalOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#lists" label="Lists" active={hash === "#lists"}> 
		{#snippet icon()}
			<ListOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#roles" label="Roles" active={hash === "#roles"}> 
		{#snippet icon()}
			<RectangleListOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#templates" label="Templates" active={hash === "#templates"}> 
		{#snippet icon()}
			<BookOpenOutline />
		{/snippet}
	</CSSidebarItem>
</CSPageSidebar>
