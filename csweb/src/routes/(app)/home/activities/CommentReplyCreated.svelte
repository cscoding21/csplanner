<script lang="ts">
	import type { Activity } from "$lib/graphql/generated/sdk";


    interface Props {
        activity: Activity
    }
    let { activity }:Props = $props()

    let detail:any = $derived(JSON.parse(activity.detail))
</script>

<div>
<b>{activity.user.firstName} {activity.user.lastName}</b> replied to a comment from 
<b>{detail.parent_user_firstname} {detail.parent_user_lastname}</b>
 on the project <b>{detail.project_name}</b>
</div>
<blockquote class="pl-4 py-2 italic text-sm">
    {detail.comment}
</blockquote>