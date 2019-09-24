import gql from 'graphql-tag';

export const QUERY_BUDGETS = gql`
  query QueryBudgets {
    budgets {
      id
      name
    }
  }
`;

export const CREATE_BUDGET = gql`
  mutation CreateBudget($name: String!) {
    createBudget(name: $name) {
      id
      name
    }
  }
`;
