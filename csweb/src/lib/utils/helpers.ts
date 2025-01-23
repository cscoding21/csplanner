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

