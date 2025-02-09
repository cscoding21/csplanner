<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/stores';

	import {
		Sidebar,
		SidebarDropdownWrapper,
		SidebarGroup,
		SidebarItem,
		SidebarWrapper
	} from 'flowbite-svelte';
	import {
		AngleDownOutline,
		AngleUpOutline,
		GithubSolid,
		ChartPieOutline,
		TableColumnSolid,

		UserOutline,

		OutdentOutline,

		AdjustmentsVerticalOutline,

		ListOutline,

		RectangleListOutline





	} from 'flowbite-svelte-icons';

	export let drawerHidden: boolean = false;

	const closeDrawer = () => {
		drawerHidden = false;
	};

	let iconClass =
		'flex-shrink-0 w-6 h-6 text-gray-500 transition duration-75 group-hover:text-gray-900 dark:text-gray-400 dark:group-hover:text-white';
	let itemClass =
		'flex items-center p-2 text-base text-gray-900 transition duration-75 rounded-lg hover:bg-gray-100 group dark:text-gray-200 dark:hover:bg-gray-700';
	let groupClass = 'pt-2 space-y-2';

	$: mainSidebarUrl = $page.url.pathname;
	let activeMainSidebar: string;

	afterNavigate((navigation) => {
		// this fixes https://github.com/themesberg/flowbite-svelte/issues/364
		document.getElementById('svelte')?.scrollTo({ top: 0 });
		closeDrawer();

		activeMainSidebar = navigation.to?.url.pathname ?? '';
	});

	let posts = [
		{ name: 'My Info', icon: UserOutline, href: '#my-info' },
		{ name: 'Org Settings', icon: AdjustmentsVerticalOutline, href: '#org-settings' },
		{ name: 'Lists', icon: ListOutline, href: '#lists' },
		{ name: 'Roles', icon: RectangleListOutline, href: '#roles' },
	];

	let links = [
		{
			label: 'GitHub Repository',
			href: 'https://github.com/cscoding21/csplanner',
			icon: GithubSolid
		}
	];
	let dropdowns = Object.fromEntries(Object.keys(posts).map((x) => [x, false]));
</script>

<!-- asideClass="fixed inset-0 z-30 flex-none h-full w-64 lg:h-auto border-e border-gray-200 dark:border-gray-600 lg:overflow-y-visible lg:pt-16 lg:block" -->
<Sidebar
	class="rounded-lg"
	activeUrl={mainSidebarUrl}
	activeClass="bg-gray-100 dark:bg-gray-700 rounded-lg"
>
	<h4 class="sr-only">Main menu</h4>
	<SidebarWrapper
		divClass="overflow-y-auto px-3 pt-20 lg:pt-5 h-full bg-white scrolling-touch max-w-2xs lg:h-[calc(100vh-4rem)] lg:block dark:bg-gray-800 lg:me-0 lg:sticky top-2"
	>
		<nav class="divide-y divide-gray-200 dark:divide-gray-700">
			<SidebarGroup ulClass={groupClass} class="mb-3">
				{#each posts as { name, icon, href } (name)}
				<SidebarItem
					label={name}
					{href}
					spanClass="ml-3"
					class={itemClass}
					active={activeMainSidebar === href}
				>
					<svelte:component this={icon} slot="icon" class={iconClass} />
				</SidebarItem>
				{/each}
			</SidebarGroup>
			<SidebarGroup ulClass={groupClass}>
				{#each links as { label, href, icon } (label)}
					<SidebarItem
						{label}
						{href}
						spanClass="ml-3"
						class={itemClass}
						active={activeMainSidebar === href}
						target="_blank"
					>
						<svelte:component this={icon} slot="icon" class={iconClass} />
					</SidebarItem>
				{/each}
			</SidebarGroup>
		</nav>
	</SidebarWrapper>
</Sidebar>
