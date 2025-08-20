<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";
	import { csJsonKeyTranlsator } from "$lib/utils/format";


    interface Props {
        activity: Activity
        object: string
    }
    let { activity, object }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
    let diffs:any = $derived(JSON.parse(detail.diffs))
</script>

Details for the {object} <b>{detail.name}</b> were updated.

<!-- <code>{JSON.stringify(diffs)}</code> --> 
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