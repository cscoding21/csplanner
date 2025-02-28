<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';

	onMount(async () => {
		console.log('layout onMount');

		goto("/home")
	});
</script>