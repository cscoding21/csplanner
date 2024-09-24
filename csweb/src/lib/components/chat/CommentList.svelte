<script lang="ts">
    import { CommentItem } from "$lib/components";
    import { findProjectComments } from "$lib/services/comment";
    import type { Comment } from "$lib/graphql/generated/sdk";
    

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    const update = () => {
        console.log("update")
    }

    let comments = $state([] as Comment[])

    const loadPage = async () => {
        await findProjectComments(id)
            .then((c) => {
            comments = c

            return comments
        })
    }

    loadPage()
</script>

{#await comments}
    <span>Loading...</span>
{:then}
<div>
    <div>
        {#each comments as comment(comment.id)}
            <CommentItem {comment} projectID={id} update={update} />
            {#if comment.replies}
                {#each comment.replies as reply(reply.id)}
                    <CommentItem comment={reply} projectID={id} update={update} />
                {/each}
            {/if}
        {/each}
    </div>
</div>
{/await}