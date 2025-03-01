<script lang="ts">
	import { particles } from "$lib/assets/js/main";
	import { onMount } from "svelte";
	import AboutCTA from "./components/AboutCTA.svelte";
    import AboutHero from "./components/AboutHero.svelte";
	import AboutRecruitment from "./components/AboutRecruitment.svelte";
	import AboutStory from "./components/AboutStory.svelte";
	import AboutTeam from "./components/AboutTeam.svelte";
	import AboutTestimonials from "./components/AboutTestimonials.svelte";

	onMount(() => {
		particles()
	})
</script>

<AboutHero />

<AboutStory />

<AboutTeam />

<AboutRecruitment />

<AboutTestimonials />

<AboutCTA />