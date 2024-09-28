<script lang="ts">
    import { Button, Toolbar, ToolbarButton } from "flowbite-svelte";
    import { PaperClipOutline, PaperPlaneOutline } from "flowbite-svelte-icons";
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
<div class="flex flex-col justify-between w-full h-full content-end">
    {#each comments as comment(comment.id)}
        <div class="px-6 pt-6 mb-4 text-base bg-white rounded-lg border border-gray-200 shadow-sm dark:bg-gray-800 dark:border-gray-700">
            <div class="mb-2">
                <CommentItem {comment} projectID={id} update={update} canReply={true} />
            </div>
            {#if comment.replies && comment.replies.length > 0}
                
                {#if showReplies }
                <div class="pl-12">
                    {#each comment.replies as reply(reply.id)}
                        <CommentItem comment={reply} projectID={id} update={update} canReply={false} />
                    {/each}
                </div>
                {/if}
            {/if}
        </div>
    {/each}
</div>

<div>
<form>
    <label for="chat" class="sr-only">Your message</label>
    <div class="flex px-3 py-2 rounded-lg bg-gray-50 dark:bg-gray-700">
        <div>
            <ToolbarButton color="dark" class="text-gray-500 dark:text-gray-400">
            <PaperClipOutline class="w-6 h-6" />
            <span class="sr-only">Upload image</span>
            </ToolbarButton>
        </div>
        <div class="w-full">
            <QuillEditor attachContext={id} bind:contents={editorContent} bind:quillEditor={qe} />
        </div>
        <div>
            <ToolbarButton type="submit" color="blue" class="rounded-full text-primary-600 dark:text-primary-500" onclick={postComment}>
            <PaperPlaneOutline class="w-6 h-6 rotate-45" />
            <span class="sr-only">Send message</span>
            </ToolbarButton>
        </div>
    </div>
</form>
</div>

{/await}