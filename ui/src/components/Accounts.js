import React, { useRef, useImperativeHandle } from 'react';
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
import { GlobalHotKeys } from 'react-hotkeys';
import { useDictionary } from './template/Utilities/Lang';

const columns = dictionary => [
  { dataField: 'name', text: dictionary.columns.name },
  {
    dataField: 'balance',
    text: dictionary.columns.balance,
    align: 'right',
    headerAlign: 'right',
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
  const { accounts } = useDictionary();
  const formData = useFormData({
    name: { $init: init.name },
  });
  return (
    <FormModal formData={formData} autoFocusRef={formData.name} {...props}>
      <FormControl
        label={accounts.modal.labels.name}
        inline={10}
        formData={formData.name}
        feedback="Provide name"
      />
    </FormModal>
  );
}

function UpdateAccountButton({ account }) {
  const [updateAccount] = useUpdateAccount();
  const { accounts } = useDictionary();
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <AccountModal
          init={account}
          title={accounts.modal.editTitle}
          onSave={input => updateAccount(account.id, input)}
          {...props}
        />
      )}
    />
  );
}

function CreateAccountButton({ openRef }) {
  const [createAccount] = useCreateAccount();
  const { accounts } = useDictionary();
  return (
    <ModalButton
      openRef={openRef}
      button={CreateButton}
      modal={props => (
        <AccountModal
          init={{ name: '' }}
          title={accounts.modal.createTitle}
          onSave={createAccount}
          {...props}
        />
      )}
    />
  );
}

const keyMap = {
  openModal: 'n',
};
const keyHandlers = openCreateModal => ({
  openModal: () => openCreateModal.current(),
});

export default function Accounts() {
  const query = useGetAccounts();
  const openCreateModal = useRef();
  const { accounts } = useDictionary();

  return (
    <GlobalHotKeys keyMap={keyMap} handlers={keyHandlers(openCreateModal)}>
      <Page>
        <PageHeader>{accounts.header}</PageHeader>
        <QueryTablePanel
          title={accounts.table.title}
          query={query}
          getData={data => data.accounts}
          buttons={<CreateAccountButton openRef={openCreateModal} />}
          columns={columns(accounts.table)}
          keyField="id"
        />
      </Page>
    </GlobalHotKeys>
  );
}
