<script lang="ts">
    import logoImage from '$lib/assets/images/csplanner-logo-white.png?enhanced'
</script>


<!-- Site header -->
<header class="absolute w-full z-30">
    <div class="max-w-6xl mx-auto px-4 sm:px-6">
        <div class="flex items-center justify-between h-16 md:h-20">

            <!-- Site branding -->
            <div class="flex-1">
                <!-- Logo -->
                <a class="inline-flex" href="/home" aria-label="csPlanner">
                    <enhanced:img class="max-w-none" width="141" src={logoImage} alt="csPlanner">
                </a>
            </div>

            <!-- Desktop navigation -->
            <nav class="hidden md:flex md:grow">

                <!-- Desktop menu links -->
                <ul class="flex grow justify-center flex-wrap items-center">
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/">Home</a>
                    </li>
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/about">About</a>
                    </li>
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/integrations">Integrations</a>
                    </li>
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/pricing">Pricing</a>
                    </li>
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/customers">Customers</a>
                    </li>
                    <li>
                        <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/changelog">Changelog</a>
                    </li>
                </ul>
            </nav>

            <!-- Desktop sign in links -->
            <ul class="flex-1 flex justify-end items-center">
                <li>
                    <a class="font-medium text-sm text-slate-300 hover:text-white whitespace-nowrap transition duration-150 ease-in-out" href="/login">Sign in</a>
                </li>
                <li class="ml-6">
                    <a class="btn-sm text-slate-300 hover:text-white transition duration-150 ease-in-out w-full group [background:linear-gradient(var(--color-slate-900),var(--color-slate-900))_padding-box,conic-gradient(var(--color-slate-400),var(--color-slate-700)_25%,var(--color-slate-700)_75%,var(--color-slate-400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none" href="/signup">
                        <span class="relative inline-flex items-center">
                            Sign up <span class="tracking-normal text-purple-500 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1">-&gt;</span>
                        </span>
                    </a>
                </li>
            </ul>
        
            <!-- Mobile menu -->
            <div class="md:hidden flex items-center ml-4">

                <!-- Hamburger button -->
                <button class="group inline-flex w-8 h-8 text-slate-300 hover:text-white text-center items-center justify-center transition" aria-controls="mobile-nav">
                    <span class="sr-only">Menu</span>
                    <svg class="w-4 h-4 fill-current pointer-events-none" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
                        <rect class="origin-center transition-all duration-300 ease-[cubic-bezier(.5,.85,.25,1.1)] -translate-y-[5px] group-aria-expanded:rotate-[315deg] group-aria-expanded:translate-y-0" y="7" width="16" height="2" rx="1"></rect>
                        <rect class="origin-center group-aria-expanded:rotate-45 transition-all duration-300 ease-[cubic-bezier(.5,.85,.25,1.8)]" y="7" width="16" height="2" rx="1"></rect>
                        <rect class="origin-center transition-all duration-300 ease-[cubic-bezier(.5,.85,.25,1.1)] translate-y-[5px] group-aria-expanded:rotate-[135deg] group-aria-expanded:translate-y-0" y="7" width="16" height="2" rx="1"></rect>
                    </svg>
                </button>                             

                <!-- Mobile navigation -->
                <nav id="mobile-nav" class="absolute top-full z-20 left-0 w-full px-4 sm:px-6 overflow-hidden transition-all duration-300 ease-in-out">
                    <ul class="border border-transparent [background:linear-gradient(var(--color-slate-900),var(--color-slate-900))_padding-box,conic-gradient(var(--color-slate-400),var(--color-slate-700)_25%,var(--color-slate-700)_75%,var(--color-slate-400)_100%)_border-box] rounded-lg px-4 py-1.5">
                        <li>
                            <a class="font-medium text-sm text-slate-300 hover:text-white mx-4 lg:mx-5 transition duration-150 ease-in-out" href="/">Home</a>
                        </li>
                        <li>
                            <a class="flex font-medium text-sm text-slate-300 hover:text-white py-1.5" href="/about">About</a>
                        </li>
                        <li>
                            <a class="flex font-medium text-sm text-slate-300 hover:text-white py-1.5" href="/integrations">Integrations</a>
                        </li>
                        <li>
                            <a class="flex font-medium text-sm text-slate-300 hover:text-white py-1.5" href="/pricing">Pricing</a>
                        </li>
                        <li>
                            <a class="flex font-medium text-sm text-slate-300 hover:text-white py-1.5" href="/customers">Customers</a>
                        </li>
                        <li>
                            <a class="flex font-medium text-sm text-slate-300 hover:text-white py-1.5" href="/changelog">Changelog</a>
                        </li>
                    </ul>
                </nav>

            </div>

        </div>
    </div>
</header>