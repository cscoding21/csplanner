<script lang="ts">
	import { particles } from "$lib/assets/js/jscript";
	import { onMount } from "svelte";
	import CTA from "../home/components/CTA.svelte";
    import IntCategories from "./components/IntCategories.svelte";
    import IntHero from "./components/IntHero.svelte";

    onMount(() => {
		particles()
	})
</script>


<IntHero />

<IntCategories />

<CTA />