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
    import { getScheduleTableByResource, getScheduleTableByTask, type ProjectScheduleTable } from "$lib/services/schedule"
	import ScheduleTable from "$lib/components/widgets/ScheduleTable.svelte";

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
    let scheduleTable = $state({} as ProjectScheduleTable|undefined)
    let startDate = $state(new Date()) 
    let view = $state("task")


    function refresh() {
		console.log('refreshing')
		load().then(p => {
			result = p

			update();
		})
	}


    const getTableByResource = (r: Schedule) => {
        scheduleTable = getScheduleTableByResource(r)
        return scheduleTable
    }


    const getTableByTask = (r: Schedule) => {
        scheduleTable = getScheduleTableByTask(r)
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

    const setView = (v:string) => {
       view = v
    }

    const loadPage = async () => {
        load().then(s => {
            result = s

            getTableByResource(s)
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
        {#snippet icon()}<InfoCircleSolid class="h-5 w-5" />{/snippet}
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

    <ScheduleTable schedule={result} view={view as "task"|"resource"|undefined} />
{/if}
{/if}

{/await}

