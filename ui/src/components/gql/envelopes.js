import gql from 'graphql-tag';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { useBudget } from './budget';

export const GET_ENVELOPES = gql`
  query GetEnvelopes($budgetID: ID!) {
    envelopes(budgetID: $budgetID) {
      id
      name
      balance
      limit
    }
  }
`;

const CREATE_ENVELOPE = gql`
  mutation CreateEnvelope($budgetID: ID!, $input: EnvelopeInput!) {
    createEnvelope(budgetID: $budgetID, in: $input) {
      id
      name
      balance
      limit
    }
  }
`;

const UPDATE_ENVELOPE = gql`
  mutation UpdateEnvelope($budgetID: ID!, $id: ID!, $input: EnvelopeUpdate!) {
    updateEnvelope(budgetID: $budgetID, id: $id, in: $input) {
      id
      name
      balance
      limit
    }
  }
`;

export function useCreateEnvelope() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(CREATE_ENVELOPE, {
    update: (cache, { data: { createEnvelope } }) => {
      const { envelopes } = cache.readQuery({
        query: GET_ENVELOPES,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_ENVELOPES,
        variables: { budgetID: selectedBudget.id },
        data: {
          envelopes: envelopes.concat([createEnvelope]),
        },
      });
    },
  });
  const wrapper = input => {
    mutation({ variables: { budgetID: selectedBudget.id, input } });
  };
  return [wrapper, ...rest];
}

export function useUpdateEnvelope() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(UPDATE_ENVELOPE);
  const wrapper = (id, input) => {
    mutation({ variables: { budgetID: selectedBudget.id, id, input } });
  };
  return [wrapper, ...rest];
}

export function useGetEnvelopes() {
  const { selectedBudget } = useBudget();
  return useQuery(GET_ENVELOPES, {
    variables: { budgetID: selectedBudget.id },
  });
}
