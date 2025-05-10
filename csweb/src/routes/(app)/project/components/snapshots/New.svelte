<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()
</script>

{#if project}
<div class="grid grid-cols-2 gap-4">
    <div>
        {@html project.projectBasics.description.replaceAll(/[\n]/g, "<br />")}
    </div>
    <div>

    </div>
</div>
{/if}