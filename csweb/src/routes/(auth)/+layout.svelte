<script lang="ts">
	import { onMount } from 'svelte';
	import { authService } from '$lib/services/auth';

	const as = authService();

	onMount(async () => {
		console.log('auth layout.svelte authenticated', as.authCheck());
	});
</script>

<div class="flex flex-row">
	<div class="basis-1/3">&nbsp;</div>
	<div class="basis-1/3"><slot /></div>
	<div class="basis-1/3">&nbsp;</div>
</div>
