import React from 'react';
import { CreateButton } from '../CreateButton';
import PropTypes from 'prop-types';

export function ListHeader({ onCreate, renderHeader }) {
  return (
    <thead className={'thead-dark'}>
      <tr>
        {renderHeader({})}
        <th>
          Actions
          <CreateButton onClick={onCreate} />
        </th>
      </tr>
    </thead>
  );
}

ListHeader.propTypes = {
  onCreate: PropTypes.func.isRequired,
  renderHeader: PropTypes.func.isRequired,
};
