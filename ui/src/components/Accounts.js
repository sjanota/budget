import React, { useState, useRef } from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import Panel from './template/Utilities/Panel';
import Spinner from './template/Utilities/Spinner';
import gql from 'graphql-tag';
import { useQuery, useMutation } from '@apollo/react-hooks';
import BootstrapTable from 'react-bootstrap-table-next';
import { useBudget } from './gql/BudgetContext';
import SplitButton from './template/Utilities/SplitButton';
import { Button, Modal, Form } from 'react-bootstrap';
import Amount from '../model/Amount';

const GET_ACCOUNTS = gql`
  query GetAccounts($budgetID: ID!) {
    accounts(budgetID: $budgetID) {
      id
      name
      balance
    }
  }
`;

const CREATE_ACCOUNT = gql`
  mutation CreateAccount($budgetID: ID!, $in: AccountInput!) {
    createAccount(budgetID: $budgetID, in: $in) {
      id
      name
      balance
    }
  }
`;

const UPDATE_ACCOUNT = gql`
  mutation UpdateAccount($budgetID: ID!, $id: ID!, $in: AccountUpdate!) {
    updateAccount(budgetID: $budgetID, id: $id, in: $in) {
      id
      name
      balance
    }
  }
`;

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: Amount.format,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateAccountButton account={row} />
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

function EditAccountModal({ title, init, show, onClose, onSave }) {
  const [validated, setValidated] = useState(false);

  const initName = init && init.name;
  const form = useRef();
  const fields = {
    name: useRef(),
  };
  const handleSave = event => {
    event.preventDefault();
    event.stopPropagation();
    const isValid = form.current.checkValidity();
    setValidated(true);

    if (!isValid === false) {
      return;
    }
    const input = {};
    if (fields.name.current.value !== initName) {
      input.name = fields.name.current.value;
    }
    onSave(input);
    onClose();
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
              Name is required
            </Form.Control.Feedback>
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

function UpdateAccountButton({ account }) {
  const [show, setShow] = useState(false);
  const { selectedBudget } = useBudget();
  const [updateAccount] = useMutation(UPDATE_ACCOUNT);
  const onClose = () => setShow(false);
  const onSave = input => {
    updateAccount({
      variables: { budgetID: selectedBudget.id, id: account.id, in: input },
    });
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
      <EditAccountModal
        title="Edit account"
        init={account}
        show={show}
        onClose={onClose}
        onSave={onSave}
      />
    </>
  );
}

function CreateAccountButton() {
  const [show, setShow] = useState(false);
  const { selectedBudget } = useBudget();
  const [createAccount] = useMutation(CREATE_ACCOUNT, {
    update: (cache, { data: { createAccount } }) => {
      const { accounts } = cache.readQuery({
        query: GET_ACCOUNTS,
        variables: { budgetID: selectedBudget.id },
      });
      cache.writeQuery({
        query: GET_ACCOUNTS,
        variables: { budgetID: selectedBudget.id },
        data: {
          accounts: accounts.concat([createAccount]),
        },
      });
    },
  });
  const onClose = () => setShow(false);
  const onSave = input => {
    createAccount({ variables: { budgetID: selectedBudget.id, in: input } });
  };
  return (
    <>
      <SplitButton faIcon="plus" size="small" onClick={() => setShow(true)}>
        Add new account
      </SplitButton>
      <EditAccountModal
        title="Add new account"
        show={show}
        onClose={onClose}
        onSave={onSave}
      />
    </>
  );
}

export default function Accounts() {
  const { selectedBudget } = useBudget();
  const { loading, error, data, refetch } = useQuery(GET_ACCOUNTS, {
    variables: { budgetID: selectedBudget.id },
  });

  return (
    <Page>
      <PageHeader>Accounts</PageHeader>
      <Panel
        header={
          <div className="d-flex justify-content-between align-items-center">
            <Panel.Title>Account list</Panel.Title>
            <div>
              <Button
                className="btn-sm btn-secondary"
                style={{ marginRight: '5px' }}
                onClick={() => refetch()}
              >
                <i className="fas fa-fw fa-sync-alt" />
              </Button>
              <CreateAccountButton />
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
              data={data.accounts}
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
