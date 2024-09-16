import type {
	UpdateResource,
	UpdateSkill,
	Resource,
	ResourceResults,
	Status,
	CreateResourceResult
} from '$lib/graphql/generated/sdk';

import {
	FindAllResourcesDocument,
	GetResourceDocument,
	CreateResourceDocument,
	DeleteResourceDocument,
	UpdateResourceSkillDocument,
	DeleteResourceSkillDocument,
	UpdateResourceDocument
} from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';

/**
 * The function retrieves all available resources
 * @returns all resources in the system
 */
export const findAllResources = async (): Promise<ResourceResults> => {
	const client = getApolloClient();
	return client
		.query({ query: FindAllResourcesDocument })
		.then((res) => {
			if (res) {
				return res.data.findAllResources;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Return a single resource based on its ID
 * @param id - the ID of the resource to get
 * @returns a single resource based on the passed-in ID
 */
export const getResource = async (id: string): Promise<Resource> => {
	const client = getApolloClient();

	return client
		.query({ query: GetResourceDocument, variables: { id } })
		.then((res) => {
			if (res) {
				return res.data.getResource;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Adds a new resources to the system
 * @param input An UpdateResource object with complete properties
 * @returns An update status to confirm creation success
 */
export const createResource = async (input: UpdateResource): Promise<CreateResourceResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: CreateResourceDocument, variables: { input } })
		.then((res) => {
			if (res) {
				return res.data.createResource;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Updates the properties of an existing resource
 * @param input An UpdateResource object with complete properties
 * @returns An update status to confirm creation success
 */
export const updateResource = async (input: UpdateResource): Promise<CreateResourceResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateResourceDocument, variables: { input } })
		.then((res) => {
			if (res) {
				return res.data.updateResource;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Removes a resource from the system, but leaves the record in tact to maintain referential integrity
 * @param id the ID of the resource to delete
 * @returns An update status to confirm deletion success
 */
export const deleteResource = async (id: string): Promise<Status> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: DeleteResourceDocument, variables: { id } })
		.then((res) => {
			if (res) {
				return res.data.deleteResource;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Creates or updates a skill entry for a given resource
 * @param input - an UpdateSkill object with the skill and proficiency specific to the resource
 * @returns An update status
 */
export const updateResourceSkill = async (input: UpdateSkill): Promise<Status> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateResourceSkillDocument, variables: { input } })
		.then((res) => {
			if (res) {
				return res.data.updateResourceSkill;
			}
		})
		.catch((err) => {
			return err;
		});
};

/**
 * Removes a skill from the resource's stack
 * @param resourceID - The resource ID to remove the skill from
 * @param skillID - The skill ID to remove
 * @returns An update status
 */
export const deleteResourceSkill = async (resourceID: string, skillID: string): Promise<Status> => {
	const client = getApolloClient();

	return client
		.mutate({
			mutation: DeleteResourceSkillDocument,
			variables: { resourceID, skillID }
		})
		.then((res) => {
			if (res) {
				return res.data.deleteResourceSkill;
			}
		})
		.catch((err) => {
			return err;
		});
};
