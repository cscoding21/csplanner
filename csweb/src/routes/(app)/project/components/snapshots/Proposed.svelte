<script lang="ts">
	import type { Project, Resource, Schedule } from "$lib/graphql/generated/sdk";
	import { CalendarWeekOutline, ClockSolid, DollarOutline } from "flowbite-svelte-icons";
	import { calculateProjectSchedule } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { DataCard, SectionSubHeading, ResourceCard } from "$lib/components";
	import { formatCurrency, pluralize } from "$lib/utils/format";
	import FinancialSummary from "./FinancialSummary.svelte";
	import BasicsSummary from "./BasicsSummary.svelte";
	import UpdateStatusButtons from "./UpdateStatusButtons.svelte";


    interface Props {
        project:Project
    }
    let { project }:Props = $props()

	let schedule:Schedule = $state({} as Schedule)

	const load = async ():Promise<Schedule> => {
		return calculateProjectSchedule(project.id as string, new Date())
			.then((s) => {
				console.log("found schedule from calculate")
				schedule = s.schedule
				return s.schedule
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project schedule (ProjectSchedule): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
    }
</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <BasicsSummary {project} />
    </div>

    <div class="col-span-2">
		<SectionSubHeading>Financials</SectionSubHeading>
        <FinancialSummary {project} />
		
		<SectionSubHeading>Cost/Schedule</SectionSubHeading>
		{#await load()}
			Loading schedule...
		{:then promiseData} 
			<div class="flex mb-8">
				<div class="flex-1 px-r mr-2">
				<DataCard dataPoint={formatCurrency.format(project.projectCost.calculated?.initialCost as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Implementation cost
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>

			<div class="flex-1 px-r mr-2">
				<DataCard dataPoint={schedule.projectActivityWeeks?.length + "" || ""} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Unadjusted {pluralize("Week", schedule.projectActivityWeeks?.length as number)} to Complete
					{/snippet}
					{#snippet indicator()}
						<CalendarWeekOutline />
					{/snippet}
				</DataCard>
			</div>
			<div class="flex-1 px-r">
				<DataCard dataPoint={project.projectCost.calculated?.hoursActualized + ""} indicatorClass="text-orange-500 dark:text-orange-500">
					{#snippet description()}
						Adjusted {pluralize("hour", project.projectCost.calculated?.hoursActualized as number)}
					{/snippet}
					{#snippet indicator()}
						<ClockSolid />
					{/snippet}
				</DataCard>
			</div>

			</div>
		{/await}
		

        <SectionSubHeading>Implementation Team</SectionSubHeading>		
        <div>
			<div class="grid grid-cols-4">
				{#each project.calculated?.teamMembers as Resource[] as tm}
					<div class="p-2"><ResourceCard resource={tm} /></div>
				{/each}
			</div>
            
        </div>

		<div class="text-center mt-8">
			<UpdateStatusButtons {project} displayStates={["draft", "rejected", "approved"]} />
		</div>
    </div>
</div>

{/if}