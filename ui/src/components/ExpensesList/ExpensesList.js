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
import { useBudget } from '../context/budget/budget';
import { ListEntry } from './ListEntry';
import { ListHeader } from './ListHeader';
import { EditModalContent } from './EditModalContent';
import { List } from '../common/List/List';

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
        editMode={List.EditMode.MODAL}
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
        emptyValue={{
          totalBalance: { integer: 0, decimal: 0 },
          title: '',
          location: '',
          date: '',
          entries: [],
        }}
        renderModalContent={props => <EditModalContent {...props} />}
      />
    </div>
  );
}
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

function prepareInput({
  title,
  date,
  totalBalance,
  location,
  account,
  entries,
}) {
  return {
    title,
    date,
    totalBalance: {
      integer: totalBalance.integer,
      decimal: totalBalance.decimal,
    },
    location,
    accountID: account ? account.ID : null,
    entries: entries.map(entry => ({
      title: '',
      categoryID: entry.categoryID || entry.category.id,
      balance: {
        integer: entry.balance.integer,
        decimal: entry.balance.decimal,
      },
    })),
  };
}
