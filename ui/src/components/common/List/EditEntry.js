import React, { useState } from 'react';
import PropTypes from 'prop-types';
import { CancelButton } from '../CancelButton';
import { SubmitButton } from '../SubmitButton';

export function EditEntry({ init, onCancel, onSubmit, renderEditEntry }) {
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
