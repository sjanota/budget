import React from 'react';
import { Panel } from '../template/Utilities/Panel';
import BootstrapTable from 'react-bootstrap-table-next';
import RefreshButton from '../template/Utilities/RefreshButton';
import { WithQuery } from '../gql/WithQuery';
import PropTypes from 'prop-types';

export function QueryTablePanel({
  title,
  query,
  buttons,
  getData,
  columns,
  ...props
}) {
  const paddedFirstColumn = {
    ...columns[0],
    classes: 'pl-3',
    headerClasses: 'pl-3',
  };
  const modifiedColumns = [
    paddedFirstColumn,
    ...columns.slice(1, columns.length),
  ];

  return (
    <Panel
      headerClassName="p-2 pl-3"
      header={
        <Panel.HeaderWithButton title={title}>
          <WithQuery query={query} size="sm" showError={false}>
            {({ refetch }) => (
              <>
                <RefreshButton onClick={() => refetch()} className="mr-1" />
                {buttons}
              </>
            )}
          </WithQuery>
        </Panel.HeaderWithButton>
      }
      bodyClassName="p-0"
      body={
        <WithQuery query={query}>
          {({ data }) => (
            <BootstrapTable
              bootstrap4
              classes="table-layout-auto table-sm m-0"
              data={getData(data)}
              striped
              hover
              bordered={false}
              columns={modifiedColumns}
              {...props}
            />
          )}
        </WithQuery>
      }
    />
  );
}

QueryTablePanel.propTypes = {
  title: PropTypes.string.isRequired,
  buttons: PropTypes.node,
  getData: PropTypes.func.isRequired,
  query: PropTypes.shape({
    data: PropTypes.any,
    refetch: PropTypes.func,
  }).isRequired,
};
