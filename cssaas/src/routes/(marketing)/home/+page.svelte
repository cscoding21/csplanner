<script lang="ts">
	import HomeClients from "./components/HomeClients.svelte";
    import HomeHero from "./components/HomeHero.svelte";
	import HomeFeatures2 from "./components/HomeFeatures2.svelte";
	import HomeFeatures3 from "./components/HomeFeatures3.svelte";
	import HomeFeatures4 from "./components/HomeFeatures4.svelte";
	import Pricing from "$lib/components/Pricing.svelte";
    import CTA from "./components/CTA.svelte";
	import { onMount } from "svelte";
	import { particles } from "$lib/assets/js/jscript";


	onMount(() => {
		particles()
	})
</script>
   

<HomeHero />

<HomeClients />

<!-- HomeFeatures  -->

<HomeFeatures2 />

<HomeFeatures3 />

<!-- HomeTestimonials  -->

<HomeFeatures4 />

<Pricing />

<!-- HomeTestimonials2  -->

<CTA />        
