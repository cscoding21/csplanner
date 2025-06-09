<script lang="ts">
	import { DataCard, SectionSubHeading } from "$lib/components";
	import { formatCurrency, formatPercent } from "$lib/utils/format";
	import { DollarOutline, SalePercentOutline } from "flowbite-svelte-icons";
	import { ProjectValueCategoryDistributionChart, ProjectValueChart } from "..";
	import type { Project } from "$lib/graphql/generated/sdk";

    interface Props {
        project: Project
        abridged?:boolean
    }
    let { project, abridged }:Props = $props()
</script>


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

{#if !abridged}
<div class="flex mb-8">			
    <div class="flex-1 px-2">
        <SectionSubHeading>Five Year Outlook</SectionSubHeading>
        <ProjectValueChart project={project}></ProjectValueChart>
    </div>
</div>
{/if}