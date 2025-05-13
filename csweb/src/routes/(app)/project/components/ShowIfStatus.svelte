<script lang="ts">
    import { type Snippet } from "svelte";

    interface Props {
		scope: string[];
		status: string; //"new"|"draft"|"proposed"|"approved"|"rejected"|"backlogged"|"scheduled"|"inflight"|"complete"|"deferred"|"abandoned";
		invert?: boolean;
        children: Snippet;
		elseRender?:Snippet;
	}
	let { 
		scope, 
		status = $bindable(), 
		invert, 
		children,
		elseRender
	}: Props = $props();
</script>


{#if scope.includes(status) && !invert}
    {@render children()}
{:else if !scope.includes(status) && invert }
    {@render children()}
{:else}
	{#if elseRender}
		{@render elseRender()}
	{/if}
{/if}


