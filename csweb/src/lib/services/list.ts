import { FindAllListsDocument, GetListDocument } from '$lib/graphql/generated/sdk';
import type { List, ListResults, ListItem } from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';

/**
 * return all lists in the system
 * @returns a promise of all lists
 */
export const findAllLists = async (): Promise<ListResults> => {
	const client = getApolloClient();

	return client
		.query({ query: FindAllListsDocument })
		.then((res) => {
			if (res) {
				return res.data.findAllLists;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * return a single list based on its name or ID
 * @param nameOrID - the name of ID of the list
 * @returns a single list object corresponding to the passed in identifier
 */
export const getList = async (nameOrID: string): Promise<List> => {
	const client = getApolloClient();

	return client
		.query({ query: GetListDocument, variables: { nameOrID } })
		.then((res) => {
			if (res) {
				return res.data.getList;
			}
		})
		.catch((err) => {
			return err;
		});
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