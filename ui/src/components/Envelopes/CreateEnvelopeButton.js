import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import CreateButton from '../template/Utilities/CreateButton';
import { useCreateEnvelope } from '../gql/envelopes';
import { EnvelopeModal } from './EnvelopeModal';

export function CreateEnvelopeButton() {
  const [createEnvelope] = useCreateEnvelope();
  return (
    <ModalButton
      button={CreateButton}
      modal={props => (
        <EnvelopeModal
          title="Add new envelope"
          init={{ name: '', limit: null }}
          onSave={createEnvelope}
          {...props}
        />
      )}
    />
  );
}
