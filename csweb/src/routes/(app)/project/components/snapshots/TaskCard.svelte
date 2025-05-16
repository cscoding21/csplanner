<script lang="ts">
	import { DirectiveLocation } from "graphql";
import type { Snippet } from "svelte";



    interface Props {
        task:string
        link:string
        children?:Snippet
        priority?:"high"|"med"|"low"
    }
    let { task, link, priority, children }:Props = $props()

    let colorClass = $state("text-red-600 dark:text-red-500")

    switch(priority){
        case("low"):
            colorClass = "text-green-600 dark:text-green-500"
            break;
        case("med"):
            colorClass = "text-yellow-600 dark:text-yellow-500"
            break;
        default:
            colorClass = "text-red-600 dark:text-red-500"
    }
</script>


<div
      class="mb-2 rounded-lg bg-white p-4 shadow-sm dark:bg-gray-900 sm:p-6"
    >
    <div class="flex items-center justify-between">
    <div class="me-4 flex items-center sm:me-0">
    <svg
        class="me-1.5 hidden h-6 w-6 shrink-0 {colorClass} sm:flex"
        aria-hidden="true"
        xmlns="http://www.w3.org/2000/svg"
        fill="none"
        viewBox="0 0 24 24"
    >
        <path
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M5 14v7M5 5v9.5c5.6-5.5 8.4 2.7 14 0V4.8c-5.6 2.7-8.4-5.5-14 0Z"
        />
    </svg>
    <h3
        class="text-sm font-medium text-gray-900 dark:text-white sm:text-base"
    >
    <a href={link}>
        {task}
    </a>
    </h3>
    </div>
    </div>
    <div class="py-4 px-8">
        {#if children}
        {@render children()}
        {/if}
    </div>
</div>