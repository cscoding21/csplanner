<script lang="ts">
  import type { Snippet } from "svelte";


  interface Props {
		children: Snippet;
	}
	let { children }: Props = $props();

</script>


<!-- Sidenav -->
<aside id="sidebar" class="hidden h-full w-80 shrink-0 overflow-y-auto border border-gray-200 bg-white p-3 shadow-sm dark:border-gray-700 dark:bg-gray-800 lg:block lg:rounded-lg">
    <ul class="space-y-2">
    {@render children()}
    </ul>
</aside>    