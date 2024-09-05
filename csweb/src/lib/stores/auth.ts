import { get } from "svelte/store";
import type { Writable } from "svelte/store"
import { persisted } from 'svelte-persisted-store'
import type { LoginResult } from '$lib/graphql/generated/sdk'


const getInitialValue = () : LoginResult => {
    const loginResult: LoginResult = {
        status: {
            success: false,
            message: []
        }, 
        token: "",
        user: {
            email: "",
            firstName: "",
            lastName: "",
            id: ""
        }
    }

    return loginResult
}

const authStore: Writable<LoginResult> = persisted('loginResult', getInitialValue());


export const setLoggedInUser = (result:LoginResult) => {
    authStore.set(result)
}

export const logout = () => {
    authStore.set(getInitialValue())
}

export const isAuthenticated = () => {
    const l = get(authStore)

    return l.status?.success.valueOf()
}

export const currentUser = ():LoginResult => {
    return get(authStore)
}