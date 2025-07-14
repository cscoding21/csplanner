<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/state';
	import { CSPageSidebar, CSSidebarItem } from '$lib/components';
	import type { Project } from '$lib/graphql/generated/sdk';

	let hash = $state(page.url.hash);

	interface Props {
		project:Project
	}
	let { project }:Props = $props()

	import {
		ChartPieOutline,
        CalendarMonthOutline,
		CalendarWeekOutline,
        ScaleBalancedOutline,
        RectangleListOutline,
        DollarOutline,
        CashOutline,
		PersonChalkboardOutline,

	} from 'flowbite-svelte-icons';
	import { scrollTop } from 'svelte-scrolling';
	import ShowIfStatus from './ShowIfStatus.svelte';
	import { ProjectStatusApproved, ProjectStatusInflight, ProjectStatusNew, ProjectStatusScheduled } from '$lib/services/project';

	afterNavigate((navigation) => {
		scrollTop()
	});
</script>

<svelte:window
	on:hashchange={() => {
		hash = page.url.hash
	}}
/>

{#if project}
<CSPageSidebar>
	<CSSidebarItem href="#snapshot" label="Snapshot" active={hash === "#snapshot"}> 
		{#snippet icon()}
			<ChartPieOutline />
		{/snippet}
		{#snippet tail()}
			<!-- <CheckCircleOutline color="green" /> -->
			<!-- <Badge>5</Badge> -->
		{/snippet}
	</CSSidebarItem>

	<CSSidebarItem href="#collab" label="Collaborate" active={hash === "#collab"}> 
		{#snippet icon()}
			<PersonChalkboardOutline />
		{/snippet}
	</CSSidebarItem>

	<CSSidebarItem href="#basics" label="Basics" active={hash === "#basics"}> 
		{#snippet icon()}
			<ScaleBalancedOutline />
		{/snippet}
	</CSSidebarItem>

	
	{#if false}
	<CSSidebarItem href="#daci" label="DACI" active={hash === "#daci"}> 
		{#snippet icon()}
			<CalendarMonthOutline />
		{/snippet}
	</CSSidebarItem>
	{/if}


	<ShowIfStatus status={project.projectStatusBlock?.status} invert={true} scope={[ProjectStatusNew]}>
	<CSSidebarItem href="#features" label="Features" active={hash === "#features"}> 
		{#snippet icon()}
			<RectangleListOutline />
		{/snippet}
	</CSSidebarItem>

	<CSSidebarItem href="#value" label="Value Prop" active={hash === "#value"}> 
		{#snippet icon()}
			<DollarOutline />
		{/snippet}
	</CSSidebarItem>

	<CSSidebarItem href="#milestones" label="Milestones" active={hash === "#milestones"}> 
		{#snippet icon()}
			<CalendarMonthOutline />
		{/snippet}
	</CSSidebarItem>
	
	<CSSidebarItem href="#costs" label="Costs" active={hash === "#costs"}> 
		{#snippet icon()}
			<CashOutline />
		{/snippet}
	</CSSidebarItem>
	
	</ShowIfStatus>

	<ShowIfStatus status={project.projectStatusBlock?.status} scope={[ProjectStatusApproved, ProjectStatusScheduled, ProjectStatusInflight]}>
	<CSSidebarItem href="#schedule" label="Schedule" active={hash === "#schedule"}> 
		{#snippet icon()}
			<CalendarWeekOutline />
		{/snippet}
	</CSSidebarItem>
	</ShowIfStatus>
	
</CSPageSidebar>
{/if}
