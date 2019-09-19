import React, {useEffect, useState} from 'react';
import {useMutation, useQuery} from "@apollo/react-hooks";
import {
  CREATE_EXPENSE,
  DELETE_EXPENSE,
  EXPENSES_EVENTS_SUBSCRIPTION,
  EXPENSES_QUERY,
  UPDATE_EXPENSE
} from "./ExpensesList.gql";
import Table from "react-bootstrap/Table";
import Octicon, {Check, Pencil, Plus, Trashcan, X,} from "@primer/octicons-react";
import {Button} from "react-bootstrap";
import {addToList, removeFromList, removeFromListByID, replaceOnListByID} from "../../util/immutable";
import {cloneDeep} from "apollo-utilities";

function handleExpenseEvent(prev, {subscriptionData}) {
  const event = subscriptionData.data.expenseEvents;
  switch (event.type) {
    case "CREATED": {
      return {expenses: addToList(prev.expenses, event.expense)};
    }
    case "DELETED": {
      return {expenses: removeFromListByID(prev.expenses, event.expense.id)}
    }
    case "UPDATED": {
      return {expenses: replaceOnListByID(prev.expenses, event.expense)}
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

function DeleteButton({onClick}) {
  return <ListButton
    icon={Trashcan}
    action={"delete"}
    onClick={onClick}
    ariaLabel={"Delete expense"}
  />
}

function DeleteExpenseButton({expense}) {
  const [deleteExpense] = useMutation(DELETE_EXPENSE);
  return <DeleteButton
    onClick={() => deleteExpense({variables: {id: expense.id}})}
  />
}

function EditButton({onClick}) {
  return <ListButton
    icon={Pencil}
    action={"edit"}
    onClick={onClick}
    ariaLabel={"Edit expense"}
  />
}

function CreateButton({onClick}) {
  return <ListButton
    icon={Plus}
    action={"create"}
    onClick={onClick}
    ariaLabel={"Create new expense"}
  />
}

function CancelButton({onClick}) {
  return <ListButton
    icon={X}
    action={"cancel"}
    onClick={onClick}
    ariaLabel={"Cancel expense creation"}
  />
}

function SubmitButton({onClick}) {
  return <ListButton
    icon={Check}
    action={"submit"}
    onClick={onClick}
    ariaLabel={"Cancel expense creation"}
  />
}

function ListHeader({onCreate}) {
  return <thead>

  <tr>
    <th>Tytuł</th>
    <th>Data</th>
    <th>Suma</th>
    <th>Miejsce</th>
    <th>Konto</th>
    <th>
      Actions
      <CreateButton onClick={onCreate}/>
    </th>
  </tr>
  </thead>
}

function ListEntry({expense, onEdit}) {
  return <tr>
    <td>{expense.title}</td>
    <td>{expense.date}</td>
    <td>{expense.total.integer}.{expense.total.decimal}</td>
    <td>{expense.location}</td>
    <td>{expense.account && expense.account.name}</td>
    <td>
      <DeleteExpenseButton expense={expense}/>
      <EditButton onClick={onEdit}/>
    </td>
  </tr>
}

function EditEntry({init, onCancel, onSubmit}) {
  const [title, setTitle] = useState(init ? init.title : "");
  const [date, setDate] = useState(init ? init.date : "");
  const [total, setTotal] = useState(init ? `${init.total.integer}.${init.total.decimal}` : "");
  const [location, setLocation] = useState(init ? init.location : "");

  function onChangeCallback(callback, modify = x => x) {
    return event => callback(modify(event.target.value))
  }

  function validateInput() {
    const [integer, decimal] = Number(total).toFixed(2).split(".");
    return {
      title,
      date,
      total: {integer, decimal},
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
      <CancelButton onClick={onCancel}/>
      <SubmitButton onClick={() => {
        onSubmit(validateInput());
        onCancel()
      }}/>
    </td>
  </tr>
}

function CreateExpenseEntry({onCancel}) {
  const [createExpense] = useMutation(CREATE_EXPENSE);
  return <EditEntry
    onCancel={onCancel}
    onSubmit={input => createExpense({variables: {input}})}
  />
}

function UpdateExpenseEntry({expense, onCancel}) {
  const [updateExpense] = useMutation(UPDATE_EXPENSE);
  return <EditEntry
    init={expense}
    onCancel={onCancel}
    onSubmit={input => updateExpense({variables: {id: expense.id, input}})}
  />
}

export default function ExpensesList() {
  const {loading, error, data, subscribeToMore} = useQuery(EXPENSES_QUERY);
  const [isCreating, setIsCreating] = useState(false);
  const [editing, setEditing] = useState([]);

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
    <ListHeader onCreate={() => setIsCreating(true)}/>
    <tbody>
    {isCreating && <CreateExpenseEntry
      onCancel={() => setIsCreating(false)}
    />}
    {data.expenses.map(expense =>
      editing.some(id => id === expense.id)
      ? <UpdateExpenseEntry
          key={expense.id}
          expense={cloneDeep(expense)}
          onCancel={() => setEditing(editing => removeFromList(editing, expense.id))}
        />
      : <ListEntry
          key={expense.id}
          expense={expense}
          onEdit={() => setEditing(editing => addToList(editing, expense.id))}
        />
    )}
    </tbody>
  </Table>
}