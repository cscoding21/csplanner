/**
 * SurrealDB creates object IDs with a prefix of the object name, a colon, and a string ID.
 * Because the colon can cause problem for Svelte arrays, this function replaces it with an undersore
 * @param id - an ID string that was created by SurrealDB
 * @returns a cleaned id value
 */
export const normalizeID = (id: string) => {
	return id.replace(':', '_');
};

/**
 * Takes a value that has previously been normalized for Svelte and returns it in its original form
 * @param id a string value that has previously been normalized
 * @returns a string ID returned to its original form
 */
export const denormalizeID = (id: string) => {
	return id.replace('_', ':');
};
