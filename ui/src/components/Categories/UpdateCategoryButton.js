import React from 'react';
import ModalButton from '../template/Utilities/ModalButton';
import EditTableButton from '../template/Utilities/EditTableButton';
import { CategoryModal } from './CategoryModal';
import PropTypes from 'prop-types';
import { useUpdateCategory } from '../gql/categories';

export function UpdateCategoryButton({ category }) {
  const [updateEnvelope] = useUpdateCategory();
  const onSave = input => {
    updateEnvelope(category.id, input);
  };
  return (
    <ModalButton
      button={EditTableButton}
      modal={props => (
        <CategoryModal
          title="Edit category"
          init={category}
          onSave={onSave}
          {...props}
        />
      )}
    />
  );
}

UpdateCategoryButton.propTypes = {
  category: PropTypes.shape({
    id: PropTypes.string.isRequired,
  }),
};
