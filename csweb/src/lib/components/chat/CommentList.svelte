<script lang="ts">
    import { Button } from "flowbite-svelte";
    import { CommentItem } from "$lib/components";
    import { findProjectComments, addComment  } from "$lib/services/comment";
    import type { Comment } from "$lib/graphql/generated/sdk";
	import QuillEditor from "../forms/QuillEditor.svelte";
    

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    let showReplies:[] = $state([])
    let qe:any;

    const update = async () => {
        console.log("update")
        return await findProjectComments(id)
            .then((c) => {
                comments = c.results as Comment[]

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
        <div class="px-6 pt-6 mb-4 text-base bg-white rounded-lg border border-gray-200 shadow-sm dark:bg-gray-800 dark:border-gray-700">
            <div class="mb-2">
                <CommentItem {comment} projectID={id} update={update} />
            </div>
            {#if comment.replies && comment.replies.length > 0}
                
                {#if showReplies }
                <div class="pl-12">
                    {#each comment.replies as reply(reply.id)}
                        <CommentItem comment={reply} projectID={id} update={update} />
                    {/each}
                </div>
                {/if}
            {/if}
        </div>
    {/each}
</div>

<QuillEditor error={""} attachContext={id} bind:contents={editorContent} bind:quillEditor={qe} />

<div class="col-span-4">
    <span class="float-right">
        <Button onclick={postComment}>Post</Button>
    </span>
    <br class="clear-both" />
</div>

{/await}