import React, {useEffect, useState} from 'react';
import {useMutation, useQuery} from "@apollo/react-hooks";
import {CREATE_EXPENSE, DELETE_EXPENSE, EXPENSES_EVENTS_SUBSCRIPTION, EXPENSES_QUERY} from "./ExpensesList.gql";
import Table from "react-bootstrap/Table";
import Octicon, {Check, Plus, Trashcan, X,} from "@primer/octicons-react";
import {Button} from "react-bootstrap";
import {addToList, removeFromListById} from "../../util/immutable";

function handleExpenseEvent(prev, {subscriptionData}) {
  const event = subscriptionData.data.expenseEvents;
  switch (event.type) {
    case "CREATED": {
      return {expenses: addToList(prev.expenses, event.expense)};
    }
    case "DELETED": {
      return {expenses: removeFromListById(prev.expenses, event.expense.id)}
    }
    default:
      return prev;
  }
}

function ListButton({icon, action, onClick, ariaLabel}) {
  return <Button
    size={"sm"}
    variant={"link"}
    data-action={action}
    onClick={onClick}
  >
    <Octicon icon={icon} size={"small"} ariaLabel={ariaLabel}/>
  </Button>
}

function DeleteButton({expense}) {
  const [deleteExpense] = useMutation(DELETE_EXPENSE);
  return <ListButton
    icon={Trashcan}
    action={"delete"}
    onClick={() => deleteExpense({variables: {id: expense.id}})}
    ariaLabel={"Delete expense"}
  />
}

function StartCreationButton({onClick}) {
  return <ListButton
    icon={Plus}
    action={"start-creation"}
    onClick={onClick}
    ariaLabel={"Create new expense"}
  />
}

function CancelCreationButton({onClick}) {
  return <ListButton
    icon={X}
    action={"cancel-creation"}
    onClick={onClick}
    ariaLabel={"Cancel expense creation"}
  />
}

function CreateButton({onClick}) {
  return <ListButton
    icon={Check}
    action={"cancel-creation"}
    onClick={onClick}
    ariaLabel={"Cancel expense creation"}
  />
}

function ListHeader({onStartCreationClick}) {
  return <thead>

  <tr>
    <th>Tytuł</th>
    <th>Data</th>
    <th>Suma</th>
    <th>Miejsce</th>
    <th>Konto</th>
    <th>
      Actions
      <StartCreationButton onClick={onStartCreationClick}/>
    </th>
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
    <td><DeleteButton expense={expense}/></td>
  </tr>
}

function NewExpenseEntry({onCancelCreationClick}) {
  const [createExpense] = useMutation(CREATE_EXPENSE);
  const [title, setTitle] = useState("");
  const [date, setDate] = useState("");
  const [total, setTotal] = useState(0.0);
  const [location, setLocation] = useState("");

  function onChangeCallback(callback) {
    return event => callback(event.target.value)
  }

  function validateInput() {
    return {
      title,
      date,
      total: Number(total),
      location,
      entries: []
    }
  }

  return <tr>
    <td>
      <input value={title} onChange={onChangeCallback(setTitle)} type={"text"}/>
    </td>
    <td>
      <input value={date} onChange={onChangeCallback(setDate)} type={"date"}/>
    </td>
    <td>
      <input value={total} onChange={onChangeCallback(setTotal)} type={"number"}/>
    </td>
    <td>
      <input value={location} onChange={onChangeCallback(setLocation)} type={"text"}/>
    </td>
    <td/>
    <td>
      <CancelCreationButton onClick={onCancelCreationClick}/>
      <CreateButton onClick={async () => {
        const input = validateInput();
        console.log(input);
        await createExpense({variables: {input}});
        onCancelCreationClick()
      }}/>
    </td>
  </tr>
}

export default function ExpensesList() {
  const {loading, error, data, subscribeToMore} = useQuery(EXPENSES_QUERY);
  const [isCreating, setIsCreating] = useState(false);

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
    <ListHeader onStartCreationClick={() => setIsCreating(true)}/>
    <tbody>
    {isCreating && <NewExpenseEntry onCancelCreationClick={() => setIsCreating(false)}/>}
    {data.expenses.map(expense =>
      <ListEntry key={expense.id} expense={expense}/>
    )}
    </tbody>
  </Table>
}