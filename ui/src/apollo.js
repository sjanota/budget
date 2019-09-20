import ApolloClient from 'apollo-client';
import { InMemoryCache } from 'apollo-cache-inmemory';
import { createHttpLink } from 'apollo-link-http';
import { ApolloLink, split } from 'apollo-link';
import { getMainDefinition } from 'apollo-utilities';
import { onError } from 'apollo-link-error';
import { WebSocketLink } from 'apollo-link-ws';

export function isSubscriptionOperation({ query }) {
    const definition = getMainDefinition(query);
    return (
        definition.kind === 'OperationDefinition' &&
        definition.operation === 'subscription'
    );
}

export default function createClient() {
    const graphqlApiUrl = "localhost:8080/query";
    const httpLink = createHttpLink({ uri: `http://${graphqlApiUrl}` });
    const wsLink = new WebSocketLink({
        uri: `ws://${graphqlApiUrl}`,
        options: {
            reconnect: true,
        },
    });
    const cache = new InMemoryCache();

    const errorLink = onError(
        ({ graphQLErrors, networkError }) => {
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
        }
    );

    const link = split(
        isSubscriptionOperation,
        wsLink,
        httpLink
    );

    return new ApolloClient({
        uri: graphqlApiUrl,
        cache,
        link: ApolloLink.from([errorLink, link]),
        connectToDevTools: true,
    });
}