import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import EditTableButton from '../template/Utilities/EditTableButton';
import { useUpdateEnvelope } from '../gql/envelopes';
import { EnvelopeModal } from './EnvelopeModal';
import PropTypes from 'prop-types';

export function UpdateEnvelopeButton({ envelope }) {
  const [updateEnvelope] = useUpdateEnvelope();
  const onSave = input => {
    updateEnvelope(envelope.id, input);
  };
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <EnvelopeModal
          title="Edit envelope"
          init={envelope}
          onSave={onSave}
          {...props}
        />
      )}
    />
  );
}

UpdateEnvelopeButton.propTypes = {
  envelope: PropTypes.shape({
    id: PropTypes.string.isRequired,
  }),
};
