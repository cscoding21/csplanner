<script lang="ts">
	import HomeHero from "$lib/components/HomeHero.svelte";
	import Section from "$lib/components/Section.svelte";



</script>

<HomeHero back="darker" />




<!-- FEATURES 2 -->
<Section back="lighter">
  <div class="max-w-screen-xl px-4 py-8 mx-auto sm:py-16 lg:py-24">
    <div class="max-w-screen-md mb-8 lg:mb-16">
      <h2 class="mb-4 text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl dark:text-white">
        Designed for business teams like yours
      </h2>
      <p class="text-gray-500 dark:text-gray-400 sm:text-xl">Here at Flowbite we focus on markets where
        technology, innovation, and capital can unlock long-term value and drive economic growth.</p>
    </div>

    <div class="space-y-8 md:grid md:grid-cols-2 lg:grid-cols-3 md:gap-8 xl:gap-12 md:space-y-0">
      <div>
        <div class="flex items-center justify-center w-16 h-16 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-12 h-12 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path d="M3 12v3c0 1.657 3.134 3 7 3s7-1.343 7-3v-3c0 1.657-3.134 3-7 3s-7-1.343-7-3z" />
            <path d="M3 7v3c0 1.657 3.134 3 7 3s7-1.343 7-3V7c0 1.657-3.134 3-7 3S3 8.657 3 7z" />
            <path d="M17 5c0 1.657-3.134 3-7 3S3 6.657 3 5s3.134-3 7-3 7 1.343 7 3z" />
          </svg>
        </div>

        <div class="mt-5">
          <h3 class="text-xl font-bold dark:text-white">
            Project Management
          </h3>

          <ul role="list" class="mt-4 space-y-4">
            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Unified Contribution Graph
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Org activity graph
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Org dependency insights
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Milestones
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 rounded-full bg-primary-100 text-primary-700 dark:bg-primary-900 dark:text-primary-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Repo insights
              </span>
            </li>
          </ul>
        </div>
      </div>

      <div>
        <div class="flex items-center justify-center w-16 h-16 bg-purple-100 rounded-lg dark:bg-purple-900">
          <svg class="w-12 h-12 text-purple-600 dark:text-purple-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z"
              clip-rule="evenodd" />
          </svg>
        </div>

        <div class="mt-5">
          <h3 class="text-xl font-bold dark:text-white">
            Collaborative Coding
          </h3>

          <ul role="list" class="mt-4 space-y-4">
            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-purple-600 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Dynamic reports and dashboards
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-purple-600 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Code review assignments
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-purple-600 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Team discussions
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-purple-600 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Protected branches
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-purple-600 bg-purple-100 rounded-full dark:bg-purple-900 dark:text-purple-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Draft pull requests
              </span>
            </li>
          </ul>
        </div>
      </div>

      <div>
        <div class="flex items-center justify-center w-16 h-16 mb-4 bg-teal-100 rounded-lg dark:bg-teal-900">
          <svg class="w-12 h-12 text-teal-600 dark:text-teal-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M6.625 2.655A9 9 0 0119 11a1 1 0 11-2 0 7 7 0 00-9.625-6.492 1 1 0 11-.75-1.853zM4.662 4.959A1 1 0 014.75 6.37 6.97 6.97 0 003 11a1 1 0 11-2 0 8.97 8.97 0 012.25-5.953 1 1 0 011.412-.088z"
              clip-rule="evenodd" />
            <path fill-rule="evenodd"
              d="M5 11a5 5 0 1110 0 1 1 0 11-2 0 3 3 0 10-6 0c0 1.677-.345 3.276-.968 4.729a1 1 0 11-1.838-.789A9.964 9.964 0 005 11zm8.921 2.012a1 1 0 01.831 1.145 19.86 19.86 0 01-.545 2.436 1 1 0 11-1.92-.558c.207-.713.371-1.445.49-2.192a1 1 0 011.144-.83z"
              clip-rule="evenodd" />
            <path fill-rule="evenodd"
              d="M10 10a1 1 0 011 1c0 2.236-.46 4.368-1.29 6.304a1 1 0 01-1.838-.789A13.952 13.952 0 009 11a1 1 0 011-1z"
              clip-rule="evenodd" />
          </svg>
        </div>

        <div class="mt-5">
          <h3 class="text-xl font-bold dark:text-white">
            Enterprise Security
          </h3>

          <ul role="list" class="mt-4 space-y-4">
            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-teal-600 bg-teal-100 rounded-full dark:bg-teal-900 dark:text-teal-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Required reviews
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-teal-600 bg-teal-100 rounded-full dark:bg-teal-900 dark:text-teal-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Dependabot security and
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-teal-600 bg-teal-100 rounded-full dark:bg-teal-900 dark:text-teal-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                Dependency graph
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-teal-600 bg-teal-100 rounded-full dark:bg-teal-900 dark:text-teal-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                GitHub Advisory Database
              </span>
            </li>

            <li class="flex space-x-2.5">
              <div
                class="inline-flex items-center justify-center w-5 h-5 text-teal-600 bg-teal-100 rounded-full dark:bg-teal-900 dark:text-teal-400">
                <svg class="w-3.5 h-3.5" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                  <path fill-rule="evenodd"
                    d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                    clip-rule="evenodd" />
                </svg>
              </div>
              <span class="leading-tight text-gray-500 dark:text-gray-400">
                GPG commit signing verification
              </span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</Section>


