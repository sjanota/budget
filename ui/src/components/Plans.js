import React, { useRef } from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import ModalButton from './template/Utilities/ModalButton';
import CreateButton from './template/Utilities/CreateButton';
import EditTableButton from './template/Utilities/EditTableButton';
import { FormControl } from './template/Utilities/FormControl';
import { OptionalFormControl } from './template/Utilities/OptionalFormControl';
import FormModal from './template/Utilities/FormModal';
import { useFormData } from './template/Utilities/useFormData';
import Amount from '../model/Amount';
import {
  useCreatePlan,
  useGetCurrentPlans,
  useUpdatePlan,
  useDeletePlan,
} from './gql/plans';
import { QueryTablePanel } from './gql/QueryTablePanel';
import { useGetEnvelopes } from './gql/envelopes';
import { WithQuery } from './gql/WithQuery';
import TableButton from './template/Utilities/TableButton';
import { GlobalHotKeys } from 'react-hotkeys';
import { InlineFormControl } from './template/Utilities/InlineFormControl';
import { Combobox } from './template/Utilities/Combobox';

const columns = [
  { dataField: 'title', text: 'Title' },
  {
    dataField: 'fromEnvelope',
    text: 'From',
    formatter: a => a && a.name,
  },
  {
    dataField: 'toEnvelope',
    text: 'To',
    formatter: a => a.name,
  },
  {
    dataField: 'currentAmount',
    text: 'Amount',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'recurringAmount',
    text: '',
    formatter: a =>
      a !== null ? <i className="fas fa-fw fa-sync-alt" /> : null,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdatePlanButton plan={row} />
        <DeletePlanButton plan={row} />
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

function PlanModal({ init, ...props }) {
  const query = useGetEnvelopes();
  const formData = useFormData({
    title: { $init: init.title },
    currentAmount: {
      $init: Amount.format(init.currentAmount, false),
      $process: Amount.parse,
    },
    recurringAmount: {
      $init: Amount.format(init.recurringAmount, false),
      $default: fd => Amount.format(fd.currentAmount.value()),
      $process: Amount.parse,
    },
    fromEnvelopeID: {
      $init: init.fromEnvelope !== null ? init.fromEnvelope.id : '',
      $process: v => (v === '' ? null : v),
    },
    toEnvelopeID: { $init: init.toEnvelope.id },
  });
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
              inline={8}
              label="Amount"
              feedback="Provide amount"
              type="number"
              required
              formData={formData.currentAmount}
              step="0.01"
            />
            <OptionalFormControl
              initEnabled={init.recurringAmount !== null}
              inline={8}
              label="Recurring"
              feedback="Provide amount for recurring plans"
              type="number"
              required
              formData={formData.recurringAmount}
              step="0.01"
            />
            <InlineFormControl size={8} label="From">
              <Combobox
                _ref={formData.fromEnvelopeID}
                defaultValue={formData.fromEnvelopeID.default()}
                allowedValues={data.envelopes.map(({ id, name }) => ({
                  id,
                  label: name,
                }))}
              />
            </InlineFormControl>
            <InlineFormControl size={8} label="To">
              <Combobox
                _ref={formData.toEnvelopeID}
                defaultValue={formData.toEnvelopeID.default()}
                allowedValues={data.envelopes.map(({ id, name }) => ({
                  id,
                  label: name,
                }))}
              />
            </InlineFormControl>
          </>
        )}
      </WithQuery>
    </FormModal>
  );
}

function UpdatePlanButton({ plan }) {
  const [updatePlan] = useUpdatePlan();
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <PlanModal
          init={plan}
          title="Edit plan"
          onSave={input => updatePlan(plan.id, input)}
          {...props}
        />
      )}
    />
  );
}

function DeletePlanButton({ plan }) {
  const [deletePlan] = useDeletePlan();
  return (
    <TableButton
      faIcon="trash-alt"
      variant="secondary"
      onClick={() => deletePlan(plan.id)}
    />
  );
}

function CreatePlanButton({ openRef }) {
  const [createPlan] = useCreatePlan();
  return (
    <ModalButton
      openRef={openRef}
      button={CreateButton}
      modal={props => (
        <PlanModal
          init={{
            title: null,
            fromEnvelope: { id: null },
            toEnvelope: { id: null },
            currentAmount: null,
            recurringAmount: null,
            date: null,
          }}
          title="Add new plan"
          onSave={createPlan}
          {...props}
        />
      )}
    />
  );
}

const keyMap = {
  create: 'n',
};

const handlers = openCreateModalRef => ({
  create: () => openCreateModalRef.current(),
});

export default function Plans() {
  const openCreateModalRef = useRef();
  const query = useGetCurrentPlans();

  return (
    <Page>
      <GlobalHotKeys keyMap={keyMap} handlers={handlers(openCreateModalRef)} />
      <PageHeader>Plans</PageHeader>
      <QueryTablePanel
        title="Plan list"
        query={query}
        getData={data => data.budget.currentMonth.plans}
        buttons={<CreatePlanButton openRef={openCreateModalRef} />}
        columns={columns}
        keyField="id"
      />
    </Page>
  );
}
