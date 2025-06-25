<script lang="ts">
	import { formatPercent } from '$lib/utils/format';

	interface Props {
		amount: number;
	}
	let { amount }: Props = $props();
</script>

{#if amount < 0.0}
	<span class="text-red-400">{formatPercent.format(amount)}</span>
{:else}
	<span class="text-green-400">{formatPercent.format(amount)}</span>
{/if}
