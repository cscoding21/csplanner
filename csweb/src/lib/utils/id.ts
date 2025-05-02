/**
 * SurrealDB creates object IDs with a prefix of the object name, a colon, and a string ID.
 * Because the colon can cause problem for Svelte arrays, this function replaces it with an undersore
 * @param id - an ID string that was created by SurrealDB
 * @returns a cleaned id value
 */
export const normalizeID = (id: string) => {
	if (!id) {
		return '';
	}
	return id.replaceAll(':', '_').replaceAll('@', '_').replaceAll('.', '_').replaceAll('/', '_').replaceAll('=', '');
};

/**
 * Replace a guid with underscores replacing dashes.  This allows it to be used as an DOM element ID
 * @param id - the GUID to normalize
 * @returns the passed in GUID without the dash
 */
export const normalizeGUID = (id: string) => {
	if (!id) {
		return '';
	}
	return id.replaceAll('-', '_').replaceAll(':', '_');
}

/**
 * Takes a value that has previously been normalized for Svelte and returns it in its original form
 * @param id a string value that has previously been normalized
 * @returns a string ID returned to its original form
 */
export const denormalizeID = (id: string) => {
	if (!id) {
		return '';
	}
	return id.replace('_', ':').replace('|', '@').replace('-', '.');
};


/**
 * take an unknown amout if inputs and generate a unique hash from them
 * @param args a list of things to create a hash ID for
 * @returns a hash ID
 */
export const getID = (...args:any):string => {
	let str = args.join('').toString()

	return normalizeID(btoa(str))
}


/**
 * Take name as a string and convert it to an ID format
 * @param name the name to be converted
 * @returns the converted name
 */
export const nameToID = (name:string):string => {
	const MAX_LENGTH:number = 64;
	let str = name.toLocaleLowerCase().trimEnd().trimStart().replaceAll(' ', '-').replace(/[^a-zA-Z0-9-]/g, '')

	if (str.length > MAX_LENGTH) {
		str = str.substring(0, MAX_LENGTH)
	}

	return normalizeID(str)
}


/**
 * get a new ID
 * @returns a string suitable for use as an ID
 */
export const newID = ():string => {
	return normalizeGUID(crypto.randomUUID())
}
