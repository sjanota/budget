import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { removeFromList, addToList } from '../../../util/immutable';
import { ListEntry } from './ListEntry';
import { CancelButton } from '../CancelButton';
import { SubmitButton } from '../SubmitButton';

export default function EditInlineList({
  emptyValue,
  entries,
  onCreate,
  onDelete,
  onUpdate,
  renderEditEntry,
  renderEntry,
  tableComponent,
}) {
  const [isCreating, setIsCreating] = useState(false);
  const [editing, setEditing] = useState([]);

  const Table = tableComponent;

  return (
    <div>
      <Table onCreate={() => setIsCreating(true)}>
        {isCreating && (
          <EditEntry
            init={emptyValue}
            renderEditEntry={renderEditEntry}
            onCancel={() => setIsCreating(false)}
            onSubmit={onCreate}
          />
        )}
        {entries.map(entry =>
          editing.some(id => id === entry.id) ? (
            <EditEntry
              key={entry.id}
              init={entry}
              renderEditEntry={renderEditEntry}
              onCancel={() =>
                setEditing(editing => removeFromList(editing, entry.id))
              }
              onSubmit={input => onUpdate(entry.id, input)}
            />
          ) : (
            <ListEntry
              key={entry.id}
              entry={entry}
              renderEntry={renderEntry}
              onEdit={() => setEditing(editing => addToList(editing, entry.id))}
              onDelete={onDelete}
            />
          )
        )}
      </Table>
    </div>
  );
}

EditInlineList.propTypes = {
  emptyValue: PropTypes.any.isRequired,
  entries: PropTypes.array.isRequired,
  onCreate: PropTypes.func.isRequired,
  onDelete: PropTypes.func,
  onUpdate: PropTypes.func.isRequired,
  renderEditEntry: PropTypes.func,
  renderEntry: PropTypes.func.isRequired,
  tableComponent: PropTypes.element.isRequired,
};

function EditEntry({ init, onCancel, onSubmit, renderEditEntry }) {
  const [entry, setEntry] = useState(init);
  return (
    <tr>
      {renderEditEntry({ entry, setEntry })}
      <td>
        <CancelButton onClick={onCancel} />
        <SubmitButton
          onClick={() => {
            onSubmit(entry);
            onCancel();
          }}
        />
      </td>
    </tr>
  );
}

EditEntry.propTypes = {
  init: PropTypes.any.isRequired,
  onCancel: PropTypes.func.isRequired,
  onSubmit: PropTypes.func.isRequired,
  renderEditEntry: PropTypes.func.isRequired,
};
