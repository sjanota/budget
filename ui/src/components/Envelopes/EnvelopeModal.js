import React from 'react';
import Amount from '../../model/Amount';
import LabeledInput from '../template/Utilities/LabeledInput';
import OptionalLabeledInput from '../template/Utilities/OptionalLabeledInput';
import FormModal from '../template/Utilities/FormModal';
import useFormData from '../template/Utilities/useFormData';
import PropTypes from 'prop-types';
import * as model from '../../model/propTypes';

export function EnvelopeModal({ init, onSave, ...props }) {
  const formData = useFormData({
    name: { $init: init.name },
    limit: { $init: Amount.format(init.limit), $process: Amount.parse },
  });
  function handleSave() {
    if (!formData.changed()) {
      return;
    }
    const input = formData.value();
    onSave(input);
  }
  return (
    <FormModal onSave={handleSave} autoFocusRef={formData.name} {...props}>
      <LabeledInput
        label="Name"
        feedback="Provide a name for the envelope"
        required
        formData={formData.name}
      />
      <OptionalLabeledInput
        initEnabled={!!formData.limit.init}
        label="Limit"
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
