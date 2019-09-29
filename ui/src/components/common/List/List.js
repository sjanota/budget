import React, { useState, useRef } from 'react';
import { Table, Modal } from 'react-bootstrap';
import PropTypes from 'prop-types';
import { cloneDeep } from 'apollo-utilities';
import { removeFromList, addToList } from '../../../util/immutable';
import { ListHeader } from './ListHeader';
import { EditEntry } from './EditEntry';
import { ListEntry } from './ListEntry';

const EDIT = {
  INLINE: 'EDIT_INLINE',
  MODAL: 'MODAL',
};

export default function List({
  emptyValue,
  entries,
  onCreate,
  onDelete,
  onUpdate,
  renderEditEntry,
  renderEntry,
  renderHeader,
  renderModalContent,
  editMode,
}) {
  const [isCreating, setIsCreating] = useState(false);
  const [editing, setEditing] = useState([]);
  const autoFocusRef1 = useRef();
  const autoFocusRef2 = useRef();

  return (
    <div className={'ExpensesList'}>
      {editMode === EDIT.MODAL && (
        <>
          <Modal
            show={isCreating}
            onHide={() => setIsCreating(false)}
            onEntered={() => autoFocusRef1.current.focus()}
          >
            {renderModalContent({
              init: cloneDeep(emptyValue),
              onCancel: () => setIsCreating(false),
              onSubmit: onCreate,
              autoFocusRef: autoFocusRef1,
            })}
          </Modal>
          <Modal
            show={editing.length > 0}
            onHide={() => setEditing([])}
            onEntered={() => autoFocusRef2.current.focus()}
          >
            {renderModalContent({
              init: entries.find(entry => entry.id === editing[0]),
              onCancel: () => setEditing([]),
              onSubmit: input => onUpdate(editing[0], input),
              autoFocusRef: autoFocusRef2,
            })}
          </Modal>
        </>
      )}
      <Table striped bordered hover size={'sm'}>
        <ListHeader
          onCreate={() => setIsCreating(true)}
          renderHeader={renderHeader}
        />
        <tbody>
          {isCreating && editMode === EDIT.INLINE && (
            <EditEntry
              init={cloneDeep(emptyValue)}
              renderEditEntry={renderEditEntry}
              onCancel={() => setIsCreating(false)}
              onSubmit={onCreate}
            />
          )}
          {entries.map(entry =>
            editing.some(id => id === entry.id) && editMode === EDIT.INLINE ? (
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
  renderEditEntry: PropTypes.func,
  renderEntry: PropTypes.func.isRequired,
  renderHeader: PropTypes.func.isRequired,
  renderModalContent: PropTypes.func,
  editMode: PropTypes.oneOf(Object.values(EDIT)).isRequired,
};

List.defaultProps = {
  editMode: EDIT.INLINE,
};

List.EditMode = EDIT;
