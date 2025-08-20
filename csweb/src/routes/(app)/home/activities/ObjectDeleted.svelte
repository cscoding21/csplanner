<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";

    interface Props {
        activity: Activity
        object: string
    }
    let { activity, object }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
</script>

The {object} <b>{detail.name}</b> was removed from project <b>{detail.project_name}</b>.
