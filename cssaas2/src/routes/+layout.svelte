<script lang="ts">
	import Footer from '$lib/components/Footer.svelte';
	import TopNav from '$lib/components/TopNav.svelte';
	import '../app.css';

	let { children } = $props();
</script>


<TopNav />

{@render children()}


<Footer />