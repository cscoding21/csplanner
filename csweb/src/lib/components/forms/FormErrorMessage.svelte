<script lang="ts">
	let { message = $bindable() } = $props();
</script>

{#if message}
	<p class="mt-4 text-sm text-red-600 dark:text-red-500">
		<svg
			class="float-left mr-3 h-4 w-4 text-red-600 dark:text-red-500"
			aria-hidden="true"
			xmlns="http://www.w3.org/2000/svg"
			fill="none"
			viewBox="0 0 20 20"
		>
			<path
				stroke="currentColor"
				stroke-linecap="round"
				stroke-linejoin="round"
				stroke-width="2"
				d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"
			/>
		</svg>
		{message}
	</p>
{/if}
