import { writable } from 'svelte/store';
import { getOrganization } from '$lib/services/organization';
import { type Organization } from '$lib/graphql/generated/sdk';

const getInitialValue = (): Organization => {
    return {} as Organization;
};

export const orgStore = writable(getInitialValue());

export const refreshOrg = async () => {
    const res = await getOrganization();

    orgStore.set(res);
};

export const teardownOrg = async () => {
    orgStore.set(getInitialValue())
}

