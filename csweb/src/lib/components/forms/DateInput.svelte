<script lang="ts">
	import { FormErrorMessage } from '$lib/components';
	import { formatDate } from '$lib/utils/format';
	import { DatePicker } from '@svelte-plugins/datepicker';
	import { CalendarMonthOutline } from 'flowbite-svelte-icons';
	import { type Snippet } from 'svelte';
    import { isDarkMode } from '$lib/utils/darkmode';
    
	interface Props {
		fieldName?: string;
		error?: string;
		value: Date|null;
		children?: Snippet;
		update?: Function;
	}
	let { 
        children, 
        fieldName, 
        error = $bindable(), 
        value = $bindable(), 
        update 
    }: Props = $props();

  const toggleDatePicker = () => (isOpen = !isOpen);

  const onChange = () => {
    startDate = new Date(value as Date)
  };

  let isOpen = $state(false);
  let startDate = $state(value || new Date())
  let formattedValue = $derived(formatDate(startDate))

  let theme = $state(isDarkMode() ? "custom-datepicker-dark" : "custom-datepicker")
</script>

<div class="mb-6">
    {#if fieldName}
    <label for="email" class="mb-2 block text-sm font-medium text-gray-900 dark:text-white"
        >{fieldName}</label
    >
    {/if}

    <DatePicker 
        bind:isOpen={isOpen}
        bind:startDate={value}
        theme={theme}
        enablePastDates={false}
        enableFutureDates={true}
        onDateChange={() => onChange()}
        >
        <div class="flex">
        <span
			class="inline-flex items-center rounded-s-md border border-e-0 border-gray-300 bg-gray-200 px-3 text-sm text-gray-900 dark:border-gray-600 dark:bg-gray-600 dark:text-gray-400"
		>
			<CalendarMonthOutline size="sm" />
		</span>
        <input 
            type="text" 
            class="block w-full min-w-0 flex-1 rounded-none rounded-e-lg border border-gray-300 bg-gray-50 p-2.5 text-sm text-gray-900 focus:border-blue-500 focus:ring-blue-500 dark:border-gray-600 dark:bg-gray-700 dark:text-white dark:placeholder-gray-400 dark:focus:border-blue-500 dark:focus:ring-blue-500"
            placeholder="Select date" 
            value={formattedValue} 
            onclick={toggleDatePicker} />
        </div>
    </DatePicker>

	<FormErrorMessage message={error} />

	{#if children}
	<small class="">{@render children()}</small>
	{/if}
</div>


<style>
    :global(.datepicker[data-picker-theme="custom-datepicker"]) {
      --datepicker-container-background: #fff;
    }

    :global(.datepicker[data-picker-theme="custom-datepicker-dark"]) {
      --datepicker-container-background: rgb(55 65 81 / var(--tw-bg-opacity, 1));
      --datepicker-container-border: 1px solid rgb(75 85 99 / var(--tw-border-opacity, 1));
  
      --datepicker-calendar-header-text-color: #fff;
      --datepicker-calendar-dow-color: #fff;
      --datepicker-calendar-day-color: #fff;
      --datepicker-calendar-day-color-disabled: #888;
      --datepicker-calendar-range-selected-background: #999;
  
      --datepicker-calendar-header-month-nav-background-hover: #999;
      --datepicker-calendar-header-month-nav-icon-next-filter: invert(100);
      --datepicker-calendar-header-month-nav-icon-prev-filter: invert(100);
      --datepicker-calendar-header-year-nav-icon-next-filter: invert(100);
      --datepicker-calendar-header-year-nav-icon-prev-filter: invert(100);
    }
  </style>
