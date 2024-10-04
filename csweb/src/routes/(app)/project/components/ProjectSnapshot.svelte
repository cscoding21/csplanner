<script lang="ts">
    import { 
        SectionHeading,
        MoneyDisplay,
        ResourceList,
    } from "$lib/components";
    import { 
        ProjectValueChart
    } from '../components'
    import type { Project, Resource } from '$lib/graphql/generated/sdk'
    import { formatPercent, formatCurrency, formatDate } from "$lib/utils/format";
    import { getProject } from "$lib/services/project";
    import { getDefaultProject } from "$lib/forms/project.validation";

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let project:Project = $state(getDefaultProject() as Project)

    const load = async (): Promise<Project> => {
		return await getProject(id)
			.then((proj) => {

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


{#await loadPage()}
	Loading...
{:then promiseData}

{#if project}
<SectionHeading>Financials: {project.projectBasics.name}</SectionHeading>
{/if}

<div class="mb-6">
    <ProjectValueChart {project}></ProjectValueChart>
</div>
<ul class="list mb-6">
    <li>
        <span>Net Present Value</span>
        <span class="float-right flex-auto font-semibold">
            <MoneyDisplay amount={project.projectValue?.netPresentValue || 0} />
        </span>
    </li>

    <li>
        <span>Internal Rate of Return</span>
        <span class="float-right flex-auto font-semibold"
            >{formatPercent.format(project.projectValue?.internalRateOfReturn || 0)}</span
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
            >{formatCurrency.format(project.projectCost?.initialCost || 0)}</span
        >
    </li>

    <li>
        <span>Development Hours</span>
        <span class="float-right flex-auto font-semibold"
            >{project.projectCost?.hourEstimate}</span
        >
    </li>

    <li>
        <span>Blended Hourly Rate</span>
        <span class="float-right flex-auto font-semibold"
            >{formatCurrency.format(project.projectCost?.blendedRate || 0)}</span
        >
    </li>
</ul>

<SectionHeading>Executive Summary</SectionHeading>
<p class="mb-6 text-xs">{project.projectBasics?.description}</p>

<SectionHeading>Details</SectionHeading>
<ul class="list mb-6">
    <li>
        <span>Owner</span>
        <span class="float-right flex-auto font-semibold">
            {#if project.projectBasics?.ownerID}
                <!-- <UserDisplay id={project.projectBasics?.ownerID as string} /> -->
            {/if}
        </span>
    </li>

    <li>
        <span>Created</span>
        <span class="float-right flex-auto font-semibold">{formatDate(project.createdAt)}</span>
    </li>

    <li>
        <span>Status</span>
        <span class="float-right flex-auto font-semibold">{project.projectBasics?.status}</span>
    </li>
</ul>

<SectionHeading>Team</SectionHeading>
<ul class="list">
    {#if project.projectDaci?.driver}
        <li>
            <div>Drivers</div>
            <div class="px-4 pb-2 text-left font-semibold">
                <ResourceList
                    size="md"
                    maxSize={4}
                    resources={project.projectDaci?.driver as Resource[]}
                />
            </div>
        </li>
    {/if}
    {#if project.projectDaci?.approver}
        <li>
            <div>Approvers</div>
            <div class="px-4 pb-2 text-left font-semibold">
                <ResourceList
                    size="md"
                    maxSize={4}
                    resources={project.projectDaci?.approver as Resource[]}
                />
            </div>
        </li>
    {/if}
    {#if project.projectDaci?.contributor}
        <li>
            <div>Contributors</div>
            <div class="px-4 pb-2 text-left font-semibold">
                <ResourceList
                    size="md"
                    maxSize={4}
                    resources={project.projectDaci?.contributor as Resource[]}
                />
            </div>
        </li>
    {/if}
    {#if project.projectDaci?.informed}
        <li>
            <div>Informed</div>
            <div class="px-4 pb-2 text-left font-semibold">
                <ResourceList
                    size="md"
                    maxSize={4}
                    resources={project.projectDaci?.informed as Resource[]}
                />
            </div>
        </li>
    {/if}
</ul>

{/await}