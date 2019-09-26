import React, { useState } from 'react';
import { Table } from 'react-bootstrap';
import { CreateButton } from '../common/CreateButton';
import PropTypes from 'prop-types';
import { CancelButton } from '../common/CancelButton';
import { SubmitButton } from '../common/SubmitButton';
import { cloneDeep } from 'apollo-utilities';
import { removeFromList, addToList } from '../../util/immutable';
import { DeleteButton } from '../common/DeleteButton';
import { EditButton } from '../common/EditButton';

export default function List({
  entries,
  onCreate,
  onUpdate,
  onDelete,
  renderHeader,
  renderEntry,
  renderEditEntry,
  emptyValue,
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
  onCreate: PropTypes.func.isRequired,
  onDelete: PropTypes.func.isRequired,
  onUpdate: PropTypes.func.isRequired,
  entries: PropTypes.array.isRequired,
  renderHeader: PropTypes.func.isRequired,
  renderEntry: PropTypes.func.isRequired,
  renderEditEntry: PropTypes.func.isRequired,
  subscriptionConfig: PropTypes.any.isRequired,
  emptyValue: PropTypes.any.isRequired,
};

function ListHeader({ onCreate, renderHeader }) {
  return (
    <thead className={'thead-dark'}>
      <tr>
        {renderHeader({})}
        <th>
          Actions
          <CreateButton onClick={onCreate} />
        </th>
      </tr>
    </thead>
  );
}

ListHeader.propTypes = {
  onCreate: PropTypes.func.isRequired,
  renderHeader: PropTypes.func.isRequired,
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

function ListEntry({ entry, onEdit, onDelete, renderEntry }) {
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
