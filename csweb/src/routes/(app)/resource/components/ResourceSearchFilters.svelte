<script lang="ts">
    import { Button, Dropdown, Checkbox, Search } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';

    let searchInput:string = $state("")

    let clearHandle:any

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
</div>