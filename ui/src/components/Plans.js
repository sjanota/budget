import React from 'react';
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
import { useCreatePlan, useGetCurrentPlans, useUpdatePlan } from './gql/plans';
import { QueryTablePanel } from './gql/QueryTablePanel';
import { useGetEnvelopes } from './gql/envelopes';
import { WithQuery } from './gql/WithQuery';

const columns = [
  { dataField: 'title', text: 'Title' },
  {
    dataField: 'currentAmount',
    text: 'Amount',
    formatter: Amount.format,
  },
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
    dataField: 'recurringAmount',
    text: 'Recurring',
    formatter: a => (
      <div className="custom-control custom-checkbox">
        <input
          className="custom-control-input"
          type="checkbox"
          checked={a != null}
          disabled
        />
        <label className="custom-control-label">{Amount.format(a)}</label>
      </div>
    ),
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdatePlanButton plan={row} />
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

function PlanModal({ init, ...props }) {
  const query = useGetEnvelopes();
  const formData = useFormData({
    title: { $init: init.title },
    currentAmount: {
      $init: Amount.format(init.currentAmount),
      $process: Amount.parse,
    },
    recurringAmount: {
      $init: Amount.format(init.recurringAmount),
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
            <FormControl
              label="From"
              inline={8}
              formData={formData.fromEnvelopeID}
              feedback="Provide from"
              as="select"
            >
              <option selected />
              {data.envelopes.map(({ id, name }) => (
                <option key={id} value={id}>
                  {name}
                </option>
              ))}
            </FormControl>
            <FormControl
              label="To"
              inline={8}
              formData={formData.toEnvelopeID}
              feedback="Provide to"
              as="select"
              required
            >
              <option disabled selected />
              {data.envelopes.map(({ id, name }) => (
                <option key={id} value={id}>
                  {name}
                </option>
              ))}
            </FormControl>
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

function CreatePlanButton() {
  const [createPlan] = useCreatePlan();
  return (
    <ModalButton
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

export default function Plans() {
  const query = useGetCurrentPlans();

  return (
    <Page>
      <PageHeader>Plans</PageHeader>
      <QueryTablePanel
        title="Plan list"
        query={query}
        getData={data => data.budget.currentMonth.plans}
        buttons={<CreatePlanButton />}
        columns={columns}
        keyField="id"
      />
    </Page>
  );
}