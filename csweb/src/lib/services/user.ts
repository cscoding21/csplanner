import { FindAllUsersDocument } from '$lib/graphql/generated/sdk';
import type { UserResults } from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';

/**
 * return all users in the system
 * @returns a promise of all users
 */
export const findAllUsers = async (): Promise<UserResults> => {
	const client = getApolloClient();

	return client
		.query({ query: FindAllUsersDocument  })
		.then((userResults) => {
			if (userResults) {
                console.log("users(res)", userResults.data.findAllUsers)
				return userResults.data.findAllUsers;
			}
		})
		.catch((err) => {
			return err;
		});
};
