<script lang="ts">
	import Section from "$lib/components/Section.svelte";


    let pricePerMonth = $state(9.99)
</script>


<Section back="lighter">
      <div class="max-w-screen-xl px-4 py-8 pt-20 sm:pt-28 mx-auto sm:py-16 lg:py-24 lg:pt-32">
    <div class="items-center gap-8 lg:grid lg:grid-cols-2 xl:gap-16">
      <div>
        <h2 class="text-3xl font-extrabold tracking-tight text-gray-900 sm:text-4xl dark:text-white">
          Get <span class="text-primary-500 font-normal">cs</span>Planner
        </h2>
        <p class="mt-4 text-base text-gray-500 sm:text-xl dark:text-gray-400">
          Here at Flowbite we focus on markets where technology, innovation, and capital can unlock long-term value.
        </p>

        <div class="grid gap-8 py-8 mt-8 border-t border-gray-200 lg:grid-cols-1 dark:border-gray-700 sm:grid-cols-2">
          <div class="flex">
            <div
              class="flex items-center justify-center w-12 h-12 mr-4 bg-white rounded-lg shadow-sm shrink-0 dark:bg-gray-700">
              <svg class="text-gray-900 w-7 h-7 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z"
                  clip-rule="evenodd"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900 dark:text-white">
                Industry analysis
              </h3>
              <p class="mt-2 text-gray-500 dark:text-gray-400">
                Benchmark your company against the industry to inform your growth strategy and analyze market trends
              </p>
            </div>
          </div>

          <div class="flex">
            <div
              class="flex items-center justify-center w-12 h-12 mr-4 bg-white rounded-lg shadow-sm shrink-0 dark:bg-gray-700">
              <svg class="text-gray-900 w-7 h-7 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M4 4a2 2 0 012-2h8a2 2 0 012 2v12a1 1 0 110 2h-3a1 1 0 01-1-1v-2a1 1 0 00-1-1H9a1 1 0 00-1 1v2a1 1 0 01-1 1H4a1 1 0 110-2V4zm3 1h2v2H7V5zm2 4H7v2h2V9zm2-4h2v2h-2V5zm2 4h-2v2h2V9z"
                  clip-rule="evenodd"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900 dark:text-white">
                Company analysis
              </h3>
              <p class="mt-2 text-gray-500 dark:text-gray-400">
                Uncover competitor digital strategies and deep dive into any company’s digital footprint.
              </p>
            </div>
          </div>

          <div class="flex">
            <div
              class="flex items-center justify-center w-12 h-12 mr-4 bg-white rounded-lg shadow-sm shrink-0 dark:bg-gray-700">
              <svg class="text-gray-900 w-7 h-7 dark:text-gray-400" fill="currentColor" viewBox="0 0 20 20"
                xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M7 2a2 2 0 00-2 2v12a2 2 0 002 2h6a2 2 0 002-2V4a2 2 0 00-2-2H7zm3 14a1 1 0 100-2 1 1 0 000 2z"
                  clip-rule="evenodd"></path>
              </svg>
            </div>
            <div>
              <h3 class="text-xl font-bold text-gray-900 dark:text-white">
                App analysis
              </h3>
              <p class="mt-2 text-gray-500 dark:text-gray-400">
                Analyze app industry trends and discover audience interests to learn which apps are gaining traction in
                the app
                ecosystem.
              </p>
            </div>
          </div>
        </div>
      </div>

      <!-- Pricing Card -->
      <div class="flex flex-col p-6 bg-white rounded-lg shadow-sm xl:p-8 dark:bg-gray-700">
        <div class="items-center justify-between md:flex">
          <div>
            <div class="flex justify-between mb-2">
              <h3 class="text-2xl font-bold text-gray-900 dark:text-white">
                Pro Plan
              </h3>
              <div class="flex items-center md:hidden">
                <div class="mr-1 text-xl font-extrabold text-gray-900 lg:text-5xl dark:text-white">
                  ${pricePerMonth}
                </div>
                <span class="text-gray-500 dark:text-gray-400">
                  /month
                </span>
              </div>
            </div>
            <p class="text-lg text-gray-500 dark:text-gray-400 md:mr-2">
              Best for large scale uses and extended redistribution rights.
            </p>
          </div>

          <div class="hidden text-right md:block">
            <div class="text-2xl font-extrabold text-gray-900 lg:text-5xl dark:text-white">
              ${pricePerMonth}
            </div>
            <span class="text-gray-500 dark:text-gray-400">
              per user / per month
            </span>
          </div>
        </div>

        <a href="#"
          class="text-white bg-primary-700 hover:bg-primary-800 focus:ring-4 focus:ring-primary-200 font-medium rounded-lg text-sm px-5 py-2.5 text-center my-5 lg:my-8 dark:bg-primary-600 dark:hover:bg-primary-700 dark:focus:ring-primary-900">
          Get csPlanner now
        </a>

        <div class="justify-between space-y-4 sm:space-y-0 sm:flex">
          <!-- List -->
          <ul role="list" class="space-y-4">
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">A/B Testing</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">24/7 Email & Chat Support</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Custom Branding</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Creative Assistant</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Website Builder</span>
            </li>
          </ul>
          <!-- List -->
          <ul role="list" class="space-y-4">
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Advanced Segmentation</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Multivariate Testing</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Comparative Reporting</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Custom Templates</span>
            </li>
            <li class="flex space-x-2.5">
              <!-- Icon -->
              <svg class="shrink-0 w-5 h-5 text-green-500 dark:text-green-400" fill="currentColor"
                viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
                <path fill-rule="evenodd"
                  d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z"
                  clip-rule="evenodd"></path>
              </svg>
              <span class="leading-tight text-gray-500 dark:text-gray-400">Dynamic Content</span>
            </li>
          </ul>
        </div>
      </div>
    </div>
  </div>

</Section>