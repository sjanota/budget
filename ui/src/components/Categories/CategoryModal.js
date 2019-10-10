import React from 'react';
import FormModal from '../template/Utilities/FormModal';
import useFormData from '../template/Utilities/useFormData';
import LabeledInput from '../template/Utilities/LabeledInput';
import PropTypes from 'prop-types';
import { Form } from 'react-bootstrap';
import { useGetEnvelopes } from '../gql/envelopes';
import WithQuery from '../gql/WithQuery';

export function CategoryModal({ init, onSave, ...props }) {
  const query = useGetEnvelopes();
  const formData = useFormData({
    name: { $init: init.name },
    envelopeID: {
      $init: init.envelope.id,
    },
  });
  function handleSave() {
    if (!formData.changed()) {
      return;
    }
    const input = formData.value();
    onSave(input);
  }
  return (
    <FormModal autoFocusRef={formData.name} onSave={handleSave} {...props}>
      <WithQuery query={query}>
        {({ data }) => (
          <>
            <LabeledInput
              label="Name"
              formData={formData.name}
              feedback="Provide name for the category"
            />
            <Form.Group>
              <Form.Label>Envelope</Form.Label>
              <Form.Control
                as="select"
                ref={formData.envelopeID}
                defaultValue={formData.envelopeID.init}
              >
                {data.envelopes.map(({ id, name }) => (
                  <option key={id} value={id}>
                    {name}
                  </option>
                ))}
              </Form.Control>
            </Form.Group>
          </>
        )}
      </WithQuery>
    </FormModal>
  );
}

CategoryModal.propTypes = {
  init: PropTypes.shape({
    name: PropTypes.string,
    envelope: PropTypes.shape({ id: PropTypes.string }).isRequired,
  }),
  onSave: PropTypes.func.isRequired,
};
