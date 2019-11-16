import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import CreateButton from '../template/Utilities/CreateButton';
import { useCreateEnvelope } from '../gql/envelopes';
import { EnvelopeModal } from './EnvelopeModal';
import { useDictionary } from '../template/Utilities/Lang';

export function CreateEnvelopeButton({ openRef }) {
  const [createEnvelope] = useCreateEnvelope();
  const { envelopes } = useDictionary();
  return (
    <ModalButton
      openRef={openRef}
      button={CreateButton}
      modal={props => (
        <EnvelopeModal
          title={envelopes.modal.createTitle}
          init={{ name: '', limit: null }}
          onSave={createEnvelope}
          {...props}
        />
      )}
    />
  );
}
