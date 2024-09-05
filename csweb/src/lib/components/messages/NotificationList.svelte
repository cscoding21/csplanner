<script lang="ts">
    import { onMount, onDestroy } from "svelte";
    import { Dropdown, DropdownItem, Avatar } from "flowbite-svelte";
    import { BellOutline, CheckCircleOutline, EyeOutline } from "flowbite-svelte-icons";
    import { subscribeToUpdates, getNotificationHeadline, setNotificationsRead } from "$lib/services/notification";
    import { notificationStore, refreshNotificationStore } from "$lib/stores/notification";
    import { formatDateTime, truncateText } from "$lib/utils/format";

    
    onMount(async () => {
        await refresh()
    })

    onDestroy(() => {
        unsubHandle.unsubscribe()
    })

    const refresh = async () => {
      await refreshNotificationStore()
    }

    const markRead = async (id:string) => {
      setNotificationsRead([id]).then(async () => {
        await refreshNotificationStore()
      })
    }
  
    let sub = subscribeToUpdates()

    let unsubHandle = sub.subscribe({
        next: ({ data }) => {
          refreshNotificationStore()
        }
    });

    $: notificationIndicator = $notificationStore.some(n => !n.isRead) 
    $: title = $notificationStore.filter(n => !n.isRead).length > 0 ? '(' + $notificationStore.filter(n => !n.isRead).length + ') csPlanner' : 'csPlanner'
</script>

<svelte:head>
  <title>{title}</title>
</svelte:head>

  {#await $notificationStore}
    <span>Loading...</span>
  {:then promiseData}

<div id="bell" class="inline-flex items-center text-sm font-medium text-center text-gray-500 hover:text-gray-900 focus:outline-none dark:hover:text-white dark:text-gray-400">
<BellOutline class="w-5 h-5kj" />
{#if notificationIndicator}
<div class="flex relative">
    <div class="inline-flex relative -top-2 end-3 w-3 h-3 bg-red-500 rounded-full border-2 border-white dark:border-gray-900"></div> 
</div>
{/if}
</div>


  <Dropdown triggeredBy="#bell" class="w-full max-w-sm rounded divide-y divide-gray-100 shadow dark:bg-gray-800 dark:divide-gray-700">
    <div slot="header" class="text-center py-2 font-bold">Notifications</div>
    {#if promiseData && promiseData.length > 0}
    {#each promiseData as n(n)}
    <DropdownItem class="flex space-x-4 rtl:space-x-reverse">
      <Avatar src={n.initiatorProfileImage || ""} dot={{ color: 'bg-gray-300' }} rounded />
      <div class="ps-3 w-full">
        <div class="text-gray-500 text-sm mb-1.5 dark:text-gray-400">{@html getNotificationHeadline(n)}</div>
        <div class="text-gray-500 text-xs mb-1.5 dark:text-gray-400">
          &quot;{truncateText(n.text || "", 120)}&quot;
        </div>
        <div class="text-xs text-primary-600 dark:text-primary-500">{formatDateTime(n.updatedAt.valueOf())}</div>
        {#if !n.isRead }
        <button on:click={() => { markRead(n.id) }} class="text-xs">mark as read</button>
        {/if}
      </div>
    </DropdownItem>
    {/each}
    {:else}
    <DropdownItem class="flex space-x-4 rtl:space-x-reverse">
      <div class="inline-flex items-center min-w-full text-sm">
        <CheckCircleOutline class="me-2 w-4 h-4 text-gray-500 dark:text-gray-400" />
        You're all caught up
      </div>
    </DropdownItem>
    {/if}
    <a slot="footer" href="/" class="block py-2 -my-1 text-sm font-medium text-center text-gray-900 bg-gray-50 hover:bg-gray-100 dark:bg-gray-800 dark:hover:bg-gray-700 dark:text-white">
      <div class="inline-flex items-center">
        <EyeOutline class="me-2 w-4 h-4 text-gray-500 dark:text-gray-400" />
        View all
      </div>
    </a>
  </Dropdown> 
    
  {/await}
  