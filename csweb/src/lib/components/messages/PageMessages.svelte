<script lang="ts">
    import { toasts } from "$lib/stores/toasts";
    import { Toast } from "$lib/components/messages";
</script>


{#if $toasts}
<div class="fixed bottom-2 right-5 left-95 z-50 flex flex-col items-right justify-right">
    {#each $toasts as toast(toast)}
     <Toast toast={toast} />
    {/each}
</div>
{/if}