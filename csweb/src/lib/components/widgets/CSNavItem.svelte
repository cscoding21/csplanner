<script lang="ts">
	import type { Snippet } from "svelte";

    interface Props {
        href: string
        active: boolean
        children: Snippet;
    }
    let { href, active, children }: Props = $props();

    let classAugment = $derived(active ? " rounded text-amber-500 dark:text-amber-500" : " text-gray-700 hover:text-primary-700 dark:text-gray-400 dark:hover:text-white")

</script>

<li>
    <a href={href} class="block {classAugment}" aria-current={active}>
        {@render children()}
    </a>
</li>