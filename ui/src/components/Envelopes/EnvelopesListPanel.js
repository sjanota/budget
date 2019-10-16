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
  },
  {
    dataField: 'balance',
    text: 'Balance',
    formatter: Amount.format,
  },
  {
    dataField: 'overlimit',
    text: 'Over limit',
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

export function EnvelopesListPanel() {
  const query = useGetEnvelopes();
  return (
    <QueryTablePanel
      title="Envelope list"
      query={query}
      buttons={<CreateEnvelopeButton />}
      getData={data => data.envelopes}
      columns={columns}
      keyField="id"
    />
  );
}
