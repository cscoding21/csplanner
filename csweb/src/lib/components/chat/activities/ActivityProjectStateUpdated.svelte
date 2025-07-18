<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	
    interface Props {
        comment: CommentEnvelope
    }
    let { comment }:Props = $props()

    let detail:any = $derived(JSON.parse(comment.data.text))
</script>

The status for project, <b>{detail.name}</b>, was moved from <b>{detail.from_state}</b> to <b>{detail.to_state}</b>.