import React from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import ModalButton from './template/Utilities/ModalButton';
import CreateButton from './template/Utilities/CreateButton';
import EditTableButton from './template/Utilities/EditTableButton';
import { FormControl } from './template/Utilities/FormControl';
import FormModal from './template/Utilities/FormModal';
import { useFormData } from './template/Utilities/useFormData';
import Amount from '../model/Amount';
import {
  useCreateAccount,
  useGetAccounts,
  useUpdateAccount,
} from './gql/accounts';
import { QueryTablePanel } from './gql/QueryTablePanel';

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

function AccountModal({ init, ...props }) {
  const formData = useFormData({
    name: { $init: init.name },
  });
  return (
    <FormModal formData={formData} autoFocusRef={formData.name} {...props}>
      <FormControl
        label="Name"
        inline={10}
        formData={formData.name}
        feedback="Provide name"
      />
    </FormModal>
  );
}

function UpdateAccountButton({ account }) {
  const [updateAccount] = useUpdateAccount();
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <AccountModal
          init={account}
          title="Edit account"
          onSave={input => updateAccount(account.id, input)}
          {...props}
        />
      )}
    />
  );
}

function CreateAccountButton() {
  const [createAccount] = useCreateAccount();
  return (
    <ModalButton
      button={CreateButton}
      modal={props => (
        <AccountModal
          init={{ name: '' }}
          title="Add new account"
          onSave={createAccount}
          {...props}
        />
      )}
    />
  );
}

export default function Accounts() {
  const query = useGetAccounts();

  return (
    <Page>
      <PageHeader>Accounts</PageHeader>
      <QueryTablePanel
        title="Accounts"
        query={query}
        getData={data => data.accounts}
        buttons={<CreateAccountButton />}
        columns={columns}
        keyField="id"
      />
    </Page>
  );
}
