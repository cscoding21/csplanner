<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
</script>

Details for the project <b>{detail.name}</b> were updated.