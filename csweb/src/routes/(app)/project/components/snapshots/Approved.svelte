<script lang="ts">
	import { PortfolioTable, SectionSubHeading } from "$lib/components";
	import type { Portfolio, Project } from "$lib/graphql/generated/sdk";
	import { buildPortfolioTable, getDraftPortfolio, getPortfolio, type ScheduleTable } from "$lib/services/portfolio";
	import { callIf } from "$lib/utils/helpers";
	import { ProjectStartDateSet } from "..";
	import BasicsSummary from "./BasicsSummary.svelte";

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
    
    const load = async () => {
        refresh()
    }

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
        {#if project.projectBasics.startDate}
            <h1 class="mb-4">Current start date is week of <b>{new Date(project.projectBasics.startDate).toLocaleDateString()}</b></h1>
        {/if}

        <ProjectStartDateSet {project} update={() => { refresh(); callIf(update)} } />
    </div>
</div>

{/if}