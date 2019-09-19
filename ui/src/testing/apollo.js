import {getMainDefinition} from "apollo-utilities";
import {MockLink, MockSubscriptionLink} from "@apollo/react-testing";
import {ApolloLink} from "apollo-link";

export function createMockLink(mocks, addTypename = true) {
  const isSubscription = ({query}) => {
    const definition = getMainDefinition(query);
    return (
      definition.kind === 'OperationDefinition' &&
      definition.operation === 'subscription'
    );
  };
  const subscriptionLink = new MockSubscriptionLink();
  const link = ApolloLink.split(
    isSubscription,
    subscriptionLink,
    new MockLink(mocks, addTypename),
  );
  return {link, sendEvent: subscriptionLink.simulateResult.bind(subscriptionLink)}
}