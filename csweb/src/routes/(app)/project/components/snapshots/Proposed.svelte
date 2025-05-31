<script lang="ts">
	import type { Project, Resource, Schedule, User } from "$lib/graphql/generated/sdk";
	import { CalendarWeekOutline, ClockSolid, DollarOutline, SalePercentOutline } from "flowbite-svelte-icons";
	import { Button } from "flowbite-svelte";
	import { calculateProjectSchedule, setProjectStatus } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { reloadPage } from "$lib/utils/helpers";
	import { DataCard, SectionSubHeading, UserCard, ResourceCard, ScheduleTable } from "$lib/components";
	import { formatCurrency, formatPercent, pluralize } from "$lib/utils/format";
	import { ProjectValueCategoryDistributionChart, ProjectValueChart } from "..";


    interface Props {
        project:Project
    }
    let { project }:Props = $props()

	let schedule:Schedule = $state({} as Schedule)

    const setStatus = async (s:string) => {
		setProjectStatus(project.id as string, s).then((res) => {
			if (res.status?.success) {
				addToast({
					message: 'Project set to ' + s,
					dismissible: true,
					type: 'success'
				});

                reloadPage()
			} else {
				addToast({
					message: 'Error setting project status to ' + s + ': ' + res.status?.message,
					dismissible: true,
					type: 'error'
				});
			}
		});
	};


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
        <SectionSubHeading>{project.projectBasics.name}</SectionSubHeading>

        <h3 class="text-gray-100 font-semibold">Owner</h3>
        <UserCard user={project.projectBasics.owner as User} />

        <h3 class="text-gray-100 font-semibold mt-6">Executive Summary</h3>
        <p class="py-2 text-gray-200">{@html project.projectBasics.description.replaceAll(/[\n]/g, "<br />")}</p>
    </div>

    <div class="col-span-2">
        <SectionSubHeading>Financials</SectionSubHeading>
        <div class="flex mb-8">
			<div class="flex-1 px-r mr-2">
				<DataCard dataPoint={formatCurrency.format(project.projectValue.calculated?.netPresentValue as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Net Present Value
					{/snippet}
					{#snippet indicator()}
						<DollarOutline />
					{/snippet}
				</DataCard>
			</div>

			<div class="flex-1 px-r">
				<DataCard dataPoint={formatPercent.format(project.projectValue.calculated?.internalRateOfReturn as number)} indicatorClass="text-green-500 dark:text-green-500">
					{#snippet description()}
						Internal Rate of Return
					{/snippet}
					{#snippet indicator()}
						<SalePercentOutline />
					{/snippet}
				</DataCard>
			</div>

			<div class="flex-1 px-2">
				<ProjectValueCategoryDistributionChart height={110} {project}></ProjectValueCategoryDistributionChart>
			</div>
        </div>

		<div class="flex mb-8">			
			<div class="flex-1 px-2">
				<SectionSubHeading>Five Year Outlook</SectionSubHeading>
				<ProjectValueChart project={project}></ProjectValueChart>
			</div>
        </div>

		
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

			<!-- <div class="mb-8">
				<div>
				<ScheduleTable view="task" schedule={schedule} />
				</div>
			</div> -->
		{/await}
		

        <SectionSubHeading>Implementation Team</SectionSubHeading>
		
        <div>
			<div class="flex">
				{#each project.calculated?.teamMembers as Resource[] as tm}
					<div class="w-1/3 p-2"><ResourceCard resource={tm} /></div>
				{/each}
			</div>
            
        </div>

		<div class="text-center mt-8">
			<Button onclick={() => setStatus("draft")} color="yellow" class="px-8 m-2 w-64">Revert to Draft</Button>
			<Button onclick={() => setStatus("reject")} color="red" class="px-8 m-2 w-64">Reject</Button>
			<Button onclick={() => setStatus("approve")} color="green" class="px-8 m-2 w-64">Approve</Button>
		</div>
    </div>
</div>

{/if}