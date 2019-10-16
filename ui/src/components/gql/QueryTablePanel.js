import React from 'react';
import { Panel } from '../template/Utilities/Panel';
import BootstrapTable from 'react-bootstrap-table-next';
import RefreshButton from '../template/Utilities/RefreshButton';
import { WithQuery } from '../gql/WithQuery';
import PropTypes from 'prop-types';

export function QueryTablePanel({ title, query, buttons, getData, ...props }) {
  return (
    <Panel
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
      body={
        <WithQuery query={query}>
          {({ data }) => (
            <BootstrapTable
              classes="table-layout-auto table-sm"
              data={getData(data)}
              striped
              hover
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
