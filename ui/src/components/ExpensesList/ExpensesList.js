import React, { useEffect } from 'react';
import { useMutation, useQuery } from '@apollo/react-hooks';
import {
  CREATE_EXPENSE,
  DELETE_EXPENSE,
  EXPENSES_EVENTS_SUBSCRIPTION,
  EXPENSES_QUERY,
  UPDATE_EXPENSE,
} from './ExpensesList.gql';
import {
  addToList,
  removeFromListByID,
  replaceOnListByID,
} from '../../util/immutable';
import './ExpensesList.css';
import { DeleteButton } from '../common/DeleteButton';
import PropTypes from 'prop-types';
import { Expense } from './ExpensesList.types';
import { useBudget } from '../context/budget/budget';
import List from '../List/List';
import * as MoneyAmount from '../../model/MoneyAmount';

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

function prepareInput({ title, date, totalBalance, location, account }) {
  return {
    title,
    date,
    totalBalance: {
      integer: totalBalance.integer,
      decimal: totalBalance.decimal,
    },
    location,
    accountID: account ? account.ID : null,
    entries: [],
  };
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

function ListHeader() {
  return (
    <>
      <th>Tytuł</th>
      <th>Data</th>
      <th>Suma</th>
      <th>Miejsce</th>
      <th>Konto</th>
    </>
  );
}

ListHeader.propTypes = {};

function ListEntry({ entry }) {
  return (
    <>
      <td>{entry.title}</td>
      <td>{entry.date}</td>
      <td>{MoneyAmount.format(entry.totalBalance)}</td>
      <td>{entry.location}</td>
      <td>{entry.account && entry.account.name}</td>
    </>
  );
}

ListEntry.propTypes = {
  entry: Expense.isRequired,
};

function EditEntry({ entry, setEntry }) {
  function setValue(value) {
    return setEntry(e => ({ ...e, ...value }));
  }

  return (
    <>
      <td>
        <input
          value={entry.title}
          onChange={event => setValue({ title: event.target.value })}
          type={'text'}
        />
      </td>
      <td>
        <input
          value={entry.date || ''}
          onChange={event => setValue({ date: event.target.value })}
          type={'date'}
        />
      </td>
      <td>
        <input
          value={MoneyAmount.format(entry.totalBalance)}
          onChange={event =>
            setValue({ totalBalance: MoneyAmount.parse(event.target.value) })
          }
          type={'number'}
        />
      </td>
      <td>
        <input
          value={entry.location || ''}
          onChange={event => setValue({ location: event.target.value })}
          type={'text'}
        />
      </td>
      <td />
    </>
  );
}

EditEntry.propTypes = {
  entry: Expense.isRequired,
  setEntry: PropTypes.func.isRequired,
};

export default function ExpensesList() {
  const { id: budgetID } = useBudget();
  const { loading, error, data, subscribeToMore } = useQuery(EXPENSES_QUERY, {
    variables: { budgetID },
  });
  const [deleteExpense] = useMutation(DELETE_EXPENSE);
  const [updateExpense] = useMutation(UPDATE_EXPENSE);
  const [createExpense] = useMutation(CREATE_EXPENSE);

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
      <List
        entries={data.expenses}
        onCreate={input =>
          createExpense({ variables: { budgetID, input: prepareInput(input) } })
        }
        onDelete={id => deleteExpense({ variables: { budgetID, id } })}
        onUpdate={(id, input) =>
          updateExpense({
            variables: { budgetID, id, input: prepareInput(input) },
          })
        }
        renderHeader={props => <ListHeader {...props} />}
        renderEntry={props => <ListEntry {...props} />}
        renderEditEntry={props => <EditEntry {...props} />}
        emptyValue={{
          totalBalance: { integer: 0, decimal: 0 },
          title: '',
          location: '',
          date: '',
        }}
      />
    </div>
  );
}
