import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import PropTypes from 'prop-types';
import { cloneDeep } from 'apollo-utilities';
import { removeFromList, addToList } from '../../../util/immutable';
import { ListHeader } from './ListHeader';
import { EditEntry } from './EditEntry';
import { ListEntry } from './ListEntry';

export default function List({
  emptyValue,
  entries,
  onCreate,
  onDelete,
  onUpdate,
  renderEditEntry,
  renderEntry,
  renderHeader,
}) {
  const [isCreating, setIsCreating] = useState(false);
  const [editing, setEditing] = useState([]);

  return (
    <div className={'ExpensesList'}>
      <Table striped bordered hover size={'sm'}>
        <ListHeader
          onCreate={() => setIsCreating(true)}
          renderHeader={renderHeader}
        />
        <tbody>
          {isCreating && (
            <EditEntry
              init={cloneDeep(emptyValue)}
              renderEditEntry={renderEditEntry}
              onCancel={() => setIsCreating(false)}
              onSubmit={onCreate}
            />
          )}
          {entries.map(entry =>
            editing.some(id => id === entry.id) ? (
              <EditEntry
                key={entry.id}
                init={cloneDeep(entry)}
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
                onEdit={() =>
                  setEditing(editing => addToList(editing, entry.id))
                }
                onDelete={onDelete}
              />
            )
          )}
        </tbody>
      </Table>
    </div>
  );
}

List.propTypes = {
  emptyValue: PropTypes.any.isRequired,
  entries: PropTypes.array.isRequired,
  onCreate: PropTypes.func.isRequired,
  onDelete: PropTypes.func,
  onUpdate: PropTypes.func.isRequired,
  renderEditEntry: PropTypes.func.isRequired,
  renderEntry: PropTypes.func.isRequired,
  renderHeader: PropTypes.func.isRequired,
};
