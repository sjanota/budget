import gql from 'graphql-tag';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { useBudget } from './budget';
import { GET_CURRENT_MONTHLY_REPORT } from './monthlyReport';
import { GET_ACCOUNTS } from './accounts';

const TRANSFER_FRAGMENT = gql`
  fragment Transfer on Transfer {
    id
    title
    fromAccount {
      id
      name
    }
    toAccount {
      id
      name
    }
    amount
    date
  }
`;

export const GET_CURRENT_TRANSFERS = gql`
  query getCurrentTransfers($budgetID: ID!) {
    budget(budgetID: $budgetID) {
      currentMonth {
        transfers {
          ...Transfer
        }
      }
    }
  }
  ${TRANSFER_FRAGMENT}
`;

export function useGetCurrentTransfers() {
  const { selectedBudget } = useBudget();
  return useQuery(GET_CURRENT_TRANSFERS, {
    variables: { budgetID: selectedBudget.id },
  });
}

const CREATE_TRANSFER = gql`
  mutation createTransfer($budgetID: ID!, $input: TransferInput!) {
    createTransfer(budgetID: $budgetID, in: $input) {
      ...Transfer
    }
  }
  ${TRANSFER_FRAGMENT}
`;

export function useCreateTransfer() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(CREATE_TRANSFER, {
    update: (cache, { data: { createTransfer } }) => {
      const { budget } = cache.readQuery({
        query: GET_CURRENT_TRANSFERS,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_CURRENT_TRANSFERS,
        variables: { budgetID: selectedBudget.id },
        data: {
          budget: {
            ...budget,
            currentMonth: {
              ...budget.currentMonth,
              transfers: budget.currentMonth.transfers.concat([createTransfer]),
            },
          },
        },
      });
    },
    refetchQueries: () => [
      { query: GET_ACCOUNTS, variables: { budgetID: selectedBudget.id } },
      {
        query: GET_CURRENT_MONTHLY_REPORT,
        variables: { budgetID: selectedBudget.id },
      },
    ],
  });
  const wrapper = input => {
    mutation({ variables: { budgetID: selectedBudget.id, input } });
  };
  return [wrapper, ...rest];
}

const UPDATE_TRANSFER = gql`
  mutation updateTransfer($budgetID: ID!, $id: ID!, $input: TransferUpdate!) {
    updateTransfer(budgetID: $budgetID, id: $id, in: $input) {
      ...Transfer
    }
  }
  ${TRANSFER_FRAGMENT}
`;

export function useUpdateTransfer() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(UPDATE_TRANSFER, {
    refetchQueries: () => [
      { query: GET_ACCOUNTS, variables: { budgetID: selectedBudget.id } },
      {
        query: GET_CURRENT_MONTHLY_REPORT,
        variables: { budgetID: selectedBudget.id },
      },
    ],
  });
  const wrapper = (id, input) => {
    mutation({ variables: { budgetID: selectedBudget.id, id, input } });
  };
  return [wrapper, ...rest];
}
