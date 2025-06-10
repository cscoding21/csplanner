<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
	import ProjectStatusBanner from "../ProjectStatusBanner.svelte";
	import BasicsSummary from "./BasicsSummary.svelte";
	import FinancialSummary from "./FinancialSummary.svelte";
	import UpdateStatusButtons from "./UpdateStatusButtons.svelte";

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

        <ProjectStatusBanner project={project} />

        <FinancialSummary {project} />

        <UpdateStatusButtons {project} />
        
    </div>
</div>

{/if}