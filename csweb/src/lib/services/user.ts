import { FindAllUsersDocument, UpdateUserDocument } from '$lib/graphql/generated/sdk';
import type { CreateUserResult, UserResults, UpdateUser } from '$lib/graphql/generated/sdk';
import { getApolloClient } from '$lib/graphql/gqlclient';
import { authService } from './auth';

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
				return userResults.data.findAllUsers;
			}
		})
		.catch((err) => {
			return err;
		});
};


/**
 * 
 * @returns a status object with the updated user results
 */
export const updateUser = async(input:UpdateUser):Promise<CreateUserResult> => {
	const client = getApolloClient();

	return client
		.mutate({ mutation: UpdateUserDocument , variables: { input }  })
		.then((userResults) => {
			if (userResults) {
				const as = authService()
				as.refresh()

				return userResults.data.updateUser
			}
		})
		.catch((err) => {
			return err;
		});

	return {}
}