<!-- PRODUCT FEATURES -->
<Section back="darker">
  <div class="max-w-screen-xl px-4 py-8 mx-auto sm:py-16 lg:py-24">
    <div class="max-w-3xl mx-auto mb-8 text-center lg:mb-16">
      <h2 class="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl dark:text-white">
        What you can do with Flowbite
      </h2>
      <p class="mt-4 text-base text-gray-500 dark:text-gray-400 sm:text-xl">
        Here at flowbite we focus on markets where technology, innovation, and capital can unlock long-term value and
        drive
        economic growth.
      </p>
    </div>

    <div class="space-y-8 md:grid md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 md:gap-8 xl:gap-8 md:space-y-0">
      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M6.672 1.911a1 1 0 10-1.932.518l.259.966a1 1 0 001.932-.518l-.26-.966zM2.429 4.74a1 1 0 10-.517 1.932l.966.259a1 1 0 00.517-1.932l-.966-.26zm8.814-.569a1 1 0 00-1.415-1.414l-.707.707a1 1 0 101.415 1.415l.707-.708zm-7.071 7.072l.707-.707A1 1 0 003.465 9.12l-.708.707a1 1 0 001.415 1.415zm3.2-5.171a1 1 0 00-1.3 1.3l4 10a1 1 0 001.823.075l1.38-2.759 3.018 3.02a1 1 0 001.414-1.415l-3.019-3.02 2.76-1.379a1 1 0 00-.076-1.822l-10-4z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Platform
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          We keep Flowbite, secure, and free of spam and abuse so that this can be the platform where developers come
          together to
          create.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Open source
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          Our Flowbite Security Lab is a world-class security R&D team. We inspire and enable the community to secure
          open source.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path
              d="M11 3a1 1 0 10-2 0v1a1 1 0 102 0V3zM15.657 5.757a1 1 0 00-1.414-1.414l-.707.707a1 1 0 001.414 1.414l.707-.707zM18 10a1 1 0 01-1 1h-1a1 1 0 110-2h1a1 1 0 011 1zM5.05 6.464A1 1 0 106.464 5.05l-.707-.707a1 1 0 00-1.414 1.414l.707.707zM5 10a1 1 0 01-1 1H3a1 1 0 110-2h1a1 1 0 011 1zM8 16v-1h4v1a2 2 0 11-4 0zM12 14c.015-.34.208-.646.477-.859a4 4 0 10-4.954 0c.27.213.462.519.476.859h4.002z" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Premium products
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          We embody the shift toward investments in safe and secure software design practices with our world-class
          front-end
          products.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M4 2a2 2 0 00-2 2v11a3 3 0 106 0V4a2 2 0 00-2-2H4zm1 14a1 1 0 100-2 1 1 0 000 2zm5-1.757l4.9-4.9a2 2 0 000-2.828L13.485 5.1a2 2 0 00-2.828 0L10 5.757v8.486zM16 18H9.071l6-6H16a2 2 0 012 2v2a2 2 0 01-2 2z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Customization
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          It's easy to customize and style Flowbite. Tweak the look and feel of your UI with CSS/Less, and add major
          features with
          HTML.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M12.395 2.553a1 1 0 00-1.45-.385c-.345.23-.614.558-.822.88-.214.33-.403.713-.57 1.116-.334.804-.614 1.768-.84 2.734a31.365 31.365 0 00-.613 3.58 2.64 2.64 0 01-.945-1.067c-.328-.68-.398-1.534-.398-2.654A1 1 0 005.05 6.05 6.981 6.981 0 003 11a7 7 0 1011.95-4.95c-.592-.591-.98-.985-1.348-1.467-.363-.476-.724-1.063-1.207-2.03zM12.12 15.12A3 3 0 017 13s.879.5 2.5.5c0-1 .5-4 1.25-4.5.5 1 .786 1.293 1.371 1.879A2.99 2.99 0 0113 13a2.99 2.99 0 01-.879 2.121z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Features
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          Start with thousands of actions and applications from our community to help you build, improve, and accelerate
          your
          automated workflows.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M6.625 2.655A9 9 0 0119 11a1 1 0 11-2 0 7 7 0 00-9.625-6.492 1 1 0 11-.75-1.853zM4.662 4.959A1 1 0 014.75 6.37 6.97 6.97 0 003 11a1 1 0 11-2 0 8.97 8.97 0 012.25-5.953 1 1 0 011.412-.088z"
              clip-rule="evenodd" />
            <path fill-rule="evenodd"
              d="M5 11a5 5 0 1110 0 1 1 0 11-2 0 3 3 0 10-6 0c0 1.677-.345 3.276-.968 4.729a1 1 0 11-1.838-.789A9.964 9.964 0 005 11zm8.921 2.012a1 1 0 01.831 1.145 19.86 19.86 0 01-.545 2.436 1 1 0 11-1.92-.558c.207-.713.371-1.445.49-2.192a1 1 0 011.144-.83z"
              clip-rule="evenodd" />
            <path fill-rule="evenodd"
              d="M10 10a1 1 0 011 1c0 2.236-.46 4.368-1.29 6.304a1 1 0 01-1.838-.789A13.952 13.952 0 009 11a1 1 0 011-1z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Data privacy
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          Flowbite is committed to developer privacy and provides a high standard of privacy protection to all our
          customers and
          team.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M18 10a8 8 0 11-16 0 8 8 0 0116 0zm-2 0c0 .993-.241 1.929-.668 2.754l-1.524-1.525a3.997 3.997 0 00.078-2.183l1.562-1.562C15.802 8.249 16 9.1 16 10zm-5.165 3.913l1.58 1.58A5.98 5.98 0 0110 16a5.976 5.976 0 01-2.516-.552l1.562-1.562a4.006 4.006 0 001.789.027zm-4.677-2.796a4.002 4.002 0 01-.041-2.08l-.08.08-1.53-1.533A5.98 5.98 0 004 10c0 .954.223 1.856.619 2.657l1.54-1.54zm1.088-6.45A5.974 5.974 0 0110 4c.954 0 1.856.223 2.657.619l-1.54 1.54a4.002 4.002 0 00-2.346.033L7.246 4.668zM12 10a2 2 0 11-4 0 2 2 0 014 0z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          Support 24/7
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          We provide high-quality services all around the world with personal assistance through our 24/7 premium
          service.
        </p>
      </div>

      <div class="p-6 bg-white rounded-lg shadow-sm dark:bg-gray-700">
        <div class="flex items-center justify-center w-12 h-12 rounded-lg bg-primary-100 dark:bg-primary-900">
          <svg class="w-8 h-8 text-primary-700 dark:text-primary-400" fill="currentColor" viewBox="0 0 20 20"
            xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M10 18a8 8 0 100-16 8 8 0 000 16zM4.332 8.027a6.012 6.012 0 011.912-2.706C6.512 5.73 6.974 6 7.5 6A1.5 1.5 0 019 7.5V8a2 2 0 004 0 2 2 0 011.523-1.943A5.977 5.977 0 0116 10c0 .34-.028.675-.083 1H15a2 2 0 00-2 2v2.197A5.973 5.973 0 0110 16v-2a2 2 0 00-2-2 2 2 0 01-2-2 2 2 0 00-1.668-1.973z"
              clip-rule="evenodd" />
          </svg>
        </div>
        <h3 class="mt-5 text-xl font-bold dark:text-white">
          GDPR
        </h3>
        <p class="mt-2 text-base text-gray-500 dark:text-gray-400">
          Our project is full GDPR compliant. GDPR compliance is shown through our strong and transparent actions, not
          through
          certifications.
        </p>
      </div>
    </div>
  </div>
