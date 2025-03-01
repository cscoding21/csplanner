<script lang="ts">
	import HomeClients from "./components/HomeClients.svelte";
    import HomeHero from "./components/HomeHero.svelte";
	import HomeFeatures2 from "./components/HomeFeatures2.svelte";
	import HomeFeatures3 from "./components/HomeFeatures3.svelte";
	import HomeFeatures4 from "./components/HomeFeatures4.svelte";
	import Pricing from "$lib/components/Pricing.svelte";
    import CTA from "./components/CTA.svelte";
	import { onMount } from "svelte";
	import { particles } from "$lib/assets/js/main";


	onMount(() => {
		particles()
		
        //@ts-expect-error
        AOS.init({
            once: true,
            disable: 'phone',
            duration: 1000,
            easing: 'ease-out-cubic',
        });
	})
</script>
   

<HomeHero />

<HomeClients />

<HomeFeatures2 />

<HomeFeatures3 />

<HomeFeatures4 />

<Pricing />

<!-- HomeTestimonials2  -->

<CTA />        
