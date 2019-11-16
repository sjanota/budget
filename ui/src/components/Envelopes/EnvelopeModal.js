import React from 'react';
import Amount from '../../model/Amount';
import { FormControl } from '../template/Utilities/FormControl';
import { OptionalFormControl } from '../template/Utilities/OptionalFormControl';
import FormModal from '../template/Utilities/FormModal';
import { useFormData } from '../template/Utilities/useFormData';
import PropTypes from 'prop-types';
import * as model from '../../model/propTypes';
import { useDictionary } from '../template/Utilities/Lang';

export function EnvelopeModal({ init, ...props }) {
  const { envelopes } = useDictionary();
  const formData = useFormData({
    name: { $init: init.name },
    limit: { $init: Amount.format(init.limit, false), $process: Amount.parse },
  });
  return (
    <FormModal formData={formData} autoFocusRef={formData.name} {...props}>
      <FormControl
        label={envelopes.modal.labels.name}
        inline={9}
        feedback="Provide a name for the envelope"
        required
        formData={formData.name}
      />
      <OptionalFormControl
        initEnabled={!!init.limit}
        inline={9}
        label={envelopes.modal.labels.limit}
        feedback="Provide a limit for the envelope"
        type="number"
        required
        formData={formData.limit}
        step="0.01"
      />
    </FormModal>
  );
}

EnvelopeModal.propTypes = {
  init: PropTypes.shape({
    name: PropTypes.string,
    limit: model.Amount,
  }).isRequired,
  onSave: PropTypes.func.isRequired,
};
