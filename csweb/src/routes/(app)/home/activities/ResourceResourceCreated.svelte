<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(activity.detail)
</script>

<b>{detail.name}</b> was added to the resource roster!