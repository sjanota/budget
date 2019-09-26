import { DELETE_EXPENSE, EXPENSES_QUERY } from '../ExpensesList.gql';
import { MOCK_BUDGET_ID } from '../../context/budget/budget.testing';

export function mockQueryExpenses(expenses) {
  return {
    request: {
      query: EXPENSES_QUERY,
      variables: { budgetID: MOCK_BUDGET_ID },
    },
    result: {
      data: {
        expenses,
      },
    },
  };
}

export function mockExpensesEvent(event) {
  return {
    result: {
      data: {
        expenseEvent: { ...event, __typename: 'ExpenseEvent' },
      },
    },
  };
}

export function mockDeleteExpense(id) {
  return {
    request: {
      query: DELETE_EXPENSE,
      variables: { budgetID: MOCK_BUDGET_ID, id },
    },
    result: jest.fn().mockReturnValueOnce({
      data: {
        deleteExpense: {
          id: id,
          __typename: 'Expense',
        },
      },
    }),
  };
}

export const expense1 = {
  id: '5d8265e4d7d8a40795fe1b31',
  title: 'Zakupy spożywcze',
  totalBalance: {
    integer: 12,
    decimal: 32,
    __typename: 'MoneyAmount',
  },
  location: 'Lidl',
  date: null,
  account: null,
  entries: [],
  __typename: 'Expense',
};

export const expense2 = {
  id: '5d826618d7d8a40795fe1b33',
  title: 'Zakupy spożywcze',
  totalBalance: {
    integer: 17,
    decimal: 56,
    __typename: 'MoneyAmount',
  },
  location: 'Lidl',
  date: null,
  account: null,
  entries: [],
  __typename: 'Expense',
};
