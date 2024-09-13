/**
 * A multi-purpose function that assess whether or not an input is a meaningful value
 * @param input any input that the user wants to determine is valuable
 * @returns true if the input looks like a meaningful value.  False otherwise
 */
export const is = (input: any): boolean => {
	if (input === null) {
		return false;
	}

	if (Array.isArray(input) && input.length > 0) {
		return true;
	} else if (Array.isArray(input) && input.length === 0) {
		return false;
	}

	if (input) {
		return true;
	}

	return false;
};

/**
 * A simple function to determine if a variable is a function or not.
 * @param functionToCheck a variable to evaluate
 * @returns true if the input is a function.  False otherwise
 */
export const isFunction = (functionToCheck: any): boolean => {
	return functionToCheck && {}.toString.call(functionToCheck) === '[object Function]';
};
