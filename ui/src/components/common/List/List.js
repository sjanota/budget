import React from 'react';
import { Table } from './Table';
import EditModalList from './EditModalList';
import EditInlineList from './EditInlineList';
import './List.css';
import PropTypes from 'prop-types';

List.EditMode = {
  MODAL: 'MODAL',
  INLINE: 'INLINE',
};

function partiallyApplyTable(renderHeader) {
  // eslint-disable-next-line react/display-name
  return props => <Table renderHeader={renderHeader} {...props} />;
}

export function List({ editMode, renderHeader, ...props }) {
  const table = partiallyApplyTable(renderHeader);
  let el;
  switch (editMode) {
    case List.EditMode.MODAL:
      el = <EditModalList tableComponent={table} {...props} />;
      break;
    case List.EditMode.INLINE:
      el = <EditInlineList tableComponent={table} {...props} />;
      break;
    default:
  }

  return <div className={'List'}>{el}</div>;
}

List.propTypes = {
  editMode: PropTypes.oneOf(Object.values(List.EditMode)),
  renderHeader: PropTypes.func.isRequired,
};

List.defaultProps = {
  editMode: List.EditMode.INLINE,
};
