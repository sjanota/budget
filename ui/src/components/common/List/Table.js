import React from 'react';
import PropTypes from 'prop-types';
import { ListHeader } from './ListHeader';
import { Table as BootstrapTable } from 'react-bootstrap';

export function Table({ children, onCreate, renderHeader }) {
  return (
    <BootstrapTable striped bordered hover size={'sm'}>
      <ListHeader onCreate={onCreate} renderHeader={renderHeader} />
      <tbody>{children}</tbody>
    </BootstrapTable>
  );
}

Table.propTypes = {
  children: PropTypes.node,
  onCreate: ListHeader.propTypes.onCreate,
  renderHeader: ListHeader.propTypes.onCreate,
};
