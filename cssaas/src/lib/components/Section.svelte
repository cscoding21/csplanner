<script lang="ts">
	import type { Snippet } from "svelte";

    interface Props {
        back: "lighter"|"darker"
        children: Snippet
    }
    let { back, children }:Props = $props()

    let bgClass = $derived(back === "lighter" ? "border-t border-b border-gray-100 bg-gray-50 dark:bg-gray-800 dark:border-gray-700" : "bg-white dark:bg-gray-900")
</script>

<section class={bgClass}>

    {@render children()}

</section>  