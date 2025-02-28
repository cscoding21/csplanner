<script lang="ts">
	import Footer from "$lib/components/Footer.svelte";
	import PageHeader from "$lib/components/PageHeader.svelte";
	import type { Snippet } from "svelte";

    interface Props {
        children: Snippet
    }
    let { children }:Props = $props();

</script>

<!-- Page wrapper -->
<div class="flex flex-col min-h-screen overflow-hidden supports-[overflow:clip]:overflow-clip">
    <PageHeader />


    <!-- Page content -->
    <main class="grow">
        {@render children()}
    </main>
    
    <Footer />
</div>