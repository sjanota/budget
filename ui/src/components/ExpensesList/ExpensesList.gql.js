import gql from 'graphql-tag';

const EXPENSE_FRAGMENT = gql`
  fragment ExpensesDetails on Expense {
    id
    title
    date
    totalBalance {
      integer
      decimal
    }
    location
    account {
      id
      name
    }
    entries {
      balance {
        integer
        decimal
      }
      category {
        id
      }
    }
  }
`;

export const EXPENSES_QUERY = gql`
  query QueryExpenses($budgetID: ID!) {
    expenses(budgetID: $budgetID) {
      ...ExpensesDetails
    }
  }
  ${EXPENSE_FRAGMENT}
`;

export const EXPENSES_EVENTS_SUBSCRIPTION = gql`
  subscription WatchExpenses($budgetID: ID!) {
    expenseEvent(budgetID: $budgetID) {
      type
      expense {
        ...ExpensesDetails
      }
    }
  }
  ${EXPENSE_FRAGMENT}
`;

export const DELETE_EXPENSE = gql`
  mutation DeleteExpense($budgetID: ID!, $id: ID!) {
    deleteExpense(budgetID: $budgetID, id: $id) {
      id
    }
  }
`;

export const UPDATE_EXPENSE = gql`
  mutation UpdateExpense($budgetID: ID!, $id: ID!, $input: ExpenseInput!) {
    updateExpense(budgetID: $budgetID, id: $id, input: $input) {
      id
    }
  }
`;

export const CREATE_EXPENSE = gql`
  mutation CreateExpense($budgetID: ID!, $input: ExpenseInput!) {
    createExpense(budgetID: $budgetID, input: $input) {
      id
    }
  }
`;
