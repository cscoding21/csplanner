<script lang="ts">
    import { Dropdown, DropdownGroup, DropdownHeader, DropdownItem, type SelectOptionType } from "flowbite-svelte";
    import { CheckCircleOutline, ChevronDownOutline, CircleMinusOutline } from "flowbite-svelte-icons";

    interface Props {
        isMulti: boolean
        filterName: string
        filterValue: string[]
        filterOpts: SelectOptionType<string>[]
        change: Function
    }
    let { 
        isMulti,
        filterName,
        filterValue = $bindable(),
        filterOpts,
        change
    }:Props = $props()

    let dropdownOpen = $state(false)
    let internalState = $state([...filterValue])

    const clearAll = () => {
        internalState = []

        change(internalState)

        dropdownOpen = false
    }

    const checkHasValues = (state:any):boolean => {
        if(isMulti) {
            return state && state.length > 0
        }

        return state[0] && state[0].length > 0
    }

    const dropDownChange = (opt:string) => {
        if (!isMulti) {
            internalState = []
        }

        if(internalState.includes(opt)) {
            internalState = internalState.filter(f => f !== opt)
        } else {
            internalState.push(opt)
        }
        change(internalState)

        dropdownOpen = false
    }

    const getFilterDisplay = ():string => {
        let out:string[] = []

        if (!filterValue.length) {
            return "All " + filterName
        }

        for(let i = 0; i < filterValue.length; i++) {
            const thisVal = filterValue[i]
            const thisOpt = filterOpts.filter(o => o.value ===thisVal ).map(m => m.name) as string[]
            out = [...out, ...thisOpt]
        }

        return out.join(", ")
    }

    let hasValues = $derived(checkHasValues(filterValue))
  </script>
  

<div>
    <div class="font-normal">
        <button
            class="text-left mt-0.5 inline-flex gap-1 rounded-lg p-2 text-sm font-medium text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white"
            >
            {filterName}: 
            {#if hasValues}
            <br />
            <span class="text-yellow-50">{getFilterDisplay()}</span>
            {:else}
            All
            {/if}
            
        <ChevronDownOutline size="lg" /></button
        >
        <Dropdown class="min-w-48">
            
            {#if hasValues}
            <DropdownGroup>
            <div role="none">
                <DropdownItem class="font-normal border-b border-gray-200 dark:border-gray-600" href="#" onclick={() => clearAll()}>
                    Clear selection <span class="float-right"><CircleMinusOutline size="sm" /></span>
                </DropdownItem>
            </div>
            </DropdownGroup>
            {/if}
            
            <DropdownGroup>
            {#if filterOpts && filterOpts.length > 0}
            {#each filterOpts as opt}
                {#if filterValue.includes(opt.value)}
                <DropdownItem class="font-bold" href="#" onclick={() => dropDownChange(opt.value)}>
                    <b>{opt.name}</b> <span class="float-right"><CheckCircleOutline size="sm" color="yellow" /></span>
                </DropdownItem>
                {:else}
                <DropdownItem class="font-normal" href="#" onclick={() => dropDownChange(opt.value)}>
                    {opt.name}
                </DropdownItem>
                {/if}
            {/each}
            {/if}
            </DropdownGroup>
        </Dropdown>
    </div>
</div>
