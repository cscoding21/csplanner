<script>
	//import '../app.postcss';
	import "../app.css";
	import { onMount } from 'svelte';
	import { isAuthenticated } from '$lib/stores/auth';
	import { goto } from '$app/navigation';
    //import { refreshListStore } from "$lib/stores/list";
    //import { refreshResourceStore, refreshResourceAllocationStore } from "$lib/stores/resource";
    //import { refreshTemplateStore } from "$lib/stores/project";

	onMount(async () => {
		if(!isAuthenticated()) {
			goto("/login");
		}

		console.log('setting default stores')
		// refreshListStore()
		// refreshResourceStore()
		// refreshResourceAllocationStore()
		// refreshTemplateStore()
	});
</script>

<slot />