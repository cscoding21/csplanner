<script lang="ts">
    import { Button, Dropdown, Checkbox, Search } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
    import { CheckBoxFilter } from '$lib/components';

    let searchInput:string = $state("")
    let status:string[] = $state([])

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

    const statusChange = (e:any) => {
        if (e.target.checked) {
            status = [...status, e.target.value]
        } else {
            status = status.filter(el => el !== e.target.value)
        }

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
            const stfil = { key: 'basics.status', value: status.join(","), operation: 'in' }
            filterArray = [...filterArray, stfil]
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

</script>

<div class="flex">
    <div class="mr-2">
        <Search slot="search" size="md" placeholder="Filter by name" bind:value={searchInput} on:keyup={searchChange} />
    </div>

    <div class="mr-2">
        <CheckBoxFilter name="Status" bind:group={status} change={statusChange} opts={statusOpts} />
    </div>
</div>

