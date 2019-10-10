import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import CreateButton from '../template/Utilities/CreateButton';
import { useCreateCategory } from '../gql/categories';
import { CategoryModal } from './CategoryModal';

export function CreateCategoryButton() {
  const [createCategory] = useCreateCategory();
  return (
    <ModalButton
      button={CreateButton}
      modal={props => (
        <CategoryModal
          title="Add new envelope"
          init={{ name: '', envelope: { id: null } }}
          onSave={createCategory}
          {...props}
        />
      )}
    />
  );
}
