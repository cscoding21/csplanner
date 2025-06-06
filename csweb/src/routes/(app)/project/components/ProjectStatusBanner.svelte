<script lang="ts">
	import type { Project, Schedule } from "$lib/graphql/generated/sdk";
    import { Alert, Button } from "flowbite-svelte";
    import { InfoCircleSolid } from "flowbite-svelte-icons";
    import { formatDate, titleCase } from "$lib/utils/format";

    interface Props {
		project: Project;
		schedule?: Schedule;
	}
    let { project, schedule }: Props = $props();

</script>


{#if schedule}
    {#if project.projectStatusBlock.status === "scheduled"}
    <Alert border color="blue" class="mt-2 mb-6">
        {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
        <span class="font-medium">Project scheduled</span>
        This project has been scheduled and is set to begin on <b>{formatDate(project.projectBasics.startDate)}</b>
    </Alert>
    {:else if project.projectStatusBlock.status === "inflight"}
    <Alert border color="green" class="mt-2 mb-6">
        {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
        <span class="font-medium">Project in-flight!</span>
        This project is currently in-flight with a scheduled completion date of <b>{formatDate(schedule?.end)}</b>
    </Alert>
    {/if}
{:else}
    <Alert border class="mt-2 mb-6">
        {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
        Project status is <span class="font-medium">{titleCase(project.projectStatusBlock.status)}</span>
    </Alert>
{/if}


