<script lang="ts">
	import { onMount } from 'svelte';
	import { goto } from '$app/navigation';
	import { authService } from '$lib/services/auth';

	const as = authService()

	onMount(async () => {
		console.log('layout onMount');
		as.authCheck().then(r => {
			if(r) {
				goto('/home');
			} else {
				goto('/login');
			}
		})
	});
</script>
