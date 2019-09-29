import React from 'react';
import PropTypes from 'prop-types';
import { useQuery } from '@apollo/react-hooks';
import { useBudget } from '../context/budget/budget';
import { Form, Row, Col } from 'react-bootstrap';
import { QUERY_CATEGORIES } from '../CategoriesList/CategoriesList.gql';
import * as MoneyAmount from '../../model/MoneyAmount';
import { MoneyAmount as MoneyAmountPropType } from '../../model/propTypes';

export function EditModalContentEntry({ entry, idx, setEntry }) {
  const { id: budgetID } = useBudget();
  const { loading, error, data } = useQuery(QUERY_CATEGORIES, {
    variables: { budgetID },
  });
  if (loading) return <p>Loading...</p>;
  if (error) {
    console.error(error);
    return <p>Error :(</p>;
  }
  return (
    <Row>
      <Col>
        <Form.Control
          as="select"
          value={entry.categoryID || (entry.category && entry.category.id)}
          onChange={e => setEntry(idx, { categoryID: e.target.value })}
        >
          <option></option>
          {data.categories
            .sort((c1, c2) => c1.name.localeCompare(c2))
            .map(category => (
              <option key={category.id} value={category.id}>
                {category.name}
              </option>
            ))}
        </Form.Control>
      </Col>
      <Col>
        <Form.Control
          type="number"
          placeholder="Kwota"
          value={MoneyAmount.format(entry.balance)}
          onChange={e => setEntry(idx, { balance: e.target.value })}
          onBlur={() =>
            setEntry(idx, { balance: MoneyAmount.parse(entry.balance) })
          }
        />
      </Col>
    </Row>
  );
}

EditModalContentEntry.propTypes = {
  entry: PropTypes.shape({
    categoryID: PropTypes.any,
    category: PropTypes.shape({
      id: PropTypes.any.isRequired,
    }),
    balance: PropTypes.oneOfType([PropTypes.string, MoneyAmountPropType]),
  }),
  idx: PropTypes.number.isRequired,
  setEntry: PropTypes.func.isRequired,
};
