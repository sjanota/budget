import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import CreateButton from '../template/Utilities/CreateButton';
import { useCreateCategory } from '../gql/categories';
import { CategoryModal } from './CategoryModal';
import { useDictionary } from '../template/Utilities/Lang';

export function CreateCategoryButton({ openRef }) {
  const [createCategory] = useCreateCategory();
  const { categories } = useDictionary();
  return (
    <ModalButton
      openRef={openRef}
      button={CreateButton}
      modal={props => (
        <CategoryModal
          title={categories.modal.createTitle}
          init={{ name: '', envelope: { id: null } }}
          onSave={createCategory}
          {...props}
        />
      )}
    />
  );
}
