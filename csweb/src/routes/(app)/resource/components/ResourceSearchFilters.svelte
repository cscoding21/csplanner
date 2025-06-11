<script lang="ts">
    import { Input, Search, type SelectOptionType } from 'flowbite-svelte';
    import type { InputFilters, InputFilter } from '$lib/graphql/generated/sdk';
	import { CSMultiFilter } from '$lib/components';
    import { getList } from '$lib/services/list';
    import { findSelectOptsFromList } from '$lib/forms/helpers';
	import { CloseCircleOutline } from 'flowbite-svelte-icons';
	import { getSetting, RESOURCE_SEARCH_INPUT_FILTER, RESOURCE_SKILLS_FILTER, RESOURCE_STATUS_FILTER, RESOURCE_TYPE_FILTER, setSetting } from '$lib/services/settings';
	import { onMount } from 'svelte';

    let searchInput:string = $state("")
    let status:string[] = $state([])
    let type:string[] = $state([])
    let skills:string = $state("")

    let clearHandle:any

    let statusOpts = [
        { value: "inhouse", name: "In-house", checked: false},
        { value: "proposed", name: "Proposed", checked: false},
        { value: "exited", name: "Exited", checked: false},
    ]

    let typeOpts = [
        { value: "human", name: "Human", checked: false},
        { value: "equipment", name: "Equipment", checked: false},
        { value: "software", name: "Software", checked: false},
    ]

    let skillsOpts:SelectOptionType<string>[] = $state([])

    const statusChange = (e:any) => {
        status = e

        setSetting(RESOURCE_STATUS_FILTER, status)

        change(getFilters())
    }

    const typeChange = (e:any) => {
        type = e

        setSetting(RESOURCE_TYPE_FILTER, type)

        change(getFilters())
    }

    const skillsChange = (e:any) => {
        skills = e.join("")

        setSetting(RESOURCE_SKILLS_FILTER, skills)

        change(getFilters())
    }

    const searchChange = () => {
        clearTimeout(clearHandle)

        clearHandle = setTimeout(() => {
            setSetting(RESOURCE_SEARCH_INPUT_FILTER, searchInput)

            change(getFilters())
        }, 500)
        
    }

    const resetFilters = () => {
        searchInput = ""
        status = []
        type = []
        skills = ""

        setSetting(RESOURCE_SEARCH_INPUT_FILTER, searchInput)
        setSetting(RESOURCE_TYPE_FILTER, type)
        setSetting(RESOURCE_STATUS_FILTER, status)
        setSetting(RESOURCE_SKILLS_FILTER, skills)

        change(getFilters())
    }

    const getFilters = ():InputFilters => {
        let out:InputFilters = { filters: [] as InputFilter[] }
        let filterArray:InputFilter[] = [] 

		const kwfil = { key: 'data.name', value: searchInput, operation: 'fl' }
        filterArray = [...filterArray, kwfil]

        if (status && status.length > 0) {
            const stfil = { key: 'data.status', value: status.join(","), operation: 'in' }
            filterArray = [...filterArray, stfil]
        }

        if (type && type.length > 0) {
            const tyfil = { key: 'data.type', value: type.join(","), operation: 'in' }
            filterArray = [...filterArray, tyfil]
        }

        if (skills) {
            const skfil = { key: 'data.skills.skill_id', value: skills, operation: 'ct' }
            filterArray = [...filterArray, skfil]
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
		getList('Skills')
			.then((l) => {
				skillsOpts = findSelectOptsFromList(l);
			})
			.then(() => {
				
			});
	};

    onMount(() => {
        searchInput = getSetting(RESOURCE_SEARCH_INPUT_FILTER, searchInput)
        type = getSetting(RESOURCE_TYPE_FILTER, type)
        status = getSetting(RESOURCE_STATUS_FILTER, status)
        skills = getSetting(RESOURCE_SKILLS_FILTER, skills)

        change(getFilters())
    })

</script>

<div class="flex">
    <div class="mr-4 text-nowrap">
        <Input slot="search" size="md" class="w-80" placeholder="Filter by name" bind:value={searchInput} onkeyup={searchChange} />
    </div>

    <div class="mr-4  text-nowrap">
        <CSMultiFilter filterOpts={statusOpts} change={statusChange} filterValue={status} filterName="Status" isMulti={true} />
    </div>  

    <div class="mr-4  text-nowrap">
        <CSMultiFilter filterOpts={typeOpts} change={typeChange} filterValue={type} filterName="Type" isMulti={true} />
    </div>

    {#await loadPage()}
        ...
    {:then}
    <div class="mr-4  text-nowrap">
        <CSMultiFilter filterOpts={skillsOpts} change={skillsChange} filterValue={[skills]} filterName="Skills" isMulti={false} />
    </div>
    {/await}

    <div class="content-right justify-end w-full">
        <button class="float-right" onclick={resetFilters} title="Clear all filters"><CloseCircleOutline size="lg" /></button>
    </div>
</div>