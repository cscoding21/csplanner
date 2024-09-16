import { writable, get } from 'svelte/store';
import type { ListResults, List, ListItem } from '$lib/graphql/generated/sdk';
import { findAllLists } from '$lib/services/list';

const getInitialValue = (): ListResults => {
	const listResult: ListResults = {
		paging: {},
		filters: {},
		results: []
	};

	return listResult;
};

const getEmptyList = (): List => {
	return { id: '', name: '', values: [] };
};

export const listStore = writable(getInitialValue());
export const skillListStore = writable(getEmptyList());
export const fundingSourceListStore = writable(getEmptyList());

/**
 * populate the list store with all list data from the backend service
 */
export const refreshListStore = async () => {
	await findAllLists().then((r) => {
		listStore.set(r);

		getListFromStore('Skills').then((sl) => {
			if (sl !== null && sl !== undefined) {
				skillListStore.set(sl);
			}
		});

		getListFromStore('Funding Source').then((fs) => {
			if (fs !== null && fs !== undefined) {
				fundingSourceListStore.set(fs);
			}
		});
	});
};

/**
 * Retrieve a single list from the store
 * @param nameOrID - the name of ID of the list to retrieve
 * @returns a list object corresponding to the passed in identifier
 */
export const getListFromStore = async (nameOrID: string) => {
	const lists = get(listStore);

	if (lists == null || lists.results == null || lists.results.length === 0) {
		return;
	}

	for (let i = 0; i < lists.results.length; i++) {
		const item = lists.results[i];

		if (item.id === nameOrID || item.name === nameOrID) {
			return item;
		}
	}

	return null;
};

/**
 * return a list of options in the list sorted by their name in ascending order
 * @param list - the list whose values to sort
 * @returns values fron the list sorted by their name
 */
export const sortListValues = (list: List): ListItem[] => {
	if (list == null || list.values == null || list.values.length === 0) {
		return list.values;
	}

	const sortedValues = JSON.parse(JSON.stringify(list.values));
	sortedValues.sort((a: ListItem, b: ListItem) => a.name.localeCompare(b.name));

	return sortedValues;
};
