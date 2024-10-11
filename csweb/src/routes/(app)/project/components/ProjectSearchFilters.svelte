<script lang="ts">
    import { Button, Dropdown, Checkbox, Search } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
    import { ChevronDownOutline } from 'flowbite-svelte-icons';

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

    <Button color="light">
        <!-- viewbox="0 0 20 20" -->
        <svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" class="w-4 h-4 mr-2 text-gray-400" fill="currentColor">
          <path fill-rule="evenodd" d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z" clip-rule="evenodd" />
        </svg>
        Status<ChevronDownOutline />
      </Button>
    <Dropdown class="w-44 p-3 space-y-3 text-sm">
    {#each statusOpts as st(st)}
    <li>
        <Checkbox bind:group={status} value={st.value} bind:checked={st.checked} on:change={statusChange}>{st.name}</Checkbox>
    </li>
    {/each}
    </Dropdown>
</div>

