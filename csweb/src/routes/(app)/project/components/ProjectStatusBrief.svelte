<script lang="ts">
	import type { Project } from "$lib/graphql/generated/sdk";
	import { getID } from "$lib/utils/id";
	import { Progressbar } from "flowbite-svelte";
	import { CalendarMonthOutline, ExclamationCircleOutline, QuestionCircleOutline, RectangleListOutline, ToolsOutline } from "flowbite-svelte-icons";
	import { ShowIfStatus } from ".";
	import { ProjectStatusApproved, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusProposed, ProjectStatusScheduled } from "$lib/services/project";
	import { formatCurrency, formatDate, formatPercent, pluralize } from "$lib/utils/format";
	import { MoneyDisplay, PercentDisplay } from "$lib/components";


    interface Props {
        project: Project
    }
    let { project }:Props = $props()

    let isAtRisk = $state(project.calculated?.unhealthyTasks as number > 0)
    let isPastDue = $state(false)

    const getBarColor = (ar:boolean, pd:boolean):"red"|"yellow"|"green" => {
        if(ar) 
            return "red"
        else if (pd)
            return "yellow"

        return "green"
    }

    let barColor = $derived(getBarColor(isAtRisk, isPastDue))


</script>

<div class="flex w-72">



<!-- In Flight -->
<ShowIfStatus scope={[ProjectStatusInflight]} status={project.projectStatusBlock?.status}>

<div class="w-full">
    <Progressbar progress={(project.calculated?.projectPercentComplete as number) * 100.0} id={getID("mil_")} size="h-4" color={barColor} labelInside={true} />
</div>

{#if isPastDue}
    <div class="ml-2" title="Has past-due tasks"><ExclamationCircleOutline color="red" size="md"  /></div>
{/if}
{#if isAtRisk}
    <div class="ml-2" title="Has at-risk tasks"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}

</ShowIfStatus>


<!-- Scheduled -->
<ShowIfStatus scope={[ProjectStatusScheduled]} status={project.projectStatusBlock?.status}>
<div class="w-full">
    NPV (IRR): <MoneyDisplay amount={project.projectValue.calculated?.netPresentValue as number} /> (<PercentDisplay amount={project.projectValue.calculated?.internalRateOfReturn as number} />)<br />
    Kickoff date: <b>{formatDate(project.projectBasics.startDate)}</b>
</div>

{#if isPastDue}
    <div class="ml-2" title="Has past-due tasks"><ExclamationCircleOutline color="red" size="md"  /></div>
{/if}
{#if isAtRisk}
    <div class="ml-2" title="Has at-risk tasks"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}

</ShowIfStatus>

<!-- Approved -->
<ShowIfStatus scope={[ProjectStatusApproved]} status={project.projectStatusBlock?.status}>
<div class="w-full">
    NPV (IRR): <MoneyDisplay amount={project.projectValue.calculated?.netPresentValue as number} /> (<PercentDisplay amount={project.projectValue.calculated?.internalRateOfReturn as number} />)<br />
    {#if project.projectBasics.startDate}
        Proposed date: <b>{formatDate(project.projectBasics.startDate)}</b>
    {:else}
        No proposed kickoff date set
    {/if}
</div>

{#if isPastDue}
    <div class="ml-2" title="Has past-due tasks"><ExclamationCircleOutline color="red" size="md"  /></div>
{/if}
{#if isAtRisk}
    <div class="ml-2" title="Has at-risk tasks"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}

</ShowIfStatus>


<!-- Draft -->
<ShowIfStatus scope={[ProjectStatusDraft]} status={project.projectStatusBlock?.status}>
<ul>

<li class="flex">
<div class="mr-2"><RectangleListOutline size="md" class="mr-2 float-left" /> {project.projectFeatures?.length} {pluralize("feature", project.projectFeatures?.length as number)}</div>
{#if project.calculated?.featureStatusProposedCount || 0 > 0}
    <div class="mx-3" title="Has unsettled features"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}
</li>

<li class="flex">
<div class="mr-2 clear-both"><CalendarMonthOutline size="md" class="mr-2 float-left" /> {project.projectMilestones?.length} {pluralize("milestone", project.projectMilestones?.length as number)}</div>
</li>

<li class="flex">
<div class="mr-2 clear-both"><ToolsOutline size="md" class="mr-2 float-left" /> {project.calculated?.totalTasks} {pluralize("task", project.calculated?.totalTasks as number)}</div>
{#if project.calculated?.unhealthyTasks || 0 > 0}
    <div class="ml-2" title="Has unsettled tasks"><QuestionCircleOutline color="yellow" size="md" /></div>
{/if}
</li>

</ul>

</ShowIfStatus>


<!-- Proposed -->
<ShowIfStatus scope={[ProjectStatusProposed]} status={project.projectStatusBlock?.status}>

<div class="mr-2">Waiting on approval decision</div>

</ShowIfStatus>


</div>

