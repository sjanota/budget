import React from 'react';
import PropTypes from 'prop-types';
import { DeleteButton } from '../common/DeleteButton';
import { EditButton } from '../common/EditButton';

export function ListEntry({ entry, onEdit, onDelete, renderEntry }) {
  return (
    <tr>
      {renderEntry({ entry })}
      <td>
        <DeleteButton onClick={() => onDelete(entry.id)} />
        <EditButton onClick={onEdit} />
      </td>
    </tr>
  );
}

ListEntry.propTypes = {
  onDelete: PropTypes.func.isRequired,
  onEdit: PropTypes.func.isRequired,
  renderEntry: PropTypes.func.isRequired,
  entry: PropTypes.any.isRequired,
};
