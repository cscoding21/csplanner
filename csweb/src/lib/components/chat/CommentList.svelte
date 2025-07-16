<script lang="ts">
    import { ToolbarButton } from "flowbite-svelte";
    import { PaperClipOutline, PaperPlaneOutline } from "flowbite-svelte-icons";
    import { CommentItem, NoResults } from "$lib/components";
    import { findProjectComments, addComment  } from "$lib/services/comment";
    import type { CommentEnvelope } from "$lib/graphql/generated/sdk";
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
                comments = c.results as CommentEnvelope[]

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

    let comments = $state([] as CommentEnvelope[])
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
    {#each comments as comment(comment.meta?.id)}
        <div class="px-2 pt-4 pb-2 mb-2 text-base bg-white rounded-lg border border-gray-200 shadow-sm dark:bg-gray-800 dark:border-gray-700 text-gray-500 dark:text-gray-400">
            <div class="mb-2">
                <CommentItem {comment} parentID={undefined} projectID={id} update={update} canReply={true} />
            </div>
            
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