<script lang="ts">
    import { Button } from "flowbite-svelte";
    import { CommentItem } from "$lib/components";
    import { findProjectComments, addComment  } from "$lib/services/comment";
    import type { Comment } from "$lib/graphql/generated/sdk";
	import QuillEditor from "../forms/QuillEditor.svelte";
	import { onMount } from "svelte";
    

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let qe:any;

    const update = async () => {
        console.log("update")
        return await findProjectComments(id)
            .then((c) => {
                comments = c.results as Comment[]

                console.log("comments loadPage", comments)

                return comments
        })
    }

    const postComment = async () => {
        return await addComment({ text: JSON.stringify(editorContent), projectId: id})
            .then(r => {
                if(r.status.success) {
                    update().then(r => {
                        qe.clear()

                        return r
                })
                } else {
                    console.error(r.status.message )
                }
            })
    }

    let comments = $state([] as Comment[])
    let editorContent = $state()
    
    onMount(() => {
        qe.toggleEnabled();
    })

    const loadPage = async () => {
        return await update()
            .then((c) => {
                return c
            })
    }
</script>

{#await loadPage()}
    <span>Loading...</span>
{:then}
<div>
    {#each comments as comment(comment.id)}
        <CommentItem {comment} projectID={id} update={update} />
        {#if comment.replies}
            <div class="pl-12">
                {#each comment.replies as reply(reply.id)}
                    <CommentItem comment={reply} projectID={id} update={update} />
                {/each}
            </div>
        {/if}
    {/each}
</div>

<QuillEditor error={""} attachContext={id} bind:contents={editorContent} bind:quillEditor={qe} />

<Button onclick={postComment}>Post</Button>

{/await}