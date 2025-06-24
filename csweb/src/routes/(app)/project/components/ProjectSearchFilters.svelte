<script lang="ts">
    import { Input } from 'flowbite-svelte';
    import type { SelectOptionType } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
    import { CSMultiFilter } from '$lib/components';
    import { findAllResources } from '$lib/services/resource';
	import { CloseCircleOutline } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { getSetting, PROJECT_SEARCH_INPUT_FILTER, PROJECT_STATUS_FILTER, PROJECT_RESOURCE_ID_FILTER, setSetting } from '$lib/services/settings';
	import { ProjectStatusNew, ProjectStatusDraft, ProjectStatusProposed, ProjectStatusApproved, ProjectStatusRejected, ProjectStatusBacklogged, ProjectStatusScheduled, ProjectStatusInflight, ProjectStatusComplete, ProjectStatusDeferred, ProjectStatusAbandoned } from '$lib/services/project';

    let searchInput:string = $state("")
    let status:string[] = $state([ProjectStatusNew, ProjectStatusDraft, ProjectStatusApproved, ProjectStatusProposed, ProjectStatusInflight, ProjectStatusScheduled])
    let resourceID:string = $state("")

    let clearHandle:any

    let statusOpts = [
        { value: ProjectStatusNew, name: "New", checked: false},
        { value: ProjectStatusDraft, name: "Draft", checked: false},
        { value: ProjectStatusProposed, name: "Proposed", checked: false},
        { value: ProjectStatusApproved, name: "Approved", checked: false},
        { value: ProjectStatusRejected, name: "Rejected", checked: false},
        { value: ProjectStatusBacklogged, name: "Backlogged", checked: false},
        { value: ProjectStatusScheduled, name: "Scheduled", checked: false},
        { value: ProjectStatusInflight, name: "In-flight", checked: false},
        { value: ProjectStatusComplete, name: "Complete", checked: false},
        { value: ProjectStatusDeferred, name: "Deferred", checked: false},
        { value: ProjectStatusAbandoned, name: "Abandoned", checked: false},
    ]

    let resourceOpts:SelectOptionType<string>[] = $state([])

    const statusChange = (e:any) => {
        status = e

        setSetting(PROJECT_STATUS_FILTER, status)

        change(getFilters())
    }

    const resourceChange = (e:any) => {
        resourceID = e.join("")

        setSetting(PROJECT_RESOURCE_ID_FILTER, resourceID)

        change(getFilters())
    }

    const searchChange = () => {
        clearTimeout(clearHandle)

        clearHandle = setTimeout(() => {
            setSetting(PROJECT_SEARCH_INPUT_FILTER, searchInput)

            change(getFilters())
        }, 500)
        
    }

    const resetFilters = () => {
        searchInput = ""
        status = [ProjectStatusNew, ProjectStatusDraft, ProjectStatusApproved, ProjectStatusProposed, ProjectStatusInflight, ProjectStatusScheduled]
        resourceID = ""

        setSetting(PROJECT_SEARCH_INPUT_FILTER, searchInput)
        setSetting(PROJECT_RESOURCE_ID_FILTER, resourceID)
        setSetting(PROJECT_STATUS_FILTER, status)

        change(getFilters())
    }

    const getFilters = ():InputFilters => {
        let out:InputFilters = { filters: [] as InputFilter[] }
        let filterArray:InputFilter[] = [] 

		const kwfil = { key: 'data.basics.name', value: searchInput, operation: 'fl' }
        filterArray = [...filterArray, kwfil]

        if (status && status.length > 0) {
            const stfil = { key: 'data.status.status', value: status.join(","), operation: 'in' }
            filterArray = [...filterArray, stfil]
        }

        if (resourceID) {
            const rsfil = { key: 'data.resourceID', value: resourceID, operation: 'custom', customName: 'resource_in_project' }
            filterArray = [...filterArray, rsfil]
        }

        out.filters = filterArray
        return out
    }

    interface Props {
        change: Function
    }
    let { 
        change = $bindable() 
    }:Props = $props()


    const loadPage = async () => {
		findAllResources()
			.then((r) => {
				resourceOpts = [...resourceOpts, ...r.results?.map((r) => ({
					name: r.data?.name,
					value: r.meta?.id as string
				})) as SelectOptionType<string>[]];
			});
	};

    onMount(() => {
        searchInput = getSetting(PROJECT_SEARCH_INPUT_FILTER, "")
        status = getSetting(PROJECT_STATUS_FILTER, [ProjectStatusNew, ProjectStatusDraft, ProjectStatusApproved, ProjectStatusProposed, ProjectStatusInflight, ProjectStatusScheduled])
        resourceID = getSetting(PROJECT_RESOURCE_ID_FILTER, "")

        change(getFilters())
    })
</script>

<div class="flex">
    <div class="mr-4 text-nowrap">
        <Input slot="search" size="md" class="w-80" placeholder="Filter by name" bind:value={searchInput} onkeyup={searchChange} />
    </div>

    <div class="mr-4 text-nowrap">
        <CSMultiFilter filterOpts={statusOpts} change={statusChange} filterValue={status} filterName="Status" isMulti={true} />
    </div>

    {#await loadPage()}
        ...
    {:then}
    <div class="mr-4 text-nowrap">
        <CSMultiFilter filterOpts={resourceOpts} change={resourceChange} filterValue={[resourceID]} filterName="Assigned Resource" isMulti={false} />
    </div>
    {/await}

    <div class="content-right justify-end w-full">
        <button class="float-right" onclick={resetFilters} title="Clear all filters"><CloseCircleOutline size="lg" /></button>
    </div>
</div>

