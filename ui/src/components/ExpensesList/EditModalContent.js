import React, { useState, useRef, useEffect } from 'react';
import PropTypes from 'prop-types';
import { replaceOnList } from '../../util/immutable';
import { Modal, Form, Button, Row, Col } from 'react-bootstrap';
import { CreateButton } from '../common/CreateButton';
import { EditModalContentEntry } from './EditModalContentEntry';
import * as MoneyAmount from '../../model/MoneyAmount';
import { MoneyAmount as MoneyAmountPropType } from '../../model/propTypes';

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
          <Row>
            <CategoriesInput state={state} setState={setState} />
          </Row>
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
  init: PropTypes.shape({
    title: PropTypes.string,
    date: PropTypes.string,
    totalBalance: MoneyAmountPropType,
    entries: PropTypes.arrayOf(EditModalContentEntry.propTypes.entry),
  }),
  onCancel: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
  autoFocusRef: PropTypes.any,
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
  state: PropTypes.shape({ title: PropTypes.string.isRequired }),
  setValue: PropTypes.func.isRequired,
  autoFocusRef: EditModalContent.propTypes.autoFocusRef,
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
  state: PropTypes.shape({ date: PropTypes.string.isRequired }),
  setValue: PropTypes.func.isRequired,
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
  setState: PropTypes.func.isRequired,
};

function CategoriesInput({ state, setState }) {
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
    <Col>
      <Form.Label>Kategorie</Form.Label>
      <CreateCategoryButton setState={setState} />
      {state.entries.map((entry, idx) => (
        <EditModalContentEntry
          key={idx}
          entry={entry}
          idx={idx}
          setEntry={setEntry}
        />
      ))}
    </Col>
  );
}
CategoriesInput.propTypes = {
  setState: PropTypes.func.isRequired,
  state: PropTypes.shape({
    entries: PropTypes.arrayOf(EditModalContentEntry.propTypes.entry),
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
          value={MoneyAmount.format(state.totalBalance)}
          readOnly={true}
        />
      </Col>
    </>
  );
}
SumOutput.propTypes = {
  state: PropTypes.shape({ totalBalance: MoneyAmountPropType }),
};

function CancelButton({ onCancel }) {
  return (
    <Button variant="secondary" onClick={onCancel}>
      Anuluj
    </Button>
  );
}
CancelButton.propTypes = {
  onCancel: EditModalContent.propTypes.onCancel,
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
  state: PropTypes.any.isRequired,
  onSubmit: EditModalContent.propTypes.onSubmit,
  onCancel: EditModalContent.propTypes.onCancel,
};
