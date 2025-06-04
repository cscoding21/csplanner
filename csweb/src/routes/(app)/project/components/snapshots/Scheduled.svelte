<script lang="ts">
	import { ScheduleTable, SectionSubHeading } from "$lib/components";
	import type { Project, Schedule } from "$lib/graphql/generated/sdk";
	import { getScheduledProjectFromPortfolio } from "$lib/services/portfolio";
	import { Alert } from "flowbite-svelte";
	import BasicsSummary from "./BasicsSummary.svelte";
	import ScheduleSummary from "./ScheduleSummary.svelte";
	import UpdateStatusButtons from "./UpdateStatusButtons.svelte";
	import { formatDate } from "$lib/utils/format";
	import { InfoCircleSolid } from "flowbite-svelte-icons";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()

    const refresh = async (): Promise<Schedule> => {
		const res = await getScheduledProjectFromPortfolio(project.id as string);

        projectSchedule = res
        console.log("projectSchedule", projectSchedule)
		return res;
	};
    
    const load = async () => {
        refresh()
    }

    let projectSchedule:Schedule|undefined = $state()
    let view:"task"|"resource"|undefined = $state("task")
</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <BasicsSummary {project} />
    </div>

    <div class="col-span-2">
        <SectionSubHeading>Current Schedule</SectionSubHeading>
        {#await load()}
            Loading...
        {:then promiseData} 
            {#if projectSchedule}
            <ScheduleSummary schedule={projectSchedule} />
            {/if} 
        {/await}
        <Alert color="blue" border class="my-4">
            {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
            This project will automatically convert to "in-flight" on its start date of <b>{formatDate(projectSchedule?.begin)}</b>
        </Alert>

        <SectionSubHeading>Plan Details</SectionSubHeading>
        <ScheduleTable schedule={projectSchedule as Schedule} view={view as "task"|"resource"|undefined} />


        <div class="text-center mt-8">
			<UpdateStatusButtons {project} />
		</div>
    </div>
</div>

{/if}