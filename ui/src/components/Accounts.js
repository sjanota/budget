import React, { useState, useRef } from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import Panel from './template/Utilities/Panel';
import Spinner from './template/Utilities/Spinner';
import gql from 'graphql-tag';
import { useQuery, useMutation } from '@apollo/react-hooks';
import BootstrapTable from 'react-bootstrap-table-next';
import { useBudget } from './contexts/BudgetContext';
import SplitButton from './template/Utilities/SplitButton';
import { Button, Modal, InputGroup, FormControl } from 'react-bootstrap';

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

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: ({ integer, decimal }) => `${integer}.${decimal}`,
  },
];

function CreateAccountModal() {
  const [show, setShow] = useState(false);
  const fields = {
    name: useRef(),
  };
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
  const onSave = () => {
    const input = { name: fields.name.current.value };
    createAccount({ variables: { budgetID: selectedBudget.id, in: input } });
    onClose();
  };
  return (
    <>
      <SplitButton faIcon="plus" size="small" onClick={() => setShow(true)}>
        Add new account
      </SplitButton>
      <Modal show={show} onHide={onClose}>
        <Modal.Header closeButton>
          <Modal.Title>Add new account</Modal.Title>
        </Modal.Header>
        <Modal.Body>
          <InputGroup className="mb-3">
            <InputGroup.Prepend>
              <InputGroup.Text className="border-0">Name</InputGroup.Text>
            </InputGroup.Prepend>
            <FormControl
              placeholder="Account name"
              className="bg-light border-0 text-dark"
              ref={fields.name}
            />
          </InputGroup>
        </Modal.Body>
        <Modal.Footer>
          <SplitButton
            variant="danger"
            faIcon="times"
            size="small"
            onClick={onClose}
          >
            Cancel
          </SplitButton>
          <SplitButton faIcon="save" size="small" onClick={onSave}>
            Save
          </SplitButton>
        </Modal.Footer>
      </Modal>
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
              <CreateAccountModal />
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
              keyField="id"
              data={data.accounts}
              columns={columns}
            />
          )
        }
      />
    </Page>
  );
}
