<script lang="ts">
    import { Button, Modal } from 'flowbite-svelte';
    import { ExclamationCircleOutline } from 'flowbite-svelte-icons';
    import { deleteComment } from '$lib/services/comment';
	import type { Snippet } from 'svelte';
    
    interface Props {
        id: string
        name: string
        children: Snippet
    }
    let { id, name, children }: Props = $props()
    let popupModal = $state(false);

    const deleteItem = async () => {
        if(!id) {
            return
        }

        const response = await deleteComment(id)
        .then((res) => {
            if(res.success) {
                console.log('success')
            } else {
                console.error(res.message)
            }
        })
        .catch((err) => {
            console.error(err)
        })
    }
  </script>
  
  <Button on:click={() => (popupModal = true)}>
    {@render children()}
  </Button>
  
  <Modal bind:open={popupModal} size="xs" autoclose>
    <div class="text-center">
      <ExclamationCircleOutline class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" />
      <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">Are you sure you want to delete this comment. {name}?</h3>
      <Button color="red" class="mr-2" on:click={deleteItem}>Yes, I'm sure</Button>
      <Button color="alternative">No, cancel</Button>
    </div>
  </Modal>