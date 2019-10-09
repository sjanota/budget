import React, { useState, useRef } from 'react';
import { useBudget } from './contexts/BudgetContext';
import gql from 'graphql-tag';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import Panel from './template/Utilities/Panel';
import Spinner from './template/Utilities/Spinner';
import { Button, Modal, InputGroup, Form } from 'react-bootstrap';
import BootstrapTable from 'react-bootstrap-table-next';
import { useQuery, useMutation } from '@apollo/react-hooks';
import SplitButton from './template/Utilities/SplitButton';

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

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'limit',
    text: 'Limit',
    formatter: ({ integer, decimal }) => `${integer}.${decimal}`,
  },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: ({ integer, decimal }) => `${integer}.${decimal}`,
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
  // const [createAccount] = useMutation(CREATE_ACCOUNT, {
  //   update: (cache, { data: { createAccount } }) => {
  //     const { accounts } = cache.readQuery({
  //       query: GET_ACCOUNTS,
  //       variables: { budgetID: selectedBudget.id },
  //     });
  //     cache.writeQuery({
  //       query: GET_ACCOUNTS,
  //       variables: { budgetID: selectedBudget.id },
  //       data: {
  //         accounts: accounts.concat([createAccount]),
  //       },
  //     });
  //   },
  // });
  const onClose = () => setShow(false);
  const onSave = input => {
    //   createAccount({ variables: { budgetID: selectedBudget.id, in: input } });
  };
  return (
    <>
      <SplitButton faIcon="plus" size="small" onClick={() => setShow(true)}>
        Add new account
      </SplitButton>
      {/* <EditAccountModal show={show} onClose={onClose} onSave={onSave} /> */}
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
            <Panel.Title>Account list</Panel.Title>
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
