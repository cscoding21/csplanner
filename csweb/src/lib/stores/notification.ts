import { writable } from 'svelte/store';
import { findNotifications } from '$lib/services/notification';
import type { Notification } from '$lib/graphql/generated/sdk';

const getInitialValue = (): Notification[] => {
	return [];
};

export const notificationStore = writable(getInitialValue());

export const refreshNotificationStore = async () => {
	const response = await findNotifications({ paging: { pageNumber: 1, resultsPerPage: 100 } });

	notificationStore.set(response.data.findUserNotifications.results);
};
