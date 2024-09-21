<script lang="ts">
    import { page } from '$app/stores';
    import { ButtonGroup, Button, Card } from "flowbite-svelte";
    import { getProject } from "$lib/services/project";
    import { ProjectActionBar, ProjectValueChart, DeleteProject } from "../../components";
    import { UserList, UserDisplay, SectionHeading, PageHeading, MoneyDisplay } from "$lib/components";
    import { formatDate, formatCurrency, formatPercent } from "$lib/utils/format"
	//import { CommentList } from "../../components";
    import { AdjustmentsHorizontalSolid, TrashBinOutline } from 'flowbite-svelte-icons';
    import { type Project } from '$lib/graphql/generated/sdk';

    const id = $page.params.id;

    let project:Project = $state({} as Project)

    const loadPage = async () => {
		getProject(id).then((p) => {
            project = p
            return p
        });
	};

	loadPage();
</script>


{#await loadPage}
    <div>Loading...</div>
{:then promiseData}
<ProjectActionBar pageDetail={project.projectBasics?.name}>
    <ButtonGroup>
        <Button href="/project/workbook/{id}">
            <AdjustmentsHorizontalSolid class="w-3 h-3 mr-2" />
            Workbook
        </Button>
        <DeleteProject id={id || ""} name={project.projectBasics?.name}>
            <TrashBinOutline class="w-3 h-3 mr-2" />
            Delete
        </DeleteProject>
    </ButtonGroup>
</ProjectActionBar>

<PageHeading title={project.projectBasics?.name} />

<div class="grid grid-cols-6 gap-4">
    <div class="col-span-2">
    <Card size="lg"> 
        <SectionHeading>Financials</SectionHeading>

        <div class="p-4 mb-6">
            <ProjectValueChart project={project}></ProjectValueChart>
        </div>
        <ul class="list mb-6">
            <li>
                <span>Net Present Value</span>
                <span class="flex-auto float-right font-semibold">
                    <MoneyDisplay amount={project.projectValue?.netPresentValue || 0} />
                </span>
            </li>

            <li>
                <span>Internal Rate of Return</span>
                <span class="flex-auto float-right font-semibold">{formatPercent.format(project.projectValue?.internalRateOfReturn || 0)}</span>
            </li>

            <li>
                <span>Funding Source</span>
                <span class="flex-auto float-right font-semibold">{project.projectValue?.fundingSource}</span>
            </li>

            <li>
                <span>Development Cost</span>
                <span class="flex-auto float-right font-semibold">{formatCurrency.format(project.projectCost?.initialCost || 0)}</span>
            </li>

            <li>
                <span>Development Hours</span>
                <span class="flex-auto float-right font-semibold">{project.projectCost?.hourEstimate}</span>
            </li>

            <li>
                <span>Blended Hourly Rate</span>
                <span class="flex-auto float-right font-semibold">{formatCurrency.format(project.projectCost?.blendedRate || 0)}</span>
            </li>
        </ul>

        <SectionHeading>Executive Summary</SectionHeading>
        <p class="mb-6 text-xs">{project.projectBasics?.description}</p>


        <SectionHeading>Details</SectionHeading>
        <ul class="list mb-6">
            <li>
                <span>Owner</span>
                <span class="flex-auto float-right font-semibold">
                    <UserDisplay id={project.projectBasics?.ownerID + ""} />
                </span>
            </li>

            <li>
                <span>Created</span>
                <span class="flex-auto float-right font-semibold">{formatDate(project.createdAt)}</span>
            </li>

            <li>
                <span>Status</span>
                <span class="flex-auto float-right font-semibold">{project.projectBasics?.status}</span>
            </li>
        </ul>

        <SectionHeading>Team</SectionHeading>

        <ul class="list">
            {#if project.projectDaci?.driver}
            <li>
                <div>Drivers</div>
                <div class="text-left font-semibold px-4 pb-2"><UserList  size="md" maxSize={4} resourceIDs={project.projectDaci?.driver as string[]} /></div>
            </li>
            {/if}
            {#if project.projectDaci?.approver}
            <li>
                <div>Approvers</div>
                <div class="text-left font-semibold px-4 pb-2"><UserList size="md" maxSize={4} resourceIDs={project.projectDaci?.approver as string[]} /></div>
            </li>
            {/if}
            {#if project.projectDaci?.contributor}
            <li>
                <div>Contributors</div>
                <div class="text-left font-semibold px-4 pb-2"><UserList size="md" maxSize={4} resourceIDs={project.projectDaci?.contributor as string[]} /></div>
            </li>
            {/if}
            {#if project.projectDaci?.informed}
            <li>
                <div>Informed</div>
                <div class="text-left font-semibold px-4 pb-2"><UserList size="md" maxSize={4} resourceIDs={project.projectDaci?.informed as string[]} /></div>
            </li>
            {/if}
        </ul>
    </Card>
    </div>

    <div class="col-span-4">
    <Card size="xl">
        <SectionHeading>Comments / Activity</SectionHeading>
        <!-- <CommentList projectID={id} /> -->
    </Card>
    </div>
</div>
{/await}