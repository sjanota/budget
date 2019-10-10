import React from 'react';
import { useGetCategories } from '../gql/categories';
import ArchiveTableButton from '../template/Utilities/ArchiveTableButton';
import { QueryTablePanel } from '../gql/QueryTablePanel';
import { CreateCategoryButton } from './CreateCategoryButton';

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'account',
    text: 'Account',
    formatter: a => a.id,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    // formatter: (cell, row) => (
    //   <span>
    //     <UpdateEnvelopeButton envelope={row} />
    //     <ArchiveTableButton />
    //   </span>
    // ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

export function CategoriesListPanel() {
  const query = useGetCategories();
  return (
    <QueryTablePanel
      buttons={<CreateCategoryButton />}
      title="Envelope list"
      query={query}
      getData={data => data.categories}
      columns={columns}
      keyField="id"
    />
  );
}
