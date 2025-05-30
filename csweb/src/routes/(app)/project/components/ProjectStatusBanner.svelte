<script lang="ts">
	import type { Project, Schedule } from "$lib/graphql/generated/sdk";
    import { Alert, Button } from "flowbite-svelte";
    import { InfoCircleSolid } from "flowbite-svelte-icons";
    import { formatDate } from "$lib/utils/format";

    interface Props {
		project: Project;
		schedule?: Schedule;
	}
    let { project, schedule }: Props = $props();

</script>


{#if schedule}
    {#if project.projectStatusBlock.status === "scheduled"}
    <Alert border color="blue" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">Project scheduled</span>
        This project has been scheduled and is set to begin on <b>{formatDate(project.projectBasics.startDate)}</b>
    </Alert>
    {:else if project.projectStatusBlock.status === "inflight"}
    <Alert border color="green" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">Project in-flight!</span>
        This project is currently in-flight with a scheduled completion date of <b>{formatDate(schedule?.end)}</b>
    </Alert>
    {/if}
{:else}
    <Alert border color="red" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">{project.projectStatusBlock.status}</span>
        <Button slot="close-button" size="xs" let:close on:click={close} class="ms-auto">Dismiss</Button>
    </Alert>
{/if}


