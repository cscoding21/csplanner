<script lang="ts">
	import { onMount } from "svelte";
	import CTA from "../home/components/CTA.svelte";
	import { particles } from "$lib/assets/js/jscript";
	
    onMount(() => {
		particles()
	})
</script>


<!-- Content -->
<section class="relative">

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
        <img src="/src/lib/assets/images/page-illustration.svg" class="max-w-none" width="1440" height="427" alt="Page Illustration">
    </div>

    <div class="max-w-6xl mx-auto px-4 sm:px-6">
        <div class="pt-32 md:pt-40">

            <!-- Section header -->
            <div class="text-center pb-12 md:pb-20">
                <div class="inline-flex font-medium bg-clip-text text-transparent bg-linear-to-r from-purple-500 to-purple-200 pb-3">Leaders love Stellar</div>
                <h1 class="h1 bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">Trusted by leading companies</h1>
                <div class="max-w-3xl mx-auto">
                    <p class="text-lg text-slate-400">Stellar powers thousands of high-impact product teams. From next-gen startups who reach for the stars to established greats who change the world.</p>
                </div>
            </div>

            <!-- Customers grid -->
            <div class="max-w-[352px] mx-auto sm:max-w-[728px] lg:max-w-none pb-12 md:pb-20">
                <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3 group [&_*:nth-child(n+5):not(:nth-child(n+12))]:order-1 [&_*:nth-child(n+10):not(:nth-child(n+11))]:order-2!" data-highlighter>
                    <!-- Item #01 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-01.png" width="352" height="198" alt="Customer Background 01" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-01.svg" width="78" height="34" alt="Customer 01" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #02 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-02.png" width="352" height="198" alt="Customer Background 02" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-02.svg" width="70" height="25" alt="Customer 02" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #03 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-03.png" width="352" height="198" alt="Customer Background 03" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-03.svg" width="107" height="33" alt="Customer 03" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #04 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-04.png" width="352" height="198" alt="Customer Background 04" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-04.svg" width="85" height="36" alt="Customer 04" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #05 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-05.png" width="352" height="198" alt="Customer Background 05" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-05.svg" width="86" height="18" alt="Customer 05" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #06 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-06.png" width="352" height="198" alt="Customer Background 06" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-06.svg" width="110" height="21" alt="Customer 06" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #07 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-07.png" width="352" height="198" alt="Customer Background 07" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-07.svg" width="83" height="23" alt="Customer 07" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #08 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-08.png" width="352" height="198" alt="Customer Background 08" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-08.svg" width="98" height="40" alt="Customer 08" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #09 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-09.png" width="352" height="198" alt="Customer Background 09" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-09.svg" width="106" height="35" alt="Customer 09" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Item #10 -->
                    <div>
                        <a href="/customers/detail">
                            <div class="relative h-full bg-slate-800 rounded-3xl p-px -m-px before:absolute before:w-64 before:h-64 before:-left-32 before:-top-32 before:bg-indigo-500 before:rounded-full before:opacity-0 before:pointer-events-none before:transition-opacity before:duration-500 before:translate-x-[var(--mouse-x)] before:translate-y-[var(--mouse-y)] hover:before:opacity-30 before:z-30 before:blur-[64px] after:absolute after:inset-0 after:rounded-[inherit] after:opacity-0 after:transition-opacity after:duration-500 after:[background:radial-gradient(250px_circle_at_var(--mouse-x)_var(--mouse-y),var(--color-slate-400),transparent)] group-hover:after:opacity-100 after:z-10 overflow-hidden">
                                <div class="relative h-full bg-slate-900 rounded-[inherit] z-20 overflow-hidden">
                                    <!-- Particles animation -->
                                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                                        <canvas data-particle-animation data-particle-quantity="3"></canvas>
                                    </div>
                                    <div class="flex items-center justify-center">
                                        <img class="w-full h-full aspect-video object-cover" src="/src/lib/assets/images/customer-bg-10.png" width="352" height="198" alt="Customer Background 10" aria-hidden="true" />
                                        <img class="absolute" src="/src/lib/assets/images/customer-10.svg" width="132" height="22" alt="Customer 10" />
                                    </div>
                                </div>
                            </div>
                        </a>
                    </div>
                    <!-- Testimonial #01 -->
                    <div class="flex flex-col items-center justify-center text-center p-4">
                        <p class="font-medium bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-3">
                            <span class="line-clamp-4">
                                “We struggled to bring all our conversations into one place until we found Stellar. The UI is very clean and we love the integration with Spark.”
                            </span>
                        </p>
                        <div class="inline-flex mb-2">
                            <img class="rounded-full" src="/src/lib/assets/images/customer-avatar-01.jpg" width="32" height="32" alt="Customer Avatar 01" />
                        </div>
                        <div class="text-sm font-medium text-slate-300">
                            Mike Hunt <span class="text-slate-700">-</span> <a class="text-purple-500 hover:underline" href="#0">Thunderbolt</a>
                        </div>
                    </div>
                    <!-- Testimonial #02 -->
                    <div class="flex flex-col items-center justify-center text-center p-4">
                        <p class="font-medium bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-3">
                            <span class="line-clamp-4">
                                “We struggled to bring all our conversations into one place until we found Stellar. The UI is very clean and we love the integration with Spark.”
                            </span>
                        </p>
                        <div class="inline-flex mb-2">
                            <img class="rounded-full" src="/src/lib/assets/images/customer-avatar-02.jpg" width="32" height="32" alt="Customer Avatar 02" />
                        </div>
                        <div class="text-sm font-medium text-slate-300">
                            Mirko Mitt <span class="text-slate-700">-</span> <a class="text-purple-500 hover:underline" href="#0">Mildrink</a>
                        </div>
                    </div>
                </div>
            </div>

        </div>
    </div>

</section>

<CTA />