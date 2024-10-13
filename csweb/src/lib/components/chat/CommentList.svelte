<script lang="ts">
    import { ToolbarButton } from "flowbite-svelte";
    import { PaperClipOutline, PaperPlaneOutline } from "flowbite-svelte-icons";
    import { CommentItem, NoResults } from "$lib/components";
    import { findProjectComments, addComment  } from "$lib/services/comment";
    import type { Comment } from "$lib/graphql/generated/sdk";
	import QuillEditor from "../forms/QuillEditor.svelte";
    import { pluralize } from "$lib/utils/format";
    import type { AssociativeArray } from '$lib/types/array'    

    interface Props {
        id: string
    }
    let { id }:Props = $props()

    // @ts-expect-error - TS parser doesn't like the instantiation is "never"
    let showReplies:AssociativeArray = $state([])
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

<!-- h-full -->
<div class="flex flex-col justify-end w-full">
{#await loadPage()}
    <span>Loading...</span>
{:then}
<div class="grow scrollbar overflow-y-auto">
    {#if comments && comments.length > 0}
    {#each comments as comment(comment.id)}
        <div class="px-2 pt-4 mb-2 text-base bg-white rounded-lg border border-gray-200 shadow-sm dark:bg-gray-800 dark:border-gray-700 text-gray-500 dark:text-gray-400">
            <div class="mb-2">
                <CommentItem {comment} projectID={id} update={update} canReply={true} />
            </div>
            {#if comment.replies && comment.replies.length > 0}
                {#if showReplies[comment.id] }
                <div class="pl-12">  
                    {#each comment.replies as reply(reply.id)}
                        <CommentItem comment={reply} projectID={id} update={update} canReply={false} />
                    {/each}
                </div>
                <button class="text-xs ml-2 mb-2 flex" onclick={() => showReplies[comment.id] = false}>
                    <svg class="mr-1.5 w-3.5 h-3.5 mt-.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                        <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
                    </svg> 
                    Hide {comment.replies.length} {pluralize("reply", comment.replies.length)}
                </button>
                {:else}
                    <button class="text-xs ml-2 mb-2 flex" onclick={() => showReplies[comment.id] = true}>
                        <svg class="mr-1.5 w-3.5 h-3.5 mt-.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
                        </svg> 
                        Show {comment.replies.length} {pluralize("reply", comment.replies.length)}
                    </button>
                {/if}
                
            {/if}
        </div>
    {/each}
    {:else}
        <NoResults title="No comments yet..." newUrl="">Jump in and get the discussion started.</NoResults>
    {/if}
</div>

<div class="mt-2">
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
</div>