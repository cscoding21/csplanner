<script lang="ts">
    import type { Pagination } from "$lib/graphql/generated/sdk";
    import { Dropdown, DropdownGroup, DropdownItem } from "flowbite-svelte";
    import { ChevronDownOutline } from "flowbite-svelte-icons";

    interface Props {
        paging: Pagination
        change: Function
    }
    let { 
        paging = $bindable(),
        change
    }:Props = $props()

    let rpp = $derived((paging.resultsPerPage || 1) as number)
    let totalResults = $derived((paging.totalResults || 0)) 
    let pageNumber = $derived(paging.pageNumber as number)
    let pagesRequired = $derived(Math.trunc(totalResults / rpp) + (totalResults % rpp > 0 ? 1 : 0))
    let rangeBegin = $derived(((pageNumber - 1) * rpp) + 1)
    let rangeEnd = $derived((rangeBegin + rpp) - 1 <= totalResults ? (rangeBegin + rpp) - 1 : totalResults)
  
    let hasPrevious = $derived((pageNumber > 1))
    let hasNext = $derived((pageNumber < pagesRequired))
    let dropdownOpen = $state(false)

    const rppOpts = [5, 10, 20, 50, 100]

    const dropDownSelect = (pn:number, rp:number) => {
        change(getPagingState(pn, rp))

        dropdownOpen = false
    }

    const getPagingState = (pn:number, rp:number):Pagination => {
        return {
            resultsPerPage: rp,
            pageNumber: pn
        }
    }
  </script>
  

<div class="grid grid-cols-2">

<div>
    <div class="font-normal">
        <button
            class="mt-0.5 inline-flex gap-1 rounded-lg p-2 text-center text-sm font-medium text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-white"
            >{rangeBegin} - {rangeEnd} of {totalResults} <ChevronDownOutline size="lg" /></button
        >
        <Dropdown class="min-w-48">
          <DropdownGroup>
            {#each rppOpts as opt}
                {#if opt === rpp}
                <DropdownItem class="font-bold" href="#">
                    Showing {opt} results
                </DropdownItem>
                {:else}
                <DropdownItem class="font-normal" href="#" onclick={() => dropDownSelect(1, opt)}>
                    Show {opt} results
                </DropdownItem>
                {/if}
            {/each}
            </DropdownGroup>
        </Dropdown>
    </div>
</div>

<div class="text-right">

    {#if pagesRequired > 1}
    <nav aria-label="Page navigation example">
      <ul class="inline-flex -space-x-px text-sm">
        {#if hasPrevious}
        <li>
          <button onclick={() => dropDownSelect(pageNumber - 1, rpp)} class="flex items-center justify-center px-3 h-8 ms-0 leading-tight text-gray-500 bg-white border border-e-0 border-gray-300 rounded-s-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">Previous</button>
        </li>
        {/if}
        {#each {length: pagesRequired} as _, index}
        {@const thisPage = (index + 1 )}
        {#if paging.pageNumber === thisPage}
          <li>
              <span aria-current="page" class="flex items-center justify-center px-3 h-8 text-gray-600 border border-gray-300 bg-blue-50 hover:bg-gray-600 hover:text-gray-100 dark:border-gray-700 dark:bg-gray-700 dark:text-white">{thisPage}</span>
          </li>
        {:else}
        <li>
          <button onclick={() => dropDownSelect(thisPage, rpp)} class="flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">{thisPage}</button>
        </li>
        {/if}
      {/each}
        {#if hasNext}
        <li>
          <button onclick={() => dropDownSelect(pageNumber + 1, rpp)} class="flex items-center justify-center px-3 h-8 leading-tight text-gray-500 bg-white border border-gray-300 rounded-e-lg hover:bg-gray-100 hover:text-gray-700 dark:bg-gray-800 dark:border-gray-700 dark:text-gray-400 dark:hover:bg-gray-700 dark:hover:text-white">Next</button>
        </li>
        {/if}
      </ul>
    </nav>
    {:else}
        <span>&nbsp;</span>
    {/if}
  </div>
</div>