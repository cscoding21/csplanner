<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
    let diffs:any = $derived(JSON.parse(detail.diffs))
</script>

Project <b>{detail.name}</b> was moved from <b>{detail.from_state}</b> to <b>{detail.to_state}</b>.