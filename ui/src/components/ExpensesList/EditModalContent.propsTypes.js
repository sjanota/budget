import { shape, string, any, oneOfType } from 'prop-types';
import { MoneyAmount } from '../../model/propTypes';

export const StateEntry = shape({
  categoryID: any,
  category: shape({
    id: any,
  }),
  balance: oneOfType([string, MoneyAmount]),
});
