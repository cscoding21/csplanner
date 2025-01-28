<script lang="ts">
	import type { ProjectActivity, Resource, Schedule } from "$lib/graphql/generated/sdk";
    import { BadgeProjectStatus, ProjectScheduleCell, ProjectStatus, type Week } from ".";
	import { calculateProjectSchedule } from "$lib/services/project";
    import { getScheduledProjectFromPortfolio } from "$lib/services/portfolio";
    import { addToast } from "$lib/stores/toasts";
	import { formatDate, formatDateNoYear, pluralize } from "$lib/utils/format";
    import { SectionHeading } from "$lib/components";
    import { Hr , Table, TableBody, TableHead, TableHeadCell, TableBodyCell, TableBodyRow, ButtonGroup, Button, Alert } from "flowbite-svelte";
	import { ResourceList } from "$lib/components";
    import { InfoCircleSolid } from "flowbite-svelte-icons";

    interface ScheduleRow {
        label: string
        resource?: Resource
        weeks: Week[]
    }

    interface Props {
		id: string;
		startDate: Date;
        update: Function;
	}
    let { id, startDate, update }: Props = $props();

    let result:Schedule = $state({} as Schedule)
    let scheduleTable = $state({header: [] as string[], body:[] as ScheduleRow[]})


    const getScheduleTableByResource = (r: Schedule) => {
        let table = {header: [] as string[], body: [] as ScheduleRow[]}
        let head = []
        let resourceSet = new Set<Resource>()

        head.push("Resources")

        if (!r.projectActivityWeeks) {
            return {}
        }

        for (let i = 0; i < r.projectActivityWeeks.length; i++) {
            const week = r.projectActivityWeeks[i]
            head.push(formatDateNoYear(week.end))

            if (week.activities) {
                for (let j = 0; j < week.activities?.length; j++) {
                    const act = week.activities[j]
                    resourceSet.add(act.resource as Resource)
                }
            }
        }

        table.header = head as string[]
        table.body = [] as ScheduleRow[]
        const resourceArray = [...resourceSet]

        for (let x = 0; x < resourceArray.length; x++) {
            const thisResource = resourceArray[x] as Resource
            let row:ScheduleRow = {label: thisResource.name, resource: thisResource, weeks: [] as Week[]}

            for (let i = 0; i < r.projectActivityWeeks.length; i++) {
                const week = r.projectActivityWeeks[i]
                let weeksActivities = { 
                    showTasks: true,
                    weekEnding: formatDate(week.end),
                    activities: [] as ProjectActivity[] 
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisResource.name == activity.resource?.name) {
                            weeksActivities.activities.push(activity)
                        }
                    }
                }
            
                row.weeks.push(weeksActivities)
            }
            
            table.body.push(row)
        }

        scheduleTable = table
        return scheduleTable
    }


    const getScheduleTableByTask = (r: Schedule) => {
        let table = {header: [] as string[], body: [] as ScheduleRow[]}
        let taskSet = new Set<string>()

        table.header.push("Task")

        if (!r.projectActivityWeeks) {
            return {}
        }

        for (let i = 0; i < r.projectActivityWeeks.length; i++) {
            const week = r.projectActivityWeeks[i]
            table.header.push(formatDateNoYear(week.end))

            if (week.activities) {
                for (let j = 0; j < week.activities?.length; j++) {
                    const act = week.activities[j]
                    taskSet.add(act.taskName as string)
                }
            }
        }

        const taskArray = [...taskSet]

        for (let x = 0; x < taskArray.length; x++) {
            const thisTask = taskArray[x] as string
            let row:ScheduleRow = {label: thisTask, resource: undefined, weeks: [] as Week[]}

            for (let i = 0; i < r.projectActivityWeeks.length; i++) {
                const week = r.projectActivityWeeks[i]
                let weeksActivities = { 
                    showTasks: false,
                    weekEnding: formatDate(week.end), 
                    activities: [] as ProjectActivity[] 
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisTask == activity.taskName) {
                            weeksActivities.activities.push(activity)
                        }
                    }
                }
            
                row.weeks.push(weeksActivities)
            }
            
            table.body.push(row)
        }

        scheduleTable = table
        return scheduleTable
    }


    const load = async ():Promise<Schedule> => {
        return await getScheduledProjectFromPortfolio(id).then(s => {
            console.log("found schedule from portfolio")

            if (s.project) {
                return s
            }

            return calculateProjectSchedule(id, startDate)
                .then((s) => {
                    console.log("found schedule from calculate")
                    return s.schedule
                })
                .catch((err) => {
                    addToast({
                        message: 'Error loading project schedule (ProjectSchedule): ' + err,
                        dismissible: true,
                        type: 'error'
                    });

                    return err
                });

        })
        .catch(err => {
            console.log("ERROR", err)
	    });
    }

    const setView = (view:string) => {
        if (view === "resource") {
            getScheduleTableByResource(result)
        } else {
            getScheduleTableByTask(result)
        }
    }

    const loadPage = async () => {
        load().then(s => {
            result = s

            getScheduleTableByResource(s)
        });
	};

</script>


{#await loadPage()}
	loading...
{:then promiseData} 
	
{#if result}


{#if result.project}
<SectionHeading>
    Schedule: {result.project.projectBasics.name}
    <span class="float-right"><BadgeProjectStatus status={result.project.projectBasics.status} /></span>
</SectionHeading>
{/if}

<ProjectStatus project={result.project} schedule={result} update={() => console.log("update")} />

{#if result.exceptions}
    <ul class="">
    {#each result.exceptions as ex}
        <li class="list-disc text-left ml-4">{ex.message}</li>
    {/each}
    </ul>
{/if}
{#if result.projectActivityWeeks && result.projectActivityWeeks.length > 0}
    {#if result.end}
    <h1>{formatDate(result.begin)} - {formatDate(result.end)}</h1>
    {#if result.project.projectBasics.startDate}
    <h2>Scheduled length: {result.projectActivityWeeks.length + " " + pluralize("week", result.projectActivityWeeks.length)}</h2>
    {:else}
    <h2>Unadjusted length: {result.projectActivityWeeks.length + " " + pluralize("week", result.projectActivityWeeks.length)}</h2>
    {/if}
    {/if}
    
    <ButtonGroup class="*:!ring-primary-700 mt-2">
        <Button onclick={() => { setView("resource") }}>Resource view</Button>
        <Button onclick={() => { setView("task") }}>Task view</Button>
    </ButtonGroup>
    <Hr />
    
    <Table>
        <TableHead>
            {#each scheduleTable.header as head, index}
            {#if index === 0}
                <TableHeadCell>{head}</TableHeadCell>
            {:else}
                <TableHeadCell class="text-center">{head}</TableHeadCell>
            {/if}
            {/each}
        </TableHead>
        <TableBody tableBodyClass="divide-y">
            {#each scheduleTable.body as row}
                <TableBodyRow>
                <TableBodyCell>
                    <div>
                        <span class="float-left mr-2">

                    {#if row.resource}
                    <ResourceList resources={[row.resource]} size="sm" maxSize={1} />
                    {/if}
                </span>
                    {row.label}
                </div>
                </TableBodyCell>
                {#each row.weeks as week}
                <TableBodyCell tdClass="text-center">
                    <ProjectScheduleCell week={week} />
                </TableBodyCell>
                {/each}
                </TableBodyRow>
            {/each}
        </TableBody>
    </Table>
{/if}
{/if}

{/await}

