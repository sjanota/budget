import React from 'react';
import { useGetCategories } from '../gql/categories';
import ArchiveTableButton from '../template/Utilities/ArchiveTableButton';
import { QueryTablePanel } from '../gql/QueryTablePanel';
import { CreateCategoryButton } from './CreateCategoryButton';
import { UpdateCategoryButton } from './UpdateCategoryButton';

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'envelope',
    text: 'Envelope',
    formatter: a => a.name,
  },
  {
    dataField: 'actions',
    text: '',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateCategoryButton category={row} />
        <ArchiveTableButton />
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

export function CategoriesListPanel({ createFunRef }) {
  const query = useGetCategories();
  return (
    <QueryTablePanel
      buttons={<CreateCategoryButton openRef={createFunRef} />}
      title="Category list"
      query={query}
      getData={data => data.categories}
      columns={columns}
      keyField="id"
    />
  );
}
