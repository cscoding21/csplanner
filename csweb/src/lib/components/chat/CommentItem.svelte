<script lang="ts">
    import { Avatar, Popover, ToolbarButton } from 'flowbite-svelte';
    import { PaperPlaneOutline, PaperClipOutline } from 'flowbite-svelte-icons';
    import type { CommentEnvelope, UpdateCommentReply, User } from "$lib/graphql/generated/sdk";
    import { addCommentReply, toggleCommentEmote, updateComment } from "$lib/services/comment";
    import { authService } from "$lib/services/auth";
    import { formatDateTime, getInitialsFromName, pluralize } from "$lib/utils/format"
    import { CommentItem, UserDisplay } from "$lib/components";
    import { EmoteButtonLike, EmoteButtonDislike, EmoteButtonLove, EmoteButtonLaugh, EmoteButtonAcknowledge, DeleteComment } from "$lib/components";
    import { DotsHorizontalOutline, CheckPlusCircleOutline, ReplyOutline } from 'flowbite-svelte-icons';
    import { is } from '$lib/utils/check';
    import { normalizeID } from '$lib/utils/id';
    import { addToast } from '$lib/stores/toasts';
    import { callIf } from '$lib/utils/helpers';
    import { QuillEditor, QuillDisplay, ReplyToComment } from '..';


    interface Props {
        comment: CommentEnvelope
        projectID: string
        update: Function
        canReply: boolean
        parentID: string|undefined
    }
    let { 
        comment, 
        projectID, 
        update,
        canReply,
        parentID
    }:Props = $props()

    const as = authService()

    const id = comment.meta.id
    const userEmail:string = as.currentUser()?.email as string
    let editorContent = $state(comment.data.text)
    let replyEditorContents = $state()

    let userCreatedComment: boolean = (userEmail === comment.meta?.createByUser?.email)
    let editMode:boolean = $state(false)
    let showReplies:boolean = $state(false)
    let showReplyForm:boolean = $state(false)

    let qe:any

    const toggleReaction = async (projectID:string, parentID:string, commentID:string, type:string) => {
        toggleCommentEmote({ projectID: projectID, parentID: parentID, commentID: commentID, emoteType: type }).then(res => {
            if(res.success) {
                callIf(update)
            } else {
                console.log(res.message)
            }
        })
    }

    const toggleEdit = () => {
        editMode = !editMode
    }

    const toggleReply = () => {
        showReplies = !showReplyForm
    }

    const postReply = async () => {
        if(!comment.meta.id) {
            return
        }

        const reply:UpdateCommentReply = {
            parentCommentID: comment.meta.id,
            text: JSON.stringify(replyEditorContents),
        }

        const response = await addCommentReply(reply)
        .then((res) => {
            if(res.status.success) {
                console.log('success')

                callIf(update)

                showReplyForm = false
            } else {
                console.error(res.status.message)
            }
        })
        .catch((err) => {
            console.error(err)
        })
    }
    
    const updateComm = async () => {
        const text = JSON.stringify(editorContent)
        const updateDelta = { projectId: projectID, id: id, text: text }

        console.log("updateDelta", updateDelta)
        updateComment(updateDelta).then(res => {
            if(res.status?.success) {
                editMode = false

                editorContent = res.comment.text

                callIf(update)
            } else {
                addToast({ 
                    message: "Error updating comment: " + res.status.message, 
                    dismissible: true, 
                    type: "error"}
                )
            }
        })
    }
</script>

