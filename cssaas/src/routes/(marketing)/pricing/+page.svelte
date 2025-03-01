<script lang="ts">
	import { particles } from "$lib/assets/js/main";
	import Pricing from "$lib/components/Pricing.svelte";
	import { onMount } from "svelte";
	import CTA from "../home/components/CTA.svelte";
	import PriceCustomers from "./components/PriceCustomers.svelte";
	import PriceFaQs from "./components/PriceFAQs.svelte";
	import PriceFeatures from "./components/PriceFeatures.svelte";

	onMount(() => {
		particles()
	})

</script>


<Pricing />

<PriceFeatures />

<PriceCustomers />

<PriceFaQs />

<CTA />