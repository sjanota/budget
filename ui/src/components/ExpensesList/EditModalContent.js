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
            ref={autoFocusRef}
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
                placeholder="Data"
                value={state.date}
                onChange={e => setValue({ date: e.target.value })}
              />
            </Col>
            <Col>
              <Form.Label>Konto</Form.Label>
              <Form.Control
                type="text"
                placeholder="Under construction"
                readOnly={true}
              />
            </Col>
          </Row>
          <hr />
          <Form.Label>Kategorie</Form.Label>
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
          {state.entries.map((entry, idx) => (
            <EditModalContentEntry
              key={idx}
              entry={entry}
              idx={idx}
              setEntry={setEntry}
            />
          ))}
          <hr />
          <Row>
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
          </Row>
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
