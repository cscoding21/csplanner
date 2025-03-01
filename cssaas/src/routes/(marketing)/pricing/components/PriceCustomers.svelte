<script lang="ts">
    import clImage1 from '$lib/assets/images/customer-01.jpg?enhanced'
    import clImage2 from '$lib/assets/images/customer-02.jpg?enhanced'
    import clImage3 from '$lib/assets/images/customer-03.jpg?enhanced'
    import clImage4 from '$lib/assets/images/customer-04.jpg?enhanced'
    import clImage5 from '$lib/assets/images/customer-05.jpg?enhanced'
    import clImage6 from '$lib/assets/images/customer-06.jpg?enhanced'
    import clImage7 from '$lib/assets/images/customer-07.jpg?enhanced'
    import clImage8 from '$lib/assets/images/customer-08.jpg?enhanced'
    import clImage9 from '$lib/assets/images/customer-09.jpg?enhanced'
</script>


<!-- Customers -->
<section class="relative">
    <!-- Radial gradient -->
    <div class="absolute inset-0 overflow-hidden pointer-events-none -z-10" aria-hidden="true">
        <div class="absolute flex items-center justify-center top-0 -translate-y-1/2 left-1/2 -translate-x-1/2 w-1/3 aspect-square">
            <div class="absolute inset-0 translate-z-0 bg-purple-500 rounded-full blur-[120px] opacity-50"></div>
        </div>
    </div>
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
        <div class="py-12 md:py-20">
            <!-- Content -->
            <div class="max-w-3xl mx-auto text-center pb-12 md:pb-20">
                <h2 class="h2 bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">Meet our customers</h2>
                <p class="text-lg text-slate-400">There are many variations available, but the majority have suffered, by injected humour, or randomised words which don't look even slightly believable.</p>
            </div>
            <!-- Customers -->
            <div class="grid lg:grid-cols-3 gap-4 sm:gap-6 max-w-xs mx-auto lg:max-w-none">
                <!-- Customer #01 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage1} width="44" height="44" alt="Customer 01" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Mike Smets</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@MikeSmets</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #02 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage2} width="44" height="44" alt="Customer 02" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Mike Bryan</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@mike0point7 </a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #03 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage3} width="44" height="44" alt="Customer 03" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Liz Park</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@liz-park82</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #04 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage4} width="44" height="44" alt="Customer 04" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Jessie Maison</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@jessiem9</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #05 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage5} width="44" height="44" alt="Customer 05" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Devani Janssen</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@deva07p</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #06 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage6} width="44" height="44" alt="Customer 06" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Mark Gerkules</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@mark-gerk</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #07 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage7} width="44" height="44" alt="Customer 07" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Andrey Propenk</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@andrey-prok</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #08 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage8} width="44" height="44" alt="Customer 08" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Jay Master</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@masterpro</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
                <!-- Customer #09 -->
                <div class="relative p-5 before:absolute before:inset-0 before:-z-10 before:border before:border-slate-300 before:bg-slate-700 before:opacity-10 before:rounded-xl">
                    <div class="flex items-center justify-between space-x-2 mb-4">
                        <div class="flex items-center space-x-4">
                            <enhanced:img class="shrink-0 rounded-full" src={clImage9} width="44" height="44" alt="Customer 09" />
                            <div class="grow truncate">
                                <div class="font-bold text-slate-100 truncate">Steve Wolf</div>
                                <a class="text-sm text-slate-500 hover:text-slate-300 font-medium truncate transition-colors" href="#0">@imsteve</a>
                            </div>
                        </div>
                        <svg class="shrink-0" xmlns="http://www.w3.org/2000/svg" width="90" height="14">
                            <path class="fill-purple-500" d="M7 0 5.006 5.368H0l4.232 3.221L2.195 14 7 10.344 11.82 14 9.768 8.589 14 5.368H8.98zM26 0l-1.994 5.368H19l4.232 3.221L21.195 14 26 10.344 30.82 14l-2.052-5.411L33 5.368h-5.02zM45 0l-1.994 5.368H38l4.232 3.221L40.195 14 45 10.344 49.82 14l-2.052-5.411L52 5.368h-5.02zM64 0l-1.994 5.368H57l4.232 3.221L59.195 14 64 10.344 68.82 14l-2.052-5.411L71 5.368h-5.02zM83 0l-1.994 5.368H76l4.232 3.221L78.195 14 83 10.344 87.82 14l-2.052-5.411L90 5.368h-5.02z" />
                        </svg>
                    </div>
                    <p class="text-sm text-slate-400">I love Stellar, is super easy to use with top notch features. My favourites are the affiliates and the lovely login box that converts like magic.</p>
                </div>
            </div>
        </div>
    </div>
</section>
