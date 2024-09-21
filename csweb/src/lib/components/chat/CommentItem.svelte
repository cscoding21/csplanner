<script lang="ts">
    import { Avatar, Popover } from 'flowbite-svelte';
    import type { Comment } from "$lib/graphql/generated/sdk";
    import { deleteComment, toggleCommentEmote, addCommentReply, updateComment } from "$lib/services/comment";
    import { authService } from "$lib/services/auth";
    import { formatDateTime, getInitialsFromName, pluralize } from "$lib/utils/format"
    import { QuillEditor, UserDisplay } from "$lib/components";
    import { EmoteButtonLike, EmoteButtonDislike, EmoteButtonLove, EmoteButtonLaugh, EmoteButtonAcknowledge } from "$lib/components";
    import { DotsHorizontalOutline, CheckPlusCircleOutline } from 'flowbite-svelte-icons';
    import { is } from '$lib/utils/check';
    import { normalizeID } from '$lib/utils/id';
    // import { commentListStateStore } from '$lib/stores/comment'
    import { addToast } from '$lib/stores/toasts';
    import { callIf } from '$lib/utils/helpers';


    interface Props {
        comment: Comment
        projectID: string
        update: Function
    }
    let { 
        comment, 
        projectID, 
        update 
    }:Props = $props()

    const as = authService()

    const id = comment.id
    const userID:string = as.currentUser()?.id as string

    let userCreatedComment: boolean = (userID === comment.controlFields?.createdBy)
    let editMode: Boolean = $state(false)
    let replyMode:Boolean = $state(false)
    let html: any

    // function setCommentRepliesOpen(id:string) {
    //     $commentListStateStore[id] = true
    // }

    // function setCommentRepliesClosed(id:string) {
    //     $commentListStateStore[id] = false
    // }

    const toggleReaction = async (commentID:string, type:string) => {
        toggleCommentEmote({ commentID: commentID, emoteType: type }).then(res => {
            if(res.data.toggleEmote.success) {
                callIf(update)
            } else {
                console.log(res.data.toggleEmote.message)
            }
        })
    }

    const openDeleteModal = async (commentID:string) => {
        await deleteComm(commentID)
    }

    const deleteComm = async (commentID:string) => {
        deleteComment(commentID || "").then(res => {
            if(res.data.deleteProjectComment.success) {
                addToast({ 
                    message: "Comment deleted successfully", 
                    dismissible: true, 
                    type: "success"}
                )

                callIf(update)
            } else {
                console.log(res.data.deleteProjectComment.message)
            }
        })
    }

    const createCommentReply = async (text:string) => {
        addCommentReply({ parentCommentID: id, text: text }).then(res => {
            if(res.data.createProjectCommentReply.status?.success) {
                replyMode = false

                //setCommentRepliesOpen(id)
                callIf(update)
            } else {
                addToast({ 
                    message: "Error creating reply: " + res.data.createProjectCommentReply.status.message, 
                    dismissible: true, 
                    type: "error"}
                )
            }
        })
    }

    
    const updateComm = async (text:string) => {
        updateComment({ projectId: projectID, id: id, text: text }).then(res => {
            if(res.data.updateProjectComment.status?.success) {
                editMode = false

                //setCommentRepliesOpen(id)
                callIf(update)
            } else {
                addToast({ 
                    message: "Error updating comment: " + res.data.updateProjectComment.status.message, 
                    dismissible: true, 
                    type: "error"}
                )
            }
        })
    }
</script>

