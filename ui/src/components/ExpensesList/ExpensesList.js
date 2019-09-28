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
  replaceOnList,
} from '../../util/immutable';
import './ExpensesList.css';
import { useBudget } from '../context/budget/budget';
import List from '../common/List/List';
import { EditEntry } from './EditEntry';
import { ListEntry } from './ListEntry';
import { ListHeader } from './ListHeader';
import { Modal, Form, Button, Row, Col } from 'react-bootstrap';
import * as MoneyAmount from '../../model/MoneyAmount';
import { QUERY_CATEGORIES } from '../CategoriesList/CategoriesList.gql';
import { CreateButton } from '../common/CreateButton';

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
          entries: [],
        }}
        editMode={List.EditMode.MODAL}
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

function EditModalContentEntry({ entry, idx, setEntry }) {
  const { id: budgetID } = useBudget();
  const { loading, error, data } = useQuery(QUERY_CATEGORIES, {
    variables: { budgetID },
  });

  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }

  return (
    <Row>
      <Col>
        <Form.Control
          as="select"
          value={entry.categoryID || (entry.category && entry.category.id)}
          placeholder="Tytuł"
          onChange={e => setEntry(idx, { categoryID: e.target.value })}
        >
          <option></option>
          {data.categories
            .sort((c1, c2) => c1.name.localeCompare(c2))
            .map(category => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
        </Form.Control>
      </Col>
      <Col>
        <Form.Control
          type="number"
          placeholder="Kwota"
          value={MoneyAmount.format(entry.balance)}
          onChange={e => setEntry(idx, { balance: e.target.value })}
          onBlur={() =>
            setEntry(idx, { balance: MoneyAmount.parse(entry.balance) })
          }
        />
      </Col>
    </Row>
  );
}

function EditModalContent({ init, onCancel, onSubmit }) {
  const [state, setState] = React.useState(init);

  console.log(state);

  function setValue(value) {
    return setState(e => ({ ...e, ...value }));
  }

  function setEntry(idx, update) {
    return setState(s => {
      const entries = replaceOnList(s.entries, idx, {
        ...s.entries[idx],
        ...update,
      });
      const totalBalance = entries.reduce(
        (acc, v) => MoneyAmount.add(acc, v.balance),
        MoneyAmount.zero()
      );

      return {
        ...s,
        entries,
        totalBalance,
      };
    });
  }

  return (
    <>
      <Modal.Header>Nowy wydatek</Modal.Header>
      <Modal.Body>
        <Form>
          <Form.Label>Tytuł</Form.Label>
          <Form.Control
            type="text"
            placeholder="Tytuł"
            value={state.title}
            onChange={e => setValue({ title: e.target.value })}
          />
          <Row>
            <Col>
              <Form.Label>Data</Form.Label>
              <Form.Control
                type="date"
                value={state.date}
                onChange={e => setValue({ date: e.target.value })}
              />
            </Col>
            <Col>
              <Form.Label>Suma</Form.Label>
              <Form.Control
                type="number"
                placeholder="Suma"
                value={MoneyAmount.format(state.totalBalance)}
                readOnly={true}
              />
            </Col>
          </Row>
          <Form.Group>
            <Form.Label>Wpisy</Form.Label>
            <CreateButton
              onClick={() =>
                setState(s => ({
                  ...s,
                  entries: [
                    ...s.entries,
                    { title: '', balance: '0.0', categoryID: '' },
                  ],
                }))
              }
            >
              Dodaj wpis
            </CreateButton>
            {state.entries.map((entry, idx) => (
              <EditModalContentEntry
                key={idx}
                entry={entry}
                idx={idx}
                setEntry={setEntry}
              />
            ))}
          </Form.Group>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <Button variant="secondary" onClick={onCancel}>
          Anuluj
        </Button>
        <Button
          variant="primary"
          onClick={() => {
            onSubmit(state);
            onCancel();
          }}
        >
          Zapisz
        </Button>
      </Modal.Footer>
    </>
  );
}
