<script lang="ts">
    import type { Snippet } from "svelte";
  
    interface Props {
          children: Snippet;
    }
    let { children }: Props = $props();
</script>
  
<aside id="page-sidebar" class="m-0 hidden h-screen w-64 shrink-0 overflow-y-auto border-gray-200 bg-white p-3 shadow-sm dark:border-gray-700 dark:bg-gray-800 lg:block">
    <ul class="text-base space-y-1">{@render children()}</ul>
</aside>    