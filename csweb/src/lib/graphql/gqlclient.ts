import { ApolloClient, InMemoryCache, split, HttpLink  } from "@apollo/client/core";
import { getMainDefinition } from '@apollo/client/utilities';
import { appConfig } from "$lib/appConfig"
import { currentUser } from '$lib/stores/auth';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';

let client: ApolloClient<any> | undefined = undefined;

const wsLink = new GraphQLWsLink(createClient({
  url: appConfig.graphqlWSEndpoint,
  connectionParams: {
    authToken: currentUser().token 
  },
}));

export const getRequestHeaders = () => {
  return {
      "Authorization": currentUser().token ? "Bearer " + currentUser().token : ""
  }
}

const httpLink = new HttpLink({
  uri: appConfig.graphqlEndpoint,
  headers: getRequestHeaders(),
});

// The split function takes three parameters:
//
// * A function that's called for each operation to execute
// * The Link to use for an operation if the function returns a "truthy" value
// * The Link to use for an operation if the function returns a "falsy" value
const splitLink = split(
  ({ query }) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  },
  wsLink,
  httpLink,
);


export const getApolloClient = () => {
  if (client) {
    return client
  }
  
  client = new ApolloClient({
    link: splitLink,
		cache: new InMemoryCache(),
    headers: getRequestHeaders()
  });

  return client
}