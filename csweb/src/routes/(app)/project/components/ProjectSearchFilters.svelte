<script lang="ts">
    import { Input } from 'flowbite-svelte';
    import type { SelectOptionType } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
    import { CSMultiFilter } from '$lib/components';
    import { findAllResources } from '$lib/services/resource';
	import { CloseCircleOutline } from 'flowbite-svelte-icons';
	import { onMount } from 'svelte';
	import { getSetting, PROJECT_SEARCH_INPUT_FILTER, PROJECT_STATUS_FILTER, PROJECT_RESOURCE_ID_FILTER, setSetting } from '$lib/services/settings';

    let searchInput:string = $state("")
    let status:string[] = $state(["new", "draft", "approved", "proposed", "inflight", "scheduled"])
    let resourceID:string = $state("")

    let clearHandle:any

    let statusOpts = [
        { value: "new", name: "New", checked: false},
        { value: "draft", name: "Draft", checked: false},
        { value: "proposed", name: "Proposed", checked: false},
        { value: "approved", name: "Approved", checked: false},
        { value: "rejected", name: "Rejected", checked: false},
        { value: "backlogged", name: "Backlogged", checked: false},
        { value: "scheduled", name: "Scheduled", checked: false},
        { value: "inflight", name: "In-flight", checked: false},
        { value: "complete", name: "Complete", checked: false},
        { value: "deferred", name: "Deferred", checked: false},
        { value: "abandoned", name: "Abandoned", checked: false},
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
        status = ["new", "draft", "approved", "proposed", "inflight", "scheduled"]
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
        status = getSetting(PROJECT_STATUS_FILTER, ["new", "draft", "approved", "proposed", "inflight", "scheduled"])
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

