/**
 * Creates a clone of the pass in object using completely new memory space
 * @param obj - an object that should be cloned
 * @returns a clone of the passed in object
 */
export const deepCopy = (obj: any): any => {
	return JSON.parse(JSON.stringify(obj));
};

/**
 * Remove all duplicate elements from an array
 * @param arr - the array to de-dupe
 * @returns - the de-duped array
 */
export const deDupeArray = (arr: any[]): any[] => {
	return [...new Set(arr)];
};

/**
 * Ensures that an array is instantiated and not null
 * @param a - an array that may be null
 * @returns the pass in array or an empty array
 */
export const safeArray = (a:any[]):any[] => {
	let out:any[] = []

	if (a !== null && a !== undefined && a.length > 0)
		out = out.concat(a)

	return out
}

/**
 * Return the value of a cookie by coercing the cookie string into kvps
 * and retrieving the on with a matching keey
 * @param cookieName the name of the cookies to be retrieved
 * @returns the value of the cookies
 */
export const getCookie = (cookieName: string): string => {
	const cookieArray = document.cookie.split('; ');

	for (let i = 0; i < cookieArray.length; i++) {
		const cookie = cookieArray[i];
		const [name, value] = cookie.split('=');

		if (name === cookieName) {
			return decodeURIComponent(value);
		}
	}

	return '';
};

/**
 * remove a cookie from the document by its name
 * @param cookieName the name of the cookie to delete
 */
export const deleteCookie = (cookieName: string) => {
	document.cookie = `${cookieName}= ; expires = Thu, 01 Jan 1970 00:00:00 GMT`;
};

/**
 * calls a function if it exists
 * @param func the function to call
 */
export const callIf = (func?: Function) => {
	if (func) {
		func();
	}
};

/**
 * ensure that the value is returned as a number as opposed to null or undefined
 * @param input - ensure a safe integer is returned from a nullable value
 * @returns an integer representing the input value
 */
export const safeInt = (input:any):number => {
	if(!input) {
		return 0
	}

	return input as number
}


/**
 * reshapes an array so that is suitable for pie charts
 * @param groupByArray the array to group by
 * @param groupByLabel the label field
 * @param valueField the value field
 * @returns an array with 
 */
export const csGroupBy = (groupByArray:any[], groupByLabel:string, valueField:string):any[] => {
	    // Initialize an empty array to store the grouped results
		const result:any[] = [];
    
		// Loop through each item in the input array
		for (const item of groupByArray) {
			// Get the label and its corresponding value from the current item
			const labelValue = item[groupByLabel];
			const currentValue = item[valueField];
			
			// If the label doesn't exist in the result, initialize it with 0 sum
			if (!result.hasOwnProperty(labelValue)) {
				result[labelValue] = 0;
			}
			
			// Add the current value to the accumulated sum for this label
			result[labelValue] += currentValue;
		}
		
		return result;
};

/**
 * return the first meaningful value in the args array
 * @param args a list of values to evaluate
 * @returns the first non null|undefined|empty value
 */
export const coalesce = (...args:any):any => {
	for (let i = 0; i < args.length; i++) {
		if (args[i]) {
			return args[i];
		}
	}

	return "";
}

/**
 * reload the current page
 */
export const reloadPage = () => {
	const thisPage = window.location.pathname;

	window.location.reload()
}

/**
 * test whether a value is contained in a set of options
 * @param val the value to check
 * @param opts the options to compare to the value
 * @returns true if the value is one of the options
 */
export const isOneOf = (val:any, opts:any[]):boolean => {
	if(!opts || opts.length == 0) {
		return false
	}

	for(let i = 0; i < opts.length; i++) {
		if(val === opts[i]) {
			return true
		}
	}

	return false
}

/**
 * check to see if the input is a number
 * @param input a string to check
 * @returns true if the passed in value resolves to a number
 */
export const isNumeric = (input:string):boolean => Number.isFinite(+input);
