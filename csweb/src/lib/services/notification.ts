import {
	FindUserNotificationsDocument,
	type NotificationResults,
	type PageAndFilter,
	SetNotificationsReadDocument,
	type Notification,
	type Status
} from '$lib/graphql/generated/sdk';
import type { FetchResult, Observable } from '@apollo/client';
import { getApolloClient } from '$lib/graphql/gqlclient';
import { gql } from '@apollo/client/core';

export const findNotifications = async (input: PageAndFilter): Promise<NotificationResults> => {
	const client = getApolloClient();

	return client
		.query({
			query: FindUserNotificationsDocument,
			variables: { input },
			fetchPolicy: 'no-cache'
		})
		.then((res) => {
			if (res) {
				return res.data.findNotifications;
			}
		})
		.catch((err) => {
			return err;
		});
};

export const subscribeToNotifications = (
	input: PageAndFilter
): Observable<FetchResult<NotificationResults>> => {
	const client = getApolloClient();

	return client.subscribe({
		query: FindUserNotificationsDocument,
		variables: { input }
	});
};

export const subscribeToUpdates = (): Observable<FetchResult<string>> => {
	const client = getApolloClient();

	const query = gql`
		subscription {
			notificationUpdate
		}
	`;

	console.log('Subscribing to updates');
	const sub = client.subscribe({
		query: query
	});

	return sub;
};

export const getNotificationHeadline = (not: Notification): string => {
	const parsed = not;

	switch (parsed.type) {
		case 1:
			return (
				'<span class="font-semibold text-gray-900 dark:text-white">' +
				parsed.initiatorName +
				'</span> mentioned you in a comment'
			);
		case 2:
			return (
				'<span class="font-semibold text-gray-900 dark:text-white">' +
				parsed.initiatorName +
				'</span>  replied to your comment'
			);
		default:
			return '';
	}
};

export const setNotificationsRead = async (ids: string[]): Promise<Status> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: SetNotificationsReadDocument, variables: { ids } })
		.then((res) => {
			if (res) {
				return res.data.setNotificationsRead;
			}
		})
		.catch((err) => {
			return err;
		});
};
