<script lang="ts">
	  import Section from "./Section.svelte";

    interface Props {
        back: "lighter"|"darker"
    }
    let { back }:Props = $props()
</script>

<Section {back}>
  <div class="max-w-screen-xl px-4 py-8 mx-auto sm:py-16 lg:py-24 pt-20 sm:pt-24 lg:pt-32">
    <div class="grid items-center gap-8 mb-8 lg:mb-16 lg:gap-12 lg:grid-cols-12">
      <div class="col-span-6 text-center sm:mb-6 lg:text-left lg:mb-0">
        <a href="/"
          class="inline-flex items-center justify-between px-1 py-1 pr-4 mb-6 text-sm text-gray-700 bg-gray-100 rounded-full dark:bg-gray-800 dark:text-white hover:bg-gray-200 dark:hover:bg-gray-700"
          >
          <span class="px-3 py-1 mr-3 text-xs text-white rounded-full bg-primary-700">New</span> <span
            class="text-sm font-medium">csPlanner is coming soon!</span>
          <svg class="w-5 h-5 ml-2" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
            <path fill-rule="evenodd"
              d="M7.293 14.707a1 1 0 010-1.414L10.586 10 7.293 6.707a1 1 0 011.414-1.414l4 4a1 1 0 010 1.414l-4 4a1 1 0 01-1.414 0z"
              clip-rule="evenodd"></path>
          </svg>
        </a>
        <h1
          class="mb-4 text-3xl font-extrabold leading-none tracking-tight text-gray-900 sm:text-4xl md:text-5xl xl:text-6xl dark:text-white">
          Hyper-optimizd software project planning
        </h1>
        <p class="max-w-xl mx-auto mb-6 text-gray-500 lg:mx-0 xl:mb-8 md:text-lg xl:text-xl dark:text-gray-400">
          csPlanner is an elegant tool that enables efficient and iterative project planning
          doing the hard work for you while providing business-critical insights.
        </p>
        <form class="max-w-lg mx-auto lg:ml-0" action="#">
          <label for="email" class="mb-2 sr-only">Email</label>
          <div class="relative">
            <div class="absolute inset-y-0 left-0 flex items-center pl-3 pointer-events-none">
              <svg aria-hidden="true" class="w-5 h-5 text-gray-500 dark:text-gray-400"
                xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20" fill="currentColor">
                <path fill-rule="evenodd" d="M10 9a3 3 0 100-6 3 3 0 000 6zm-7 9a7 7 0 1114 0H3z" clip-rule="evenodd" />
              </svg>
            </div>

            <input type="email" id="email"
              class="block w-full py-4 pl-10 pr-4 text-base text-gray-900 border border-gray-300 rounded-lg bg-gray-50 focus:ring-primary-500 focus:border-primary-500 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500"
              placeholder="Enter your email here" required={undefined}>
          </div>
        </form>
      </div>
      <div class="col-span-6">
        <img src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/hero/search-mockup.png" class="dark:hidden"
          alt="mockup">
        <img src="https://flowbite.s3.amazonaws.com/blocks/marketing-ui/hero/search-mockup-dark.png"
          class="hidden dark:block" alt="mockup dark">
      </div>
    </div>

    <div class="grid gap-8 sm:gap-12 md:grid-cols-3">
      <div class="flex justify-center">
        <svg class="w-6 h-6 mr-3 text-primary-700 dark:text-primary-500 shrink-0" fill="currentColor"
          viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd"
            d="M4 2a2 2 0 00-2 2v11a3 3 0 106 0V4a2 2 0 00-2-2H4zm1 14a1 1 0 100-2 1 1 0 000 2zm5-1.757l4.9-4.9a2 2 0 000-2.828L13.485 5.1a2 2 0 00-2.828 0L10 5.757v8.486zM16 18H9.071l6-6H16a2 2 0 012 2v2a2 2 0 01-2 2z"
            clip-rule="evenodd"></path>
        </svg>
        <div>
          <h3 class="mb-1 text-lg font-semibold leading-tight text-gray-900 dark:text-white">
            Project Scheduling
          </h3>
          <p class="text-gray-500 dark:text-gray-400">
            Auto-scheduling algorithm calculates end dates for each project considering
            resource skills, contention, and focus.
          </p>
        </div>
      </div>
      <div class="flex justify-center">
        <svg class="w-6 h-6 mr-3 text-primary-700 dark:text-primary-500 shrink-0" fill="currentColor"
          viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd"
            d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
            clip-rule="evenodd"></path>
        </svg>
        <div>
          <h3 class="mb-1 text-lg font-semibold leading-tight text-gray-900 dark:text-white">
            Resource Management
          </h3>
          <p class="text-gray-500 dark:text-gray-400">
            Understand the skills, cost, and capacity of the individual resources in your roster.
          </p>
        </div>
      </div>
      <div class="flex justify-center">
        <svg class="w-6 h-6 mr-3 text-primary-700 dark:text-primary-500 shrink-0" fill="currentColor"
          viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg">
          <path fill-rule="evenodd"
            d="M3 3a1 1 0 000 2v8a2 2 0 002 2h2.586l-1.293 1.293a1 1 0 101.414 1.414L10 15.414l2.293 2.293a1 1 0 001.414-1.414L12.414 15H15a2 2 0 002-2V5a1 1 0 100-2H3zm11 4a1 1 0 10-2 0v4a1 1 0 102 0V7zm-3 1a1 1 0 10-2 0v3a1 1 0 102 0V8zM8 9a1 1 0 00-2 0v2a1 1 0 102 0V9z"
            clip-rule="evenodd"></path>
        </svg>
        <div>
          <h3 class="mb-1 text-lg font-semibold leading-tight text-gray-900 dark:text-white">
            ROI Tracking
          </h3>
          <p class="text-gray-500 dark:text-gray-400">
            Understand the cost and value of your projects across your portfolio, 
            informing priority and go/no-go decisions.
          </p>
        </div>
      </div>
    </div>
  </div>
</Section>