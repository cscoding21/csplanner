<script lang="ts">
	import type { Snippet } from "svelte";

    interface Props {
        children: Snippet
    }
    let { children }:Props = $props()
</script>

<!-- overflow-y-auto  -->
<div class="h-dvh overflow-y-auto p-4 w-full rounded-lg bg-white dark:bg-gray-800 text-gray-500 dark:text-gray-400 border-gray-200 dark:border-gray-700 divide-gray-200 dark:divide-gray-700">
    <div class="growx">
        {@render children()}
    </div>
</div>