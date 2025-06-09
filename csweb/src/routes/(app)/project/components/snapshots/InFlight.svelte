<script lang="ts">
	import { SectionSubHeading } from "$lib/components";
	import type { Project, ProjectMilestoneTask, Resource, Schedule, UpdateProjectTemplateTask } from "$lib/graphql/generated/sdk";
	import { getScheduledProjectFromPortfolio } from "$lib/services/portfolio";
	import { Accordion, AccordionItem, Table, TableBody, TableBodyCell, TableBodyRow, TableHead, TableHeadCell } from "flowbite-svelte";
	import BasicsSummary from "./BasicsSummary.svelte";
	import ScheduleSummary from "./ScheduleSummary.svelte";
	import UpdateStatusButtons from "./UpdateStatusButtons.svelte";
	import { CheckCircleSolid, CirclePlusOutline } from "flowbite-svelte-icons";
	import FinancialSummary from "./FinancialSummary.svelte";
	import ResourceList from "$lib/components/widgets/ResourceList.svelte";
	import { updateProjectTaskStatus } from "$lib/services/project";
	import { addToast } from "$lib/stores/toasts";
	import { callIf } from "$lib/utils/helpers";

    interface Props {
        project:Project
        update:Function
    }
    let { project, update }:Props = $props()

    const refresh = async (): Promise<Schedule> => {
		const res = await getScheduledProjectFromPortfolio(project.id as string);

        projectSchedule = res
        console.log("projectSchedule", projectSchedule)
		return res;
	};

    const setTaskStatus = async (milestoneID:string, taskID:string, status:string) => {
        updateProjectTaskStatus(project.id as string, milestoneID, taskID, status).then(res => {
            if (res.status?.success) {
                    addToast({
                        message: 'Project task status updated successfully',
                        dismissible: true,
                        type: 'success'
                    });

                    callIf(update)
                    refresh()
                } else {
                    addToast({
                        message: 'Error updating project task status: ' + res.status?.message,
                        dismissible: true,
                        type: 'error'
                    });
                }
            });
    }
    
    const load = async () => {
        refresh()
    }

    let projectSchedule:Schedule|undefined = $state()
</script>

{#if project}

<div class="grid grid-cols-3 gap-16">
    <div>
        <BasicsSummary {project} />
    </div>

    <div class="col-span-2">
        <SectionSubHeading>Summary</SectionSubHeading>

        <!-- <FinancialSummary {project} abridged={true} /> -->
        
        {#await load()}
            Loading...
        {:then promiseData} 
            <ScheduleSummary schedule={projectSchedule as Schedule} />    
        {/await}
         


        <SectionSubHeading>Tasks</SectionSubHeading>
        {#if project.projectMilestones && project.projectMilestones.length > 0}

        <Accordion>
        {#each project.projectMilestones as m, mi}
        <AccordionItem open={m.calculated?.isInFlight || true}>
            {#snippet header()}
            <div class="grid grid-cols-2 gap-4">
                <div>{m.phase.name}</div>
                <div>{m.phase.description}</div>
<!-- 
                <div>{m.calculated?.completedTasks}</div>
                <div>{m.calculated?.hoursRemaining}</div>

                <div>{m.calculated?.isInFlight}</div>
                <div>{m.calculated?.isComplete}</div> -->
            </div>
            
            {/snippet}

            <Table>
                <TableHead>
                    <TableHeadCell>Owner(s)</TableHeadCell>
                    <TableHeadCell>Task</TableHeadCell>
                    <TableHeadCell class="text-center">Hours</TableHeadCell>
                    <TableHeadCell class="text-center">Status</TableHeadCell>
                </TableHead>
                <TableBody>

                {#each m.tasks as t, ti}
                <TableBodyRow>
                <TableBodyCell>
                    <ResourceList size="sm" maxSize={3} resources={t.resources as Resource[]} />
                </TableBodyCell>

                <TableBodyCell>
                    {#if t.skills && t.skills.length > 0}
                    <b>{t.skills[0].name}</b> :
                    {/if}
                    {t.name}
                </TableBodyCell>
                
                <TableBodyCell class="text-center">
                    {t.calculated?.actualizedHoursToComplete}
                </TableBodyCell>
                
                <TableBodyCell class="text-center">
                    {#if project.projectMilestones[mi].tasks[ti].status === "done"}
                    <button onclick={() => setTaskStatus(m.id, t.id, "accepted")} title="Mark active">
                        <CheckCircleSolid color="green" />
                    </button>
                    {:else}
                        <button onclick={() => setTaskStatus(m.id, t.id, "done")} title="Mark done">
                            <CirclePlusOutline color="yellow" />
                        </button>
                    {/if}
                </TableBodyCell>
                </TableBodyRow>
                {/each}

                </TableBody>
            </Table>
        </AccordionItem>
            
        {/each}
        </Accordion>

        {/if}


        <div class="text-center mt-8">
			<UpdateStatusButtons {project} displayStates={["deferred"]} />
		</div>
    </div>
</div>

{/if}