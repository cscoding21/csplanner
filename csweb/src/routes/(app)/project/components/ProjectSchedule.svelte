<script lang="ts">
	import type { ProjectActivity, ProjectScheduleResult } from "$lib/graphql/generated/sdk";
    import { ProjectScheduleCell, type Week } from ".";
	import { calculateProjectSchedule } from "$lib/services/project";
    import { addToast } from "$lib/stores/toasts";
	import { formatDate } from "$lib/utils/format";
    import { Hr , Table, TableBody, TableHead, TableHeadCell, TableBodyCell, TableBodyRow } from "flowbite-svelte";

    interface ScheduleRow {
        resource: string
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

    const getScheduleTable = (r: ProjectScheduleResult) => {
        let table = {header: [] as string[], body: [] as ScheduleRow[]}
        let head = []
        let resourceSet = new Set()

        head.push("Resources")

        for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
            const week = r.schedule.projectActivityWeeks[i]
            head.push(formatDate(week.end))

            if (week.activities) {
                for (let j = 0; j < week.activities?.length; j++) {
                    const act = week.activities[j]
                    resourceSet.add(act.resourceName)
                }
            }
        }

        table.header = head as string[]
        table.body = [] as ScheduleRow[]
        const resourceArray = [...resourceSet]

        for (let x = 0; x < resourceArray.length; x++) {
            const thisResource = resourceArray[x] as string
            let row:ScheduleRow = {resource: thisResource, weeks: [] as Week[]}

            for (let i = 0; i < r.schedule.projectActivityWeeks.length; i++) {
                const week = r.schedule.projectActivityWeeks[i]
                let weeksActivities = { 
                    weekEnding: formatDate(week.end), 
                    resourceName: thisResource, 
                    activities: [] as ProjectActivity[] 
                }

                if (week.activities) {
                    for(let j = 0; j < week.activities.length; j++) {
                        const activity = week.activities[j]
                        
                        if (thisResource == activity.resourceName) {
                            weeksActivities.activities.push(activity)
                        }
                    }
                }
            
                row.weeks.push(weeksActivities)
            }
            
            table.body.push(row)
        }

        console.log(table)

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

    const loadPage = async () => {
        load().then(s => {
            result = s

            getScheduleTable(s)
        });
	};

</script>


{#await loadPage()}
	loading...
{:then promiseData} 
	
{#if result && result.schedule}
    <h1>Name: {result.schedule.projectName} {formatDate(result.schedule.begin)} - {formatDate(result.schedule.end)}</h1>
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
                <TableBodyCell>{row.resource}</TableBodyCell>
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

{/await}

