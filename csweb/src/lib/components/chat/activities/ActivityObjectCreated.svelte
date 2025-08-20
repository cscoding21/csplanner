<script lang="ts">
	import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
	
    interface Props {
        comment: CommentEnvelope
        object: string
    }
    let { comment, object }:Props = $props()

    let detail:any = $derived(JSON.parse(comment.data.text))
</script>


The {object} <b>{detail.name}</b> was created!!