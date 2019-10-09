import React from 'react';
import Panel from '../template/Utilities/Panel';
import BootstrapTable from 'react-bootstrap-table-next';
import RefreshButton from '../template/Utilities/RefreshButton';
import WithQuery from '../gql/WithQuery';
import { useGetEnvelopes } from '../gql/envelopes';
import { CreateEnvelopeButton } from './CreateEnvelopeButton';
import Amount from '../../model/Amount';
import ArchiveTableButton from '../template/Utilities/ArchiveTableButton';
import { UpdateEnvelopeButton } from './UpdateEnvelopeButton';

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
    <Panel
      header={
        <div className="d-flex justify-content-between align-items-center">
          <Panel.Title>Envelope list</Panel.Title>
          <div>
            <WithQuery query={query} size="sm">
              {({ refetch }) => (
                <>
                  <RefreshButton
                    style={{ marginRight: '5px' }}
                    onClick={() => refetch()}
                  />
                  <CreateEnvelopeButton />
                </>
              )}
            </WithQuery>
          </div>
        </div>
      }
      body={
        <WithQuery query={query}>
          {({ data }) => (
            <BootstrapTable
              classes="table-layout-auto"
              keyField="id"
              data={data.envelopes}
              columns={columns}
              striped
              hover
            />
          )}
        </WithQuery>
      }
    />
  );
}
