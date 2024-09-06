/**
 * Creates a clone of the pass in object using completely new memory space
 * @param obj - an object that should be cloned
 * @returns a clone of the passed in object
 */
export const deepCopy = (obj: any):any => {
    return JSON.parse(JSON.stringify(obj))
}

/**
 * Remove all duplicate elements from an array
 * @param arr - the array to de-dupe
 * @returns - the de-duped array
 */
export const deDupeArray = (arr: any[]) : any[] => {
    return [...new Set(arr)]
}