<script lang="ts">
	import type { Resource, Schedule } from "$lib/graphql/generated/sdk";
    import { BadgeProjectStatus, ProjectStartDateSet, ShowIfStatus } from ".";
	import { calculateProjectSchedule, ProjectStatusAbandoned, ProjectStatusApproved, ProjectStatusDeferred, ProjectStatusInflight, ProjectStatusScheduled } from "$lib/services/project";
    import { getScheduledProjectFromPortfolio } from "$lib/services/portfolio";
    import { addToast } from "$lib/stores/toasts";
	import { formatDate, pluralize } from "$lib/utils/format";
    import { SectionHeading, SectionSubHeading } from "$lib/components";
    import { Alert } from "flowbite-svelte";
	import { InfoCircleSolid } from "flowbite-svelte-icons";
    import { getScheduleTableByResource, getScheduleTableByTask, type ProjectScheduleTable } from "$lib/services/schedule"
	import ScheduleTable from "$lib/components/widgets/ScheduleTable.svelte";
	import ScheduleSummary from "./snapshots/ScheduleSummary.svelte";

    interface ScheduleRow {
        label: string
        resource?: Resource
        weeks: []
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


<ShowIfStatus scope={[ProjectStatusScheduled]} status={result.project.projectStatusBlock?.status}>
    <ScheduleSummary schedule={result} />
</ShowIfStatus>

<ShowIfStatus scope={[ProjectStatusApproved, ProjectStatusScheduled, ProjectStatusInflight, ProjectStatusDeferred, ProjectStatusAbandoned]} status={result.project.projectStatusBlock?.status}>
    <ProjectStartDateSet project={result.project} update={() => refresh()} />

    {#snippet elseRender()}
        <h3>{result.project.projectBasics.startDate}</h3>
	{/snippet}
</ShowIfStatus>
{/if}

{#if result.exceptions && result.exceptions.length > 0}
    <div class="my-4">
    <Alert class="items-start!" border>
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
    <SectionSubHeading cssClass="my-2">{formatDate(result.begin)} - {formatDate(result.end)}</SectionSubHeading>
    {#if result.project.projectBasics.startDate}
    <h2>Scheduled length: {result.projectActivityWeeks.length + " " + pluralize("week", result.projectActivityWeeks.length)}</h2>
    {:else}
    <h2>Unadjusted length: {result.projectActivityWeeks.length + " " + pluralize("week", result.projectActivityWeeks.length)}</h2>
    {/if}
    {/if}


    <div class="block w-full items-center justify-between border-t py-3 dark:border-gray-800 sm:flex border-gray-200 mt-4">
        <div class="flex flex-wrap gap-4">
          <div class="flex items-center text-sm font-medium text-gray-900 dark:text-white">View by:</div>
          <div class="flex items-center">
            <input
              id="sort-role"
              type="radio"
              value=""
              name="show-only"
              checked={true}
              onclick={() => { setView("task") } }
              class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="sort-role" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Task</label>
          </div>
          <div class="flex items-center">
            <input
              id="all-users"
              type="radio"
              value=""
              name="show-only"
              checked={false}
              onclick={() => { setView("resource") } }
              class="h-4 w-4 border-gray-300 bg-gray-100 text-primary-600 focus:ring-2 focus:ring-primary-500 dark:border-gray-600 dark:bg-gray-700 dark:ring-offset-gray-800 dark:focus:ring-primary-600"
            />
            <label for="all-users" class="ml-2 text-sm font-medium text-gray-900 dark:text-gray-300">Resource</label>
          </div>
        </div>
    </div>

    <ScheduleTable schedule={result} view={view as "task"|"resource"|undefined} />
{/if}
{/if}

{/await}

