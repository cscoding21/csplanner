<script lang="ts">
    import { Search, Select } from 'flowbite-svelte';
    import type { SelectOptionType } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
    import { CheckBoxFilter } from '$lib/components';
    import { findAllResources } from '$lib/services/resource';

    let searchInput:string = $state("")
    let status:string[] = $state([])
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
        if (e.target.checked) {
            status = [...status, e.target.value]
        } else {
            status = status.filter(el => el !== e.target.value)
        }

        change(getFilters())
    }

    const resourceChange = (e:any) => {
        change(getFilters())
    }

    const searchChange = () => {
        clearTimeout(clearHandle)

        clearHandle = setTimeout(() => {
            change(getFilters())
        }, 500)
        
    }

    const getFilters = ():InputFilters => {
        let out:InputFilters = { filters: [] as InputFilter[] }
        let filterArray:InputFilter[] = [] 

		const kwfil = { key: 'basics.name', value: searchInput, operation: 'fl' }
        filterArray = [...filterArray, kwfil]

        if (status && status.length > 0) {
            const stfil = { key: 'status.status', value: status.join(","), operation: 'in' }
            filterArray = [...filterArray, stfil]
        }

        if (resourceID) {
            const rsfil = { key: 'resourceID', value: resourceID, operation: 'custom', customName: 'resource_in_project' }
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
                resourceOpts.push({name: "All", value: ""})
				resourceOpts = [...resourceOpts, ...r.results?.map((r) => ({
					name: r.name,
					value: r.id as string
				})) as SelectOptionType<string>[]];
			});
	};
</script>

<div class="flex">
    <div class="mr-2">
        <Search slot="search" size="md" placeholder="Filter by name" bind:value={searchInput} on:keyup={searchChange} />
    </div>

    <div class="mr-2">
        <CheckBoxFilter name="Status" bind:group={status} change={statusChange} opts={statusOpts} />
    </div>

    {#await loadPage()}
        ...
    {:then}
    <div class="mr-2">
        <!-- <SelectInput name="Status" bind:group={status} change={statusChange} opts={statusOpts} /> -->
        <Select items={resourceOpts} bind:value={resourceID} placeholder="Assigned Resource" on:change={resourceChange} />
    </div>
    {/await}
</div>

