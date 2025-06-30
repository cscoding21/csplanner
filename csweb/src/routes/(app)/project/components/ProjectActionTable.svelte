<script lang="ts">
    import type { Portfolio, ProjectEnvelope, Resource, Schedule } from "$lib/graphql/generated/sdk";
	import { Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
	import { ResourceList } from "$lib/components";
	import { ProjectStatusBrief, BadgeProjectStatus } from ".";
	import { ProjectStatusApproved, ProjectStatusDraft, ProjectStatusInflight, ProjectStatusNew, ProjectStatusProposed, ProjectStatusScheduled } from "$lib/services/project";
	import { extractProjectScheduleFromPortfolio } from "$lib/services/portfolio";

    interface Props {
        projects:ProjectEnvelope[]
        portfolio:Portfolio
    }
    let { projects, portfolio }:Props = $props()

    let projectFilter = $state("all")
    let displayProjects:ProjectEnvelope[] = $derived(projects?.filter(p => projectFilter === "all" || p.data.projectStatusBlock.status === projectFilter))
</script>


<div class="block w-full items-center justify-between border-t py-3 dark:border-gray-800 sm:flex border-gray-200 mt-4 mb-2">
    <div class="flex flex-wrap gap-4">
        <div class="flex items-center text-sm font-medium text-gray-900 dark:text-white">Status:</div>
        <div class="flex items-center">
            <input
                id="sort-role"
                type="radio"
                value="all"
                name="status"
                checked={true}
                onclick={() => { projectFilter = "all" } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="sort-role" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">All</label>
        </div>
        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusNew}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusNew } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">New</label>
        </div>

        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusDraft}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusDraft } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Draft</label>
        </div>

        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusProposed}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusProposed } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Proposed</label>
        </div>

        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusApproved}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusApproved } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Approved</label>
        </div>

        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusScheduled}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusScheduled } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Scheduled</label>
        </div>

        <div class="flex items-center">
            <input
                id="all-users"
                type="radio"
                value={ProjectStatusInflight}
                name="status"
                checked={false}
                onclick={() => { projectFilter = ProjectStatusInflight } }
                class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">In-Flight</label>
        </div>
        
    </div>
</div>

<Table>
    <TableHead>
        <TableHeadCell>Project</TableHeadCell>
        <TableHeadCell>Status</TableHeadCell>
        <TableHeadCell>Team</TableHeadCell>
        <TableHeadCell></TableHeadCell>
    </TableHead>
    <TableBody class="divide-y">
        {#each displayProjects as project}
            <TableBodyRow>
            <TableBodyCell>
            <a href="/project/detail/{project?.meta.id}#snapshot" class="text-md dark:text-gray-300">
                {project.data.projectBasics.name}
            </a>
            </TableBodyCell>
            <TableBodyCell><BadgeProjectStatus status={project?.data?.projectStatusBlock.status} /></TableBodyCell>
            <TableBodyCell>
                {#if project?.data.calculated?.teamMembers}
                <ResourceList size="md" maxSize={4} resources={project?.data.calculated?.teamMembers as Resource[]} />
                {:else}
                    No team yet
                {/if}
            </TableBodyCell>
            <TableBodyCell>
                <ProjectStatusBrief project={project?.data} {portfolio} />
            </TableBodyCell>
            </TableBodyRow>
        {/each}
    </TableBody>
</Table>



