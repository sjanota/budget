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
import { QueryTablePanel } from './gql/QueryTablePanel';
import {
  useCreateExpense,
  useGetCurrentExpenses,
  useUpdateExpense,
  useDeleteExpense,
} from './gql/expenses';
import { useGetAccounts } from './gql/accounts';
import { useGetCategories } from './gql/categories';
import { WithQuery } from './gql/WithQuery';
import { useBudget } from './gql/budget';
import Month from '../model/Month';
import { Form, Row, Col } from 'react-bootstrap';
import TableButton from './template/Utilities/TableButton';
import { useDictionary, withColumnNames } from './template/Utilities/Lang';
import { Combobox } from './template/Utilities/Combobox';
import { InlineFormControl } from './template/Utilities/InlineFormControl';

const columns = [
  { dataField: 'title' },
  { dataField: 'date' },
  {
    dataField: 'account',
    formatter: a => a.name,
  },
  {
    dataField: 'totalAmount',
    text: 'Amount',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateExpenseButton expense={row} />
        <DeleteExpenseButton expense={row} />
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

const rowClasses = (row, rowIndex) => {
  return rowIndex % 2 === 0 && 'table-striped';
};

const expandRow = {
  className: 'background-color-white',
  renderer: row => (
    <table className="table table-sm mb-0">
      <tbody>
        {row.categories.map((category, idx) => (
          <tr key={idx}>
            <td className="pl-3">{category.category.name}</td>
            <td>{Amount.format(category.amount)}</td>
          </tr>
        ))}
      </tbody>
    </table>
  ),
};

function CategoriesInput({ formData }) {
  const query = useGetCategories();
  const { expenses } = useDictionary();

  return (
    <WithQuery query={query}>
      {({ data }) => (
        <>
          <small className="d-flex align-items-center">
            {expenses.modal.labels.categories}
            <TableButton
              faIcon="plus"
              variant="primary"
              size="sm"
              onClick={() =>
                formData.push({
                  category: { id: null },
                  amount: null,
                })
              }
            />
          </small>
          {formData.map((categoryFormData, idx) => (
            <Form.Group
              as={Row}
              key={categoryFormData.categoryID.init() || idx}
              className="d-flex align-items-center"
            >
              <Col sm={7}>
                <Combobox
                  _ref={categoryFormData.categoryID}
                  defaultValue={categoryFormData.categoryID.init()}
                  allowedValues={data.categories.map(({ id, name }) => ({
                    id,
                    label: name,
                  }))}
                />
              </Col>
              <Col>
                <Form.Control
                  type="number"
                  required
                  placeholder={expenses.modal.labels.amount}
                  defaultValue={categoryFormData.amount.init()}
                  ref={categoryFormData.amount}
                  step="0.01"
                />
              </Col>
              <Col sm={1}>
                <TableButton
                  faIcon="minus"
                  variant="danger"
                  size="sm"
                  onClick={() => {
                    console.log(categoryFormData, idx);
                    formData.removeAt(idx);
                  }}
                />
              </Col>
            </Form.Group>
          ))}
        </>
      )}
    </WithQuery>
  );
}

function ExpenseModal({ init, ...props }) {
  const { selectedBudget } = useBudget();
  const { expenses } = useDictionary();
  const accountsQuery = useGetAccounts();
  const formData = useFormData({
    title: { $init: init.title },
    date: { $init: init.date },
    accountID: { $init: init.account.id },
    categories: {
      $init: init.categories,
      $model: c => ({
        categoryID: { $init: c.category.id },
        amount: {
          $init: Amount.format(c.amount, false),
          $process: Amount.parse,
        },
        $includeAllValues: true,
      }),
    },
  });
  const month = Month.parse(selectedBudget.currentMonth.month);
  const first = month.firstDay();
  const last = month.lastDay();
  return (
    <FormModal formData={formData} autoFocusRef={formData.title} {...props}>
      <WithQuery query={accountsQuery}>
        {({ data: accountsData }) => (
          <>
            <FormControl
              label={expenses.modal.labels.title}
              inline={10}
              formData={formData.title}
              feedback="Provide name"
              required
            />
            <FormControl
              label={expenses.modal.labels.date}
              inline={10}
              formData={formData.date}
              feedback="Provide date"
              type="date"
              required
              min={first.format()}
              max={last.format()}
            />
            <InlineFormControl label={expenses.modal.labels.account} size={9}>
              <Combobox
                _ref={formData.accountID}
                defaultValue={formData.accountID.init()}
                allowedValues={accountsData.accounts.map(({ id, name }) => ({
                  id,
                  label: name,
                }))}
              />
            </InlineFormControl>
            <CategoriesInput formData={formData.categories} />
          </>
        )}
      </WithQuery>
    </FormModal>
  );
}

function DeleteExpenseButton({ expense }) {
  const [deleteExpense] = useDeleteExpense();
  return (
    <TableButton
      faIcon="trash-alt"
      variant="secondary"
      onClick={() => deleteExpense(expense.id)}
    />
  );
}

function UpdateExpenseButton({ expense }) {
  const [updateExpense] = useUpdateExpense();
  const { expenses } = useDictionary();

  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <ExpenseModal
          init={expense}
          title={expenses.modal.editTitle}
          onSave={input => updateExpense(expense.id, input)}
          {...props}
        />
      )}
    />
  );
}

function CreateExpenseButton() {
  const [createExpense] = useCreateExpense();
  const { expenses } = useDictionary();

  return (
    <ModalButton
      button={CreateButton}
      modal={props => (
        <ExpenseModal
          init={{
            name: null,
            account: {},
            date: null,
            categories: [],
          }}
          title={expenses.modal.createTitle}
          onSave={createExpense}
          {...props}
        />
      )}
    />
  );
}

export default function Expenses() {
  const query = useGetCurrentExpenses();
  const { expenses, sidebar } = useDictionary();

  return (
    <Page>
      <PageHeader>{sidebar.pages.expenses}</PageHeader>
      <QueryTablePanel
        title={expenses.table.title}
        query={query}
        getData={data => data.budget.currentMonth.expenses}
        buttons={<CreateExpenseButton />}
        columns={withColumnNames(columns, expenses.table.columns)}
        keyField="id"
        expandRow={expandRow}
        rowClasses={rowClasses}
        striped={false}
      />
    </Page>
  );
}
