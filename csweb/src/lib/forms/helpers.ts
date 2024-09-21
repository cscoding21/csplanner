import type { ObjectSchema, Maybe, AnyObject } from 'yup';
import type { List, ListItem } from '$lib/graphql/generated/sdk';
import type { SelectOptionType } from 'flowbite-svelte';
import { sortListValues } from '$lib/services/list';

function isObjKey<T>(key: any, obj: any): key is keyof T {
	return key in obj;
}

export function coalesceToType<T>(input: any, schema: ObjectSchema<Maybe<AnyObject>>): T {
	const out: T = <T>{};
	const describe = schema.describe();

	Object.keys(input).forEach((key) => {
		if (isObjKey<T>(key, describe.fields)) {
			out[key] = input[key];
		}
	});

	return out;
}

export const parseErrors = (err: any) => {
	return err.inner.reduce((acc: any, err: any) => {
		return { ...acc, [err.path]: err.message };
	}, {});
};

export const mergeErrors = (err1: any, err2: any) => {
	return {
		...err1,
		...err2
	};
};

/**
 * Helper method to turn the values from a list into select options
 * @param list the list object to extract values from
 * @returns a select-box-friendly list of options
 */
export const findSelectOptsFromList = (list: List): SelectOptionType<string>[] => {
	const vals = sortListValues(list);
	let opts: SelectOptionType<string>[] = [];

	opts = vals.map((r: ListItem) => ({
		name: r.name,
		value: r.value as string
	})) as SelectOptionType<string>[];

	return opts;
};
