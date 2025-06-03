<script lang="ts">
	import { PortfolioTable, SectionSubHeading, UserCard } from "$lib/components";
	import type { Portfolio, Project, Schedule, User } from "$lib/graphql/generated/sdk";
	import { type ProjectRow, getPortfolio, buildPortfolioTable, type ScheduleTable } from "$lib/services/portfolio";

    interface Props {
        project:Project
    }
    let { project }:Props = $props()

    let portfolioTable:ScheduleTable = $state({} as ScheduleTable)
    let portfolio:Portfolio = $state({} as Portfolio)

    const refresh = async (): Promise<Portfolio> => {
		const res = await getPortfolio();

		portfolioTable = buildPortfolioTable(res, undefined, undefined)

		return res;
	};
    
    const load = async () => {
        refresh()
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

        This project is approved and needs to be scheduled.

        {#if portfolio}
				<div class="mb-2">
					<PortfolioTable {portfolio}></PortfolioTable>
				</div>
        {/if}
        
    </div>
</div>

{/if}