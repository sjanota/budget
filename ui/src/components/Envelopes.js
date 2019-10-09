import React, { useState, useRef } from 'react';
import { useBudget } from './contexts/BudgetContext';
import gql from 'graphql-tag';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import Panel from './template/Utilities/Panel';
import Spinner from './template/Utilities/Spinner';
import {
  Button,
  Modal,
  InputGroup,
  Form,
  Row,
  Collapse,
} from 'react-bootstrap';
import BootstrapTable from 'react-bootstrap-table-next';
import { useQuery, useMutation } from '@apollo/react-hooks';
import SplitButton from './template/Utilities/SplitButton';
import Amount from '../model/Amount';

const GET_ENVELOPES = gql`
  query GetEnvelopes($budgetID: ID!) {
    envelopes(budgetID: $budgetID) {
      id
      name
      balance
      limit
    }
  }
`;

const CREATE_ENVELOPE = gql`
  mutation CreateEnvelope($budgetID: ID!, $in: EnvelopeInput!) {
    createEnvelope(budgetID: $budgetID, in: $in) {
      id
      name
      balance
      limit
    }
  }
`;

// const UPDATE_ACCOUNT = gql`
//   mutation UpdateAccount($budgetID: ID!, $id: ID!, $in: AccountUpdate!) {
//     updateAccount(budgetID: $budgetID, id: $id, in: $in) {
//       id
//       name
//       balance
//     }
//   }
// `;

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'limit',
    text: 'Limit',
    formatter: Amount.Formatter,
  },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: Amount.Formatter,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateEnvelopeButton envelope={row} />
        <span style={{ cursor: 'pointer' }}>
          <i className="fas fa-archive fa-fw" />
        </span>
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

function EditEnvelopeModal({ title, init, show, onClose, onSave }) {
  const [validated, setValidated] = useState(false);
  const initName = init && init.name;
  const initLimit = init && Amount.format(init.limit);
  const [limitEnabled, setLimitEnabled] = useState(!!initLimit);
  const form = useRef();
  const fields = {
    name: useRef(),
    limit: useRef(),
  };
  const handleSave = event => {
    event.preventDefault();
    event.stopPropagation();
    const isValid = form.current.checkValidity();
    setValidated(true);

    if (!isValid) {
      return;
    }
    const input = {};
    if (fields.name.current.value !== initName) {
      input.name = fields.name.current.value;
    }
    if (!limitEnabled && initLimit !== null) {
      input.limit = null;
    } else if (fields.limit.current.value !== initLimit) {
      input.limit = Amount.parse(fields.limit.current.value);
    }
    onSave(input);
    onClose();
    setValidated(false);
  };
  return (
    <Modal
      show={show}
      onHide={onClose}
      onEntered={() => fields.name.current.focus()}
    >
      <Form validated={validated} ref={form} onSubmit={handleSave}>
        <Modal.Header
          closeButton
          className="m-0 font-weight-bold text-primary bg-light"
        >
          <Modal.Title>{title}</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <Form.Group className="mb-3">
            <Form.Label>Name</Form.Label>
            <Form.Control required defaultValue={initName} ref={fields.name} />
            <Form.Control.Feedback type="invalid">
              Provide a name for the envelope
            </Form.Control.Feedback>
          </Form.Group>

          <Form.Group className="mb-3">
            <Form.Check custom type="switch">
              <Form.Check.Input checked={limitEnabled} onChange={() => {}} />
              <Form.Check.Label
                onClick={() => {
                  setLimitEnabled(v => !v);
                }}
              >
                <Form.Label>Limit</Form.Label>
              </Form.Check.Label>
            </Form.Check>
            {limitEnabled && (
              <>
                <Form.Control
                  required={limitEnabled}
                  type="number"
                  defaultValue={initLimit}
                  ref={fields.limit}
                  step="0.01"
                />
                <Form.Control.Feedback type="invalid">
                  Provide a limit for the envelope
                </Form.Control.Feedback>
              </>
            )}
            <Form.Control.Feedback type="invalid"></Form.Control.Feedback>
          </Form.Group>
        </Modal.Body>
        <Modal.Footer className=" bg-light">
          <SplitButton
            variant="danger"
            faIcon="times"
            size="small"
            onClick={onClose}
          >
            Cancel
          </SplitButton>
          <SplitButton
            faIcon="save"
            size="small"
            type="submit"
            onClick={handleSave}
          >
            Save
          </SplitButton>
        </Modal.Footer>
      </Form>
    </Modal>
  );
}

function UpdateEnvelopeButton({ envelope }) {
  const [show, setShow] = useState(false);
  const { selectedBudget } = useBudget();
  // const [updateAccount] = useMutation(UPDATE_ACCOUNT);
  const onClose = () => setShow(false);
  const onSave = input => {
    //   updateAccount({
    //     variables: { budgetID: selectedBudget.id, id: account.id, in: input },
    //   });
  };
  return (
    <>
      <span
        style={{ cursor: 'pointer', marginRight: '5px' }}
        onClick={() => {
          setShow(true);
        }}
      >
        <i className="fas fa-edit fa-fw text-primary" />
      </span>
      {/* <EditAccountModal
          init={account}
          show={show}
          onClose={onClose}
          onSave={onSave}
        /> */}
    </>
  );
}

function CreateEnvelopeButton() {
  const [show, setShow] = useState(false);
  const { selectedBudget } = useBudget();
  const [createEnvelope] = useMutation(CREATE_ENVELOPE, {
    update: (cache, { data: { createEnvelope } }) => {
      const { envelopes } = cache.readQuery({
        query: GET_ENVELOPES,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_ENVELOPES,
        variables: { budgetID: selectedBudget.id },
        data: {
          envelopes: envelopes.concat([createEnvelope]),
        },
      });
    },
  });
  const onClose = () => setShow(false);
  const onSave = input => {
    createEnvelope({ variables: { budgetID: selectedBudget.id, in: input } });
  };
  return (
    <>
      <SplitButton faIcon="plus" size="small" onClick={() => setShow(true)}>
        Add new envelope
      </SplitButton>
      <EditEnvelopeModal
        title="Add new envelope"
        show={show}
        onClose={onClose}
        onSave={onSave}
      />
    </>
  );
}

export default function Envelopes() {
  const { selectedBudget } = useBudget();
  const { loading, error, data, refetch } = useQuery(GET_ENVELOPES, {
    variables: { budgetID: selectedBudget.id },
  });

  return (
    <Page>
      <PageHeader>Envelopes</PageHeader>
      <Panel
        header={
          <div className="d-flex justify-content-between align-items-center">
            <Panel.Title>Envelope list</Panel.Title>
            <div>
              <Button
                className="btn-sm btn-secondary"
                style={{ marginRight: '5px' }}
                onClick={() => refetch()}
              >
                <i className="fas fa-fw fa-sync-alt" />
              </Button>
              <CreateEnvelopeButton />
            </div>
          </div>
        }
        body={
          loading ? (
            <Spinner />
          ) : error ? (
            <i className="fas fa-fw fa-exclamation-triangle text-secondary" />
          ) : (
            <BootstrapTable
              classes="table-layout-auto"
              keyField="id"
              data={data.envelopes}
              columns={columns}
              striped
              hover
            />
          )
        }
      />
    </Page>
  );
}
