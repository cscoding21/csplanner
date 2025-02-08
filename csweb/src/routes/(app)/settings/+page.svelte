<script lang="ts">
    import { page } from "$app/stores";
    import { CSSection } from "$lib/components";
    import { OrgSettings } from "./components";

    let hash = $state($page.url.hash);
</script>


<div class="grid grid-cols-5">
	<div class="flex w-full col-span-5">
		{#if hash == "#org-settings"}
			<CSSection>
				<OrgSettings  />
			</CSSection>
		{:else if hash == '#basics'}
			<CSSection>
				bloop
			</CSSection>
		{/if}
	</div>
	
	</div>

