<script lang="ts">
	import { afterNavigate } from '$app/navigation';
	import { page } from '$app/state';
	import { CSPageSidebar, CSSidebarItem } from '$lib/components';

	let hash = $state(page.url.hash);

	import {
		ChartPieOutline,
        CalendarMonthOutline,
		CalendarWeekOutline,
        ScaleBalancedOutline,
        RectangleListOutline,
        DollarOutline,
        CashOutline,

		CheckCircleOutline

	} from 'flowbite-svelte-icons';
	import { scrollTop } from 'svelte-scrolling';

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
	<CSSidebarItem href="#snapshot" label="Snapshot" active={hash === "#snapshot"}> 
		{#snippet icon()}
			<ChartPieOutline />
		{/snippet}
		<!-- {#snippet tail()}
			<CheckCircleOutline color="green" />
		{/snippet} -->
	</CSSidebarItem>
	<CSSidebarItem href="#basics" label="Basics" active={hash === "#basics"}> 
		{#snippet icon()}
			<ScaleBalancedOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#daci" label="DACI" active={hash === "#daci"}> 
		{#snippet icon()}
			<CalendarMonthOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#value" label="Value Prop" active={hash === "#value"}> 
		{#snippet icon()}
			<DollarOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#costs" label="Costs" active={hash === "#costs"}> 
		{#snippet icon()}
			<CashOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#features" label="Features" active={hash === "#features"}> 
		{#snippet icon()}
			<RectangleListOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#milestones" label="Milestones" active={hash === "#milestones"}> 
		{#snippet icon()}
			<CalendarMonthOutline />
		{/snippet}
	</CSSidebarItem>
	<CSSidebarItem href="#schedule" label="Schedule" active={hash === "#schedule"}> 
		{#snippet icon()}
			<CalendarWeekOutline />
		{/snippet}
	</CSSidebarItem>
</CSPageSidebar>
