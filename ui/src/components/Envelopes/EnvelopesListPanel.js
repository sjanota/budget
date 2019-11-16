import React from 'react';
import { useGetEnvelopes } from '../gql/envelopes';
import { CreateEnvelopeButton } from './CreateEnvelopeButton';
import Amount from '../../model/Amount';
import ArchiveTableButton from '../template/Utilities/ArchiveTableButton';
import { UpdateEnvelopeButton } from './UpdateEnvelopeButton';
import { QueryTablePanel } from '../gql/QueryTablePanel';
import { withColumnNames, useDictionary } from '../template/Utilities/Lang';

const columns = [
  { dataField: 'name' },
  {
    dataField: 'limit',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'balance',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'overLimit',
    align: 'right',
    headerAlign: 'right',
    formatter: (cell, row) =>
      row.limit !== null && row.limit < row.balance
        ? Amount.format(row.balance - row.limit)
        : '',
  },
  {
    dataField: 'actions',
    isDummyColumn: true,
    formatter: (cell, row) => (
      <span>
        <UpdateEnvelopeButton envelope={row} />
        <ArchiveTableButton />
      </span>
    ),
    style: {
      whiteSpace: 'nowrap',
      width: '1%',
    },
  },
];

export function EnvelopesListPanel({ createFunRef }) {
  const query = useGetEnvelopes();
  const { envelopes } = useDictionary();
  return (
    <QueryTablePanel
      title={envelopes.table.title}
      query={query}
      buttons={<CreateEnvelopeButton openRef={createFunRef} />}
      getData={data => data.envelopes}
      columns={withColumnNames(columns, envelopes.table.columns)}
      keyField="id"
    />
  );
}
