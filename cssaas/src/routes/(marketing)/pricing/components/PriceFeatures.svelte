<script lang="ts">
    import pi from '$lib/assets/images/pricing-illustration.svg'
    import pit from '$lib/assets/images/pricing-illustration-top.svg'
    import pi1 from '$lib/assets/images/pricing-icon-01.svg'
    import pi2 from '$lib/assets/images/pricing-icon-02.svg'
    import pi3 from '$lib/assets/images/pricing-icon-03.svg'
    import pi4 from '$lib/assets/images/pricing-icon-04.svg'
    import pi5 from '$lib/assets/images/pricing-icon-05.svg'
    import pi6 from '$lib/assets/images/pricing-icon-06.svg'
    import pi7 from '$lib/assets/images/pricing-icon-07.svg'
</script>


<!-- Features -->
<section class="relative">
            
    <!-- Particles animation -->
    <div class="absolute left-1/2 -translate-x-1/2 top-0 -z-10 w-64 h-64 -mt-24">
        <div class="absolute inset-0 -z-10" aria-hidden="true">
            <canvas data-particle-animation data-particle-quantity="6" data-particle-staticity="30"></canvas>
        </div>
    </div>

    <div class="max-w-6xl mx-auto px-4 sm:px-6">
        <div class="pb-12 md:pb-20">

            <!-- Section header -->
            <div class="max-w-3xl mx-auto text-center pb-12">
                <h2 class="h2 bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">Stop overpaying for software</h2>
                <p class="text-lg text-slate-400">There are many variations available, but the majority have suffered alteration in some form, by injected humour.</p>
            </div>

            <!-- Rings illustration -->
            <div class="text-center pb-8">
                <div class="inline-flex items-center justify-center relative">
                    <!-- Particles animation -->
                    <div class="absolute inset-0 -z-10" aria-hidden="true">
                        <canvas data-particle-animation data-particle-quantity="10"></canvas>
                    </div>
                    <div class="inline-flex">
                        <img src={pi} width="334" height="334" alt="Features illustration" />
                    </div>
                    <img class="absolute -mt-[40%]" src={pit} width="396" height="328" alt="Features illustration top" aria-hidden="true" />
                    <div class="absolute w-[200%] flex items-center justify-center space-x-5">
                        <div class="-rotate-[4deg]">
                            <img class="w-11 h-11 drop-shadow-lg animate-[float_2.6s_ease-in-out_infinite] opacity-10" src={pi1} width="80" height="80" alt="Pricing icon 01" />
                        </div>
                        <div class="rotate-[4deg]">
                            <img class="w-14 h-14 drop-shadow-lg animate-[float_2.4s_ease-in-out_infinite] opacity-25" src={pi2} width="80" height="80" alt="Pricing icon 02" />
                        </div>
                        <div class="-rotate-[4deg]">
                            <img class="w-16 h-16 drop-shadow-lg animate-[float_2.2s_ease-in-out_infinite] opacity-60" src={pi3} width="80" height="80" alt="Pricing icon 03" />
                        </div>
                        <img class="drop-shadow-lg animate-float" src={pi4} width="80" height="80" alt="Pricing icon 04" />
                        <div class="rotate-[4deg]">
                            <img class="w-16 h-16 drop-shadow-lg animate-[float_2.2s_ease-in-out_infinite] opacity-60" src={pi5} width="80" height="80" alt="Pricing icon 05" />
                        </div>
                        <div class="-rotate-[4deg]">
                            <img class="w-14 h-14 drop-shadow-lg animate-[float_2.4s_ease-in-out_infinite] opacity-25" src={pi6} width="80" height="80" alt="Pricing icon 06" />
                        </div>
                        <div class="rotate-[4deg]">
                            <img class="w-11 h-11 drop-shadow-lg animate-[float_2.6ås_ease-in-out_infinite] opacity-10" src={pi7} width="80" height="80" alt="Pricing icon 07" />
                        </div>
                    </div>
                </div>
            </div>

            <!-- Features list -->
            <div class="grid gap-6 sm:grid-cols-2 lg:grid-cols-3">
                <!-- Card #1 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            Purpose-built for company that requires more than a <strong class="text-slate-300 font-medium">simple plan</strong> with security infrastructure.
                        </p>
                    </div>
                </div>
                <!-- Card #2 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            AI-powered to <strong class="text-slate-300 font-medium">remove the burdens</strong> of tedious knowledge management and security tasks.
                        </p>
                    </div>
                </div>
                <!-- Card #3 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            There's no prioritized support in Stellar. You can use email or live chat and you will hear back in a <strong class="text-slate-300 font-medium">couple of hours</strong>.
                        </p>
                    </div>
                </div>
                <!-- Card #4 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            Comprehensive <strong class="text-slate-300 font-medium">developer docs</strong> and a centralized support center packed many resources.
                        </p>
                    </div>
                </div>
                <!-- Card #5 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            No upcharges—and we'd never upsell you to a higher plan or a <strong class="text-slate-300 font-medium">dedicated IP</strong> to improve deliverability.
                        </p>
                    </div>
                </div>
                <!-- Card #6 -->
                <div class="bg-linear-to-tr from-slate-800/50 to-slate-800/10 rounded-3xl border border-slate-800">
                    <div class="flex items-center space-x-4 p-4 h-full">
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="20" height="20">
                            <path class="fill-purple-500" fill-opacity=".24" d="M10 0C4.5 0 0 4.5 0 10s4.5 10 10 10 10-4.5 10-10S15.5 0 10 0Z" />
                            <path class="fill-purple-400" fill-rule="nonzero" d="M13 6.586 14.414 8l-5.747 5.748-3.081-3.081L7 9.252l1.667 1.667z" />
                        </svg>
                        <p class="text-sm text-slate-400">
                            Tool training, dedicated resources, and <strong class="text-slate-300 font-medium">regular updates</strong> are available for both small and large teams.
                        </p>
                    </div>
                </div>
            </div>

        </div>
    </div>
</section>