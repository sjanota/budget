import React, { useEffect, useState } from 'react';
import { useMutation, useQuery } from '@apollo/react-hooks';
import {
  CREATE_EXPENSE,
  DELETE_EXPENSE,
  EXPENSES_EVENTS_SUBSCRIPTION,
  EXPENSES_QUERY,
  UPDATE_EXPENSE,
} from './ExpensesList.gql';
import Table from 'react-bootstrap/Table';
import {
  addToList,
  removeFromList,
  removeFromListByID,
  replaceOnListByID,
} from '../../util/immutable';
import { cloneDeep } from 'apollo-utilities';
import './ExpensesList.css';
import { DeleteButton } from '../common/DeleteButton';
import { EditButton } from '../common/EditButton';
import { CreateButton } from '../common/CreateButton';
import { CancelButton } from '../common/CancelButton';
import { SubmitButton } from '../common/SubmitButton';
import PropTypes from 'prop-types';
import { Expense } from './ExpensesList.types';
import { useBudget } from '../context/budget/budget';

function handleExpenseEvent(prev, { subscriptionData }) {
  const event = subscriptionData.data.expenseEvent;
  switch (event.type) {
    case 'CREATED': {
      return { expenses: addToList(prev.expenses, event.expense) };
    }
    case 'DELETED': {
      return { expenses: removeFromListByID(prev.expenses, event.expense.id) };
    }
    case 'UPDATED': {
      return { expenses: replaceOnListByID(prev.expenses, event.expense) };
    }
    default:
      return prev;
  }
}

function DeleteExpenseButton({ expense }) {
  const { id: budgetID } = useBudget();
  const [deleteExpense] = useMutation(DELETE_EXPENSE, {
    variables: { budgetID },
  });
  return (
    <DeleteButton
      onClick={() => deleteExpense({ variables: { id: expense.id } })}
    />
  );
}

DeleteExpenseButton.propTypes = {
  expense: Expense,
};

function ListHeader({ onCreate }) {
  return (
    <thead className={'thead-dark'}>
      <tr>
        <th>Tytuł</th>
        <th>Data</th>
        <th>Suma</th>
        <th>Miejsce</th>
        <th>Konto</th>
        <th>
          Actions
          <CreateButton onClick={onCreate} />
        </th>
      </tr>
    </thead>
  );
}

ListHeader.propTypes = {
  onCreate: PropTypes.func,
};

function ListEntry({ expense, onEdit }) {
  return (
    <tr>
      <td>{expense.title}</td>
      <td>{expense.date}</td>
      <td>
        {expense.totalBalance.integer}.{expense.totalBalance.decimal}
      </td>
      <td>{expense.location}</td>
      <td>{expense.account && expense.account.name}</td>
      <td>
        <DeleteExpenseButton expense={expense} />
        <EditButton onClick={onEdit} />
      </td>
    </tr>
  );
}

ListEntry.propTypes = {
  expense: Expense,
  onEdit: PropTypes.func,
};

function EditEntry({ init, onCancel, onSubmit }) {
  const [title, setTitle] = useState(init ? init.title : '');
  const [date, setDate] = useState(init ? init.date : '');
  const [totalBalance, setTotalBalance] = useState(
    init ? `${init.totalBalance.integer}.${init.totalBalance.decimal}` : ''
  );
  const [location, setLocation] = useState(init ? init.location : '');

  function onChangeCallback(callback, modify = x => x) {
    return event => callback(modify(event.target.value));
  }

  function validateInput() {
    const [integer, decimal] = Number(totalBalance)
      .toFixed(2)
      .split('.');
    return {
      title,
      date,
      totalBalance: { integer, decimal },
      location,
      entries: [],
    };
  }

  return (
    <tr>
      <td>
        <input
          value={title}
          onChange={onChangeCallback(setTitle)}
          type={'text'}
        />
      </td>
      <td>
        <input
          value={date}
          onChange={onChangeCallback(setDate)}
          type={'date'}
        />
      </td>
      <td>
        <input
          value={totalBalance}
          onChange={onChangeCallback(setTotalBalance)}
          type={'number'}
        />
      </td>
      <td>
        <input
          value={location}
          onChange={onChangeCallback(setLocation)}
          type={'text'}
        />
      </td>
      <td />
      <td>
        <CancelButton onClick={onCancel} />
        <SubmitButton
          onClick={() => {
            onSubmit(validateInput());
            onCancel();
          }}
        />
      </td>
    </tr>
  );
}

EditEntry.propTypes = {
  init: Expense,
  onCancel: PropTypes.func,
  onSubmit: PropTypes.func,
};

function CreateExpenseEntry({ onCancel }) {
  const { id: budgetID } = useBudget();
  const [createExpense] = useMutation(CREATE_EXPENSE);
  return (
    <EditEntry
      onCancel={onCancel}
      onSubmit={input => createExpense({ variables: { input, budgetID } })}
    />
  );
}

CreateExpenseEntry.propTypes = {
  onCancel: PropTypes.func,
};

function UpdateExpenseEntry({ expense, onCancel }) {
  const { id: budgetID } = useBudget();
  const [updateExpense] = useMutation(UPDATE_EXPENSE);
  return (
    <EditEntry
      init={expense}
      onCancel={onCancel}
      onSubmit={input =>
        updateExpense({ variables: { id: expense.id, input, budgetID } })
      }
    />
  );
}

UpdateExpenseEntry.propTypes = {
  expense: Expense,
  onCancel: PropTypes.func,
};

export default function ExpensesList() {
  const { id: budgetID } = useBudget();
  const { loading, error, data, subscribeToMore } = useQuery(EXPENSES_QUERY, {
    variables: { budgetID },
  });
  const [isCreating, setIsCreating] = useState(false);
  const [editing, setEditing] = useState([]);

  useEffect(() => {
    if (loading) return;
    return subscribeToMore({
      document: EXPENSES_EVENTS_SUBSCRIPTION,
      variables: { budgetID },
      updateQuery: handleExpenseEvent,
      onError: console.error,
    });
  }, [loading, subscribeToMore, budgetID]);

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return (
    <div className={'ExpensesList'}>
      <Table striped bordered hover size={'sm'}>
        <ListHeader onCreate={() => setIsCreating(true)} />
        <tbody>
          {isCreating && (
            <CreateExpenseEntry onCancel={() => setIsCreating(false)} />
          )}
          {data.expenses.map(expense =>
            editing.some(id => id === expense.id) ? (
              <UpdateExpenseEntry
                key={expense.id}
                expense={cloneDeep(expense)}
                onCancel={() =>
                  setEditing(editing => removeFromList(editing, expense.id))
                }
              />
            ) : (
              <ListEntry
                key={expense.id}
                expense={expense}
                onEdit={() =>
                  setEditing(editing => addToList(editing, expense.id))
                }
              />
            )
          )}
        </tbody>
      </Table>
    </div>
  );
}
