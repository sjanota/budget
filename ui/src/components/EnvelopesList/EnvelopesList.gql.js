import gql from 'graphql-tag';

export const QUERY_ENVELOPES = gql`
  query QueryEnvelopes($budgetID: ID!) {
    envelopes(budgetID: $budgetID) {
      id
      name
      balance {
        integer
        decimal
      }
    }
  }
`;

export const UPDATE_ENVELOPE = gql`
  mutation UpdateEnvelope($budgetID: ID!, $id: ID!, $input: EnvelopeInput!) {
    updateEnvelope(budgetID: $budgetID, id: $id, input: $input) {
      id
    }
  }
`;

export const CREATE_ENVELOPE = gql`
  mutation CreateEnvelope($budgetID: ID!, $input: EnvelopeInput!) {
    createEnvelope(budgetID: $budgetID, input: $input) {
      id
    }
  }
`;