<article class="">
    <footer class="flex justify-between items-center">
        <div class="flex items-center">
            <Avatar class="mr-2 w-8 h-8 rounded-lg" src={comment.meta.createByUser?.profileImage || ""}>{getInitialsFromName(comment.meta.createByUser?.firstName + " " + comment.meta.createByUser?.lastName)}</Avatar>
            <div>
                <span class="text-sm font-semibold text-gray-900 dark:text-white"><UserDisplay user={comment.meta.createByUser as User} /></span> 
                <p class="text-xs text-gray-500 dark:text-gray-400">
                    <time datetime="2022-02-08" title={formatDateTime(comment.meta?.createdAt)}>{formatDateTime(comment.meta?.createdAt)}</time>
                    {#if comment.data?.isEdited}
                        <span class="text-xs text-yellow-700 dark:text-yellow-200"> (edited {formatDateTime(comment.meta?.updatedAt)}) </span>
                    {/if}
                </p>
            </div>    
        </div>
        <div class="flex items-center space-x-2">
            {#if is(comment.data?.likes)}
            <EmoteButtonLike 
                        userID={userEmail}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.likes??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.data?.laughsAt)}
            <EmoteButtonLaugh 
                        userID={userEmail}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.laughsAt??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.data?.dislikes)}
            <EmoteButtonDislike 
                        userID={userEmail}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.dislikes??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.data?.loves)}
            <EmoteButtonLove 
                        userID={userEmail}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.loves??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.data?.acknowledges)}
            <EmoteButtonAcknowledge 
                        userID={userEmail}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.acknowledges??[]}
                        size="sm"
                        />
            {/if}

            <button type="button" id="emote_{normalizeID(id)}" class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
                <CheckPlusCircleOutline size="sm" />
            </button>
            <Popover triggeredBy="#emote_{normalizeID(id)}">
                <div class="flex">
                <span class="p-1">
                <EmoteButtonLike 
                        userID={userEmail}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={id}
                        {projectID}
                        {parentID}
                        users={comment.data?.likes??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonLaugh 
                        userID={userEmail}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.meta?.id}
                        {projectID}
                        {parentID}
                        users={comment.data?.laughsAt??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonDislike 
                        userID={userEmail}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.meta?.id}
                        {projectID}
                        {parentID}
                        users={comment.data?.dislikes??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonLove 
                        userID={userEmail}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.meta?.id}
                        {projectID}
                        {parentID}
                        users={comment.data?.loves??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonAcknowledge 
                        userID={userEmail}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.meta?.id}
                        {projectID}
                        {parentID}
                        users={comment.data?.acknowledges??[]}
                        />
                </span>
                </div>
            </Popover>

            <!-- {#if canReply}
            <ReplyToComment comment={comment.data} update={update}>
                <ReplyOutline />
            </ReplyToComment>
            {/if} -->

            {#if userCreatedComment}
            <button id="user_{normalizeID(id)}" class="inline-flex items-center text-sm font-medium text-center text-gray-500 dark:text-gray-400 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-gray-50 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
                <DotsHorizontalOutline />
                <span class="sr-only">Comment settings</span>
            </button>
            <!-- Dropdown menu -->
            <Popover triggeredBy="#user_{normalizeID(id)}" trigger="click">
            <ul class="py-1 text-sm text-gray-700 dark:text-gray-200"
                aria-labelledby="commentUserMenu">
                <li>
                    <button onclick={() => toggleEdit()} class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Edit</button>
                </li>
                <li>
                    <DeleteComment id={id} update={update}>Delete</DeleteComment>
                </li>
            </ul>
            </Popover>
            {/if}
        </div>
    </footer>

    {#if editMode }
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
                        <QuillEditor attachContext={comment.meta?.id} bind:contents={editorContent} quillEditor={qe} />
                    </div>
                    <div>
                        <ToolbarButton type="submit" color="blue" class="rounded-full text-primary-600 dark:text-primary-500" onclick={updateComm}>
                        <PaperPlaneOutline class="w-6 h-6 rotate-45" />
                        <span class="sr-only">Send message</span>
                        </ToolbarButton>
                    </div>
                </div>
            </form>
            </div>
    {:else}
        <QuillDisplay attachContext={id} bind:contents={editorContent} />
    {/if}

    {#if comment.data?.replies && comment.data.replies.length > 0}
        {#if showReplies }
        <div class="pl-4">  
            {#each comment.data.replies as reply(reply.meta?.id)}
                <CommentItem comment={reply} {parentID} projectID={projectID} update={update} canReply={false} />
            {/each}
        </div>
        <button class="text-xs ml-2 mb-2 flex" onclick={() => showReplies = false}>
            <svg class="mr-1.5 w-3 h-3 mt-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
            </svg> 
            <span class="inline-flex cursor-pointer items-center gap-1 text-xs font-medium text-primary-700 hover:underline dark:text-primary-500 sm:text-sm">Hide {comment.data.replies.length} {pluralize("reply", comment.data.replies.length)}</span>
        </button>
        {:else}
            <button class="text-xs ml-2 mb-2 flex" onclick={() => showReplies = true}>
                <svg class="mr-1.5 w-3 h-3 mt-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
                </svg> 
                <span class="inline-flex cursor-pointer items-center gap-1 text-xs font-medium text-primary-700 hover:underline dark:text-primary-500 sm:text-sm">Show {comment.data.replies.length} {pluralize("reply", comment.data.replies.length)}
            </span>
            </button>
        {/if}
    {:else if canReply}
        <button class="text-xs ml-2 mb-2 flex" onclick={() => toggleReply()}>
            <svg class="mr-1.5 w-3 h-3 mt-1" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
                    <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
                </svg> 
                <span class="inline-flex cursor-pointer items-center gap-1 text-xs font-medium text-primary-700 hover:underline dark:text-primary-500 sm:text-sm">Reply</span>
        </button>
    {/if}


    {#if showReplies && canReply}
    <div class="pl-8">
    <form>
        <label for="chat" class="sr-only">Your message</label>
        <div class="flex px-3 py-2 rounded-lg bg-gray-50 dark:bg-gray-700">
            <!-- <div>
                <ToolbarButton color="dark" class="text-gray-500 dark:text-gray-400">
                <PaperClipOutline class="w-6 h-6" />
                <span class="sr-only">Upload image</span>
                </ToolbarButton>
            </div> -->
            <div class="w-full">
                <QuillEditor attachContext={comment.meta.id} bind:contents={replyEditorContents} quillEditor={qe} />
            </div>
            <div>
                <ToolbarButton type="submit" color="blue" class="rounded-full text-primary-600 dark:text-primary-500" onclick={postReply}>
                <PaperPlaneOutline class="w-6 h-6 rotate-45" />
                <span class="sr-only">Send message</span>
                </ToolbarButton>
            </div>
        </div>
    </form>
    </div>
    {/if}
</article>