<script lang="ts">
	import type { ProjectActivity, Resource, Schedule } from "$lib/graphql/generated/sdk";
    import { BadgeProjectStatus, ProjectScheduleCell, ProjectStatusBanner, ProjectStartDateSet, ShowIfStatus, type Week, RiskLegend } from ".";
	import { calculateProjectSchedule } from "$lib/services/project";
    import { findWeekRisks, getScheduledProjectFromPortfolio } from "$lib/services/portfolio";
    import { addToast } from "$lib/stores/toasts";
	import { formatDate, formatDateNoYear, pluralize } from "$lib/utils/format";
    import { SectionHeading, SectionSubHeading } from "$lib/components";
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
        update: Function;
	}
    let { id, update }: Props = $props();

    let result:Schedule = $state({} as Schedule)
    let scheduleTable = $state({header: [] as string[], body:[] as ScheduleRow[]})
    let startDate = $state(new Date()) 


    function refresh() {
		console.log('refreshing')
		load().then(p => {
			result = p

			update();
		})
	}


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
                    activities: [] as ProjectActivity[], 
                    risks: [] as string[]
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisResource.name == activity.resource?.name) {
                            weeksActivities.activities.push(activity)
                            weeksActivities.risks = findWeekRisks(week.activities, activity.resourceID)
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
                    activities: [] as ProjectActivity[],
                    risks: [] as string[]
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                    
                        if (thisTask == activity.taskName) {
                            weeksActivities.activities.push(activity)
                        }
                    }

                    weeksActivities.risks = findWeekRisks(weeksActivities.activities, undefined)
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
                startDate = s.begin
                return s
            }

            return calculateProjectSchedule(id, startDate)
                .then((s) => {
                    console.log("found schedule from calculate")
                    result = s.schedule
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
    <span class="float-right"><BadgeProjectStatus status={result.project.projectStatusBlock?.status} /></span>
</SectionHeading>

<ProjectStatusBanner project={result.project} schedule={result} />

<ShowIfStatus scope={["approved", "scheduled", "inflight", "deferred", "abandoned"]} status={result.project.projectStatusBlock?.status}>
    <ProjectStartDateSet project={result.project} update={() => refresh()} />

    {#snippet elseRender()}
        <h3>{result.project.projectBasics.startDate}</h3>
	{/snippet}
</ShowIfStatus>
{/if}

{#if result.exceptions}
    <div class="my-4">
    <Alert class="items-start!" color="yellow" border>
        <span slot="icon">
          <InfoCircleSolid class="w-5 h-5" />
          <span class="sr-only">Warning</span>
        </span>
        <p class="font-medium">The following issues were found and could affect execution:</p>
        <ul class="mt-1.5 ms-4 list-disc list-inside">
            {#each result.exceptions as ex}
                <li>{ex.message}</li>
            {/each}
        </ul>
      </Alert> 
    </div>
{/if}

{#if result.projectActivityWeeks && result.projectActivityWeeks.length > 0}
    {#if result.end}
    <SectionSubHeading>{formatDate(result.begin)} - {formatDate(result.end)}</SectionSubHeading>
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
                    {#if row.resource}
                    <span class="float-left mr-2">
                            <ResourceList resources={[row.resource]} size="sm" maxSize={1} />
                    </span>      
                    <a href="/resource/detail/{row.resource.id}">{row.label}</a>
                    {:else}
                        {row.label} 
                    {/if}
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

    <div class="mt-4">
        <RiskLegend />
    </div>
{/if}
{/if}

{/await}

