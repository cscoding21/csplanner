<script lang="ts">
	import type { Project, Schedule } from "$lib/graphql/generated/sdk";
    import { Alert, Button } from "flowbite-svelte";
    import { InfoCircleSolid } from "flowbite-svelte-icons";
    import { formatDate } from "$lib/utils/format";

    interface Props {
		project: Project;
		schedule?: Schedule;
        update: Function;
	}
    let { project, schedule, update }: Props = $props();

</script>


{#if schedule}
    {#if project.projectBasics.status === "scheduled"}
    <Alert border color="blue" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">Project scheduled</span>
        This project has been scheduled and is set to begin on <b>{formatDate(project.projectBasics.startDate)}</b>
    </Alert>
    {:else if project.projectBasics.status === "inflight"}
    <Alert border color="green" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">Project in-flight!</span>
        This project is currently in-flight with a scheduled completion date of <b>{formatDate(schedule?.end)}</b>
    </Alert>
    {/if}
{:else}
    <Alert border color="green" class="mt-2 mb-6">
        <InfoCircleSolid slot="icon" class="w-5 h-5" />
        <span class="font-medium">{project.projectBasics.status}</span>
        <Button slot="close-button" size="xs" let:close on:click={close} class="ms-auto">Dismiss</Button>
    </Alert>
{/if}


