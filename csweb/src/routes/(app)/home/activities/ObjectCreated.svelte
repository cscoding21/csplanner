<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";

    interface Props {
        activity: Activity
        object: string
    }
    let { activity, object }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
</script>

A new {object}, <b>{detail.name}</b>, was added to the project <b>{detail.project_name}</b>.
