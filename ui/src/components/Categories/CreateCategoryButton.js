import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import CreateButton from '../template/Utilities/CreateButton';
import { useCreateCategory } from '../gql/categories';
import FormModal from '../template/Utilities/FormModal';
import useFormData from '../template/Utilities/useFormData';
import LabeledInput from '../template/Utilities/LabeledInput';

export function CreateCategoryButton() {
  const [createCategory] = useCreateCategory();
  return (
    <ModalButton
      button={CreateButton}
      modal={props => (
        <CategoryModal
          title="Add new envelope"
          init={{ name: '', limit: null }}
          onSave={createCategory}
          {...props}
        />
      )}
    />
  );
}

function CategoryModal({ init, ...props }) {
  const formData = useFormData({
    name: { $init: init.name },
  });
  return (
    <FormModal autoFocusRef={formData.name} {...props}>
      <LabeledInput
        label="Name"
        formData={formData.name}
        feedback="Provide name for the category"
      />
    </FormModal>
  );
}