<article class="p-6 mb-6 text-base bg-white rounded-lg border border-gray-200 shadow-sm dark:bg-gray-800 dark:border-gray-700">
    <footer class="flex justify-between items-center mb-2">
        <div class="flex items-center">
            <Avatar class="mr-2 w-8 h-8 rounded-lg" src={comment.resource?.profileImage || ""}>{getInitialsFromName(comment.resource?.name || "")}</Avatar>
            <div>
                <span class="text-sm font-semibold text-gray-900 dark:text-white"><UserDisplay id={comment.resource?.id as string} /></span> 
                <p class="text-xs text-gray-500 dark:text-gray-400">
                    <time datetime="2022-02-08" title={formatDateTime(comment.controlFields?.createdAt.valueOf())}>{formatDateTime(comment.controlFields?.createdAt.valueOf())}</time>
                    {#if comment.isEdited}
                        <span class="text-xs text-yellow-700 dark:text-yellow-200"> (edited {formatDateTime(comment.controlFields?.updatedAt.valueOf())}) </span>
                    {/if}
                </p>
            </div>    
        </div>
        <div class="flex items-center space-x-2">
            {#if is(comment.likes)}
            <EmoteButtonLike 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.likes??[]}
                        size="xs"
                        />
            {/if}

            {#if is(comment.laughsAt)}
            <EmoteButtonLaugh 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.laughsAt??[]}
                        size="xs"
                        />
            {/if}

            {#if is(comment.dislikes)}
            <EmoteButtonDislike 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.dislikes??[]}
                        size="xs"
                        />
            {/if}

            {#if is(comment.loves)}
            <EmoteButtonLove 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.loves??[]}
                        size="xs"
                        />
            {/if}

            {#if is(comment.acknowledges)}
            <EmoteButtonAcknowledge 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.acknowledges??[]}
                        size="xs"
                        />
            {/if}

            <button type="button" id="emote_{normalizeID(id)}" class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-400 font-medium">
                <CheckPlusCircleOutline size="sm" />
            </button>
            <Popover triggeredBy="#emote_{normalizeID(id)}">
                <div class="flex">
                <span class="p-1">
                <EmoteButtonLike 
                        userID={userID}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.likes??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonLaugh 
                        userID={userID}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.id}
                        users={comment.laughsAt??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonDislike 
                        userID={userID}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.id}
                        users={comment.dislikes??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonLove 
                        userID={userID}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.id}
                        users={comment.loves??[]}
                        />
                </span>

                <span class="p-1">
                <EmoteButtonAcknowledge 
                        userID={userID}
                        size="sm"
                        toggleReaction={toggleReaction}
                        commentID={comment.id}
                        users={comment.acknowledges??[]}
                        />
                </span>
                </div>
            </Popover>
            {#if userCreatedComment}
            <button id="user_{normalizeID(id)}" class="inline-flex items-center p-2 text-sm font-medium text-center text-gray-500 dark:text-gray-400 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-gray-50 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
                <DotsHorizontalOutline />
                <span class="sr-only">Comment settings</span>
            </button>
            <!-- Dropdown menu -->
            <Popover triggeredBy="#user_{normalizeID(id)}" trigger="click">
            <ul class="py-1 text-sm text-gray-700 dark:text-gray-200"
                aria-labelledby="commentUserMenu">
                <li>
                    <button onclick={() => editMode = !editMode } class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Edit</button>
                </li>
                <li>
                    <button onclick={() => openDeleteModal(id)} class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">Remove</button>
                </li>
            </ul>
            </Popover>
            {/if}
        </div>
    </footer>
    {#if userCreatedComment && editMode}
        <QuillEditor post={updateComm} attachContext={id} contents={comment.text} />
    {:else}
        <p class="text-gray-500 dark:text-gray-400 text-sm mb-1">{@html html}</p>
    {/if}

    <button type="button"
        onclick={() => { replyMode=!replyMode }}
        class="flex items-center text-sm text-gray-500 hover:underline dark:text-gray-300 font-medium">
        <svg class="mr-1.5 w-3.5 h-3.5" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 18">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 5h5M5 8h2m6-3h2m-5 3h6m2-7H2a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h3v5l5-5h8a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/>
        </svg>                
        Reply
    </button>

    {#if replyMode}
    <div class="mt-4">
        <QuillEditor post={createCommentReply} attachContext={id} contents="" />
    </div>
    {/if}
</article>