import React, {useEffect} from 'react';
import {useQuery} from "@apollo/react-hooks";
import {EXPENSES_EVENTS_SUBSCRIPTION, EXPENSES_QUERY} from "./ExpensesList.gql";
import Table from "react-bootstrap/Table";

function handleExpenseEvent(prev, {subscriptionData}) {
  const event = subscriptionData.data.expenses;
  switch (event.type) {
    case "ADDED": {
      const newExpense = event.expense;
      return {expenses: [...prev.expenses, newExpense]};
    }
    default:
      return prev;
  }
}

function ListHeader() {
  return <thead>
  <tr>
    <th>Tytuł</th>
    <th>Data</th>
    <th>Suma</th>
    <th>Miejsce</th>
    <th>Konto</th>
  </tr>
  </thead>
}

function ListEntry({expense}) {
  return <tr>
    <td>{expense.title}</td>
    <td>{expense.date}</td>
    <td>{expense.total}</td>
    <td>{expense.location}</td>
    <td>{expense.account && expense.account.name}</td>
  </tr>
}

export default function ExpensesList() {
  const {loading, error, data, subscribeToMore} = useQuery(EXPENSES_QUERY);

  useEffect(() => {
    if (loading) return;
    return subscribeToMore({
      document: EXPENSES_EVENTS_SUBSCRIPTION,
      updateQuery: handleExpenseEvent,
      onError: console.error
    });
  }, [loading, subscribeToMore]);

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return <Table striped bordered hover>
    <ListHeader/>
    <tbody>
    {data.expenses.map(expense =>
      <ListEntry key={expense.id} expense={expense}/>
    )}
    </tbody>
  </Table>
}