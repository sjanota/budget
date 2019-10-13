import React from 'react';
import Page from './template/Page/Page';
import PageHeader from './template/Page/PageHeader';
import ModalButton from './template/Utilities/ModalButton';
import CreateButton from './template/Utilities/CreateButton';
import EditTableButton from './template/Utilities/EditTableButton';
import { FormControl } from './template/Utilities/FormControl';
import FormModal from './template/Utilities/FormModal';
import useFormData from './template/Utilities/useFormData';
import Amount from '../model/Amount';
import { QueryTablePanel } from './gql/QueryTablePanel';
import { useCreateExpense, useGetCurrentExpenses } from './gql/expenses';

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

function ExpenseModal({ init, ...props }) {
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
          init={{ name: '' }}
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
