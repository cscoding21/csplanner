import { writable } from "svelte/store";
import type { Resource, ResourceSnapshot } from "$lib/graphql/generated/sdk";
import { findAllResources } from "$lib/services/resource";
//import { resourceSnapshot } from "$lib/services/portfolio";


const getInitialValue = () : Resource[] => {
    const resources: Resource[] = []

    return resources
}

const getAllocationInitialValue = () : ResourceSnapshot => {
    const resourceAllocation: ResourceSnapshot = { scheduledResources: [] }

    return resourceAllocation
}

export const resourceStore = writable(getInitialValue());
export const resourceAllocationStore = writable(getAllocationInitialValue());

export const refreshResourceStore = async () => {
    findAllResources().then(r => {
        resourceStore.set(r.data.findAllResources.results)
    })
}

/*
export const refreshResourceAllocationStore = async () => {
    resourceSnapshot().then(r => {
        resourceAllocationStore.set(r.data.resourceSnapshot)
    })
}
*/