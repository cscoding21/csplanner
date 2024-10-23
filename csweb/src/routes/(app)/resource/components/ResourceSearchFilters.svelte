<script lang="ts">
    import { Search } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
	import { CheckBoxFilter } from '$lib/components';

    let searchInput:string = $state("")
    let status:string[] = $state([])
    let type:string[] = $state([])

    let clearHandle:any

    let statusOpts = [
        { value: "inhouse", name: "In-house", checked: false},
        { value: "proposed", name: "Proposed", checked: false},
    ]

    let typeOpts = [
        { value: "human", name: "Human", checked: false},
        { value: "equipment", name: "Equipment", checked: false},
        { value: "software", name: "Software", checked: false},
    ]

    const statusChange = (e:any) => {
        if (e.target.checked) {
            status = [...status, e.target.value]
        } else {
            status = status.filter(el => el !== e.target.value)
        }

        change(getFilters())
    }

    const typeChange = (e:any) => {
        if (e.target.checked) {
            type = [...type, e.target.value]
        } else {
            type = type.filter(el => el !== e.target.value)
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

		const kwfil = { key: 'name', value: searchInput, operation: 'fl' }
        filterArray = [...filterArray, kwfil]

        if (status && status.length > 0) {
            const stfil = { key: 'status', value: status.join(","), operation: 'in' }
            filterArray = [...filterArray, stfil]
        }

        if (type && type.length > 0) {
            const tyfil = { key: 'type', value: type.join(","), operation: 'in' }
            filterArray = [...filterArray, tyfil]
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

    <div class="mr-2">
        <CheckBoxFilter name="Type" bind:group={type} change={typeChange} opts={typeOpts} />
    </div>
</div>