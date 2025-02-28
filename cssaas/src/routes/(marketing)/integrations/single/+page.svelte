<script lang="ts">

</script>

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

            <div class="md:flex md:justify-between">

                <!-- Page content -->
                <div class="md:grow pb-12 md:pb-20">
                    <div class="max-w-[720px]">

                        <div class="flex flex-col lg:flex-row space-y-6 lg:space-y-0 lg:space-x-16">

                            <!-- Back button -->
                            <div class="shrink-0">
                                <div class="sticky top-6">
                                    <a class="flex items-center justify-center w-9 h-9 group border border-transparent rounded-full [background:linear-gradient(var(--color-slate-900),var(--color-slate-900))_padding-box,conic-gradient(var(--color-slate-400),var(--color-slate-700)_25%,var(--color-slate-700)_75%,var(--color-slate-400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/30 before:rounded-full before:pointer-events-none" href="/integrations">
                                        <span class="sr-only">Go back</span>
                                        <svg class="w-4 h-4 fill-purple-500" viewBox="0 0 16 16" xmlns="http://www.w3.org/2000/svg">
                                            <path d="M6.7 14.7l1.4-1.4L3.8 9H16V7H3.8l4.3-4.3-1.4-1.4L0 8z" />
                                        </svg>
                                    </a>
                                </div>
                            </div>

                            <!-- Content -->
                            <div>
                                <article class="pb-12 mb-12 border-b [border-image:linear-gradient(to_right,transparent,var(--color-slate-800),transparent)1]">

                                    <figure class="bg-slate-700/20 border border-slate-300/10 p-4 rounded-3xl mb-8">
                                        <img class="w-full rounded-2xl" src="/src/lib/assets/images/integration-image.png" width="586" height="316" alt="Integration image" />
                                    </figure>

                                    <h1 class="sr-only">GitHub</h1>

                                    <!-- Post content -->
                                    <div class="prose max-w-none text-slate-400 prose-headings:text-slate-50 prose-h2:text-xl prose-h2:mt-8 prose-h2:mb-4 prose-p:leading-relaxed prose-a:text-purple-500 prose-a:no-underline hover:prose-a:underline prose-strong:text-slate-50 prose-strong:font-medium prose-blockquote:pl-5 prose-blockquote:xl:-ml-5 prose-blockquote:border-l-2 prose-blockquote:border-purple-500 prose-blockquote:font-medium prose-blockquote:text-slate-300 prose-blockquote:italic">
                                        <h2>Overview</h2>
                                        <p>
                                            This powerful GitHub integration keeps your work in sync in both applications. It links features to Pull Requests so that details update automatically from In Progress to Done as the PR moves from drafted to merged - there is no need to update the issue in Stellar at all.
                                        </p>
                                        <p>
                                            Move even faster by using a keyboard shortcut that creates the issue's git branch name, assigns the issue and moves the issue to In Progress in one step. This integration supports <a href="#0">GitHub Enterprise</a>.
                                        </p>
                                        <h2>How it works</h2>
                                        <p>
                                            Configure the GitHub integration <strong class="text-slate-50 font-medium">in a couple of minutes</strong> by authenticating and selecting which repos to connect to Stellar. Once you do, anyone in the workspace can start linking PRs.
                                        </p>
                                        <p>
                                            Link PRs to issues in a few ways. Use the issue ID in the branch name or PR title before creating the PR. To link multiple PRs, or link issues after creating the PR, use magic words along with the issue ID in the PR description.
                                        </p>
                                    </div>
                                </article>

                                <aside class="pl-6 border-l-2 border-purple-500">
                                    <p class="inline-flex font-medium italic text-lg bg-clip-text text-transparent bg-linear-to-r from-slate-200/60 via-slate-200 to-slate-200/60 pb-4">“ This integration is so perfect it tickles my brain. It ensures that your work is in sync across your entire team. ”</p>
                                    <footer class="flex items-center space-x-4">
                                        <img class="shrink-0 rounded-full" src="/src/lib/assets/images/post-avatar.jpg" width="32" height="32" alt="Author" />
                                        <div class="text-sm font-medium text-slate-300">
                                            Mike Hunt <span class="text-slate-700">-</span> <a class="text-purple-500 hover:underline" href="#0">Thunderbolt</a>
                                        </div>
                                    </footer>
                                </aside>
                            </div>
            
                        </div>

                    </div>
                </div>

                <!-- Sidebar -->
                <aside class="md:w-64 lg:w-80 md:shrink-0 md:pt-[3.75rem] lg:pt-0 pb-12 md:pb-20">
                    <div class="sticky top-6 md:pl-6 lg:pl-10">

                        <!-- Sidebar content -->
                        <div class="space-y-6">

                            <!-- Widget -->
                            <div class="bg-linear-to-tr from-slate-800 to-slate-800/25 rounded-3xl border border-slate-800">
                                <div class="px-5 py-6">
                                    <div class="text-center mb-5">
                                        <div class="mb-4">
                                            <div class="relative inline-flex">
                                                <img src="/src/lib/assets/images/integrations-08.svg" width="80" height="80" alt="Icon 08" />
                                                <img class="absolute top-0 -right-1" src="/src/lib/assets/images/star.svg" width="24" height="24" alt="Star" aria-hidden="true" />
                                            </div>
                                        </div>
                                        <button class="btn-sm text-slate-300 hover:text-white transition duration-150 ease-in-out group [background:linear-gradient(var(--color-slate-900),var(--color-slate-900))_padding-box,conic-gradient(var(--color-slate-400),var(--color-slate-700)_25%,var(--color-slate-700)_75%,var(--color-slate-400)_100%)_border-box] relative before:absolute before:inset-0 before:bg-slate-800/70 before:rounded-full before:pointer-events-none">
                                            <span class="relative inline-flex items-center">
                                                Install Integration <span class="tracking-normal text-purple-500 group-hover:translate-x-0.5 transition-transform duration-150 ease-in-out ml-1">-&gt;</span>
                                            </span>
                                        </button>
                                    </div>
                                    <ul class="text-sm">
                                        <li class="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,--theme(--color-slate-700/.3),--theme(--color-slate-700),--theme(--color-slate-700/.3))1]">
                                            <span class="text-slate-400">Company</span>
                                            <span class="text-slate-300 font-medium">Github</span>
                                        </li>
                                        <li class="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,--theme(--color-slate-700/.3),--theme(--color-slate-700),--theme(--color-slate-700/.3))1]">
                                            <span class="text-slate-400">Website</span>
                                            <a class="text-purple-500 font-medium flex items-center space-x-1" href="#0">
                                                <span>github.com</span>
                                                <svg class="fill-current" xmlns="http://www.w3.org/2000/svg" width="9" height="9">
                                                    <path d="m1.285 8.514-.909-.915 5.513-5.523H1.663l.01-1.258h6.389v6.394H6.794l.01-4.226z" />
                                                </svg>
                                            </a>
                                        </li>
                                        <li class="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,--theme(--color-slate-700/.3),--theme(--color-slate-700),--theme(--color-slate-700/.3))1]">
                                            <span class="text-slate-400">Industry</span>
                                            <span class="text-slate-300 font-medium">Technology</span>
                                        </li>
                                        <li class="flex items-center justify-between space-x-2 py-3 border-t [border-image:linear-gradient(to_right,--theme(--color-slate-700/.3),--theme(--color-slate-700),--theme(--color-slate-700/.3))1]">
                                            <span class="text-slate-400">Installs</span>
                                            <span class="text-slate-300 font-medium">100K+</span>
                                        </li>
                                    </ul>
                                </div>
                            </div>

                        </div>

                    </div>
                </aside>

            </div>

        </div>
    </div>
</section>