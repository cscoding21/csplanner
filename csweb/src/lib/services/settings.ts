interface UserSetting {
    value: any
}


/**
 * store a setting to be persisted for the user
 * @param key the key of the setting
 * @param value an object representation of the setting
 */
export const setSetting = (key:string, value:any) => {
    const setting:UserSetting = { value }

    localStorage.setItem(key, JSON.stringify(value))
}

/**
 * return an object from local storage based on the passed in key
 * @param key the keyof the object to retrieve
 * @param def what to return if the requested key does not exist
 * @returns the requested key or defined value
 */
export const getSetting = (key:string, def:any):any|null => {
    const out = localStorage.getItem(key)

    if(out !== null) {
        const outVal = JSON.parse(out) as UserSetting

        return outVal
    }

    return def
}


export const PROJECT_SEARCH_INPUT_FILTER = "projectSearchInputFilter"
export const PROJECT_STATUS_FILTER = "projectStatusFilter"
export const PROJECT_RESOURCE_ID_FILTER = "projectResourceIDFilter"

export const RESOURCE_SEARCH_INPUT_FILTER = "resourceSearchInputFilter"
export const RESOURCE_TYPE_FILTER = "resourceTypeFilter"
export const RESOURCE_STATUS_FILTER = "resourceStatusFilter"
export const RESOURCE_SKILLS_FILTER = "resourceSkillsFilter"