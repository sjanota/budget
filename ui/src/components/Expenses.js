import React, { useState, useEffect } from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import ModalButton from './template/Utilities/ModalButton';
import CreateButton from './template/Utilities/CreateButton';
import EditTableButton from './template/Utilities/EditTableButton';
import { FormControl } from './template/Utilities/FormControl';
import FormModal from './template/Utilities/FormModal';
import { createFormData } from './template/Utilities/createFormData';
import Amount from '../model/Amount';
import { QueryTablePanel } from './gql/QueryTablePanel';
import { useCreateExpense, useGetCurrentExpenses } from './gql/expenses';
import { useGetAccounts } from './gql/accounts';
import { useGetCategories } from './gql/categories';
import { WithQuery } from './gql/WithQuery';
import { useBudget } from './gql/BudgetContext';
import Month from '../model/Month';
import { Form, Row, Col } from 'react-bootstrap';

const columns = [
  { dataField: 'title', text: 'Title' },
  {
    dataField: 'date',
    text: 'Date',
  },
  {
    dataField: 'totalAmount',
    text: 'Amount',
    formatter: Amount.format,
  },
  {
    dataField: 'account',
    text: 'Account',
    formatter: a => a.name,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        {/* <UpdateExpenseButton account={row} /> */}
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

const rowClasses = (row, rowIndex) => {
  return rowIndex % 2 === 0 && 'table-striped';
};

const expandRow = {
  className: 'background-color-white',
  renderer: row => (
    <table className="table table-sm">
      <tbody>
        {row.categories.map((category, idx) => (
          <tr key={idx}>
            <td>{category.category.name}</td>
            <td>{Amount.format(category.amount)}</td>
          </tr>
        ))}
      </tbody>
    </table>
  ),
};

function CategoriesInput({ formData }) {
  const query = useGetCategories();
  const [categories, setCategories] = useState(
    formData.init.map(c =>
      createFormData({
        categoryID: { $init: c.category.id },
        amount: { $init: Amount.format(c.amount), $process: Amount.parse },
      })
    )
  );
  useEffect(() => {
    formData.current = { value: categories };
  }, [categories]);
  return (
    <WithQuery query={query}>
      {({ data }) => (
        <>
          <small>
            Categories
            <i
              className="fas fa-fw fa-plus text-primary"
              onClick={() =>
                setCategories(c => [
                  ...c,
                  createFormData({
                    categoryID: { $init: null },
                    amount: { $init: null, $process: Amount.parse },
                  }),
                ])
              }
            />
          </small>
          {categories.map((categoryFormData, idx) => (
            <Form.Group as={Row} key={idx}>
              <Col>
                <Form.Control
                  placeholder="Category"
                  defaultValue={categoryFormData.categoryID.init}
                  ref={categoryFormData.categoryID}
                  as="select"
                  required
                >
                  <option></option>
                  {data.categories.map(({ id, name }) => (
                    <option key={id} value={id}>
                      {name}
                    </option>
                  ))}
                </Form.Control>
              </Col>
              <Col>
                <Form.Control
                  type="number"
                  required
                  placeholder="Amount"
                  defaultValue={categoryFormData.amount.init}
                  ref={categoryFormData.amount}
                  step="0.01"
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
  const accountsQuery = useGetAccounts();
  const formData = createFormData({
    title: { $init: init.title },
    date: { $init: init.date },
    accountID: { $init: init.account.id },
    categories: { $init: init.categories },
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
              label="Title"
              inline={10}
              formData={formData.title}
              feedback="Provide name"
              required
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
              label="Account"
              inline={9}
              formData={formData.accountID}
              feedback="Provide account"
              as="select"
              required
            >
              <option></option>
              {accountsData.accounts.map(({ id, name }) => (
                <option key={id} value={id}>
                  {name}
                </option>
              ))}
            </FormControl>
            <CategoriesInput formData={formData.categories} />
          </>
        )}
      </WithQuery>
    </FormModal>
  );
}

// function UpdateExpenseButton({ expense }) {
//   const [updateExpense] = useUpdateExpense();
//   return (
//     <ModalButton
//       button={EditTableButton}
//       modal={props => (
//         <ExpenseModal
//           init={expense}
//           title="Edit expense"
//           onSave={input => updateExpense(expense.id, input)}
//           {...props}
//         />
//       )}
//     />
//   );
// }

function CreateExpenseButton() {
  const [createExpense] = useCreateExpense();
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
          title="Add new expense"
          onSave={createExpense}
          {...props}
        />
      )}
    />
  );
}

export default function Expenses() {
  const query = useGetCurrentExpenses();

  return (
    <Page>
      <PageHeader>Expenses</PageHeader>
      <QueryTablePanel
        title="Expense list"
        query={query}
        getData={data => data.budget.currentMonth.expenses}
        buttons={<CreateExpenseButton />}
        columns={columns}
        keyField="id"
        expandRow={expandRow}
        rowClasses={rowClasses}
        striped={false}
      />
    </Page>
  );
}
