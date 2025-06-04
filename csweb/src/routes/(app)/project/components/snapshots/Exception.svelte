<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
	import BasicsSummary from "./BasicsSummary.svelte";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()


</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <BasicsSummary {project} />
    </div>

    <div class="col-span-2">

        All your error are belong to us!
        
    </div>
</div>

{/if}