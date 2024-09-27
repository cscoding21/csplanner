<script lang="ts">
    import { Button, Modal } from 'flowbite-svelte';
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
  
  <button onclick={() => (popupModal = true)} class="block py-2 px-4 hover:bg-gray-100 dark:hover:bg-gray-600 dark:hover:text-white">
    {@render children()}
  </button>
  
  <Modal bind:open={popupModal} size="lg" autoclose>
    <h3>Reply to comment</h3>
    <QuillDisplay attachContext={comment.id} bind:contents={comment.text} />
    <div>
        <QuillEditor attachContext={comment.id} error="" bind:contents={editorContents} quillEditor={qe} />

        <div class="col-span-4">
            <span class="float-right">
                <Button color="alternative">Cancel</Button>
                <Button class="mr-2" on:click={postReply}>Reply</Button>
            </span>
            <br class="clear-both" />
        </div>
    </div>
  </Modal>