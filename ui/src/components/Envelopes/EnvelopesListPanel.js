import React from 'react';
import { useGetEnvelopes } from '../gql/envelopes';
import { CreateEnvelopeButton } from './CreateEnvelopeButton';
import Amount from '../../model/Amount';
import ArchiveTableButton from '../template/Utilities/ArchiveTableButton';
import { UpdateEnvelopeButton } from './UpdateEnvelopeButton';
import { QueryTablePanel } from '../gql/QueryTablePanel';

const columns = [
  { dataField: 'name', text: 'Name' },
  {
    dataField: 'limit',
    text: 'Limit',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: Amount.format,
    align: 'right',
    headerAlign: 'right',
  },
  {
    dataField: 'overlimit',
    text: 'Over limit',
    align: 'right',
    headerAlign: 'right',
    formatter: (cell, row) =>
      row.limit !== null && row.limit < row.balance
        ? Amount.format(row.balance - row.limit)
        : '',
  },
  {
    dataField: 'actions',
    text: '',
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
  return (
    <QueryTablePanel
      title="Envelope list"
      query={query}
      buttons={<CreateEnvelopeButton openRef={createFunRef} />}
      getData={data => data.envelopes}
      columns={columns}
      keyField="id"
    />
  );
}