</Section>


<!-- PRICING & FEATURES -->
<Section back="lighter">
  <div class="max-w-screen-xl px-4 py-8 mx-auto sm:py-16 lg:py-24">
    <div class="bg-white rounded-lg shadow-sm lg:grid lg:grid-cols-3 dark:bg-gray-700">
      <div class="col-span-2 p-6 lg:p-8">
        <h2 class="mb-1 text-2xl font-bold text-gray-900 sm:text-3xl dark:text-white">
          Pricing built for all businesses.
        </h2>
        <p class="text-lg text-gray-500 dark:text-gray-400">
          Best for large scale uses and extended
          redistribution rights.
        </p>

        <div class="grid gap-4 mt-4 lg:mt-6 sm:grid-cols-2 md:grid-cols-3">
          <!-- List -->
          <ul role="list" class="space-y-4 dark:text-white">
            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">A/B Testing</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">24/7 Chat Support</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Custom Branding</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Creative Assistant</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Website Builder</span>
            </li>
          </ul>

          <!-- List -->
          <ul role="list" class="hidden space-y-4 dark:text-white sm:block">
            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Customer Builder</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Marketing CRM</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Custom Templates</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Creative Assistant</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Multivariate Testing</span>
            </li>
          </ul>

          <!-- List -->
          <ul role="list" class="hidden space-y-4 dark:text-white lg:block">
            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Advanced Tools</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Multivariate Testing</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Reporting</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Custom Templates</span>
            </li>

            <li class="flex space-x-2.5">
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M16.707 5.293a1 1 0 010 1.414l-8 8a1 1 0 01-1.414 0l-4-4a1 1 0 011.414-1.414L8 12.586l7.293-7.293a1 1 0 011.414 0z"
                  clip-rule="evenodd" />
              </svg>

              <span class="leading-tight text-gray-500 dark:text-gray-400">Dynamic Content</span>
            </li>
          </ul>
        </div>
      </div>

      <div class="flex p-6 text-center bg-gray-50 lg:p-8 dark:bg-gray-700">
        <div class="self-center w-full">
          <div class="text-5xl font-extrabold text-gray-900 dark:text-white">$99</div>
          <div class="mt-1 mb-4 text-gray-500 text-light dark:text-gray-400">per month</div>
          <a href="#"
            class="flex justify-center text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-bue-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-primary-600 dark:hover:bg-primary-700">Buy
            now</a>
          <a href="#"
            class="flex items-center justify-center mt-4 hover:underline font-medium text-gray-500 hover:text-gray-900 dark:text-gray-400 dark:hover:text-gray-300">
            View team pricing
            <svg class="w-5 h-5 ml-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
              <path fill-rule="evenodd"
                d="M12.293 5.293a1 1 0 011.414 0l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414-1.414L14.586 11H3a1 1 0 110-2h11.586l-2.293-2.293a1 1 0 010-1.414z"
                clip-rule="evenodd"></path>
            </svg>
          </a>
        </div>
      </div>
    </div>
  </div>
</Section>


