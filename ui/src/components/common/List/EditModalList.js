import React, { useState, useRef } from 'react';
import { Modal } from 'react-bootstrap';
import PropTypes from 'prop-types';
import { ListEntry } from './ListEntry';
import './List.css';

export default function EditModalList({
  emptyValue,
  entries,
  onCreate,
  onDelete,
  onUpdate,
  renderEntry,
  tableComponent,
  renderModalContent,
}) {
  const [editing, setEditing] = useState(null);
  const autoFocusRef = useRef();

  const Table = tableComponent;
  return (
    <>
      <Modal
        show={editing !== null}
        onHide={() => setEditing(null)}
        onEntered={() => autoFocusRef.current.focus()}
      >
        {renderModalContent({
          init: editing,
          onCancel: () => setEditing(null),
          onSubmit:
            editing === emptyValue
              ? onCreate
              : input => onUpdate(editing.id, input),
          autoFocusRef,
        })}
      </Modal>
      <Table onCreate={() => setEditing(emptyValue)}>
        {entries.map(entry => (
          <ListEntry
            key={entry.id}
            entry={entry}
            renderEntry={renderEntry}
            onEdit={() => setEditing(entry)}
            onDelete={onDelete}
          />
        ))}
      </Table>
    </>
  );
}

EditModalList.propTypes = {
  emptyValue: PropTypes.any.isRequired,
  entries: PropTypes.array.isRequired,
  onCreate: PropTypes.func.isRequired,
  onDelete: PropTypes.func,
  onUpdate: PropTypes.func.isRequired,
  renderEntry: PropTypes.func.isRequired,
  renderHeader: PropTypes.func.isRequired,
  renderModalContent: PropTypes.func.isRequired,
  tableComponent: PropTypes.elementType.isRequired,
};
