<script lang="ts">
	import { onMount } from 'svelte';
	import { page } from '$app/state';
	import logoImage from '$lib/assets/images/csplanner-logo-white.png?enhanced'
	import logoImageLightBG from '$lib/assets/images/csplanner-logo-dark.png?enhanced'

	import {
		Navbar,
		NavBrand,
		Avatar,
		Dropdown,
		DropdownItem,
		DropdownHeader,
		DropdownDivider,
		DarkMode,
		Button,
		Input,

		DropdownGroup

	} from 'flowbite-svelte';
	import { authService } from '$lib/services/auth';
	import { SearchOutline } from 'flowbite-svelte-icons';
	import { goto } from '$app/navigation';
	import { getInitialsFromName } from '$lib/utils/format';
	import { PageMessages, CSNavItem, NotificationList, OrgStateChecker } from '$lib/components';
	import { isDarkMode } from '$lib/utils/darkmode';
	import type { Organization } from '$lib/graphql/generated/sdk';
	import { getOrganization } from '$lib/services/organization';

	const as = authService();
	const cu = as.currentUser();

	let { children } = $props();

	let pageCat = $derived(page.url.pathname)
	let showDarkModeLogo = $derived(isDarkMode())
	let org = $state({} as Organization)

	const logoutUser = () => {
		as.signout().then((r) => {
			if (r) {
				goto('/login');
			}
		});
	};

	const isTLPage = (token:string):boolean => {
		return pageCat.indexOf(token) > -1
	}

	onMount(async () => {
		console.log('layout onMount');
		if (!as.authCheck()) {
			goto('/login');
		}

		as.refreshCycle();

		await getOrganization().then(o => {
			org = o
		})
	});
</script>

<div class="w-full h-screen">

<OrgStateChecker invert={false} stateToCheck="isReadyForProjects">
<Navbar class="dark:bg-gray-800 bg-gray-100" fluid={true}>
	<NavBrand href="/">
		{#if showDarkModeLogo}
		<enhanced:img src={logoImage} width="141" alt="csPlanner" />
		{:else}
		<enhanced:img src={logoImageLightBG} width="141" alt="csPlanner" />
		{/if}
	</NavBrand>
	<div class="flex items-center md:order-3">
		<DarkMode class="mr-2 text-2xl" />
		<NotificationList />
		<span class="ml-4 text-gray-800 dark:text-gray-200">{org.name}</span>
		
		<Avatar id="avatar-menu" src={cu?.profileImage || ''} class="ml-6 cursor-pointer"
			>{getInitialsFromName(cu?.firstName + ' ' + cu?.lastName || '')}</Avatar
		>
	</div>

	<Dropdown placement="bottom" triggeredBy="#avatar-menu">
		<DropdownHeader>
			<span class="block text-sm">{cu?.firstName}</span>
			<span class="block truncate text-sm font-medium">{cu?.email}</span>
		</DropdownHeader>
		<DropdownGroup>
		<DropdownItem>
			<a href="http://localhost:3006" target="_blank">Account</a>
		</DropdownItem>
		<DropdownItem>
			<a href="/settings">Settings</a>
		</DropdownItem>
		<DropdownDivider />
		<DropdownItem>
			<button class="gb-initial" onclick={logoutUser}>Sign out</button>
		</DropdownItem>
		</DropdownGroup>
	</Dropdown>

	<!-- <div class="flex md:order-2">
		<!-- color="none" ->
		<Button
			data-collapse-toggle="mobile-menu-3"
			aria-controls="mobile-menu-3"
			aria-expanded="false"
			class="mr-1 rounded-lg p-2.5 text-sm text-gray-500 hover:bg-gray-100 focus:outline-none focus:ring-4 focus:ring-gray-200 dark:text-gray-400 dark:hover:bg-gray-700 dark:focus:ring-gray-700 md:hidden"
		>
			<SearchOutline class="h-5 w-5" />
		</Button>
		<div class="relative hidden md:block">
			<div class="pointer-events-none absolute inset-y-0 left-0 flex items-center pl-3">
				<SearchOutline class="h-4 w-4" />
			</div>
			<Input id="search-navbar" class="pl-10" placeholder="Search..." />
		</div>
	</div> -->

	<div class="hidden justify-between items-center w-full lg:flex lg:w-auto lg:order-1">
		<ul class="flex flex-col mt-4 space-x-6 font-medium lg:flex-row xl:space-x-8 lg:mt-0">
			<CSNavItem   href="/home" active={isTLPage("/home")}>Home</CSNavItem>	
			<CSNavItem   href="/project" active={isTLPage("/project")}>Projects</CSNavItem>	
			<CSNavItem   href="/resource" active={isTLPage("/resource")}>Resources</CSNavItem>	
			<CSNavItem   href="/roadmap" active={isTLPage("/roadmap")}>Roadmap</CSNavItem>	
			<CSNavItem   href="/insight" active={isTLPage("/insight")}>Insights</CSNavItem>	
		</ul>
	</div>	
</Navbar>
</OrgStateChecker>	

<div>
	{@render children()}
</div>

</div>

<PageMessages />

