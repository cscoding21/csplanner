<script lang="ts">
    import { Modal, ToolbarButton } from 'flowbite-svelte';
    import { PaperClipOutline, PaperPlaneOutline } from 'flowbite-svelte-icons';
    import { type Comment, type UpdateCommentReply } from '$lib/graphql/generated/sdk';
    import { addCommentReply } from '$lib/services/comment';
	import type { Snippet } from 'svelte';
	import { callIf } from '$lib/utils/helpers';

	import { QuillEditor, QuillDisplay } from '$lib/components';

    
    interface Props {
        comment: Comment
        update?: Function
        children: Snippet
    }
    let { 
        comment, 
        update,
        children 
    }: Props = $props()

    let popupModal = $state(false);
    let editorContents = $state()
    let qe:any

    const postReply = async () => {
        if(!comment.id) {
            return
        }

        const reply:UpdateCommentReply = {
            parentCommentID: comment.id,
            text: JSON.stringify(editorContents),
        }

        const response = await addCommentReply(reply)
        .then((res) => {
            if(res.status.success) {
                console.log('success')

                callIf(update)
            } else {
                console.error(res.status.message)
            }
        })
        .catch((err) => {
            console.error(err)
        })
    }
  </script>
  
  <button onclick={() => (popupModal = true)} class="inline-flex items-center text-sm font-medium text-center text-gray-500 dark:text-gray-400 bg-white rounded-lg hover:bg-gray-100 focus:ring-4 focus:outline-none focus:ring-gray-50 dark:bg-gray-800 dark:hover:bg-gray-700 dark:focus:ring-gray-600">
    {@render children()}
  </button>
  
  <Modal bind:open={popupModal} size="lg" autoclose>
    <h3>Reply to comment</h3>
    <QuillDisplay attachContext={comment.id} bind:contents={comment.text} />

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
                <QuillEditor attachContext={comment.id} bind:contents={editorContents} quillEditor={qe} />
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

  </Modal>