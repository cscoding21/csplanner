/**
 * Applies consistent string formatting for date values
 * @param input The date to be formatted
 * @returns A cformatted date value
 */
export const formatDate = (input: any): string => {
	if (!input) {
		return '';
	}

	if (input instanceof Date) {
		return input.toLocaleDateString();
	}

	const o = new Date(Date.parse(input.toString()));
	return o.toLocaleDateString();
};

/**
 * Applies consistent formatting rules to date and time values
 * @param input A date time value to be formatted
 * @returns the formatted date/time
 */
export const formatDateTime = (input: any): string => {
	if (!input) {
		return '';
	}

	if (input instanceof Date) {
		return getDateDisplay(input) + ' at ' + input.toLocaleTimeString();
	}

	const o = new Date(Date.parse(input.toString()));
	return getDateDisplay(o) + ' at ' + o.toLocaleTimeString();
};

/**
 * Takes a date that something occurred and returns a more human-friendly representation.
 * Today or Yesterday will be returned if appropriate.  Otherwise, the date in dd/mm/yyyy
 * @param input the date of the event
 * @returns A user-friendly representation of when the event occured
 */
const getDateDisplay = (input: Date): string => {
	const today = new Date();
	const yesterday = new Date(new Date().setDate(new Date().getDate() - 1));

	if (input.toLocaleDateString() === today.toLocaleDateString()) {
		return 'Today';
	}

	if (input.toLocaleDateString() === yesterday.toLocaleDateString()) {
		return 'Yesterday';
	}

	return input.toLocaleDateString();
};

/**
 * A formatter that is used to parse a number into a currency using standard US settings.  Syntax for use is:
 * formatCurrency.format(input)
 */
export const formatCurrency = new Intl.NumberFormat('en-US', {
	style: 'currency',
	currency: 'USD',

	// These options are needed to round to whole numbers if that's what you want.
	//minimumFractionDigits: 0, // (this suffices for whole numbers, but will print 2500.10 as $2,500.1)
	maximumFractionDigits: 0 // (causes 2500.99 to be printed as $2,501)
});

/**
 * A formatter that is used to parse a number into a percent using standard settings.  Syntax for use is:
 * formatPercent.format(input)
 */
export const formatPercent = new Intl.NumberFormat('default', {
	style: 'percent',
	minimumFractionDigits: 2,
	maximumFractionDigits: 2
});

/**
 * Takes a user's name and returns the initials to be used as backups for profile pics
 * @param name The name of the user
 * @returns a two-letter abbreviation of the passed in name
 */
export const getInitialsFromName = (name: string) => {
	if (!name) {
		return 'XX';
	}

	if (name.indexOf(' ') === -1) {
		if (name.length === 1) {
			return name.substring(0, 1).toUpperCase() + 'X';
		}

		return name.substring(0, 2).toUpperCase();
	}

	const names = name.split(' ');

	return (
		names[0].substring(0, 1).toUpperCase() + names[names.length - 1].substring(0, 1).toUpperCase()
	);
};

/**
 * Takes an array of strings and returns a comma separated version based on the count
 * @param list the list that should be returned as comma separated
 * @returns the string representation of the list
 */
export const formatToCommaSepList = (list: string[]): string => {
	if (list.length === 0) {
		return '';
	} else if (list.length === 1) {
		return list[0];
	}

	return list.join(', ');
};

/**
 * Takes a word that is describing a list of elements and returns a plural version of that word if appropriate
 * @param input A string that should be converted to plural if appropriate.  This should be a single wword
 * @param count the number of items that the string is describing.
 * @returns A string that is plural based on engligh rules if count is greater than one.  Otherwise the input is sent back unmodified
 */
export const pluralize = (input: string, count: number): string => {
	if (count === 1) {
		return input;
	}

	if (input.endsWith('y')) {
		return input.substring(0, input.length - 1) + 'ies';
	}

	return input + 's';
};

/**
 * Returns a shortened version of the passed in test based on the length
 * @param text The text to shorten
 * @param length The length you'd like the test to be truncated to
 * @returns a shortened version of the text with an elipses at the end if appropriate
 */
export const truncateText = (text: string, length: number): string => {
	if (text.length > length) {
		return text.substring(0, length) + '...';
	}

	return text;
};
