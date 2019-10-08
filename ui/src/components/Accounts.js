import React from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import Panel from './template/Utilities/Panel';
import Spinner from './template/Utilities/Spinner';
import gql from 'graphql-tag';
import { useQuery } from '@apollo/react-hooks';
import BootstrapTable from 'react-bootstrap-table-next';
import { useBudget } from './contexts/BudgetContext';
import SplitButton from './template/Utilities/SplitButton';

const GET_ACCOUNTS = gql`
  query GetAccounts($budgetID: ID!) {
    budget(id: $budgetID) {
      accounts {
        id
        name
        balance
      }
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

export default function Accounts() {
  const { selectedBudget } = useBudget();
  const { loading, error, data } = useQuery(GET_ACCOUNTS, {
    variables: { budgetID: selectedBudget.id },
  });

  return (
    <Page>
      <PageHeader>Accounts</PageHeader>
      <Panel
        header={
          <div className="d-flex justify-content-between align-items-center">
            <Panel.Title>Account list</Panel.Title>
            <SplitButton faIcon="plus" size="small">
              Add new account
            </SplitButton>
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
              data={data.budget.accounts}
              columns={columns}
            />
          )
        }
      />
    </Page>
  );
}
