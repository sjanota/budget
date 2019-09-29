import React, { useState, useRef, useEffect } from 'react';
import { shape, string, arrayOf, any, func, number } from 'prop-types';
import { replaceOnList } from '../../util/immutable';
import { Modal, Form, Button, Row, Col } from 'react-bootstrap';
import { CreateButton } from '../common/CreateButton';
import * as MoneyAmounts from '../../model/MoneyAmount';
import { MoneyAmount, Expense } from '../../model/propTypes';
import { useQuery } from '@apollo/react-hooks';
import { useBudget } from '../context/budget/budget';
import { QUERY_CATEGORIES } from '../CategoriesList/CategoriesList.gql';
import { StateEntry } from './EditModalContent.propsTypes';

export function EditModalContent({ init, onCancel, onSubmit, autoFocusRef }) {
  const [state, setState] = useState(init);
  const titleRef = useRef();

  useEffect(() => {
    if (titleRef.current) titleRef.focus();
  }, [titleRef]);

  function setValue(value) {
    return setState(e => ({ ...e, ...value }));
  }

  const inputProps = {
    state,
    setValue,
  };

  return (
    <>
      <Modal.Header>Nowy wydatek</Modal.Header>
      <Modal.Body>
        <Form>
          <Row>
            <TitleInput {...inputProps} autoFocusRef={autoFocusRef} />
          </Row>
          <Row>
            <DateInput {...inputProps} />
            <AccountInput {...inputProps} />
          </Row>
          <hr />
          <CategoriesList state={state} setState={setState} />
          <hr />
          <Row>
            <SumOutput state={state} />
          </Row>
        </Form>
      </Modal.Body>
      <Modal.Footer>
        <CancelButton onCancel={onCancel} />
        <SubmitButton state={state} onSubmit={onSubmit} onCancel={onCancel} />
      </Modal.Footer>
    </>
  );
}
EditModalContent.propTypes = {
  init: Expense,
  onCancel: func.isRequired,
  onSubmit: func.isRequired,
  autoFocusRef: any,
};

function TitleInput({ state, autoFocusRef, setValue }) {
  return (
    <Col>
      <Form.Label>Tytuł</Form.Label>
      <Form.Control
        ref={autoFocusRef}
        type="text"
        placeholder="Tytuł"
        value={state.title}
        onChange={e => setValue({ title: e.target.value })}
      />
    </Col>
  );
}
TitleInput.propTypes = {
  state: shape({ title: string.isRequired }),
  setValue: func.isRequired,
  autoFocusRef: EditModalContent.autoFocusRef,
};

function DateInput({ state, setValue }) {
  return (
    <Col>
      <Form.Label>Data</Form.Label>
      <Form.Control
        type="date"
        placeholder="Data"
        value={state.date}
        onChange={e => setValue({ date: e.target.value })}
      />
    </Col>
  );
}
DateInput.propTypes = {
  state: shape({ date: string.isRequired }),
  setValue: func.isRequired,
};

function AccountInput() {
  return (
    <Col>
      <Form.Label>Konto</Form.Label>
      <Form.Control
        type="text"
        placeholder="Under construction"
        readOnly={true}
      />
    </Col>
  );
}
AccountInput.propTypes = {};

function CreateCategoryButton({ setState }) {
  return (
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
    />
  );
}
CreateCategoryButton.propTypes = {
  setState: func.isRequired,
};

function CategoriesList({ state, setState }) {
  function setEntry(idx, update) {
    return setState(s => {
      const entries = replaceOnList(s.entries, idx, {
        ...s.entries[idx],
        ...update,
      });
      const totalBalance = entries.reduce(
        (acc, v) => MoneyAmounts.add(acc, v.balance),
        MoneyAmounts.zero()
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
      <Form.Label>Kategorie</Form.Label>
      <CreateCategoryButton setState={setState} />
      {state.entries.map((entry, idx) => (
        <Row key={idx}>
          <CategoryRow entry={entry} idx={idx} setEntry={setEntry} />
        </Row>
      ))}
    </>
  );
}
CategoriesList.propTypes = {
  setState: func.isRequired,
  state: shape({
    entries: arrayOf(StateEntry),
  }),
};

function SumOutput({ state }) {
  return (
    <>
      <Col>
        <Form.Label>Suma</Form.Label>
      </Col>
      <Col>
        <Form.Control
          type="number"
          placeholder="Suma"
          value={MoneyAmounts.format(state.totalBalance)}
          readOnly={true}
        />
      </Col>
    </>
  );
}
SumOutput.propTypes = {
  state: shape({ totalBalance: MoneyAmount }),
};

function CancelButton({ onCancel }) {
  return (
    <Button variant="secondary" onClick={onCancel}>
      Anuluj
    </Button>
  );
}
CancelButton.propTypes = {
  onCancel: EditModalContent.onCancel,
};

function SubmitButton({ state, onSubmit, onCancel }) {
  return (
    <Button
      variant="primary"
      onClick={() => {
        onSubmit(state);
        onCancel();
      }}
    >
      Zapisz
    </Button>
  );
}
SubmitButton.propTypes = {
  state: any.isRequired,
  onSubmit: EditModalContent.onSubmit,
  onCancel: EditModalContent.onCancel,
};

function CategoryRow(props) {
  return (
    <>
      <Col>
        <CategoryInput {...props} />
      </Col>
      <Col>
        <CategoryBalanceInput {...props} />
      </Col>
    </>
  );
}
CategoryRow.propTypes = {
  entry: StateEntry,
  idx: number.isRequired,
  setEntry: func.isRequired,
};

function CategoryInput({ entry, idx, setEntry }) {
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
    <Form.Control
      as="select"
      value={entry.categoryID || (entry.category && entry.category.id)}
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
  );
}
CategoryInput.propTypes = {
  entry: StateEntry,
  idx: number.isRequired,
  setEntry: func.isRequired,
};

function CategoryBalanceInput({ entry, idx, setEntry }) {
  return (
    <Form.Control
      type="number"
      placeholder="Kwota"
      value={MoneyAmounts.format(entry.balance)}
      onChange={e => setEntry(idx, { balance: e.target.value })}
      onBlur={() =>
        setEntry(idx, { balance: MoneyAmounts.parse(entry.balance) })
      }
    />
  );
}
CategoryBalanceInput.propTypes = {
  entry: StateEntry,
  idx: number.isRequired,
  setEntry: func.isRequired,
};
