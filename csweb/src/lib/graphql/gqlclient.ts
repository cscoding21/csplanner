import { ApolloClient, InMemoryCache, split, HttpLink } from '@apollo/client/core';
import { onError } from '@apollo/client/link/error';
import { getMainDefinition } from '@apollo/client/utilities';
import { appConfig } from '$lib/appConfig';
import { authService } from '$lib/services/auth';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';
import { goto } from '$app/navigation';

export const getApolloClient = () => {
	const as = authService();

	const httpLink = new HttpLink({
		uri: appConfig.graphqlEndpoint,
		headers: as.getAuthHeaders()
	});

	const wsLink = new GraphQLWsLink(
		createClient({
			url: appConfig.graphqlWSEndpoint,
			connectionParams: {
				authToken: as.getAccessToken()
			}
		})
	);

	// The split function takes three parameters:
	//
	// * A function that's called for each operation to execute
	// * The Link to use for an operation if the function returns a "truthy" value
	// * The Link to use for an operation if the function returns a "falsy" value
	const splitLink = split(
		({ query }) => {
			const definition = getMainDefinition(query);
			return definition.kind === 'OperationDefinition' && definition.operation === 'subscription';
		},
		wsLink,
		httpLink
	);

	// const client = new ApolloClient({
	// 	link: logoutLink.concat(splitLink),
	// 	cache: new InMemoryCache(),
	// 	headers: as.getAuthHeaders()
	// });

	const client = new ApolloClient({
		link: logoutLink.concat(splitLink),
		cache: new InMemoryCache(),
		headers: as.getAuthHeaders()
	});

	return client;
};

const logoutLink = onError(({ networkError }) => {
	// @ts-expect-error : TS doesn't recognize statusCode and must be silenced
	if (networkError?.statusCode === 403 || networkError?.statusCode === 401) {
		goto('/login');
	}
});
