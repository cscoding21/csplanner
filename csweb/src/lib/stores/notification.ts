import { writable } from 'svelte/store';
import { findNotifications } from '$lib/services/notification';
import type { Notification } from '$lib/graphql/generated/sdk';

const getInitialValue = (): Notification[] => {
	return [];
};

export const notificationStore = writable(getInitialValue());

export const refreshNotificationStore = async () => {
	const res = await findNotifications({ paging: { pageNumber: 1, resultsPerPage: 100 } });

	if (res && res.results) {
		notificationStore.set(res.results);
	}
};
