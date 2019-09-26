import React from 'react';
import PropTypes from 'prop-types';
import { DeleteButton } from '../DeleteButton';
import { EditButton } from '../EditButton';

export function ListEntry({ entry, onEdit, onDelete, renderEntry }) {
  return (
    <tr>
      {renderEntry({ entry })}
      <td>
        {onDelete && <DeleteButton onClick={() => onDelete(entry.id)} />}
        <EditButton onClick={onEdit} />
      </td>
    </tr>
  );
}

ListEntry.propTypes = {
  onDelete: PropTypes.func,
  onEdit: PropTypes.func.isRequired,
  renderEntry: PropTypes.func.isRequired,
  entry: PropTypes.any.isRequired,
};
