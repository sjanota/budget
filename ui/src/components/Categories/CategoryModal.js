import React from 'react';
import FormModal from '../template/Utilities/FormModal';
import { useFormData } from '../template/Utilities/useFormData';
import { FormControl } from '../template/Utilities/FormControl';
import PropTypes from 'prop-types';
import { useGetEnvelopes } from '../gql/envelopes';
import { WithQuery } from '../gql/WithQuery';
import { Combobox } from '../template/Utilities/Combobox';
import { InlineFormControl } from '../template/Utilities/InlineFormControl';
import { useDictionary } from '../template/Utilities/Lang';

export function CategoryModal({ init, ...props }) {
  const query = useGetEnvelopes();
  const { categories } = useDictionary();
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
              label={categories.modal.labels.name}
              inline={9}
              formData={formData.name}
              feedback="Provide name"
            />
            <InlineFormControl
              size={9}
              label={categories.modal.labels.envelope}
            >
              <Combobox
                allowedValues={data.envelopes.map(({ id, name }) => ({
                  id,
                  label: name,
                }))}
                _ref={formData.envelopeID}
                defaultValue={formData.envelopeID.default()}
              />
            </InlineFormControl>
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
