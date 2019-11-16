import React, { useRef, useState, useEffect } from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import ModalButton from './template/Utilities/ModalButton';
import CreateButton from './template/Utilities/CreateButton';
import EditTableButton from './template/Utilities/EditTableButton';
import { FormControl } from './template/Utilities/FormControl';
import FormModal from './template/Utilities/FormModal';
import { useFormData } from './template/Utilities/useFormData';
import Amount from '../model/Amount';
import Month from '../model/Month';
import {
  useCreateTransfer,
  useGetCurrentTransfers,
  useUpdateTransfer,
  useDeleteTranfer,
} from './gql/transfers';
import { QueryTablePanel } from './gql/QueryTablePanel';
import { useGetAccounts } from './gql/accounts';
import { useBudget } from './gql/budget';
import { WithQuery } from './gql/WithQuery';
import TableButton from './template/Utilities/TableButton';
import { HotKeys } from 'react-hotkeys';
import classnames from 'classnames';

const columns = [
  { dataField: 'title', text: 'Title' },
  {
    dataField: 'amount',
    text: 'Amount',
    align: 'right',
    formatter: Amount.format,
  },
  {
    dataField: 'fromAccount',
    text: 'From',
    formatter: a => a && a.name,
  },
  {
    dataField: 'toAccount',
    text: 'To',
    formatter: a => a.name,
  },
  {
    dataField: 'date',
    text: 'Date',
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateTransferButton transfer={row} />
        <DeleteTransferButton transfer={row} />
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

function Combobox({ allowedValues, _ref, defaultValue, className }) {
  const defaultValueObject = allowedValues.find(v => v.id === defaultValue);
  const [show, setShow] = useState(false);
  const [filter, setFilter] = useState(
    defaultValueObject ? defaultValueObject.label : ''
  );
  const [selectedIdx, setSelectedIdx] = useState(0);

  const menuRef = useRef();
  const classNames = classnames('input-group', className);
  const filtered = allowedValues.filter(v =>
    v.label.toLowerCase().includes(filter.toLowerCase())
  );

  useEffect(() => {
    _ref.current = { value: defaultValue };
  }, []);

  function onClick(id) {
    const selectedLabel = allowedValues.find(v => v.id === id).label;
    setShow(false);
    setFilter(selectedLabel);
    _ref.current = { value: id };
  }

  function onInputChange(e) {
    const value = e.target.value;
    setShow(true);
    setFilter(value);
    setSelectedIdx(0);
    const selected = allowedValues.find(v => v.label === value);
    if (selected) {
      _ref.current = { value: selected.id };
    }
  }

  function onInputBlur() {
    setShow(false);
    const selected = allowedValues.find(v => v.label === filter);
    if (!selected) {
      setFilter('');
    }
  }

  function onKeyDown(e) {
    if (e.keyCode === 40) {
      // up
      e.preventDefault();
      if (selectedIdx >= filtered.length - 1) {
        setSelectedIdx(0);
      } else {
        setSelectedIdx(v => v + 1);
      }
    } else if (e.keyCode === 38) {
      // down
      e.preventDefault();
      if (selectedIdx <= 0) {
        setSelectedIdx(filtered.length - 1);
      } else {
        setSelectedIdx(v => v - 1);
      }
    } else if (e.keyCode === 13 && show) {
      e.preventDefault();
      onClick(filtered[selectedIdx].id);
    }
  }

  return (
    <div className={classNames}>
      <input
        className="form-control"
        value={filter}
        type="text"
        onChange={onInputChange}
        onBlur={onInputBlur}
        onKeyDown={onKeyDown}
      />
      <button
        className="btn btn-secondary dropdown-toggle dropdown-toggle-split no-arrow"
        data-toggle="dropdown"
        data-reference="parent"
        data-flip="false"
        style={{
          maxWidth: '2rem',
          borderTopLeftRadius: 0,
          borderBottomLeftRadius: 0,
          margin: -1,
        }}
      />
      <ul
        ref={menuRef}
        role="menu"
        className={classnames('dropdown-menu', { show })}
        style={{
          maxHeight: '200px',
          overflowY: 'auto',
        }}
      >
        {filtered.map((v, idx) => (
          <li
            className={`dropdown-item ${idx === selectedIdx ? 'active' : ''}`}
            onClick={() => onClick(v.id)}
            key={v.id}
          >
            {v.label}
          </li>
        ))}
      </ul>
    </div>
  );
}

function TransferModal({ init, ...props }) {
  const { selectedBudget } = useBudget();
  const query = useGetAccounts();
  const formData = useFormData({
    title: { $init: init.title },
    date: { $init: init.date },
    amount: { $init: Amount.format(init.amount), $process: Amount.parse },
    fromAccountID: {
      $init: init.fromAccount && init.fromAccount.id,
      $process: v => (v === '' ? null : v),
    },
    toAccountID: { $init: init.toAccount.id },
  });
  const month = Month.parse(selectedBudget.currentMonth.month);
  const first = month.firstDay();
  const last = month.lastDay();
  return (
    <FormModal formData={formData} autoFocusRef={formData.title} {...props}>
      <WithQuery query={query}>
        {({ data }) => (
          <>
            <FormControl
              required
              label="Title"
              inline={10}
              formData={formData.title}
              feedback="Provide title"
            />
            <FormControl
              label="Date"
              inline={10}
              formData={formData.date}
              feedback="Provide date"
              type="date"
              required
              min={first.format()}
              max={last.format()}
            />
            <FormControl
              inline={10}
              label="Amount"
              feedback="Provide amount"
              type="number"
              required
              formData={formData.amount}
              step="0.01"
            />
            <div className="form-group row">
              <label className="col col-form-label">From</label>
              <div className="col-sm-10">
                <Combobox
                  allowedValues={data.accounts.map(({ id, name }) => ({
                    id,
                    label: name,
                  }))}
                  _ref={formData.fromAccountID}
                  defaultValue={formData.fromAccountID.default()}
                />
              </div>
            </div>
            <div className="form-group row">
              <label className="col col-form-label">To</label>
              <div className="col-sm-10">
                <Combobox
                  allowedValues={data.accounts.map(({ id, name }) => ({
                    id,
                    label: name,
                  }))}
                  _ref={formData.toAccountID}
                  defaultValue={formData.toAccountID.default()}
                />
              </div>
            </div>
          </>
        )}
      </WithQuery>
    </FormModal>
  );
}

function DeleteTransferButton({ transfer }) {
  const [deleteTransfer] = useDeleteTranfer();
  return (
    <TableButton
      faIcon="trash-alt"
      variant="secondary"
      onClick={() => deleteTransfer(transfer.id)}
    />
  );
}

function UpdateTransferButton({ transfer }) {
  const [updateTransfer] = useUpdateTransfer();
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <TransferModal
          init={transfer}
          title="Edit transfer"
          onSave={input => updateTransfer(transfer.id, input)}
          {...props}
        />
      )}
    />
  );
}

function CreateTransferButton({ openRef }) {
  const [createTransfer] = useCreateTransfer();
  return (
    <ModalButton
      openRef={openRef}
      button={CreateButton}
      modal={props => (
        <TransferModal
          init={{
            title: null,
            fromAccount: { id: null },
            toAccount: { id: null },
            amount: null,
            date: null,
          }}
          title="Add new transfer"
          onSave={createTransfer}
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

export default function Transfers() {
  const query = useGetCurrentTransfers();
  const openCreateModal = useRef();

  return (
    <HotKeys keyMap={keyMap} handlers={keyHandlers(openCreateModal)}>
      <Page>
        <PageHeader>Transfers</PageHeader>
        <QueryTablePanel
          title="Transfer list"
          query={query}
          getData={data => data.budget.currentMonth.transfers}
          buttons={<CreateTransferButton openRef={openCreateModal} />}
          columns={columns}
          keyField="id"
        />
      </Page>
    </HotKeys>
  );
}
