<script lang="ts">
    import { 
        SectionHeading,
        MoneyDisplay,
        ResourceList,
        UserCard
    } from "$lib/components";
    import { BadgeFeaturePriority, BadgeProjectStatus, BadgeFeatureStatus } from ".";
    import { Table, TableBody, TableHead, TableBodyRow, TableHeadCell, TableBodyCell, Popover, Hr } from "flowbite-svelte";
    import { QuestionCircleSolid } from "flowbite-svelte-icons";
    import { 
        ProjectValueChart
    } from '../components'
    import type { Project, Resource, User } from '$lib/graphql/generated/sdk'
    import { formatPercent, formatCurrency } from "$lib/utils/format";
    import { getProject } from "$lib/services/project";
    import { getDefaultProject } from "$lib/forms/project.validation";
    import { normalizeGUID } from "$lib/utils/id";

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let project:Project = $state(getDefaultProject() as Project)

    const load = async (): Promise<Project> => {
		return await getProject(id)
			.then((proj) => {
                //proj.projectFeatures?.sort((a, b) => a.priority.localeCompare(b.priority))

				return proj;
			})
			.catch((err) => {
				return err;
			});
	};

    const loadPage = async () => {
        load().then(p => {
            project = p
        });
	};

</script>

<div class="grid grid-cols-3 gap-8">


{#await loadPage()}
	Loading...
{:then promiseData}

{#if project}
<!-- col span 3 -->
<div class="col-span-3">
<SectionHeading>
    Financials: {project.projectBasics.name}
    <span class="float-right"><BadgeProjectStatus status={project.projectBasics.status} /></span>
</SectionHeading>
</div>
{/if}

<div class="col-span-1">
    <!-- col span 1 -->
    <SectionHeading>Executive Summary</SectionHeading>
    <p class="mb-6 text-sm">{project.projectBasics?.description}</p>
</div>


<div class="col-span-1">
    <!-- col span 1 -->
    <SectionHeading>Financials</SectionHeading>
    <ProjectValueChart {project}></ProjectValueChart>
    <ul class="list mb-6 col-span-2 text-xs p-2">
        <li>
            <span>Net Present Value</span>
            <span class="float-right flex-auto font-semibold">
                <MoneyDisplay amount={project.projectValue?.calculated?.netPresentValue || 0} />
            </span>
        </li>

        <li>
            <span>Internal Rate of Return</span>
            <span class="float-right flex-auto font-semibold"
                >{formatPercent.format(project.projectValue?.calculated?.internalRateOfReturn || 0)}</span
            >
        </li>

        <li>
            <span>Funding Source</span>
            <span class="float-right flex-auto font-semibold"
                >{project.projectValue?.fundingSource}</span
            >
        </li>

        <li>
            <span>Development Cost</span>
            <span class="float-right flex-auto font-semibold"
                >{formatCurrency.format(project.projectCost?.calculated?.initialCost || 0)}</span
            >
        </li>

        <li>
            <span>Development Hours</span>
            <span class="float-right flex-auto font-semibold"
                >{project.projectCost?.calculated?.hourEstimate}</span
            >
        </li>

        <li>
            <span>Blended Hourly Rate</span>
            <span class="float-right flex-auto font-semibold"
                >{formatCurrency.format(project.projectCost?.blendedRate || 0)}</span
            >
        </li>
    </ul>
</div>


<div class="col-span-1">
    <!-- col span 1 -->
<SectionHeading>Team</SectionHeading>

<div class="font-medium text-gray-900 dark:text-white">Owner</div>
<UserCard user={project.projectBasics.owner as User} />
<Hr />
<ul class="-m-4 p-2 divide-y divide-gray-200 bg-white dark:divide-gray-700 dark:bg-gray-800">
    {#if project.projectDaci?.driver}
        {@render teamSection("Drivers", project.projectDaci?.driver as Resource[])}
    {/if}
    {#if project.projectDaci?.approver}
        {@render teamSection("Approvers", project.projectDaci?.approver as Resource[])}
    {/if}
    {#if project.projectDaci?.contributor}
        {@render teamSection("Contributors", project.projectDaci?.contributor as Resource[])}
    {/if}
    {#if project.projectDaci?.informed}
        {@render teamSection("Informed", project.projectDaci?.informed as Resource[])}
    {/if}
</ul>
</div>



<div class="col-span-2">
    <!-- col span 2 -->
    <SectionHeading>Features</SectionHeading>
    <Table>
        <TableHead>
          <TableHeadCell>Feature</TableHeadCell>
          <TableHeadCell>Priority</TableHeadCell>
          <TableHeadCell>Status</TableHeadCell>
        </TableHead>
        <TableBody tableBodyClass="divide-y">
        {#if project.projectFeatures}
            {#each project.projectFeatures as feature}
          <TableBodyRow>
            <TableBodyCell>
                {feature.name}
                <button id={'feat' + normalizeGUID(feature.id)}>
                    <QuestionCircleSolid class="w-3 h-3 ms-1.5" />
                    <span class="sr-only">Show information</span>
                </button>
                  <Popover triggeredBy={'#feat' + normalizeGUID(feature.id)} class="w-72 text-sm font-light text-gray-500 bg-white dark:bg-gray-800 dark:border-gray-600 dark:text-gray-400" placement="bottom-start">
                    <div class="p-3 space-y-2">
                      {feature.description}
                    </div>
                  </Popover>
            </TableBodyCell>
            <TableBodyCell><BadgeFeaturePriority priority={feature.priority} /></TableBodyCell>
            <TableBodyCell><BadgeFeatureStatus status={feature.status} /></TableBodyCell>
          </TableBodyRow>
          {/each}
        {/if}
        </TableBody>
    </Table>
</div>

{/await}

</div>

{#snippet teamSection(title:string, resources:Resource[])}
<li class="py-3 sm:py-3.5">
<div class="flex items-center justify-between">
<div class="flex min-w-0 items-center">
    <div class="ml-3">
        <p class="truncate font-medium text-gray-900 dark:text-white">
            {title}
        </p>
    </div>
    
</div>
<div class="inline-flex items-center text-base font-semibold text-gray-900 dark:text-white">
    <ResourceList
        size="md"
        maxSize={4}
        {resources}
    />
</div>
</div>
</li>
{/snippet}