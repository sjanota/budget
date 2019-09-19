import {EXPENSES_QUERY} from "./ExpensesList.gql";

export function mockQueryExpenses(expenses) {
  return {
    request: {
      query: EXPENSES_QUERY,
    },
    result: {
      data: {
        expenses
      }
    },
  }
}

export function mockExpensesEvent(event) {
  return {
    result: {
      data: {
        expenses: event
      }
    }
  }
}

export const expense1 = {
  id: "5d8265e4d7d8a40795fe1b31",
  title: "Zakupy spożywcze",
  total: 12.32,
  location: "Lidl",
  date: null,
  account: null,
  entries: [],
  __typename: "Expense"
};

export const expense2 = {
  id: "5d826618d7d8a40795fe1b33",
  title: "Zakupy spożywcze",
  total: 12.32,
  location: "Lidl",
  date: null,
  account: null,
  entries: [],
  __typename: "Expense"
};