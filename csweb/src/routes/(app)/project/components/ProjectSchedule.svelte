<script lang="ts">
	import type { ProjectActivity, ProjectScheduleResult, Resource } from "$lib/graphql/generated/sdk";
    import { ProjectScheduleCell, type Week } from ".";
	import { calculateProjectSchedule } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
	import { formatDate, pluralize } from "$lib/utils/format";
    import { SectionHeading } from "$lib/components";
    import { Hr , Table, TableBody, TableHead, TableHeadCell, TableBodyCell, TableBodyRow, ButtonGroup, Button } from "flowbite-svelte";
	import { ResourceList } from "$lib/components";

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

    let result:ProjectScheduleResult = $state({} as ProjectScheduleResult)
    let scheduleTable = $state({header: [] as string[], body:[] as ScheduleRow[]})


    const getScheduleTableByResource = (r: ProjectScheduleResult) => {
        let table = {header: [] as string[], body: [] as ScheduleRow[]}
        let head = []
        let resourceSet = new Set<Resource>()

        head.push("Resources")

        if (!r.schedule.projectActivityWeeks) {
            return {}
        }

        for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
            const week = r.schedule.projectActivityWeeks[i]
            head.push(formatDate(week.end))

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

            for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
                const week = r.schedule.projectActivityWeeks[i]
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


    const getScheduleTableByTask = (r: ProjectScheduleResult) => {
        let table = {header: [] as string[], body: [] as ScheduleRow[]}
        let taskSet = new Set<string>()

        table.header.push("Task")

        if (!r.schedule.projectActivityWeeks) {
            return {}
        }

        for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
            const week = r.schedule.projectActivityWeeks[i]
            table.header.push(formatDate(week.end))

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

            for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
                const week = r.schedule.projectActivityWeeks[i]
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


    const load = async ():Promise<ProjectScheduleResult> => {
		return await calculateProjectSchedule(id, startDate)
			.then((s) => {
				return s
			})
			.catch((err) => {
				addToast({
					message: 'Error loading project schedule (ProjectSchedule): ' + err,
					dismissible: true,
					type: 'error'
				});

				return err
			});
	};

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
	
{#if result && result.schedule}

<SectionHeading>Schedule: {result.schedule.project.projectBasics.name}</SectionHeading>

{#if result.schedule.exceptions}
    <ul class="">
    {#each result.schedule.exceptions as ex}
        <li class="list-disc text-left ml-4">{ex.message}</li>
    {/each}
    </ul>
{/if}
{#if result.schedule.projectActivityWeeks && result.schedule.projectActivityWeeks.length > 0}
    {#if result.schedule.end}
    <h1>{formatDate(result.schedule.begin)} - {formatDate(result.schedule.end)}</h1>
    <h2>Unadjusted length: {result.schedule.projectActivityWeeks.length + " " + pluralize("week", result.schedule.projectActivityWeeks.length)}</h2>
    {/if}
    
    <ButtonGroup class="*:!ring-primary-700 mt-2">
        <Button onclick={() => { setView("resource") }}>Resource view</Button>
        <Button onclick={() => { setView("task") }}>Task view</Button>
    </ButtonGroup>
    <Hr />
    
    <Table>
        <TableHead>
            {#each scheduleTable.header as head}
            <TableHeadCell>{head}</TableHeadCell>
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

