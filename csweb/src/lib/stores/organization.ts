import { writable, type Writable } from 'svelte/store';
import { getOrganization } from '$lib/services/organization';
import { type Organization } from '$lib/graphql/generated/sdk';

const getInitialValue = (): Organization => {
    return {} as Organization;
};

export const orgStore = writable(getInitialValue());

export const refreshOrg = async () => {
    console.log("refreshOrg")
    const res = await getOrganization();

    orgStore.set(res);
};

export const teardownOrg = async () => {
    orgStore.set(getInitialValue())
}


export const getOrg = ():Writable<Organization> => {
    if(orgStore) {
        return orgStore
    }

    refreshOrg()

    return orgStore
}

