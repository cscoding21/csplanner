<script lang="ts">
	import { PortfolioTable, SectionSubHeading } from "$lib/components";
	import type { Portfolio, Project, Schedule } from "$lib/graphql/generated/sdk";
	import { buildPortfolioTable, getDraftPortfolio, getPortfolio, type ScheduleTable } from "$lib/services/portfolio";
	import { pluralize } from "$lib/utils/format";
	import { callIf } from "$lib/utils/helpers";
	import { ProjectStartDateSet } from "..";
	import BasicsSummary from "./BasicsSummary.svelte";
	import UpdateStatusButtons from "./UpdateStatusButtons.svelte";

    interface Props {
        project:Project
        update?:Function
    }
    let { project, update }:Props = $props()

    let portfolioTable:ScheduleTable = $state({} as ScheduleTable)
    let portfolio:Portfolio = $state({} as Portfolio)

    const refresh = async (): Promise<Portfolio> => {
		const res = await getDraftPortfolio(project.id as string);

        portfolio = res
		portfolioTable = buildPortfolioTable(res, new Date(2025, 5, 1), new Date(2025, 8, 30), project.id as string)

        console.log("portfolio", portfolio)

		return res;
	};

    const getDraftSchedule = (p:Portfolio):Schedule|undefined => {
        if (!p || !p.schedule)
            return 

        for (let i = 0; i < p.schedule.length; i++) {
            const sch = p.schedule[i]

            if (sch.project.id === project.id) {
                return sch
            }
        }
    }
    
    const load = async () => {
        refresh()
    }

    let draftSchedule:Schedule|undefined = $derived(getDraftSchedule(portfolio))
</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <BasicsSummary {project} />
    </div>

    <div class="col-span-2">

        <SectionSubHeading>Current Roadmap</SectionSubHeading>
        {#await load()}
            Loading...
        {:then promiseData} 
            {#if portfolio}
				<div class="mb-2">
					<PortfolioTable {portfolio} highlightProjectID={project.id as string}></PortfolioTable>
				</div>
            {/if}
        {/await}

        <SectionSubHeading cssClass="mt-8">Set Start Date </SectionSubHeading>
        {#if project.projectBasics.startDate && draftSchedule}
            <h1 class="mb-2">Current start date is week of <b>{new Date(draftSchedule.begin).toLocaleDateString()}</b> and end the week of <b>{new Date(draftSchedule.end).toLocaleDateString()}</b></h1>
            <h1 class="mb-4">Duration <b>{draftSchedule.projectActivityWeeks?.length} {pluralize("week", draftSchedule.projectActivityWeeks?.length as number)}</b></h1>
        {/if}

        <ProjectStartDateSet {project} update={() => { refresh(); callIf(update)} } />

        <div class="text-center mt-8">
			<UpdateStatusButtons {project} displayStates={["draft", "backlogged", "scheduled"]} />
		</div>
    </div>
</div>

{/if}