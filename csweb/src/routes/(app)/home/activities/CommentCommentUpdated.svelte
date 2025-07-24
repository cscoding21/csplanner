<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
    let diffs:any = $derived(JSON.parse(detail.diffs))
</script>

<div>
<b>{activity.user.firstName} {activity.user.lastName}</b> updated their comment on the project
<b>{detail.project_name}</b> 
</div>
<blockquote class="pl-4 py-2 italic text-sm">
    "{detail.comment}"
</blockquote>
<!-- <code>{JSON.stringify(diffs)}</code> -->