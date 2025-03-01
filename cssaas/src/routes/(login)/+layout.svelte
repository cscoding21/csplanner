<script lang="ts">
	import type { Snippet } from "svelte";

    interface Props {
        children: Snippet
    }
    let { children }:Props = $props();
</script>


<div class="flex flex-col min-h-screen overflow-hidden supports-[overflow:clip]:overflow-clip">
    <main class="grow">
        {@render children()}
    </main>
</div>  