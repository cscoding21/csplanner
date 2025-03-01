<script lang="ts">
    import pi2 from '$lib/assets/images/page-illustration-02.svg'
    import pill from '$lib/assets/images/page-illustration.svg'

	import { onMount } from "svelte";

    onMount(() => {

    })
</script>


<!-- Hero -->
<section class="relative">

    <!-- Illustration 02 -->
    <div class="md:block absolute left-1/2 -translate-x-1/2 bottom-0 -mb-16 blur-2xl opacity-90 pointer-events-none -z-10" aria-hidden="true">
        <img src={pi2} class="max-w-none" width="1440" height="427" alt="Page Illustration 02">
    </div>
    
    <!-- Opacity layer -->
    <div class="absolute inset-0 bg-slate-900 opacity-60 -z-10" aria-hidden="true"></div>

    <!-- Radial gradient -->
    <div class="absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 pointer-events-none -z-10 w-[800px] aspect-square" aria-hidden="true">
        <div class="absolute inset-0 translate-z-0 bg-purple-500 rounded-full blur-[120px] opacity-30"></div>
        <div class="absolute w-64 h-64 translate-z-0 bg-purple-400 rounded-full blur-[80px] opacity-70"></div>
    </div>

    <!-- Particles animation -->
    <div class="absolute inset-0 h-96 -z-10" aria-hidden="true">
        <canvas data-particle-animation data-particle-quantity="15"></canvas>
    </div>

    <!-- Illustration -->
    <div class="md:block absolute left-1/2 -translate-x-1/2 -mt-16 blur-2xl opacity-90 pointer-events-none -z-10" aria-hidden="true">
        <img src={pill} class="max-w-none" width="1440" height="427" alt="Page Illustration">
    </div>

    <div class="max-w-6xl mx-auto px-4 sm:px-6">
        <div class="pt-32 md:pt-40">

            <!-- Hero content -->
            <div class="text-center pb-12 md:pb-20">
                <div class="inline-flex font-medium bg-clip-text text-transparent bg-linear-to-r from-purple-500 to-purple-200 pb-3">Integrations &amp; Add-ons</div>
                <h1 class="h1 bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">Make Stellar uniquely yours</h1>
                <div class="max-w-3xl mx-auto">
                    <p class="text-lg text-slate-400">Our landing page template works on all devices, so you only have to set it up once, and get beautiful results forever.</p>
                </div>
            </div>
        </div>
    </div>
</section>