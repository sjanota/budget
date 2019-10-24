import gql from 'graphql-tag';
import { useMutation, useQuery } from '@apollo/react-hooks';
import { useBudget } from './budget';
import { removeFromListByID } from '../../util/immutable';

export const GET_CURRENT_EXPENSES = gql`
  query getCurrentExpenses($budgetID: ID!) {
    budget(budgetID: $budgetID) {
      currentMonth {
        expenses {
          id
          title
          account {
            id
            name
          }
          categories {
            category {
              id
              name
            }
            amount
          }
          totalAmount
          date
        }
      }
    }
  }
`;

const CREATE_EXPENSE = gql`
  mutation createExpense($budgetID: ID!, $input: ExpenseInput!) {
    createExpense(budgetID: $budgetID, in: $input) {
      id
      title
      account {
        id
        name
      }
      categories {
        category {
          id
          name
        }
        amount
      }
      totalAmount
      date
    }
  }
`;

const UPDATE_EXPENSE = gql`
  mutation updateExpense($budgetID: ID!, $id: ID!, $input: ExpenseUpdate!) {
    updateExpense(budgetID: $budgetID, id: $id, in: $input) {
      id
      title
      account {
        id
        name
      }
      categories {
        category {
          id
          name
        }
        amount
      }
      totalAmount
      date
    }
  }
`;

export function useCreateExpense() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(CREATE_EXPENSE, {
    update: (cache, { data: { createExpense } }) => {
      const { budget } = cache.readQuery({
        query: GET_CURRENT_EXPENSES,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_CURRENT_EXPENSES,
        variables: { budgetID: selectedBudget.id },
        data: {
          budget: {
            ...budget,
            currentMonth: {
              ...budget.currentMonth,
              expenses: budget.currentMonth.expenses.concat([createExpense]),
            },
          },
        },
      });
    },
  });
  const wrapper = input => {
    mutation({ variables: { budgetID: selectedBudget.id, input } });
  };
  return [wrapper, ...rest];
}

export function useUpdateExpense() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(UPDATE_EXPENSE);
  const wrapper = (id, input) => {
    mutation({ variables: { budgetID: selectedBudget.id, id, input } });
  };
  return [wrapper, ...rest];
}

export function useGetCurrentExpenses() {
  const { selectedBudget } = useBudget();
  return useQuery(GET_CURRENT_EXPENSES, {
    variables: { budgetID: selectedBudget.id },
  });
}

const DELETE_EXPENSE = gql`
  mutation deleteExpense($budgetID: ID!, $id: ID!) {
    deleteExpense(budgetID: $budgetID, id: $id) {
      id
    }
  }
`;

export function useDeleteExpense() {
  const { selectedBudget } = useBudget();
  const [mutation, ...rest] = useMutation(DELETE_EXPENSE, {
    update: (cache, { data: { deleteExpense } }) => {
      const { budget } = cache.readQuery({
        query: GET_CURRENT_EXPENSES,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_CURRENT_EXPENSES,
        variables: { budgetID: selectedBudget.id },
        data: {
          budget: {
            ...budget,
            currentMonth: {
              ...budget.currentMonth,
              expenses: removeFromListByID(
                budget.currentMonth.expenses,
                deleteExpense.id
              ),
            },
          },
        },
      });
    },
  });
  const wrapper = id => {
    mutation({ variables: { budgetID: selectedBudget.id, id } });
  };
  return [wrapper, ...rest];
}
