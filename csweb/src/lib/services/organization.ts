import { GetOrganizationDocument, UpdateOrganizationDocument } from '$lib/graphql/generated/sdk';
import type { UpdateOrganization, CreateOrganizationResult, Organization } from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';

/**
 * return the default organization with settings
 * @returns a promise with an organization object
 */
export const getOrganization = async (): Promise<Organization> => {
    const client = getApolloClient();

    return client
        .query({ query:  GetOrganizationDocument })
        .then((res) => {
            if (res) {
                return res.data.getOrganization;
            }
        })
        .catch((err) => {
            return err;
        });
};

/**
 * Updates the properties of an existing organization
 * @param input An UpdateOrganization object with complete properties
 * @returns An update status to confirm creation success
 */
export const updateOrganization = async (input: UpdateOrganization): Promise<CreateOrganizationResult> => {
    const client = getApolloClient();

    console.log(input)

    return client
        .mutate({ mutation: UpdateOrganizationDocument, variables: { input } })
        .then((res) => {
            if (res) {
                return res.data.updateOrganization;
            }
        })
        .catch((err) => {
            return err;
        });
};


