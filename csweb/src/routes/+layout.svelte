<script>
	import '../app.postcss';
	import '../app.css';
	import { onMount } from 'svelte';
	import { refreshResourceStore /* refreshResourceAllocationStore */ } from '$lib/stores/resource';
	import { refreshTemplateStore } from '$lib/stores/project';

	onMount(async () => {
		console.log('setting default stores');
		refreshResourceStore();
		// refreshResourceAllocationStore()
		refreshTemplateStore();
	});
</script>

<slot />
