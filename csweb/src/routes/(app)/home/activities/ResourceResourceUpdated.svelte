<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
    let diffs:any = $derived(JSON.parse(detail.diffs))
</script>

The details for resource <b>{detail.name}</b> were updated.

<ul>
{#each diffs as d}
    <li>{d.path} is now <b>{d.value}</b></li>
{/each}
</ul>