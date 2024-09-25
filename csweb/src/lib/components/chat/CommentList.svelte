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
        return await findProjectComments(id)
            .then((c) => {
            comments = c.results as Comment[]

            console.log("comments loadPage", comments)

            return comments
        })
    }
</script>

{#await loadPage()}
    <span>Loading...</span>
{:then}
<div>
    {#each comments as comment(comment.id)}
        <CommentItem {comment} projectID={id} update={update} />
        <h1>{comment.text}</h1>
        {#if comment.replies}
            <div class="pl-12">
                {#each comment.replies as reply(reply.id)}
                    <CommentItem comment={reply} projectID={id} update={update} />
                {/each}
            </div>
        {/if}
    {/each}
</div>
{/await}