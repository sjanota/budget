import gql from 'graphql-tag';

export const QUERY_CATEGORIES = gql`
  query QueryCategories($budgetID: ID!) {
    categories(budgetID: $budgetID) {
      id
      name
      envelope {
        id
        name
      }
    }
  }
`;

export const UPDATE_CATEGORY = gql`
  mutation UpdateCategory($budgetID: ID!, $id: ID!, $input: CategoryInput!) {
    updateCategory(budgetID: $budgetID, id: $id, input: $input) {
      id
    }
  }
`;

export const CREATE_CATEGORY = gql`
  mutation CreateCategory($budgetID: ID!, $input: CategoryInput!) {
    createCategory(budgetID: $budgetID, input: $input) {
      id
    }
  }
`;
