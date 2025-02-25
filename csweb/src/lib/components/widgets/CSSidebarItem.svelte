<script lang="ts">
	import { callIf } from "$lib/utils/helpers";
	import type { Snippet } from "svelte";


    interface Props {
		  label: string;
      active: boolean;
      onclick?: Function;
      href?: string;
      icon?: Snippet;
      tail?: Snippet;
	}
	let { label, active, onclick, href, icon, tail } : Props = $props();

  let classAugment = $derived(active ? " bg-gray-100 dark:text-white dark:bg-gray-700" : " hover:bg-gray-100 dark:text-white dark:hover:bg-gray-700")
</script>

<li>
    {#if href}
    <a href={href} onclick={() => callIf(onclick)} class="group w-full flex items-center rounded-lg p-2 text-base font-medium text-gray-900 {classAugment}">
      {#if icon}
        {@render icon()}
      {/if}
      <span class="flex items-center p-2 text-base text-gray-900 transition duration-75 rounded-lg hover:bg-gray-100 group dark:text-gray-200 dark:hover:bg-gray-700 whitespace-nowrap">{label}</span>
      {#if tail}
      <span class="content-end">
        {@render tail()}
      </span>
      {/if}
    </a>

    {:else}
    <button onclick={() => callIf(onclick)} class="group  w-full flex items-center rounded-lg p-2 text-base font-medium text-gray-900 {classAugment}">
      {#if icon}
        {@render icon()}
      {/if}
      <span class="text-left flex-1 ml-3 whitespace-nowrap">{label}</span>
      {#if tail}
      <span class="float-right">
        {@render tail()}
      </span>
      {/if}
    </button>
    {/if}
</li>