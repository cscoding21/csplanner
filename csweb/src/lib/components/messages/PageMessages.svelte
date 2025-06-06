<script lang="ts">
	import { toasts } from '$lib/stores/toasts';
	import { Toast } from '$lib/components';
</script>

{#if $toasts}
	<div class="items-right justify-right fixed bottom-2 right-5 z-50 flex flex-col">
		{#each $toasts as toast (toast)}
			<Toast {toast} />
		{/each}
	</div>
{/if}
