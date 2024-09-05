<script lang="ts">
	import { onMount } from 'svelte';
	import { isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';

	onMount(async () => {
		if(!isAuthenticated()) {
			goto("/login");
		} else {
			goto("/home")
		}
	});
</script>