<script lang="ts">
    import { page } from "$app/state";
    import { CSSection } from "$lib/components";
    import { OrgSettings, Lists, Roles, MyInfo, Templates } from "./components";

    let hash = $state(page.url.hash);
</script>

<svelte:window
	on:hashchange={() => {
		hash = page.url.hash
	}}
/>

<div class="grid grid-cols-5">
	<div class="flex w-full col-span-5">
	<CSSection>
		{#if hash == "#org-settings"}
			<OrgSettings  />
		{:else if hash == '#lists'}
			<Lists />	
		{:else if hash == '#roles'}
			<Roles />	
		{:else if hash == '#templates'}
			<Templates />
		{:else if hash == '#my-info' || hash == ''}
			<MyInfo />
		{/if}
	</CSSection>
	</div>
	
	</div>

