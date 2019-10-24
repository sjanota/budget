import gql from 'graphql-tag';
import { useQuery, useMutation } from '@apollo/react-hooks';
import { useBudget } from './budget';

export const GET_ACCOUNTS = gql`
  query GetAccounts($budgetID: ID!) {
    accounts(budgetID: $budgetID) {
      id
      name
      balance
    }
  }
`;

export function useGetAccounts() {
  const { selectedBudget } = useBudget();
  return useQuery(GET_ACCOUNTS, {
    variables: { budgetID: selectedBudget.id },
  });
}

const CREATE_ACCOUNT = gql`
  mutation CreateAccount($budgetID: ID!, $input: AccountInput!) {
    createAccount(budgetID: $budgetID, in: $input) {
      id
      name
      balance
    }
  }
`;

export function useCreateAccount() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(CREATE_ACCOUNT, {
    update: (cache, { data: { createAccount } }) => {
      const { accounts } = cache.readQuery({
        query: GET_ACCOUNTS,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_ACCOUNTS,
        variables: { budgetID: selectedBudget.id },
        data: {
          accounts: accounts.concat([createAccount]),
        },
      });
    },
  });
  const wrapper = input => {
    mutation({ variables: { budgetID: selectedBudget.id, input } });
  };
  return [wrapper, ...rest];
}

const UPDATE_ACCOUNT = gql`
  mutation UpdateAccount($budgetID: ID!, $id: ID!, $input: AccountUpdate!) {
    updateAccount(budgetID: $budgetID, id: $id, in: $input) {
      id
      name
      balance
    }
  }
`;

export function useUpdateAccount() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(UPDATE_ACCOUNT);
  const wrapper = (id, input) => {
    mutation({ variables: { budgetID: selectedBudget.id, id, input } });
  };
  return [wrapper, ...rest];
}
