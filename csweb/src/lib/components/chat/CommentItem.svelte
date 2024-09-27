<script lang="ts">
    import { Avatar, Popover } from 'flowbite-svelte';
    import type { Comment, User } from "$lib/graphql/generated/sdk";
    import { toggleCommentEmote, updateComment } from "$lib/services/comment";
    import { authService } from "$lib/services/auth";
    import { formatDateTime, getInitialsFromName, pluralize } from "$lib/utils/format"
    import { UserDisplay } from "$lib/components";
    import { EmoteButtonLike, EmoteButtonDislike, EmoteButtonLove, EmoteButtonLaugh, EmoteButtonAcknowledge, DeleteComment } from "$lib/components";
    import { DotsHorizontalOutline, CheckPlusCircleOutline, ReplyOutline } from 'flowbite-svelte-icons';
    import { is } from '$lib/utils/check';
    import { normalizeID } from '$lib/utils/id';
    // import { commentListStateStore } from '$lib/stores/comment'
    import { addToast } from '$lib/stores/toasts';
    import { callIf } from '$lib/utils/helpers';
    import { QuillEditor, QuillDisplay, ReplyToComment } from '..';


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
    const userEmail:string = as.currentUser()?.email as string
    let editorContent = $state(comment.text)

    let userCreatedComment: boolean = (userEmail === comment.user.email)
    let editMode:boolean = $state(false)

    let qe:any

    const toggleReaction = async (commentID:string, type:string) => {
        toggleCommentEmote({ commentID: commentID, emoteType: type }).then(res => {
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
    
    const updateComm = async () => {
        const text = JSON.stringify(editorContent)
        const updateDelta = { projectId: projectID, id: id, text: text }

        console.log("updateDelta", updateDelta)
        updateComment(updateDelta).then(res => {
            if(res.status?.success) {
                editMode = false

                console.log("updatedText", res.comment.text)
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
            <Avatar class="mr-2 w-8 h-8 rounded-lg" src={comment.user?.profileImage || ""}>{getInitialsFromName(comment.user?.firstName + " " + comment.user?.lastName)}</Avatar>
            <div>
                <span class="text-sm font-semibold text-gray-900 dark:text-white"><UserDisplay user={comment.user as User} /></span> 
                <p class="text-xs text-gray-500 dark:text-gray-400">
                    <time datetime="2022-02-08" title={formatDateTime(comment.createdAt)}>{formatDateTime(comment.createdAt)}</time>
                    {#if comment.isEdited}
                        <span class="text-xs text-yellow-700 dark:text-yellow-200"> (edited {formatDateTime(comment.updatedAt)}) </span>
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
                        size="sm"
                        />
            {/if}

            {#if is(comment.laughsAt)}
            <EmoteButtonLaugh 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.laughsAt??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.dislikes)}
            <EmoteButtonDislike 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.dislikes??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.loves)}
            <EmoteButtonLove 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.loves??[]}
                        size="sm"
                        />
            {/if}

            {#if is(comment.acknowledges)}
            <EmoteButtonAcknowledge 
                        userID={userID}
                        toggleReaction={toggleReaction}
                        commentID={id}
                        users={comment.acknowledges??[]}
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

            <ReplyToComment comment={comment} update={update}>
                <ReplyOutline />
            </ReplyToComment>

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
        <QuillEditor error={""} attachContext={id} bind:contents={editorContent} bind:quillEditor={qe} />
        <button onclick={() => updateComm() }>Save</button>
    {:else}
        <QuillDisplay attachContext={id} bind:contents={editorContent} />
    {/if}
</article>