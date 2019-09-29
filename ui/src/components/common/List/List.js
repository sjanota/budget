import React from 'react';
import { Table } from './Table';
import EditModalList from './EditModalList';
import EditInlineList from './EditInlineList';
import './List.css';

List.EditMode = {
  MODAL: 'MODAL',
  INLINE: 'INLINE',
};

export function List({ editMode, renderHeader, ...props }) {
  function PartiallyAppliedTable({ children, onCreate }) {
    return (
      <Table renderHeader={renderHeader} onCreate={onCreate}>
        {children}
      </Table>
    );
  }

  let el;
  switch (editMode) {
    case List.EditMode.MODAL:
      el = <EditModalList tableComponent={PartiallyAppliedTable} {...props} />;
      break;
    case List.EditMode.INLINE:
      el = <EditInlineList tableComponent={PartiallyAppliedTable} {...props} />;
      break;
    default:
  }

  return <div className={'List'}>{el}</div>;
}

List.defaultProps = {
  editMode: List.EditMode.INLINE,
};
