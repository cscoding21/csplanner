<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import { Dropdown, DropdownItem, Avatar, DropdownHeader, DropdownGroup } from 'flowbite-svelte';
	import { BellOutline, CheckCircleOutline, EyeOutline } from 'flowbite-svelte-icons';
	import {
		subscribeToUpdates,
		getNotificationHeadline,
		setNotificationsRead
	} from '$lib/services/notification';
	import { notificationStore, refreshNotificationStore } from '$lib/stores/notification';
	import { formatDateTime, truncateText } from '$lib/utils/format';
	import { orgStore } from '$lib/stores/organization';

	onMount(async () => {
		await refresh();
	});

	onDestroy(() => {
		unsubHandle.unsubscribe();
	});

	const refresh = async () => {
		await refreshNotificationStore();
	};

	const markRead = async (id: string) => {
		setNotificationsRead([id]).then(async () => {
			await refreshNotificationStore();
		});
	};

	let sub = subscribeToUpdates();

	let unsubHandle = sub.subscribe({
		next: () => {
			refreshNotificationStore();
		}
	});

	let notificationIndicator = $state($notificationStore?.some((n) => !n.isRead));
	let orgIdentifier = $state($orgStore.name ? ": " + $orgStore.name : "")
	let title = $state(
		$notificationStore?.filter((n) => !n.isRead).length > 0
			? '(' + $notificationStore?.filter((n) => !n.isRead).length + ') csPlanner' + orgIdentifier
			: 'csPlanner' + orgIdentifier
	);
</script>

<svelte:head>
	<title>{title}</title>
</svelte:head>

{#await $notificationStore}
	<span>Loading...</span>
{:then promiseData}
	<div
		id="bell"
		class="inline-flex items-center text-center text-sm font-medium text-gray-500 hover:text-gray-900 focus:outline-none dark:text-gray-400 dark:hover:text-white"
	>
		<BellOutline class="h-5kj w-5" />
		{#if notificationIndicator}
			<div class="relative flex">
				<div
					class="relative -top-2 end-3 inline-flex h-3 w-3 rounded-full border-2 border-white bg-red-500 dark:border-gray-900"
				></div>
			</div>
		{/if}
	</div>

	<Dropdown
		triggeredBy="#bell"
		class="w-full max-w-sm divide-y divide-gray-100 rounded shadow dark:divide-gray-700 dark:bg-gray-800"
	>
		<DropdownHeader>Notifications</DropdownHeader>
		<DropdownGroup>
		{#if promiseData && promiseData.length > 0}
			{#each promiseData as n (n)}
				<DropdownItem class="flex space-x-4 rtl:space-x-reverse">
					<Avatar src={n.initiatorProfileImage || ''} dot={{ color: 'bg-gray-300' }} />
					<div class="w-full ps-3">
						<div class="mb-1.5 text-sm text-gray-500 dark:text-gray-400">
							{@html getNotificationHeadline(n)}
						</div>
						<div class="mb-1.5 text-xs text-gray-500 dark:text-gray-400">
							&quot;{truncateText(n.text || '', 120)}&quot;
						</div>
						<div class="text-xs text-primary-600 dark:text-primary-500">
							{formatDateTime(n.updatedAt.valueOf())}
						</div>
						{#if !n.isRead}
							<button
								onclick={() => {
									markRead(n.id);
								}}
								class="text-xs">mark as read</button
							>
						{/if}
					</div>
				</DropdownItem>
			{/each}
		{:else}
			<DropdownItem class="flex space-x-4 rtl:space-x-reverse">
				<div class="inline-flex min-w-full items-center text-sm">
					<CheckCircleOutline class="me-2 h-4 w-4 text-gray-500 dark:text-gray-400" />
					You're all caught up
				</div>
			</DropdownItem>
		{/if}
		</DropdownGroup>
		<DropdownHeader>
		<a
			href="/"
			class="-my-1 block bg-gray-50 py-2 text-center text-sm font-medium text-gray-900 hover:bg-gray-100 dark:bg-gray-800 dark:text-white dark:hover:bg-gray-700"
		>
			<div class="inline-flex items-center">
				<EyeOutline class="me-2 h-4 w-4 text-gray-500 dark:text-gray-400" />
				View all
			</div>
		</a>
		</DropdownHeader>
	</Dropdown>
{/await}
