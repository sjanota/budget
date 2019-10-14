import React from 'react';
import FormModal from '../template/Utilities/FormModal';
import { useFormData } from '../template/Utilities/createFormData';
import { FormControl } from '../template/Utilities/FormControl';
import PropTypes from 'prop-types';
import { useGetEnvelopes } from '../gql/envelopes';
import { WithQuery } from '../gql/WithQuery';

export function CategoryModal({ init, ...props }) {
  const query = useGetEnvelopes();
  const formData = useFormData({
    name: { $init: init.name },
    envelopeID: {
      $init: init.envelope.id,
    },
  });
  return (
    <FormModal autoFocusRef={formData.name} formData={formData} {...props}>
      <WithQuery query={query}>
        {({ data }) => (
          <>
            <FormControl
              label="Name"
              inline={9}
              formData={formData.name}
              feedback="Provide name"
            />
            <FormControl
              label="Envelope"
              inline={9}
              formData={formData.envelopeID}
              feedback="Provide envelope"
              as="select"
            >
              {data.envelopes.map(({ id, name }) => (
                <option key={id} value={id}>
                  {name}
                </option>
              ))}
            </FormControl>
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
