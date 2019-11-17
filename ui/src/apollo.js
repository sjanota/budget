import ApolloClient from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { createHttpLink } from 'apollo-link-http';
import { ApolloLink } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { onError } from 'apollo-link-error';
import { setContext } from 'apollo-link-context';

import { IntrospectionFragmentMatcher } from 'apollo-cache-inmemory';
import introspectionQueryResultData from './fragmentTypes.json';
import { useAuth0 } from './react-auth0-spa.js';
import React, { useEffect, useState } from 'react';
import { ApolloProvider } from '@apollo/react-hooks';

const fragmentMatcher = new IntrospectionFragmentMatcher({
  introspectionQueryResultData,
});

export function isSubscriptionOperation({ query }) {
  const definition = getMainDefinition(query);
  return (
    definition.kind === 'OperationDefinition' &&
    definition.operation === 'subscription'
  );
}

export function createClient(token) {
  const graphqlApiUrl = 'localhost:8080/query';
  const httpLink = createHttpLink({ uri: `http://${graphqlApiUrl}` });
  const authLink = setContext((_, { headers }) => {
    // get the authentication token from local storage if it exists
    // return the headers to the context so httpLink can read them
    return {
      headers: {
        ...headers,
        authorization: token ? `Bearer ${token}` : '',
      },
    };
  });
  const cache = new InMemoryCache({ fragmentMatcher });

  const errorLink = onError(({ graphQLErrors, networkError }) => {
    if (process.env.REACT_APP_ENV !== 'production') {
      if (graphQLErrors) {
        graphQLErrors.map(({ message, locations, path }) =>
          console.log(
            `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
          )
        );
      }

      if (networkError) console.log(`[Network error]: ${networkError}`);
    }
  });

  // const link = split(isSubscriptionOperation, wsLink, httpLink);
  const link = authLink.concat(httpLink);

  return new ApolloClient({
    uri: graphqlApiUrl,
    cache,
    link: ApolloLink.from([errorLink, link]),
    connectToDevTools: true,
  });
}

export function AuthApolloProvider({ children }) {
  const {
    isAuthenticated,
    loading,
    loginWithRedirect,
    getTokenSilently,
  } = useAuth0();
  const [token, setToken] = useState();

  useEffect(() => {
    if (loading) {
      return;
    }
    if (!isAuthenticated) {
      loginWithRedirect({});
      return;
    }

    getTokenSilently().then(setToken);
  }, [isAuthenticated, loginWithRedirect, loading, getTokenSilently]);

  if (loading) {
    return <div>Loading...</div>;
  }

  if (!isAuthenticated || !token) {
    return <div />;
  }

  return (
    <ApolloProvider client={createClient(token)}>{children}</ApolloProvider>
  );
}
