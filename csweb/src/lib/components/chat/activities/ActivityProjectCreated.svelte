<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	
    interface Props {
        comment: CommentEnvelope
    }
    let { comment }:Props = $props()

    let detail:any = $derived(JSON.parse(comment.data.text))
</script>


The project <b>{detail.name}</b> was created!!