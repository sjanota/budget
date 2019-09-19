import React, {useEffect} from 'react';
import {useMutation, useQuery} from "@apollo/react-hooks";
import {DELETE_EXPENSE, EXPENSES_EVENTS_SUBSCRIPTION, EXPENSES_QUERY} from "./ExpensesList.gql";
import Table from "react-bootstrap/Table";
import Octicon, {Trashcan} from "@primer/octicons-react";
import {Button} from "react-bootstrap";

function handleExpenseEvent(prev, {subscriptionData}) {
  const event = subscriptionData.data.expenseEvents;
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
    <th>Actions</th>
  </tr>
  </thead>
}

function DeleteButton({expense}) {
  const [deleteExpense] = useMutation(DELETE_EXPENSE);
  return <Button
    size={"sm"}
    variant={"danger"}
    data-action={"delete"}
    onClick={() => deleteExpense({variables: {id: expense.id}})}
  >
    <Octicon icon={Trashcan} size={"small"} ariaLabel={"Delete expense"}/>
  </Button>
}

function ListEntry({expense}) {
  return <tr>
    <td>{expense.title}</td>
    <td>{expense.date}</td>
    <td>{expense.total}</td>
    <td>{expense.location}</td>
    <td>{expense.account && expense.account.name}</td>
    <td><DeleteButton expense={expense}/></td>
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