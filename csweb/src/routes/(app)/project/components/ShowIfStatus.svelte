<script lang="ts">
    import { type Snippet } from "svelte";

    interface Props {
		scope: string[];
		status: string; //"new"|"draft"|"proposed"|"approved"|"rejected"|"backlogged"|"scheduled"|"inflight"|"complete"|"deferred"|"abandoned";
		invert?: boolean;
        children: Snippet
	}
	let { 
		scope, 
		status = $bindable(), 
		invert, 
		children 
	}: Props = $props();

    console.log("status check", scope, status)
</script>


{#if scope.includes(status)}
    {@render children()}
{:else if !scope.includes(status) && invert }
    {@render children()}
{/if}


