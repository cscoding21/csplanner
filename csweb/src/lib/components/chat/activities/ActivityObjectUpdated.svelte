<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	import { csJsonKeyTranlsator } from "$lib/utils/format";
	
    interface Props {
        comment: CommentEnvelope
        object: string
    }
    let { comment, object }:Props = $props()

    let detail:any = $derived(JSON.parse(comment.data.text))
    let diffs:any = $derived(JSON.parse(detail.diffs))
</script>

Details for the {object} were updated.  
{#if diffs && diffs.length > 0}
<ul>
{#each diffs as d}
    {@const key = csJsonKeyTranlsator(d.path)}
    {#if key}
    <li>{key} is now <b>{d.value}</b></li>
    {/if}
{/each}
</ul>
{/if}